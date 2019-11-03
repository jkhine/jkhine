package main

import (
	"context"
	"fmt"
	pb "joek-go/grpc/simple_client_server/protos"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("10.200.0.103:50051", grpc.WithInsecure())
	fmt.Printf("%+v\n\n", conn)

	time.Sleep(5 * time.Second)

	c := pb.NewHelloWorldClient(conn)

	fmt.Println("New ProtoBuf, before hello")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("i said hello")

	r, _ := c.SayHello(ctx, &pb.HelloRequest{Name: "Georgi"})
	log.Printf("Greeting .. " + r.Message)

}
