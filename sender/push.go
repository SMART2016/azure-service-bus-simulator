package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"log"
	"os"
)

func main() {
	// Get connection string from environment variable
	connectionString := os.Getenv("SERVICEBUS_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("Missing SERVICEBUS_CONNECTION_STRING environment variable")
	}

	// Get queue or topic name
	queueName := os.Getenv("SERVICEBUS_QUEUE_NAME")
	if queueName == "" {
		log.Fatal("Missing SERVICEBUS_QUEUE_NAME environment variable")
	}

	// Set up a context with timeout
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	//defer cancel()

	// Create a Service Bus client
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("Failed to create Service Bus client: %v", err)
	}
	defer client.Close(ctx)

	// Create a sender
	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		log.Fatalf("Failed to create sender: %v", err)
	}
	defer sender.Close(ctx)

	// Send a message
	sessionID := "session-1"
	message := azservicebus.Message{
		Body:      []byte("Hello from local Azure Service Bus Emulator!"),
		SessionID: &sessionID, // Required for session-enabled queues
	}
	err = sender.SendMessage(ctx, &message, nil)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("✅ Successfully sent message to queue: %s\n", queueName)
}
