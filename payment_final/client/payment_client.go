package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	paymentpb "payment_final"

	"github.com/streadway/amqp"
)

// RabbitMQ configuration
const (
	rabbitMQURL     = "amqp://guest:guest@localhost:5672/"
	rabbitMQQueue   = "payment_queue"
	rabbitMQDurable = true
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := paymentpb.NewPaymentServiceClient(conn)

	// Payment req
	req := &paymentpb.PaymentRequest{
		UserId: "Kanych",
		Amount: 100.0,
	}

	// Call the ProcessPayment gRPC method
	res, err := client.ProcessPayment(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to process payment: %v", err)
	}

	// Print the response from the server
	log.Println("Payment of $%.2f has been made by user %s with status %s", req.Amount, req.UserId, res.Status)

	isSuccessful := simulatePaymentProcessing()

	// Prepare payment details for storage and publishing
	paymentDetails := fmt.Sprintf("Payment of $%.2f has been made by user %s with status %s", req.Amount, req.UserId, getPaymentStatus(isSuccessful))

	// Store payment details in the database
	storePaymentInDatabase(paymentDetails)

	// Publish payment details to RabbitMQ
	publishPaymentToRabbitMQ(paymentDetails)

	log.Println("Payment processing completed.")
}

func simulatePaymentProcessing() bool {
	// Just simulation of the payment
	// Here, we randomly generate a boolean value to simulate a successful or failed payment
	rand.Seed(time.Now().UnixNano())
	isSuccessful := rand.Float32() < 0.8 // 80% chance of success

	return isSuccessful
}

func getPaymentStatus(isSuccessful bool) string {
	if isSuccessful {
		return "success"
	}
	return "failure"
}

func storePaymentInDatabase(paymentDetails string) {
	log.Printf("Payment details stored in the database: %s", paymentDetails)
}

func publishPaymentToRabbitMQ(paymentDetails string) {
	// Create a connection to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	_, err = ch.QueueDeclare(
		rabbitMQQueue,   // queue name
		rabbitMQDurable, // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Publish the payment details to the queue
	err = ch.Publish(
		"",            // exchange
		rabbitMQQueue, // routing key (queue name)
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(paymentDetails),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message to RabbitMQ: %v", err)
	}

	log.Printf("Payment details published to RabbitMQ: %s", paymentDetails)
}
