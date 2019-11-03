package main

import (
	"context"
	"fmt"
	pb "joek-go/grpc/simple_client_server/protos"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":50051"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("request came with name -> " + in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, _ := net.Listen("tcp", grpcPort)
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
