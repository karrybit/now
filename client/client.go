package main

import (
	"context"
	"io"
	"log"
	"now/pb/now"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := now.NewNowClient(conn)
	stream, err := client.Tick(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response)
	}
}
