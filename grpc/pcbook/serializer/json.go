package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "    ",
		OrigName:     true,
	}

	return marshaler.MarshalToString(message)
}

func JsonToProtobuf(json string, message proto.Message) error {
	return jsonpb.UnmarshalString(json, message)
}
