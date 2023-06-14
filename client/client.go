package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "Final/protoc" // Update with your actual package path

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new UserService client
	client := pb.NewUserServiceClient(conn)

	// Perform user registration
	registerReq := &pb.RegistrationRequest{
		Username: "slava1",
		Password: promptPassword("Enter the password: "),
		Email:    "slava1@example.com",
	}
	registerRes, err := client.Register(context.Background(), registerReq)
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}
	log.Printf("Registration successful. User ID: %s", registerRes.GetUserId())

	// Perform user login
	loginReq := &pb.LoginRequest{
		Username: "slava1",
		Password: registerReq.GetPassword(),
	}
	loginRes, err := client.Login(context.Background(), loginReq)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login successful. User ID: %s", loginRes.GetUserId())

	// Request password reset
	// Request password reset
	// Request password reset
	resetPasswordReq := &pb.ResetPasswordRequest{
		Email:       "slava1@example.com",
		NewPassword: promptPassword("Enter the new password: "),
	}
	_, err = client.ResetPassword(context.Background(), resetPasswordReq)
	if err != nil {
		log.Fatalf("Password reset failed: %v", err)
	}

	// Perform user login with the new password
	loginReqNew := &pb.LoginRequest{
		Username: "slava1",
		Password: promptPassword("Login with the new password: "),
	}
	_, err = client.Login(context.Background(), loginReqNew)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login successful with new password.")
	log.Println("Password reset successful.")
}

func promptPassword(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}

	// Remove the newline character from the password
	password = password[:len(password)-1]

	return password
}
