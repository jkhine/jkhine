package main

import (
	"context"
	pb "joek-go/grpc/simple_server/protos"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":50051"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "hello " + in.Name}, nil
}

func main() {
	listen, _ := net.Listen("tcp", grpcPort)
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
