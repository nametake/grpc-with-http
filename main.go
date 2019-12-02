package main

import (
	"context"
	"log"
	"net"

	"github.com/nametake/grpc-with-http/pb"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	pb.RegisterPingAPIServer(s, &PingAPIServer{})

	lis, err := net.Listen("tcp", ":9998")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}

type PingAPIServer struct{}

func (p *PingAPIServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Msg: "pong",
	}, nil
}
