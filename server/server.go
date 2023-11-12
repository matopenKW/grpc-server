package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"

	pb "github.com/matopenKW/grpc-server/pb"
	"github.com/matopenKW/grpc-server/service"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &service.MyCatService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(server, catService)

	fmt.Println("start grpc server, Port: 19003")
	server.Serve(listenPort)
}
