package sql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type stmtRes struct {
	stmt string
	args []any
}

func TestInsertStmt(t *testing.T) {
	tcs := []struct {
		name    string
		arg     any
		wantErr error
		wantRes stmtRes
	}{
		{
			name:    "nil",
			arg:     nil,
			wantErr: NilEntityErr,
		},
		{
			name:    "multi ptr",
			arg:     toPtr(&simple{}),
			wantErr: InvalidEntityErr,
		},
		{
			name: "simple struct",
			arg: simple{
				RealName: "jcHong",
				NickName: "jrmarcco",
				Age:      30,
				Num:      nil, // 通过反射获取到的值为指向 nil 的指针：(*int64)(nil)
			},
			wantErr: nil,
			wantRes: stmtRes{
				stmt: "INSERT INTO `simple`(`real_name`,`nick_name`,`age`,`num`) VALUES (?,?,?,?);",
				args: []any{"jcHong", "jrmarcco", 30, (*int64)(nil)},
			},
		},
		{
			name: "simple ptr",
			arg: &simple{
				RealName: "jcHong",
				NickName: "jrmarcco",
				Age:      30,
				Num:      toPtr(int64(1001)),
			},
			wantErr: nil,
			wantRes: stmtRes{
				stmt: "INSERT INTO `simple`(`real_name`,`nick_name`,`age`,`num`) VALUES (?,?,?,?);",
				args: []any{"jcHong", "jrmarcco", 30, toPtr(int64(1001))},
			},
		},
		{
			name: "composition struct",
			arg: compositionStruct{
				RealName: "jcHong",
				NickName: "jrmarcco",
				sub: sub{
					NickName:  "foobar",
					FirstStr:  "first",
					SecondInt: 18,
				},
			},
			wantErr: nil,
			wantRes: stmtRes{
				stmt: "INSERT INTO `composition_struct`(`real_name`,`nick_name`,`first_str`,`second_int`) VALUES (?,?,?,?);",
				args: []any{"jcHong", "jrmarcco", "first", 18},
			},
		},
		{
			name: "composition ptr",
			arg: compositionPtr{
				PtrSub:   &PtrSub{},
				RealName: "jcHong",
				NickName: "jrmarcco",
			},
			wantErr: nil,
			wantRes: stmtRes{
				stmt: "INSERT INTO `composition_ptr`(`ptr_sub`,`real_name`,`nick_name`) VALUES (?,?,?);",
				args: []any{&PtrSub{}, "jcHong", "jrmarcco"},
			},
		},
		{
			name: "not embed field",
			arg: notEmbed{
				Sub: sub{},
			},
			wantErr: nil,
			wantRes: stmtRes{
				stmt: "INSERT INTO `not_embed`(`sub`) VALUES (?);",
				args: []any{sub{}},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			stmt, args, err := InsertStmt(tc.arg)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Equal(t, tc.wantRes, stmtRes{stmt: stmt, args: args})
		})
	}
}

func TestCamelToUnderLine(t *testing.T) {
	tcs := []struct {
		name    string
		arg     string
		wantRes string
	}{
		{
			name:    "empty str",
			arg:     "",
			wantRes: "",
		},
		{
			name:    "single",
			arg:     "single",
			wantRes: "single",
		},
		{
			name:    "normal case",
			arg:     "normalCase",
			wantRes: "normal_case",
		},
		{
			name:    "UpperStart",
			arg:     "UpperStart",
			wantRes: "upper_start",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := camelToUnderLine(tc.arg)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

type simple struct {
	RealName string
	NickName string
	Age      int
	Num      *int64
}

type compositionStruct struct {
	RealName string
	NickName string
	sub
}

type compositionPtr struct {
	*PtrSub
	RealName string
	NickName string
}

type sub struct {
	NickName  string
	FirstStr  string
	SecondInt int
}

type PtrSub struct {
	NickName  string
	FirstStr  string
	SecondInt int
}

type notEmbed struct {
	Sub sub
}

func toPtr[T any](t T) *T {
	return &t
}
