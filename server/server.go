package main

import (
	"log"
	"net"
	"now/pb/ticker"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TickServerImpl struct {
	ticker.UnimplementedTickServer
}

func (t *TickServerImpl) Now(req *emptypb.Empty, stream ticker.Tick_NowServer) error {
	log.Printf("requested %s\n", req.String())
	for _ = range time.NewTicker(1 * time.Second).C {
		if err := stream.Send(&timestamppb.Timestamp{}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.Println("run server")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	var tickServerImple TickServerImpl
	ticker.RegisterTickServer(server, &tickServerImple)

	if err := server.Serve(lis); err != nil {
		log.Fatalln(err)
	}

	log.Println("stop server")
}
