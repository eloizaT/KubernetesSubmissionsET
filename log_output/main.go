package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

func generateRandomMessage() string {
	message := os.Getenv("LOG_MESSAGE")

	if message == "" {
		message = uuid.NewString()
	}

	return message
}

func main() {
	interval := 5 * time.Second

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Println("Application started")

	log.Println(generateRandomMessage())

	for {
		select {
		case <-ticker.C:
			log.Println(generateRandomMessage())

		case <-ctx.Done():
			log.Println("Application shutting down")
			return
		}
	}
}