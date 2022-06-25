package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/surajkumarsinha/go_grpc_demo/pb/messages"
	"github.com/surajkumarsinha/go_grpc_demo/sample"
	"github.com/surajkumarsinha/go_grpc_demo/serializer"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop, binaryFile)
	require.NoError(t, err)

	laptop2 := &messages.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop, jsonFile)
	require.NoError(t, err)
}