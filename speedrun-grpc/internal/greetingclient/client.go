package greetingclient

import (
	"context"
	"errors"
	"fmt"
	"io"

	greetingv1 "github.com/IAmRiteshKoushik/speedrun-grpc/gen/greeting/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	service greetingv1.GreeterServiceClient
}

func New(conn grpc.ClientConnInterface) *Client {
	return &Client{
		service: greetingv1.NewGreeterServiceClient(conn),
	}
}

func (c *Client) GreetMany(ctx context.Context, name string, requestID string) ([]string, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-request-id", requestID)

	stream, err := c.service.GreetMany(ctx, &greetingv1.GreetRequest{
		Name: name,
	})
	if err != nil {
		return nil, fmt.Errorf("start greeting stream: %w", err)
	}

	var messages []string

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return messages, nil
		}
		if err != nil {
			return nil, fmt.Errorf("receive greeting stream: %w", err)
		}
		messages = append(messages, response.GetMessage())
	}
}
