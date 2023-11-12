package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	pb "github.com/matopenKW/grpc-server/pkg/go/cat/v1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{TargetCat: "tama"}


	ctx := context.Background()
	res, err := client.GetMyCat(ctx, message)
	if err != nil {
		log.Fatal("client request error:", err)
	}
	fmt.Println(res)
}
