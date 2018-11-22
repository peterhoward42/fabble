package server

import (
	"context"
	"log"

	pb "github.com/peterhoward42/fabble/store/protocol"
	"github.com/peterhoward42/fabble/store/usr"
)

// Server an applications-specific wrapper around a gRPC server.
type Server struct {
	Database map[int]*usr.User
}

// Store is the gRPC handler for the gRPC *Store* RPC call.
func (s Server) Store(ctx context.Context, req *pb.StoreRequest) (
	*pb.StoreResponse, error) {

	user := usr.User{
		ID:        int(req.GetId()),
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Telephone: req.GetPhone(),
	}

	// Store or overwrite the user of this ID
	s.Database[user.ID] = &user

	log.Printf("Server Store() event")
	return &pb.StoreResponse{}, nil
}
