package main

import (
	"context"
	"log"
	"time"

	pb "order_management_final"

	"google.golang.org/grpc"
)

func createOrder(client pb.OrderManagementServiceClient) {
	req := &pb.Order{
		Id:      1,
		Name:    "Tair",
		Address: "Mangilik",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("CreateOrder failed: %v", err)
	}

	log.Printf("Created Order: %v", res)
}

func updateOrder(client pb.OrderManagementServiceClient) {
	req := &pb.Order{
		Id:      1,
		Name:    "Darkhan",
		Address: "Sultan Hazret",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateOrder(ctx, req)
	if err != nil {
		log.Fatalf("UpdateOrder failed: %v", err)
	}

	log.Printf("Updated Order: %v", res)
}

func readOrder(client pb.OrderManagementServiceClient) {
	req := &pb.Order{
		Id: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ReadOrder(ctx, req)
	if err != nil {
		log.Fatalf("ReadOrder failed: %v", err)
	}

	log.Printf("Read Order: %v", res)
}

func deleteOrder(client pb.OrderManagementServiceClient) {
	req := &pb.Order{
		Id: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteOrder(ctx, req)
	if err != nil {
		log.Fatalf("DeleteOrder failed: %v", err)
	}

	log.Printf("Deleted Order: %v", res)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementServiceClient(conn)

	createOrder(client)
	updateOrder(client)
	readOrder(client)
	deleteOrder(client)
}
