package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/peterhoward42/fabble/store/protocol"
	"github.com/peterhoward42/fabble/store/server"
	"github.com/peterhoward42/fabble/store/usr"
)

// This program creates and launches a gRPC service that accepts requests
// to store *User*s. This first POC implementation uses a trivial (and of course
// volatile) Go map for storage. A proper solution would like use as Redis
// service.
func main() {

	database := map[int]*usr.User{} // Keyed on ID

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	const host string = "localhost:9876"
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}
	s := server.Server{Database: database}
	pb.RegisterStoreUserServer(grpcServer, s)
	// Listen forever.
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve: %v", err)
	}
}
