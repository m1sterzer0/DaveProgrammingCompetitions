package stack

type DATATYPE int

// START HERE
type STACK struct {
	buf []DATATYPE
	l   int
}

func NewSTACK() *STACK           { buf := make([]DATATYPE, 0); return &STACK{buf, 0} }
func (q *STACK) IsEmpty() bool   { return q.l == 0 }
func (q *STACK) Clear()          { q.buf = q.buf[:0]; q.l = 0 }
func (q *STACK) Len() int        { return q.l }
func (q *STACK) Push(x DATATYPE) { q.buf = append(q.buf, x); q.l++ }
func (q *STACK) Pop() DATATYPE {
	if q.l == 0 {
		panic("Empty STACK Pop()")
	}
	v := q.buf[q.l-1]
	q.l--
	q.buf = q.buf[:q.l]
	return v
}
func (q *STACK) Head() DATATYPE {
	if q.l == 0 {
		panic("Empty STACK Head()")
	}
	return q.buf[q.l-1]
}
func (q *STACK) Top() DATATYPE { return q.Head() }
