package ref

import "reflect"

type FuncInfo struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type

	Res []any
}

func IterateFunc(entity any) (map[string]*FuncInfo, error) {

	if entity == nil {
		return nil, NilErr
	}

	typ := reflect.TypeOf(entity)
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Pointer {
		return nil, InvalidTypErr
	}

	numMethod := typ.NumMethod()
	fm := make(map[string]*FuncInfo, numMethod)

	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		methodTyp := method.Type

		// 输入参数
		// 注意第一个参数一定是接收器本身
		numIn := methodTyp.NumIn()
		in := make([]reflect.Type, 0, numIn)
		inVals := make([]reflect.Value, 0, numIn)

		in = append(in, reflect.TypeOf(entity))
		inVals = append(inVals, reflect.ValueOf(entity))

		for j := 1; j < numIn; j++ {
			inTyp := methodTyp.In(j)

			in = append(in, inTyp)
			inVals = append(inVals, reflect.Zero(inTyp))
		}

		// 输出参数
		numOut := methodTyp.NumOut()
		out := make([]reflect.Type, 0, numOut)

		for j := 0; j < numOut; j++ {
			out = append(out, methodTyp.Out(j))
		}

		// 处理输出
		resVals := method.Func.Call(inVals)
		res := make([]any, 0, len(resVals))

		for _, val := range resVals {
			res = append(res, val.Interface())
		}

		fm[method.Name] = &FuncInfo{
			Name: method.Name,
			In:   in,
			Out:  out,
			Res:  res,
		}

	}
	return fm, nil
}
