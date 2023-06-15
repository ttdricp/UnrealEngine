package main

import (
	"context"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"

	paymentpb "payment_final"
)

type Payment struct {
	gorm.Model
	PaymentID string
	UserID    string
	Amount    float64
	Status    string
}

type paymentServer struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {
	paymentID := req.PaymentId
	userID := req.UserId

	status := "success"

	res := &paymentpb.PaymentResponse{
		PaymentId: paymentID,
		UserId:    userID,
		Status:    status,
	}

	return res, nil
}

func main() {

	db, err := gorm.Open(sqlite.Open("payments.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(&Payment{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate table: %v", err)
	}

	// Set up RabbitMQ connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Start gRPC server
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(s, &paymentServer{})

	log.Println("Starting Payment Service server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Println("Starting Payment Service server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)

	}
}
