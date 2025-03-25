package copy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRefCopier_Copy(t *testing.T) {
	tcs := []struct {
		name     string
		copyFunc func() (any, error)
		wantRes  any
		wantErr  error
	}{
		{
			name: "basic",
			copyFunc: func() (any, error) {
				copier, err := NewRefCopier[BasicSrc, BasicDst]()
				if err != nil {
					return nil, err
				}
				return copier.Copy(&BasicSrc{
					Name: "Foo",
					Age:  18,
				})
			},
			wantErr: nil,
			wantRes: &BasicDst{
				Name: "Foo",
				Age:  18,
			},
		},
		{
			name: "diff type",
			copyFunc: func() (any, error) {
				copier, err := NewRefCopier[DiffTypeSrc, DiffTypeDst]()
				if err != nil {
					return nil, err
				}
				return copier.Copy(&DiffTypeSrc{
					Name: "Foo",
					Sex:  "male",
				})
			},
			wantErr: nil,
			wantRes: &DiffTypeDst{
				Name: "Foo",
			},
		},
		{
			name: "diff ptr type",
			copyFunc: func() (any, error) {
				copier, err := NewRefCopier[DiffTypeSrc, DiffTypeDst]()
				if err != nil {
					return nil, err
				}
				return copier.Copy(&DiffTypeSrc{
					Name:   "Foo",
					PtrSex: toPtr("male"),
				})
			},
			wantErr: nil,
			wantRes: &DiffTypeDst{
				Name: "Foo",
			},
		},
		{
			name: "struct field",
			copyFunc: func() (any, error) {
				copier, err := NewRefCopier[BasicSrc, BasicDst]()
				if err != nil {
					return nil, err
				}
				return copier.Copy(&BasicSrc{
					Name: "Foo",
					Age:  18,
					StructAddr: Address{
						Province: "Fujian",
						City:     "Xiamen",
					},
				})
			},
			wantErr: nil,
			wantRes: &BasicDst{
				Name: "Foo",
				Age:  18,
				StructAddr: Address{
					Province: "Fujian",
					City:     "Xiamen",
				},
			},
		},
		{
			name: "ptr field",
			copyFunc: func() (any, error) {
				copier, err := NewRefCopier[BasicSrc, BasicDst]()
				if err != nil {
					return nil, err
				}
				return copier.Copy(&BasicSrc{
					Name: "Foo",
					Age:  18,
					PtrAddr: &Address{
						Province: "Fujian",
						City:     "Xiamen",
						Code:     toPtr(361000),
					},
				})
			},
			wantErr: nil,
			wantRes: &BasicDst{
				Name: "Foo",
				Age:  18,
				PtrAddr: &Address{
					Province: "Fujian",
					City:     "Xiamen",
					Code:     toPtr(361000),
				},
			},
		},
		{
			name: "str map field",
			copyFunc: func() (any, error) {
				copier, err := NewRefCopier[BasicSrc, BasicDst]()
				if err != nil {
					return nil, err
				}
				return copier.Copy(&BasicSrc{
					Name: "Foo",
					StrMap: map[string]string{
						"first":  "abc",
						"second": "efg",
					},
				})
			},
			wantErr: nil,
			wantRes: &BasicDst{
				Name: "Foo",
				StrMap: map[string]string{
					"first":  "abc",
					"second": "efg",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.copyFunc()
			assert.Equal(t, tc.wantErr, err)

			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

type BasicSrc struct {
	Name       string
	Age        int
	StructAddr Address
	PtrAddr    *Address

	StrMap map[string]string
}

type BasicDst struct {
	Name     string
	Nickname string
	Age      int

	StructAddr Address
	PtrAddr    *Address

	StrMap map[string]string
}

type DiffTypeSrc struct {
	Name   string
	Sex    string
	PtrSex *string
}

type DiffTypeDst struct {
	Name   string
	Sex    int
	PtrSex *int
}

type Address struct {
	Province string
	City     string

	Code *int
}

func toPtr[T any](t T) *T {
	return &t
}
