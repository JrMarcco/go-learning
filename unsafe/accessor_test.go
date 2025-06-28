package unsafe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAccessor(t *testing.T) {

	tcs := []struct {
		name    string
		arg     any
		wantErr error
	}{
		{
			name:    "nil",
			arg:     nil,
			wantErr: errors.New("can not be nil"),
		}, {
			name:    "basic type",
			arg:     "string type",
			wantErr: errors.New("invalid type"),
		}, {
			name: "multi pointer",
			arg: func() any {
				s := &struct{}{}
				return &s
			}(),
			wantErr: errors.New("invalid type"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fa, err := NewAccessor(tc.arg)
			assert.Equal(t, tc.wantErr, err)

			if err == nil {
				require.NotNil(t, fa)
			}
		})
	}
}

type fmArg struct {
	Name string
	Age  int32
}

func TestFieldAccessor_Get(t *testing.T) {
	tcs := []struct {
		name    string
		fd      string
		wantRes any
		wantErr error
	}{
		{
			name:    "basic",
			fd:      "Name",
			wantRes: "jrmarcco",
		}, {
			name:    "invalid field",
			fd:      "invalid",
			wantErr: errors.New("invalid field"),
		},
	}

	arg := &fmArg{
		Name: "jrmarcco",
		Age:  18,
	}

	fa, err := NewAccessor(arg)
	require.NoError(t, err)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := fa.Get(tc.fd)
			assert.Equal(t, tc.wantErr, err)

			if err == nil {
				assert.Equal(t, tc.wantRes, res)
			}
		})
	}
}

func TestFieldAccessor_Set(t *testing.T) {

	tcs := []struct {
		name    string
		fd      string
		val     any
		wantRes any
		wantErr error
	}{
		{
			name:    "basic",
			fd:      "Age",
			val:     20,
			wantRes: 20,
		}, {
			name:    "invalid field",
			fd:      "invalid",
			val:     "",
			wantErr: errors.New("invalid field"),
		},
	}

	arg := &fmArg{
		Name: "jrmarcco",
		Age:  18,
	}

	fa, err := NewAccessor(arg)
	require.NoError(t, err)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err = fa.Set(tc.fd, tc.val)
			assert.Equal(t, tc.wantErr, err)

			if err == nil {
				res, err := fa.Get(tc.fd)
				require.NoError(t, err)

				assert.Equal(t, tc.wantRes, res)
			}
		})
	}
}
