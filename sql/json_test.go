package sql

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	Name string
}

func TestJsonColumn_Value(t *testing.T) {
	js := JsonColumn[user]{
		Val:   user{Name: "jrmarcco"},
		Valid: true,
	}
	value, err := js.Value()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`{"Name":"jrmarcco"}`), value)

	js = JsonColumn[user]{}
	value, err = js.Value()
	assert.Nil(t, err)
	assert.Nil(t, value)
}

func TestJsonColumn_Scan(t *testing.T) {
	tcs := []struct {
		name    string
		arg     any
		wantErr error
		wantRes user
	}{
		{
			name:    "nil",
			arg:     nil,
			wantErr: errors.New("UnSupport type: <nil>"),
		},
		{
			name:    "string",
			arg:     `{"Name": "jrmarcco"}`,
			wantErr: nil,
			wantRes: user{
				Name: "jrmarcco",
			},
		},
		{
			name:    "bytes",
			arg:     []byte(`{"Name": "jrmarcco"}`),
			wantErr: nil,
			wantRes: user{
				Name: "jrmarcco",
			},
		},
		{
			name: "bytes pointer",
			arg: func() *[]byte {
				bs := []byte(`{"Name": "jrmarcco"}`)
				return &bs
			}(),
			wantErr: nil,
			wantRes: user{
				Name: "jrmarcco",
			},
		},
		{
			name:    "sql.RawBytes",
			arg:     sql.RawBytes(`{"Name": "jrmarcco"}`),
			wantErr: nil,
			wantRes: user{
				Name: "jrmarcco",
			},
		},
		{
			name: "sql.RawBytes pointer",
			arg: func() *sql.RawBytes {
				srb := sql.RawBytes(`{"Name": "jrmarcco"}`)
				return &srb
			}(),
			wantErr: nil,
			wantRes: user{
				Name: "jrmarcco",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			js := &JsonColumn[user]{}
			err := js.Scan(tc.arg)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, js.Val)
			assert.True(t, js.Valid)
		})
	}
}
