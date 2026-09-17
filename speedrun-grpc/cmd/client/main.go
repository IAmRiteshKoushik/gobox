package main

import (
	"context"
	"fmt"
	"log"
	"time"

	greetingv1 "github.com/IAmRiteshKoushik/speedrun-grpc/gen/greeting/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	client := greetingv1.NewGreeterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(
		ctx, metadata.Pairs("x-request-id", "practice-001"),
	)

	response, err := client.Greet(ctx, &greetingv1.GreetRequest{
		Name: "Ritesh",
	})
	if err != nil {
		switch status.Code(err) {
		case codes.Unavailable:
			return fmt.Errorf("Greet failed: code=%s, server may not be reachable: %w", status.Code(err), err)
		case codes.DeadlineExceeded:
			return fmt.Errorf("Greet failed: code=%s, call exceeded its two-second budget: %w", status.Code(err), err)
		default:
			return fmt.Errorf("Greet failed: code=%s: %w", status.Code(err), err)
		}
	}

	fmt.Println(response.GetMessage())
	return nil
}
