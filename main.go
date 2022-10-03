package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string `json:"name"`
}

type MyResponse struct {
        Name string `json:"name"`
	Time int `json:"time"`
}

func handleRequest(ctx context.Context, event MyEvent) (MyResponse, error) {
	return MyResponse{Name: fmt.Sprintf("Hello World: %s!", event.Name ), Time: 1234}, nil
}

func main() {
        lambda.Start(handleRequest)
}
