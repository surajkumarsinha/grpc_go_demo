package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/surajkumarsinha/go_grpc_demo/api"
	"github.com/surajkumarsinha/go_grpc_demo/pb/messages"
	"github.com/surajkumarsinha/go_grpc_demo/pb/services"
	"github.com/surajkumarsinha/go_grpc_demo/repository"
	"github.com/surajkumarsinha/go_grpc_demo/sample"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := repository.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)

	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *messages.Laptop
		store  repository.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  repository.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  repository.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidID,
			store:  repository.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &services.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := api.NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}