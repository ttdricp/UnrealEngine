package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "order_management_final"
)

type orderManagementServer struct {
	pb.UnimplementedOrderManagementServiceServer // Embed the generated UnimplementedOrderManagementServiceServer
	channel                                      *amqp.Channel
	queue                                        amqp.Queue
}

var orders []*pb.Order

func (s *orderManagementServer) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	orders = append(orders, req)
	s.publishOrder(req)
	return req, nil
}

func (s *orderManagementServer) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	for _, order := range orders {
		if order.Id == req.Id {
			order.Name = req.Name
			order.Address = req.Address
			return order, nil
		}
	}
	return nil, fmt.Errorf("Order not found")
}

func (s *orderManagementServer) ReadOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	for _, order := range orders {
		if order.Id == req.Id {
			return order, nil
		}
	}
	return nil, fmt.Errorf("Order not found")
}

func (s *orderManagementServer) DeleteOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	for i, order := range orders {
		if order.Id == req.Id {
			orders = append(orders[:i], orders[i+1:]...)
			return order, nil
		}
	}
	return nil, fmt.Errorf("Order not found")
}

func (s *orderManagementServer) publishOrder(order *pb.Order) {
	msg := fmt.Sprintf("Order ID: %d, Name: %s, Address: %s", order.Id, order.Name, order.Address)
	err := s.channel.Publish(
		"",           // exchange
		s.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		log.Printf("Failed to publish order: %v", err)
	}
}

func setupRabbitMQ() (*amqp.Channel, amqp.Queue, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("Failed to connect to RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("Failed to open a RabbitMQ channel: %v", err)
	}

	queue, err := channel.QueueDeclare(
		"orders", // queue name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("Failed to declare a RabbitMQ queue: %v", err)
	}

	return channel, queue, nil
}

func main() {
	channel, queue, err := setupRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	orderManagementServer := &orderManagementServer{
		channel: channel,
		queue:   queue,
	}
	pb.RegisterOrderManagementServiceServer(s, orderManagementServer)

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
