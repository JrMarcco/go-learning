package ref

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFields(t *testing.T) {
	tcs := []struct {
		name    string
		arg     any
		wantErr error
		wantRes map[string]any
	}{
		{
			name:    "nil",
			arg:     nil,
			wantErr: NilErr,
		},
		{
			name:    "invalid type",
			arg:     []string{},
			wantErr: InvalidTypErr,
		},
		{
			name: "base struct",
			arg: Base{
				Name: "name struct",
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "name struct",
			},
		},
		{
			name: "base pointer",
			arg: &Base{
				Name: "name pointer",
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "name pointer",
			},
		},
		{
			name: "multi pointer",
			arg: toPtr(&Base{
				Name: "name multi pointer",
			}),
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "name multi pointer",
			},
		},
		{
			name: "private field",
			arg: Base{
				Name: "name private field",
				age:  18,
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "name private field",
			},
		},
		{
			name: "nest struct field",
			arg: Base{
				Name: "name nest struct field",
				NestFd: Nest{
					Name: "Name Nest",
					Type: "Type Nest",
				},
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "name nest struct field",
				"NestFd": map[string]any{
					"Name": "Name Nest",
					"Type": "Type Nest",
				},
			},
		},
		{
			name: "nest pointer field",
			arg: Base{
				Name: "name nest pointer field",
				NestFdPtr: &Nest{
					Name: "Name Nest Pointer",
					Type: "Type Nest Pointer",
				},
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "name nest pointer field",
				"NestFdPtr": map[string]any{
					"Name": "Name Nest Pointer",
					"Type": "Type Nest Pointer",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fm, err := iterateFields(tc.arg)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, fm)
		})
	}
}

type setFieldArg struct {
	entity any
	filed  string
	val    any
}

func TestSetField(t *testing.T) {

	tcs := []struct {
		name    string
		arg     setFieldArg
		wantErr error
	}{
		{
			name: "nil entity",
			arg: setFieldArg{
				entity: nil,
			},
			wantErr: NilErr,
		},
		{
			name: "base struct",
			arg: setFieldArg{
				entity: Base{},
			},
			wantErr: InvalidTypErr,
		},
		{
			name: "multi ptr entity",
			arg: setFieldArg{
				entity: toPtr(&Base{}),
			},
			wantErr: InvalidTypErr,
		},
		{
			name: "base pointer",
			arg: setFieldArg{
				entity: &Base{
					Name: "base name",
				},
				filed: "Name",
				val:   "new name",
			},
		},
		{
			name: "not exist field",
			arg: setFieldArg{
				entity: &Base{
					Name: "base name",
				},
				filed: "NotExist",
			},
			wantErr: NotExistFieldErr,
		},
		{
			name: "can't set field",
			arg: setFieldArg{
				entity: &Base{
					age: 18,
				},
				filed: "age",
				val:   20,
			},
			wantErr: UnsetFiledErr,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.arg.entity, tc.arg.filed, tc.arg.val)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.arg.val, getFiledValByName(tc.arg.entity, tc.arg.filed))
		})
	}
}

func getFiledValByName(entity any, field string) any {
	val := reflect.ValueOf(entity).Elem()
	return val.FieldByName(field).Interface()
}

type Base struct {
	Name      string
	age       int
	NestFd    Nest
	NestFdPtr *Nest
}

type Nest struct {
	Name string
	Type string
}

func toPtr[T any](t T) *T {
	return &t
}
