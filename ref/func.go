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

	mcnt := typ.NumMethod()
	fm := make(map[string]*FuncInfo, mcnt)

	for i := 0; i < mcnt; i++ {
		method := typ.Method(i)
		mTyp := method.Type

		// 输入参数
		// 注意第一个参数一定是接收器本身
		inCnt := mTyp.NumIn()
		in := make([]reflect.Type, 0, inCnt)
		for j := 0; j < inCnt; j++ {
			in = append(in, mTyp.In(j))
		}

		// 输出参数
		outCnt := mTyp.NumOut()
		out := make([]reflect.Type, 0, outCnt)
		for j := 0; j < outCnt; j++ {
			out = append(out, mTyp.Out(j))
		}

		// 处理输出
		callRes := method.Func.Call([]reflect.Value{reflect.ValueOf(entity)})
		res := make([]any, 0, len(callRes))

		for _, cr := range callRes {
			res = append(res, cr.Interface())
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
