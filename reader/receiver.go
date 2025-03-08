package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func main() {
	// Load environment variables
	connectionString := os.Getenv("SERVICEBUS_CONNECTION_STRING")
	queueName := os.Getenv("SERVICEBUS_QUEUE_NAME")

	if connectionString == "" || queueName == "" {
		log.Fatal("Missing SERVICEBUS_CONNECTION_STRING or SERVICEBUS_QUEUE_NAME")
	}

	// Create client
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("❌ Failed to create Service Bus client: %v", err)
	}
	defer client.Close(context.Background())

	// Create receiver
	receiver, err := client.NewReceiverForQueue(queueName, nil)
	if err != nil {
		log.Fatalf("❌ Failed to create receiver: %v", err)
	}
	defer receiver.Close(context.Background())

	// Set up a context that listens for system interrupts (Ctrl+C)
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Goroutine to handle shutdown signals
	go func() {
		<-sigCh
		fmt.Println("\n⚠️  Received termination signal. Shutting down...")
		cancel() // Cancel the context to exit message loop
	}()

	fmt.Println("🚀 Listening for messages. Press Ctrl+C to stop.")

	// Continuous message processing loop
	for {
		// Check if context is canceled
		if ctx.Err() != nil {
			fmt.Println("🛑 Stopping message receiver...")
			break
		}

		// Receive messages with a small batch size to ensure real-time processing
		messages, err := receiver.ReceiveMessages(ctx, 1, nil)
		if err != nil {
			log.Printf("⚠️  Error receiving messages: %v", err)
			continue
		}

		// Process messages
		for _, msg := range messages {
			fmt.Printf("📩 Received message: %s\n", string(msg.Body))

			// Complete message after successful processing
			err := receiver.CompleteMessage(ctx, msg, nil)
			if err != nil {
				log.Printf("⚠️  Failed to complete message: %v", err)
			}
		}
	}

	fmt.Println("✅ Message receiver stopped gracefully.")
}
