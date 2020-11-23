package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	pb "pancake/maker/gen/api"
)

func main() {
	port := 50051
	conn, err := grpc.Dial(fmt.Sprint("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewPancakeBakerServiceClient(conn)
	message := &pb.BakeRequest{}
	res, err := client.Bake(context.TODO(), message)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
