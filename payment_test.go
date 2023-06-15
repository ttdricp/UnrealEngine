package main

import (
	"context"
	"testing"

	paymentpb "payment_final"
)

func TestProcessPayment(t *testing.T) {
	// Create a new payment server
	server := &paymentServer{}

	// Create a test request
	req := &paymentpb.PaymentRequest{
		PaymentId: "payment_123",
		UserId:    "user_456",
	}

	// Invoke the ProcessPayment method
	res, err := server.ProcessPayment(context.Background(), req)
	if err != nil {
		t.Errorf("ProcessPayment returned an error: %v", err)
	}

	// Verify the response
	if res.PaymentId != req.PaymentId {
		t.Errorf("Payment ID mismatch, expected %s but got %s", req.PaymentId, res.PaymentId)
	}

	if res.UserId != req.UserId {
		t.Errorf("User ID mismatch, expected %s but got %s", req.UserId, res.UserId)
	}

	if res.Status != "success" {
		t.Errorf("Status mismatch, expected 'success' but got %s", res.Status)
	}
}
