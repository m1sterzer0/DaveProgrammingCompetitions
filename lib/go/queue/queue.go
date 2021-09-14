package queue

type DATATYPE int

// START HERE
type QUEUE struct {
	buf                   []DATATYPE
	head, tail, sz, bm, l int
}

func NewQUEUE() *QUEUE         { buf := make([]DATATYPE, 8); return &QUEUE{buf, 0, 0, 8, 7, 0} }
func (q *QUEUE) IsEmpty() bool { return q.l == 0 }
func (q *QUEUE) Clear()        { q.head = 0; q.tail = 0; q.l = 0 }
func (q *QUEUE) Len() int      { return q.l }
func (q *QUEUE) Push(x DATATYPE) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.head = (q.head - 1) & q.bm
	}
	q.l++
	q.buf[q.head] = x
}
func (q *QUEUE) Pop() DATATYPE {
	if q.l == 0 {
		panic("Empty QUEUE Pop()")
	}
	v := q.buf[q.tail]
	q.l--
	if q.l > 0 {
		q.tail = (q.tail - 1) & q.bm
	} else {
		q.Clear()
	}
	return v
}
func (q *QUEUE) Head() DATATYPE {
	if q.l == 0 {
		panic("Empty QUEUE Head()")
	}
	return q.buf[q.head]
}
func (q *QUEUE) Tail() DATATYPE {
	if q.l == 0 {
		panic("Empty QUEUE Tail()")
	}
	return q.buf[q.tail]
}
func (q *QUEUE) sizeup() {
	buf := make([]DATATYPE, 2*q.sz)
	for i := 0; i < q.l; i++ {
		buf[i] = q.buf[(q.head+i)&q.bm]
	}
	q.buf = buf
	q.head = 0
	q.tail = q.sz - 1
	q.sz = 2 * q.sz
	q.bm = q.sz - 1
}
