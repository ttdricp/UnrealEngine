package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "Final/protoc"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func (s *userServiceServer) Register(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	userID := generateUserID()

	// Store the user information in the database
	err := s.storeUser(userID, req.GetUsername(), req.GetPassword(), req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err)
	}

	// Prepare the registration response
	res := &pb.RegistrationResponse{
		UserId: userID,
	}

	return res, nil
}

func (s *userServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	// Retrieve the user from the database
	user, err := s.getUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	if user == nil || user.Password != password {
		return &pb.LoginResponse{
			Success: false,
			Message: "Invalid username or password",
		}, nil
	}

	return &pb.LoginResponse{
		Success: true,
		Message: "User logged in successfully",
	}, nil
}

func (s *userServiceServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	// Retrieve the email address and new password from the request
	email := req.GetEmail()
	newPassword := req.GetNewPassword()

	err := s.resetUserPassword(email, newPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to reset password: %v", err)
	}

	res := &pb.ResetPasswordResponse{
		Message: "Password reset successfully",
	}

	return res, nil
}

func (s *userServiceServer) resetUserPassword(email, newPassword string) error {
	// Check if the user exists in the database
	stmtExists, err := s.db.Prepare("SELECT id FROM users WHERE email = $1")
	if err != nil {
		return err
	}
	defer stmtExists.Close()

	var userID string
	err = stmtExists.QueryRow(email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user with email '%s' not found", email)
		}
		return err
	}

	// Prepare the SQL statement for updating the user's password
	stmtUpdate, err := s.db.Prepare("UPDATE users SET password = $1 WHERE email = $2")
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	// Execute the SQL statement to update the user's password
	_, err = stmtUpdate.Exec(newPassword, email)
	if err != nil {
		return err
	}

	return nil
}

func generateUserID() string {
	id := uuid.New()
	userID := id.String()

	return userID
}

func (s *userServiceServer) storeUser(userID, username, password, email string) error {
	stmt, err := s.db.Prepare("INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, username, password, email)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceServer) getUserByUsernameAndPassword(username, password string) (*pb.User, error) {
	stmt, err := s.db.Prepare("SELECT id, username, password, email FROM users WHERE username = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	var userID, retrievedUsername, retrievedPassword, email string
	err = row.Scan(&userID, &retrievedUsername, &retrievedPassword, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if retrievedPassword != password {
		return nil, nil
	}

	user := &pb.User{
		Id:       userID,
		Username: retrievedUsername,
		Password: retrievedPassword,
		Email:    email,
	}

	return user, nil
}

func startGRPCServer() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=12345678 dbname=server sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	server := &userServiceServer{
		db: db,
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to create listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, server)

	log.Println("Server started. Listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := startGRPCServer(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
