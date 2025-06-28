package gpc

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/JrMarcco/go-learning/gpc/message"
)

func setProxyFunc(service Service, proxy Proxy) {
	val := reflect.ValueOf(service)
	elem := val.Elem()
	typ := elem.Type()

	numField := typ.NumField()
	for i := 0; i < numField; i++ {
		fieldVal := elem.Field(i)
		if fieldVal.CanSet() {
			field := typ.Field(i)

			fn := func(args []reflect.Value) []reflect.Value {
				in := args[1].Interface()
				out := reflect.New(field.Type.Out(0).Elem()).Interface()

				inData, err := json.Marshal(in)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				req := &message.Req{
					ServiceName: service.Name(),
					MethodName:  field.Name,
					Arg:         inData,
				}

				resp, err := proxy.Call(args[0].Interface().(context.Context), req)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				err = json.Unmarshal(resp.Data, out)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				return []reflect.Value{reflect.ValueOf(out), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fieldVal.Set(reflect.MakeFunc(field.Type, fn))
		}
	}
}
