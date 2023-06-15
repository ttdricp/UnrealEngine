package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	pb "order_management_final"
)

type orderManagementServer struct {
	pb.UnimplementedOrderManagementServiceServer // Embed the generated UnimplementedOrderManagementServiceServer
	channel                                      *amqp.Channel
	queue                                        amqp.Queue
}

func TestServer(t *testing.T) {
	go func() {
		if err := runServer(); err != nil {
			t.Fatalf("Failed to start server: %v", err)
		}
	}()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Test connection to RabbitMQ
	channel, _, err := setupRabbitMQ()
	if err != nil {
		t.Fatalf("Failed to setup RabbitMQ: %v", err)
	}
	defer channel.Close()

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

func runServer() error {
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
	return nil
}

func TestCreateOrder(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementServiceClient(conn)

	order := &pb.Order{
		Id:      3,
		Name:    "Test Order",
		Address: "Test Address",
	}

	createdOrder, err := client.CreateOrder(context.Background(), order)
	assert.NoError(t, err)
	assert.Equal(t, order, createdOrder)
}

func TestReadOrder(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementServiceClient(conn)

	order := &pb.Order{
		Id:      1,
		Name:    "Test Order",
		Address: "Test Address",
	}

	_, err = client.CreateOrder(context.Background(), order)
	assert.NoError(t, err)

	readOrder, err := client.ReadOrder(context.Background(), &pb.Order{Id: 1})
	assert.NoError(t, err)
	assert.Equal(t, order, readOrder)
}

func TestUpdateOrder(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementServiceClient(conn)

	order := &pb.Order{
		Id:      1,
		Name:    "Test Order",
		Address: "Test Address",
	}

	_, err = client.CreateOrder(context.Background(), order)
	assert.NoError(t, err)

	updatedOrder := &pb.Order{
		Id:      1,
		Name:    "Updated Order",
		Address: "Updated Address",
	}

	_, err = client.UpdateOrder(context.Background(), updatedOrder)
	assert.NoError(t, err)

	readOrder, err := client.ReadOrder(context.Background(), &pb.Order{Id: 1})
	assert.NoError(t, err)
	assert.Equal(t, updatedOrder, readOrder)
}

func TestDeleteOrder(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementServiceClient(conn)

	order := &pb.Order{
		Id:      1,
		Name:    "Test Order",
		Address: "Test Address",
	}

	_, err = client.CreateOrder(context.Background(), order)
	assert.NoError(t, err)

	_, err = client.DeleteOrder(context.Background(), &pb.Order{Id: 1})
	assert.NoError(t, err)

	readOrder, err := client.ReadOrder(context.Background(), &pb.Order{Id: 1})
	assert.Error(t, err)
	assert.Nil(t, readOrder)
}
