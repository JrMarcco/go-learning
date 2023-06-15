package orm

type Column struct {
	name string
}

func (c Column) expr() {}

func Col(name string) Column {
	return Column{name: name}
}

func (c Column) Eq(val any) Predicate {
	return Predicate{
		left:  c,
		op:    opEq,
		right: valOf(val),
	}
}

func (c Column) Gt(val any) Predicate {
	return Predicate{
		left:  c,
		op:    opGt,
		right: valOf(val),
	}
}

func (c Column) Lt(val any) Predicate {
	return Predicate{
		left:  c,
		op:    opLt,
		right: valOf(val),
	}
}

type Value struct {
	val any
}

func (v Value) expr() {}

func valOf(val any) Value {
	return Value{val: val}
}
