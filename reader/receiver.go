package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	// Receive messages
	messages, err := receiver.ReceiveMessages(context.Background(), 5, nil)
	if err != nil {
		log.Fatalf("❌ Failed to receive messages: %v", err)
	}

	// Process messages
	for _, msg := range messages {
		fmt.Printf("📩 Received message: %s\n", string(msg.Body))
		err := receiver.CompleteMessage(context.Background(), msg, nil)
		if err != nil {
			log.Fatalf("❌ Failed to complete message: %v", err)
		}
	}

	fmt.Println("✅ Finished processing messages")
}
