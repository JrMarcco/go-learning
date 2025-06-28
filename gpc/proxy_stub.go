package gpc

import (
	"context"
	"encoding/json"
	"reflect"
)

var _ ProxyStub = (*DefaultProxyStub)(nil)

type DefaultProxyStub struct {
	service Service
	refVal  reflect.Value
}

func (r *DefaultProxyStub) Call(ctx context.Context, methodName string, arg []byte) ([]byte, error) {
	method := r.refVal.MethodByName(methodName)

	inTyp := method.Type().In(1)
	in := reflect.New(inTyp.Elem())

	err := json.Unmarshal(arg, in.Interface())
	if err != nil {
		return nil, err
	}

	out := method.Call([]reflect.Value{reflect.ValueOf(ctx), in})
	if len(out) > 1 && !out[1].IsZero() {
		return nil, out[1].Interface().(error)
	}

	return json.Marshal(out[0].Interface())
}
