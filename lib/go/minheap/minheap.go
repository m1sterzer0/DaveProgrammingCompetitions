package minheap

type DATATYPE int

// START HERE
type MINHEAP struct {
	buf  []DATATYPE
	less func(DATATYPE, DATATYPE) bool
}

func NewMINHEAP(f func(DATATYPE, DATATYPE) bool) *MINHEAP {
	buf := make([]DATATYPE, 0)
	return &MINHEAP{buf, f}
}
func (q *MINHEAP) IsEmpty() bool   { return len(q.buf) == 0 }
func (q *MINHEAP) Clear()          { q.buf = q.buf[:0] }
func (q *MINHEAP) Len() int        { return len(q.buf) }
func (q *MINHEAP) Push(v DATATYPE) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *MINHEAP) Head() DATATYPE  { return q.buf[0] }
func (q *MINHEAP) Pop() DATATYPE {
	v1 := q.buf[0]
	l := len(q.buf)
	if l == 1 {
		q.buf = q.buf[:0]
	} else {
		l--
		q.buf[0] = q.buf[l]
		q.buf = q.buf[:l]
		q.siftup(0)
	}
	return v1
}
func (q *MINHEAP) Heapify(pri []DATATYPE) {
	q.buf = append(q.buf, pri...)
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}
func (q *MINHEAP) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1
		p := q.buf[ppos]
		if !q.less(newitem, p) {
			break
		}
		q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *MINHEAP) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1
		if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) {
			chpos = rtpos
		}
		q.buf[pos], pos = q.buf[chpos], chpos
		chpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}
