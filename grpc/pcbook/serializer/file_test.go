package serializer

import (
	"github.com/stretchr/testify/require"
	"go-learning/grpc/pcbook/pb"
	"go-learning/grpc/pcbook/sample"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binFile := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop1, binFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	jsonFile := "../tmp/laptop.json"
	err = WriteProtobufToJsonFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop3 := &pb.Laptop{}
	err = ReadProtobufFromJsonFile(jsonFile, laptop3)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop3))

}
