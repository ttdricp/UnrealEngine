package main

import (
	pb "Final/protoc"
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=12345678 dbname=server sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	assert.Nil(t, err, "Failed to ping the database")
}

func TestGRPCServer(t *testing.T) {
	// Start the gRPC server in a separate goroutine
	go func() {
		if err := startGRPCServer(); err != nil {
			t.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a UserService client
	client := pb.NewUserServiceClient(conn)

	// Example: Test user registration
	registerReq := &pb.RegistrationRequest{
		Username: "slava",
		Password: "password123",
		Email:    "slava@example.com",
	}
	_, err = client.Register(context.Background(), registerReq)
	if err != nil {
		t.Fatalf("Registration failed: %v", err)
	}
}
