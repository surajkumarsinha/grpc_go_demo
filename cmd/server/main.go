package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/surajkumarsinha/go_grpc_demo/api"
	"github.com/surajkumarsinha/go_grpc_demo/pb/services"
	"github.com/surajkumarsinha/go_grpc_demo/repository"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopServer := api.NewLaptopServer(repository.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	services.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
	err = grpcServer.Serve(listener)

	log.Printf("server started on port %d", *port)
	
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
   
}