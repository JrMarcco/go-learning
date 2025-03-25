package ref

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPrintFieldOffset(t *testing.T) {
	fmt.Println(unsafe.Sizeof(User{}))
	PrintFieldOffset(User{})

	fmt.Println(unsafe.Sizeof(UserV1{}))
	PrintFieldOffset(UserV1{})

	fmt.Println(unsafe.Sizeof(UserV2{}))
	PrintFieldOffset(UserV2{})
}

type User struct {
	Name    string
	age     int32
	Alias   []byte
	Address string
}

type UserV1 struct {
	Name    string
	age     int32
	ageV1   int32
	Alias   []byte
	Address string
}

type UserV2 struct {
	Name    string
	Alias   []byte
	Address string
	age     int32
}
