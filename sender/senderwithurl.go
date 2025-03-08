package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func sendwithurl() {
	// Provided Service Bus URL
	serviceBusURL := "servicebus+https://localhost:5672/queues/control-plane-notifications"

	// Parse the URL to extract components
	parsedURL, err := url.Parse(serviceBusURL)
	if err != nil {
		log.Fatalf("❌ Failed to parse Service Bus URL: %v", err)
	}

	// Extract namespace and queue name
	namespace := parsedURL.Hostname()
	queueName := parsedURL.Path[len("/queues/"):] // Extract queue name

	if namespace == "" || queueName == "" {
		log.Fatal("❌ Invalid Service Bus URL format")
	}

	// Construct connection string (if required)
	connectionString := fmt.Sprintf("Endpoint=sb://%s:5672/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;", namespace)

	// Create Service Bus client
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("❌ Failed to create Service Bus client: %v", err)
	}
	defer client.Close(context.Background())

	// Create sender
	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		log.Fatalf("❌ Failed to create sender: %v", err)
	}
	defer sender.Close(context.Background())

	// Define message
	message := &azservicebus.Message{
		Body: []byte("Hello, Azure Service Bus Emulator!"),
	}

	// Send message
	ctx := context.Background()
	err = sender.SendMessage(ctx, message, nil)
	if err != nil {
		log.Fatalf("❌ Failed to send message: %v", err)
	}

	fmt.Println("✅ Successfully sent message to Azure Service Bus Queue (Emulator)")
}
