package api_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/surajkumarsinha/go_grpc_demo/api"
	"github.com/surajkumarsinha/go_grpc_demo/pb/messages"
	"github.com/surajkumarsinha/go_grpc_demo/pb/services"
	"github.com/surajkumarsinha/go_grpc_demo/repository"
	"github.com/surajkumarsinha/go_grpc_demo/sample"
	"github.com/surajkumarsinha/go_grpc_demo/serializer"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &services.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// check that the laptop is saved to the store
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}


func startTestLaptopServer(t *testing.T) (*api.LaptopServer, string) {
	laptopServer := api.NewLaptopServer(repository.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	services.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) services.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return services.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *messages.Laptop, laptop2 *messages.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}