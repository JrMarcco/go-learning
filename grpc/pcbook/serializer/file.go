package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can not marshal proto message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can not write binary to file: %w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can not read binary from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can not unmarshal proto message from binary: %w", err)
	}
	return nil
}

func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	json, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("can not marshal proto message to json: %w", err)
	}

	err = os.WriteFile(filename, []byte(json), 0644)
	if err != nil {
		return fmt.Errorf("can not write json to file: %w", err)
	}
	return nil
}

func ReadProtobufFromJsonFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can not read json from file: %w", err)
	}

	err = JsonToProtobuf(string(data), message)
	if err != nil {
		return fmt.Errorf("can not unmarshal proto message from json: %w", err)
	}
	return nil
}
