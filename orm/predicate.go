package orm

type op string

const (
	opEq  = "="
	opGt  = ">"
	opLt  = "<"
	opAnd = "AND"
	opOr  = "OR"
	opNot = "NOT"
)

type Predicate struct {
	left  Expression
	op    op
	right Expression
}

func (p Predicate) expr() {}

func (p Predicate) And(right Predicate) Predicate {
	return Predicate{
		left:  p,
		op:    opAnd,
		right: right,
	}
}

func (p Predicate) Or(right Predicate) Predicate {
	return Predicate{
		left:  p,
		op:    opOr,
		right: right,
	}
}

func Not(right Predicate) Predicate {
	return Predicate{
		op:    opNot,
		right: right,
	}
}
