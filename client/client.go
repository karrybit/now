package main

import (
	"context"
	"io"
	"log"
	"now/pb/ticker"

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

	client := ticker.NewTickClient(conn)
	stream, err := client.Now(context.Background(), &emptypb.Empty{})
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
