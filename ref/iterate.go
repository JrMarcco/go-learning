package ref

import "reflect"

func Iterate(in any) ([]any, error) {
	if in == nil {
		return []any{}, nil
	}

	val := reflect.ValueOf(in)
	kind := val.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		return nil, InvalidTypErr
	}

	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		res = append(res, val.Index(i).Interface())
	}

	return res, nil
}

func IterateMap(in any) ([]any, []any, error) {

	if in == nil {
		return []any{}, []any{}, nil
	}

	val := reflect.ValueOf(in)
	kind := val.Type().Kind()
	if kind != reflect.Map {
		return nil, nil, InvalidTypErr
	}

	valLen := val.Len()

	keys := make([]any, 0, valLen)
	vals := make([]any, 0, valLen)

	it := val.MapRange()
	for it.Next() {
		keys = append(keys, it.Key().Interface())
		vals = append(vals, it.Value().Interface())
	}

	return keys, vals, nil
}
