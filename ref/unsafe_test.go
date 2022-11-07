package ref

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type unsafeArg struct {
	entity *UnsafeObj
	field  string
}

func TestUnsafeAccessor_Field(t *testing.T) {
	intTcs := []struct {
		name    string
		arg     unsafeArg
		wantErr error
		wantRes int
	}{
		{
			name: "nil",
			arg: unsafeArg{
				entity: nil,
			},
			wantErr: NilErr,
		},
		{
			name: "invalid field",
			arg: unsafeArg{
				entity: &UnsafeObj{},
				field:  "Invalid",
			},
			wantErr: InvalidFieldErr,
		},
		{
			name: "base int case",
			arg: unsafeArg{
				entity: &UnsafeObj{
					Age: 18,
				},
				field: "Age",
			},
			wantErr: nil,
			wantRes: 18,
		},
	}

	for _, tc := range intTcs {
		t.Run(tc.name, func(t *testing.T) {
			ua, err := NewUnsafeAccessor[int](tc.arg.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			res, err := ua.Field(tc.arg.field)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, res)
		})
	}

	strTcs := []struct {
		name    string
		arg     unsafeArg
		wantErr error
		wantRes string
	}{
		{
			name: "base str case",
			arg: unsafeArg{
				entity: &UnsafeObj{
					Name: "name",
				},
				field: "Name",
			},
			wantErr: nil,
			wantRes: "name",
		},
	}

	for _, tc := range strTcs {
		t.Run(tc.name, func(t *testing.T) {
			ua, err := NewUnsafeAccessor[string](tc.arg.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			res, err := ua.Field(tc.arg.field)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestUnsafeAccessor_SetField(t *testing.T) {
	intTcs := []struct {
		name    string
		arg     unsafeArg
		val     int
		wantErr error
		wantRes int
	}{
		{
			name: "base int case",
			arg: unsafeArg{
				entity: &UnsafeObj{},
				field:  "Age",
			},
			val:     20,
			wantErr: nil,
			wantRes: 20,
		},
	}

	for _, tc := range intTcs {
		t.Run(tc.name, func(t *testing.T) {
			ua, err := NewUnsafeAccessor[int](tc.arg.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			err = ua.SetField(tc.arg.field, tc.val)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.val, tc.arg.entity.Age)
		})
	}

	strTcs := []struct {
		name    string
		arg     unsafeArg
		val     string
		wantErr error
		wantRes string
	}{
		{
			name: "base str case",
			arg: unsafeArg{
				entity: &UnsafeObj{},
				field:  "Name",
			},
			val:     "name",
			wantErr: nil,
			wantRes: "name",
		},
	}

	for _, tc := range strTcs {
		t.Run(tc.name, func(t *testing.T) {
			ua, err := NewUnsafeAccessor[string](tc.arg.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			err = ua.SetField(tc.arg.field, tc.val)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.val, tc.arg.entity.Name)
		})
	}
}

type UnsafeObj struct {
	Name string
	Age  int
}
