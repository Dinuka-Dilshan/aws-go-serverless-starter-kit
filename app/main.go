package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event json.RawMessage) (events.LambdaFunctionURLResponse, error) {
	res, err := json.Marshal(map[string]string{
		"message": "hello there!",
	})

	if err != nil {
		slog.Error(err.Error())
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-type": "application/json",
		},
		Body: string(res),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
