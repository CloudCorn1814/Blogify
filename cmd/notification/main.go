package main

import (
	"Blogify/internal/notification/consumer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	consumer, err := consumer.NewConsumer(os.Getenv("KAFKA"))
	if err != nil {
		log.Fatalf("consumer init error: %v", err)
	}
	log.Println("Listening for messages...")

	go func() {
		err = consumer.ConsumePartition("user.registered")
		if err != nil {
			log.Fatalf("failed to start partition consumer: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down...")

	consumer.Close()
}
