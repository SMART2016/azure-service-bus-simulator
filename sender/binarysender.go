package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func main() {
	connectionString := os.Getenv("SERVICEBUS_CONNECTION_STRING")
	queueName := os.Getenv("SERVICEBUS_QUEUE_NAME")

	if connectionString == "" || queueName == "" {
		log.Fatal("Missing SERVICEBUS_CONNECTION_STRING or SERVICEBUS_QUEUE_NAME")
	}

	// Sample binary data (this should be replaced with actual Protobuf-encoded data)
	protoBinaryData := []byte{0x0A, 0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6E, 0x67} // Example Protobuf binary

	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("Failed to create Service Bus client: %v", err)
	}
	defer client.Close(context.Background())

	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		log.Fatalf("Failed to create sender: %v", err)
	}
	defer sender.Close(context.Background())

	// Create message with binary data and set the content type
	contentType := "application/x-protobuf"
	message := &azservicebus.Message{
		Body:        protoBinaryData, // Set the binary data
		ContentType: &contentType,    // Set Content-Type as Protobuf
	}

	err = sender.SendMessage(context.Background(), message, nil)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("✅ Successfully sent Protobuf message to queue: %s\n", queueName)
}
