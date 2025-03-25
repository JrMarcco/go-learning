package unsafe

import (
	"fmt"
	"reflect"
)

func PrintFieldOffset(entity any) {
	typ := reflect.TypeOf(entity)

	numField := typ.NumField()

	for i := 0; i < numField; i++ {
		fd := typ.Field(i)

		fmt.Printf("Field: %s, Offset: %d\n", fd.Name, fd.Offset)
	}
}
