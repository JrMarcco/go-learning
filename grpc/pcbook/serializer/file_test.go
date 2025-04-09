package serializer

import (
	"testing"

	"github.com/JrMarcco/go-learning/grpc/pcbook/pb"
	"github.com/JrMarcco/go-learning/grpc/pcbook/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
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
