package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	// pb "github.com/matopenKW/grpc-server/pb"
	"golang.org/x/net/context"
	// "google.golang.org/grpc"
)

// var (
// 	userServiceEndpoint = flag.String("user_service_endpoint", "localhost:8080", "Users sample")
// )

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
