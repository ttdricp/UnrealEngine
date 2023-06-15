package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "Final/protoc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	registerReq := &pb.RegistrationRequest{
		Username: "slava",
		Password: promptPassword("Enter the password: "),
		Email:    "slava@example.com",
	}
	registerRes, err := client.Register(context.Background(), registerReq)
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}
	log.Printf("Registration successful. User ID: %s", registerRes.GetUserId())

	loginReq := &pb.LoginRequest{
		Username: "slava",
		Password: registerReq.GetPassword(),
	}
	loginRes, err := client.Login(context.Background(), loginReq)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login successful. User ID: %s", loginRes.GetUserId())

	// Request password reset
	resetPasswordReq := &pb.ResetPasswordRequest{
		Email:       "slava@example.com",
		NewPassword: promptPassword("Enter the new password: "),
	}
	_, err = client.ResetPassword(context.Background(), resetPasswordReq)
	if err != nil {
		log.Fatalf("Password reset failed: %v", err)
	}

	loginReqNew := &pb.LoginRequest{
		Username: "slava",
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
