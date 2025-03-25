package unsafe

import "testing"

func TestPrintFieldOffset(t *testing.T) {

	tcs := []struct {
		name string
		arg  any
	}{
		{
			name: "v1",
			arg:  v1{},
		}, {
			name: "v2",
			arg:  v2{},
		}, {
			name: "v3",
			arg:  v3{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			PrintFieldOffset(tc.arg)
		})
	}
}

type v1 struct {
	Name    string
	Age     int32
	Alias   []string
	Address string
}

type v2 struct {
	Name    string
	Age     int32
	Height  int16
	Alias   []string
	Address string
}

type v3 struct {
	Name    string
	Age     int32
	Height  int16
	Weight  int16
	Alias   []string
	Address string
}
