package main

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

	"github.com/notify-vital/backend/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.Firebase.CredentialsFile)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	// Initialize Firestore client
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}
	defer firestoreClient.Close()

	// Try a simple Firestore operation
	_, err = firestoreClient.Collection("test").Doc("connection-test").Set(ctx, map[string]interface{}{
		"message":   "Firebase connection successful",
		"timestamp": time.Now(),
	})
	if err != nil {
		log.Fatalf("Error writing to Firestore: %v", err)
	}

	// Read the document to verify
	docRef, err := firestoreClient.Collection("test").Doc("connection-test").Get(ctx)
	if err != nil {
		log.Fatalf("Error reading from Firestore: %v", err)
	}

	data := docRef.Data()
	fmt.Println("✅ Firebase connection successful")
	fmt.Println("✅ Firestore write operation successful")
	fmt.Printf("📄 Document data: %v\n", data)
}
