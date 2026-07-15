package main

import (
	"Blogify/internal/notification/consumer"
	"log"
	"os"
)

func main() {
	consumer, err := consumer.NewConsumer(os.Getenv("KAFKA"))
	if err != nil {
		log.Fatalf("consumer init error: %v", err)
	}
	defer consumer.Close()
	log.Println("Listening for messages...")

	err = consumer.ConsumePartition("user.registered")
	if err != nil {
		log.Fatalf("failed to start partition consumer: %v", err)
	}
}
