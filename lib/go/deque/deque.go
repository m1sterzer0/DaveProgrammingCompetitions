package deque

type DATATYPE int

// START HERE
type DEQUE struct {
	buf                   []DATATYPE
	head, tail, sz, bm, l int
}

func NewDEQUE() *DEQUE         { buf := make([]DATATYPE, 8); return &DEQUE{buf, 0, 0, 8, 7, 0} }
func (q *DEQUE) IsEmpty() bool { return q.l == 0 }
func (q *DEQUE) Clear()        { q.head = 0; q.tail = 0; q.l = 0 }
func (q *DEQUE) PushFront(x DATATYPE) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.head = (q.head - 1) & q.bm
	}
	q.l++
	q.buf[q.head] = x
}
func (q *DEQUE) PushBack(x DATATYPE) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.tail = (q.tail + 1) & q.bm
	}
	q.l++
	q.buf[q.tail] = x
}
func (q *DEQUE) PopFront() DATATYPE {
	if q.l == 0 {
		panic("Empty DEQUE PopFront()")
	}
	v := q.buf[q.head]
	q.l--
	if q.l > 0 {
		q.head = (q.head + 1) & q.bm
	} else {
		q.Clear()
	}
	return v
}
func (q *DEQUE) PopBack() DATATYPE {
	if q.l == 0 {
		panic("Empty DEQUE PopBack()")
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
func (q *DEQUE) Len() int { return q.l }
func (q *DEQUE) Head() DATATYPE {
	if q.l == 0 {
		panic("Empty DEQUE Head()")
	}
	return q.buf[q.head]
}
func (q *DEQUE) Tail() DATATYPE {
	if q.l == 0 {
		panic("Empty DEQUE Tail()")
	}
	return q.buf[q.tail]
}
func (q *DEQUE) sizeup() {
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
