package main

import (
	"log"
	"net"
	"now/pb/now"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NowServerImpl struct {
	now.UnimplementedNowServer
}

func (t *NowServerImpl) Tick(req *emptypb.Empty, stream now.Now_TickServer) error {
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
	var nowServerImple NowServerImpl
	now.RegisterNowServer(server, &nowServerImple)

	if err := server.Serve(lis); err != nil {
		log.Fatalln(err)
	}

	log.Println("stop server")
}
