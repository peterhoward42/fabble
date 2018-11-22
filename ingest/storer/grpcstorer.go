package storer

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "github.com/peterhoward42/fabble/store/protocol"
	"github.com/peterhoward42/fabble/store/usr"
)

// GRPCStorer exposes a *Store* method - which sends a gRPC message to
// a storage service.
type GRPCStorer struct {
	Stub pb.StoreUserClient
}

// NewGRPCStorer instantiates and initialises a GRPCStorer.
func NewGRPCStorer() (*GRPCStorer, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	const host string = ":9876"

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial: %v", err)
	}
	stub := pb.NewStoreUserClient(conn)
	return &GRPCStorer{Stub: stub}, nil
}

// Store saves the given *User* by sending a message to the storage service
// using the gRPC stub.
func (storer *GRPCStorer) Store(user *usr.User) error {

	timeout := time.Duration(500 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	storeRequest := &pb.StoreRequest{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Telephone,
	}
	_, err := storer.Stub.Store(ctx, storeRequest)
	if err != nil {
		return fmt.Errorf("stub.Store: %v", err)
	}
	return nil
}
