package ref

import (
	"reflect"
	"unsafe"
)

/*
读：*(*T)(ptr):
	|-- T 是目标类型，如果类型不知道，那么可以用 reflect.NewAt(typ, ptr).Elem()。
写：*(*T)(ptr) = T
	|-- T 是目标类型
*/

type FieldAccessor[T any] interface {
	Filed(field string) (T, error)
	SetField(field string, val T) error
}

type UnsafeAccessor[T any] struct {
	fm   map[string]fieldMeta
	addr unsafe.Pointer
}

type fieldMeta struct {
	offset uintptr
}

func NewUnsafeAccessor[T any](entity any) (*UnsafeAccessor[T], error) {
	if entity == nil {
		return nil, NilErr
	}

	// 只能是指针
	typ := reflect.TypeOf(entity)
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return nil, InvalidTypErr
	}

	val := reflect.ValueOf(entity)
	if val.IsNil() {
		// 处理指向 nil 的指针
		return nil, NilErr
	}

	typ = typ.Elem()
	fdCnt := typ.NumField()
	fm := make(map[string]fieldMeta, fdCnt)
	for i := 0; i < fdCnt; i++ {
		fd := typ.Field(i)
		fm[fd.Name] = fieldMeta{offset: fd.Offset}
	}

	return &UnsafeAccessor[T]{
		fm:   fm,
		addr: val.UnsafePointer(),
	}, nil
}

func (u *UnsafeAccessor[T]) Field(field string) (any, error) {
	meta, ok := u.fm[field]
	if !ok {
		return nil, InvalidFieldErr
	}

	ptr := unsafe.Pointer(uintptr(u.addr) + meta.offset)
	if ptr == nil {
		return nil, InvalidFieldAddrErr
	}

	return *(*T)(ptr), nil
}

func (u *UnsafeAccessor[T]) SetField(field string, val T) error {
	meta, ok := u.fm[field]
	if !ok {
		return InvalidFieldErr
	}

	ptr := unsafe.Pointer(uintptr(u.addr) + meta.offset)
	if ptr == nil {
		return InvalidFieldAddrErr
	}

	*(*T)(ptr) = val
	return nil
}
