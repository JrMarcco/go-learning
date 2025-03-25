package ref

import "errors"

var (
	NilErr              = errors.New("can not be nil")
	ZeroValErr          = errors.New("can not be zero")
	InvalidTypErr       = errors.New("invalid type")
	InvalidFieldErr     = errors.New("invalid field")
	InvalidFieldAddrErr = errors.New("invalid field address")
	NotExistFieldErr    = errors.New("field not exist")
	UnsetFiledErr       = errors.New("field can not be set")
)
