package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "sample/proto"
)

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// gRPC-Gateway を gRPC サーバーに接続
	err := pb.RegisterHelloServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to start gRPC-Gateway: %v", err)
	}

	log.Println("gRPC-Gateway listening on port 8080...")
	http.ListenAndServe(":8080", mux)
}
