package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IAmRiteshKoushik/speedrun-grpc/internal/greetingclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const target = "localhost:50051"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("create gRPC client: %w", err)
	}
	defer conn.Close()

	client := greetingclient.New(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	messages, err := client.GreetMany(ctx, "Ritesh", "practice-001")
	if err != nil {
		return fmt.Errorf("greet many: %w", err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}

	return nil
}
