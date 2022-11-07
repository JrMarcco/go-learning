package ref

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterate(t *testing.T) {
	tcs := []struct {
		name    string
		arg     any
		wantErr error
		wantRes []any
	}{
		{
			name:    "nil",
			arg:     nil,
			wantErr: nil,
			wantRes: []any{},
		},
		{
			name: "invalid type",
			arg: struct {
			}{},
			wantErr: InvalidTypErr,
		},
		{
			name:    "string",
			arg:     "hello",
			wantErr: nil,
			wantRes: []any{uint8('h'), uint8('e'), uint8('l'), uint8('l'), uint8('o')},
		},
		{
			name:    "int array",
			arg:     [...]int{1, 5, 6, 9},
			wantErr: nil,
			wantRes: []any{1, 5, 6, 9},
		},
		{
			name:    "string array",
			arg:     [...]string{"123", "abc", "zxc"},
			wantErr: nil,
			wantRes: []any{"123", "abc", "zxc"},
		},
		{
			name:    "int slice",
			arg:     []int{1, 5, 6, 9},
			wantErr: nil,
			wantRes: []any{1, 5, 6, 9},
		},
		{
			name:    "string slice",
			arg:     []string{"123", "abc", "zxc"},
			wantErr: nil,
			wantRes: []any{"123", "abc", "zxc"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Iterate(tc.arg)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, res)
		})
	}
}

type iterateRes struct {
	keys []any
	vals []any
}

func TestIterateMap(t *testing.T) {
	tcs := []struct {
		name    string
		arg     any
		wantErr error
		wantRes iterateRes
	}{
		{
			name:    "nil",
			arg:     nil,
			wantErr: nil,
			wantRes: iterateRes{
				keys: []any{},
				vals: []any{},
			},
		},
		{
			name: "invalid type",
			arg: struct {
			}{},
			wantErr: InvalidTypErr,
		},
		{
			name: "base map",
			arg: map[string]int{
				"abc": 123,
				"xyz": 456,
			},
			wantErr: nil,
			wantRes: iterateRes{
				keys: []any{"abc", "xyz"},
				vals: []any{123, 456},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			keys, vals, err := IterateMap(tc.arg)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, iterateRes{keys: keys, vals: vals})
		})
	}
}
