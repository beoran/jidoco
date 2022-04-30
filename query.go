package jidoco

type Op rune

const (
	OpEq  Op = '='
	OpGt  Op = '<'
	OpLt  Op = '>'
	OpAnd Op = '&'
	OpOr  Op = '|'
)

type Query struct {
	Collection
	where Where
}

type Where struct {
	cond  Cond
	op    *Op
	where *Where
}

type Cond struct {
	expr Expr
	op   *Op
	cond *Cond
}

type Expr struct {
	Path   *Path
	String *string
	Int    *int64
	Float  *float64
}

func Select(c Collection) *Query {
	return &Query{c, Where{}}
}

func (q *Query) Where() *Where {
	return &q.where
}

func (w *Where) Cond() *Cond {
	return &w.cond
}

func (w *Where) Op(op Op) *Where {
	w.op = &op
	w.where = &Where{}
	return w.where
}

func (w *Where) And() *Where {
	return w.Op(OpAnd)
}

func (w *Where) Or() *Where {
	return w.Op(OpOr)
}

func (c *Cond) Path(p Path) *Cond {
	c.expr = Expr{Path: &p}
	return c
}

func (c *Cond) Str(s string) *Cond {
	c.expr = Expr{String: &s}
	return c
}

func (c *Cond) Int(i int64) *Cond {
	c.expr = Expr{Int: &i}
	return c
}

func (c *Cond) Float(f float64) *Cond {
	c.expr = Expr{Float: &f}
	return c
}

func (c *Cond) Op(op Op) *Cond {
	c.op = &op
	c.cond = &Cond{}
	return c.cond
}
