package main

import (
	"context"
	"fmt"
	"golang-mono/godoc-test/pb"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDR = "127.0.0.1:9000"
)

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Could not connect to the server", "Error", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Todd"})
	if err != nil {
		slog.Error("Could not connect", "Error", err)
	}
	fmt.Printf("Message recieved", response.Message)
}
