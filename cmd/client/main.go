package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/surajkumarsinha/go_grpc_demo/pb/services"
	"github.com/surajkumarsinha/go_grpc_demo/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal("cannot dial to the server")
	}

	laptopClient := services.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := services.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := laptopClient.CreateLaptop(ctx, &req)

	defer cancel()

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}

		return 
	}

	log.Printf("Laptop Created with id: %s", res.Id)
}