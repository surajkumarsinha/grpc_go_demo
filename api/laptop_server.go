package api

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/surajkumarsinha/go_grpc_demo/pb/services"
	"github.com/surajkumarsinha/go_grpc_demo/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct{
	services.UnimplementedLaptopServiceServer
	Store repository.LaptopStore
}

// Return a laptop server
func NewLaptopServer(laptopStore repository.LaptopStore) *LaptopServer {
	laptopServer := LaptopServer{}
	laptopServer.Store = laptopStore
	return &laptopServer
}

// Create Unimplemented rpc method
func (server *LaptopServer) CreateLaptop(
	ctx context.Context, 
	req *services.CreateLaptopRequest,
	) (*services.CreateLaptopResponse, error) {
		laptop := req.GetLaptop()
		log.Printf("recieve a create laptop request with id : %s", laptop.Id)

		if(len(laptop.Id) > 0) {
				// Check if it is a valid id or not
				_, err := uuid.Parse(laptop.Id)
				if err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not a valid UUID: %v", err)
				}
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Laptop Id could not be generated: %v", err)
			}
			laptop.Id = id.String()
		}

		// save the laptop in-memory
		err := server.Store.Save(laptop)
		if err != nil {
			code := codes.Internal
			if errors.Is(err, repository.ErrAlreadyExists) {
				code = codes.AlreadyExists
			}
			return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
		}

		log.Printf("saved laptop with id: %v", laptop.Id)

		res := &services.CreateLaptopResponse{
			Id: laptop.Id,
		}

		return res, nil
}