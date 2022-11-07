package ref

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFunc(t *testing.T) {

	tcs := []struct {
		name    string
		arg     any
		wantErr error
		wantRes map[string]*FuncInfo
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
			arg: BaseFunc{
				name: "base name",
			},
			wantErr: nil,
			wantRes: map[string]*FuncInfo{
				"GetName": {
					Name: "GetName",
					In: []reflect.Type{
						reflect.TypeOf(BaseFunc{}),
					},
					Out: []reflect.Type{reflect.TypeOf("base name")},
					Res: []any{"base name"},
				},
			},
		},
		{
			name: "base struct with ptr",
			arg: &BaseFunc{
				name: "base name",
			},
			wantErr: nil,
			wantRes: map[string]*FuncInfo{
				"GetName": {
					Name: "GetName",
					In: []reflect.Type{
						reflect.TypeOf(&BaseFunc{}),
					},
					Out: []reflect.Type{reflect.TypeOf("base name")},
					Res: []any{"base name"},
				},
			},
		},
		{
			name: "base ptr",
			arg: &PtrFunc{
				name: "ptr name",
			},
			wantErr: nil,
			wantRes: map[string]*FuncInfo{
				"GetName": {
					Name: "GetName",
					In: []reflect.Type{
						reflect.TypeOf(&PtrFunc{}),
					},
					Out: []reflect.Type{reflect.TypeOf("ptr name")},
					Res: []any{"ptr name"},
				},
			},
		},
		{
			name: "base ptr with struct",
			arg: PtrFunc{
				name: "ptr name",
			},
			wantErr: nil,
			wantRes: map[string]*FuncInfo{},
		},
		{
			name: "multi ptr",
			arg: toPtr(&PtrFunc{
				name: "ptr name",
			}),
			wantErr: nil,
			wantRes: map[string]*FuncInfo{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fm, err := IterateFunc(tc.arg)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, fm)
		})
	}
}

type BaseFunc struct {
	name string
}

func (b BaseFunc) GetName() string {
	return b.name
}

type PtrFunc struct {
	name string
}

func (p *PtrFunc) GetName() string {
	return p.name
}
