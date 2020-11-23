package main

import (
	"context"
	"fmt"
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"pancake/maker/gen/api"
	"pancake/maker/handler"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	zaplogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	grpc_zap.ReplaceGrpcLogger(zaplogger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(zaplogger),
		),
	)
	api.RegisterPancakeBakerServiceServer(server, handler.NewBakerHandler())
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}

func myCustomInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		//req = modifyRequest(req)
		return handler(ctx, req)
	}
}
