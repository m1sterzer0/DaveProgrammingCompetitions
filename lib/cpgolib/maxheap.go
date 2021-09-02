package cpgolib

////////////////////////////////////////////////////////////////////////////////
// Maxheap 1 -- Works with an integer key for priority value
////////////////////////////////////////////////////////////////////////////////
type Maxheapelement1 struct {
	pri int
	val interface{}
}

type Maxheap1 struct {
	buf []Maxheapelement1
}

func NewMaxheap1() *Maxheap1 {
	buf := make([]Maxheapelement1, 0)
	return &Maxheap1{buf}
}

func (q *Maxheap1) Empty() bool {
	return len(q.buf) == 0
}

func (q *Maxheap1) Clear() {
	q.buf = q.buf[:0]
}

func (q *Maxheap1) Len() int {
	return len(q.buf)
}

func (q *Maxheap1) Push(pri int, val interface{}) {
	q.buf = append(q.buf, Maxheapelement1{pri, val})
	q.siftdown(0, len(q.buf)-1)
}

func (q *Maxheap1) Head() (int, interface{}) {
	return q.buf[0].pri, q.buf[0].val
}

func (q *Maxheap1) Pop() (int, interface{}) {
	v1 := q.buf[0].pri
	v2 := q.buf[0].val
	if len(q.buf) == 1 {
		q.buf = q.buf[:0]
	} else {
		l := len(q.buf)
		q.buf[0] = q.buf[l-1]
		q.buf = q.buf[:l-1]
		q.siftup(0)
	}
	return v1, v2
}

func (q *Maxheap1) Heapify(pri []int, val []interface{}) {
	if len(pri) != len(val) {
		panic("In Heapify, need the length of pri and val arguments to match")
	}
	for i, p := range pri {
		v := val[i]
		q.buf = append(q.buf, Maxheapelement1{p, v})
	}
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}

func (q *Maxheap1) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if newitem.pri <= parent.pri {
			break
		}
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}

func (q *Maxheap1) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !(q.buf[childpos].pri > q.buf[rightpos].pri) {
			childpos = rightpos
		}
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

////////////////////////////////////////////////////////////////////////////////
// Maxheap 2 -- Works with all ints
////////////////////////////////////////////////////////////////////////////////

type Maxheap2 struct {
	buf []int
}

func NewMaxheap2() *Maxheap2 {
	buf := make([]int, 0)
	return &Maxheap2{buf}
}

func (q *Maxheap2) Empty() bool {
	return len(q.buf) == 0
}

func (q *Maxheap2) Clear() {
	q.buf = q.buf[:0]
}

func (q *Maxheap2) Len() int {
	return len(q.buf)
}

func (q *Maxheap2) Push(v int) {
	q.buf = append(q.buf, v)
	q.siftdown(0, len(q.buf)-1)
}

func (q *Maxheap2) Head() int {
	return q.buf[0]
}

func (q *Maxheap2) Pop() int {
	v1 := q.buf[0]

	if len(q.buf) == 1 {
		q.buf = q.buf[:0]
	} else {
		l := len(q.buf)
		q.buf[0] = q.buf[l-1]
		q.buf = q.buf[:l-1]
		q.siftup(0)
	}
	return v1
}

func (q *Maxheap2) Heapify(pri []int) {
	q.buf = append(q.buf, pri...)
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}

func (q *Maxheap2) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if newitem <= parent {
			break
		}
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}

func (q *Maxheap2) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !(q.buf[childpos] > q.buf[rightpos]) {
			childpos = rightpos
		}
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

////////////////////////////////////////////////////////////////////////////////
// Maxheap 3 -- Create a custom class with a Less function
////////////////////////////////////////////////////////////////////////////////

type Maxheap3Element struct {
	x int
	y int
}

func Maxheap3Less(a, b Maxheap3Element) bool {
	return a.x < b.x || a.x == b.x && a.y < b.y
}

type Maxheap3 struct {
	buf []Maxheap3Element
}

func NewMaxheap3() *Maxheap3 {
	buf := make([]Maxheap3Element, 0)
	return &Maxheap3{buf}
}

func (q *Maxheap3) Empty() bool {
	return len(q.buf) == 0
}

func (q *Maxheap3) Clear() {
	q.buf = q.buf[:0]
}

func (q *Maxheap3) Len() int {
	return len(q.buf)
}

func (q *Maxheap3) Push(v Maxheap3Element) {
	q.buf = append(q.buf, v)
	q.siftdown(0, len(q.buf)-1)
}

func (q *Maxheap3) Head() Maxheap3Element {
	return q.buf[0]
}

func (q *Maxheap3) Pop() Maxheap3Element {
	v1 := q.buf[0]
	if len(q.buf) == 1 {
		q.buf = q.buf[:0]
	} else {
		l := len(q.buf)
		q.buf[0] = q.buf[l-1]
		q.buf = q.buf[:l-1]
		q.siftup(0)
	}
	return v1
}

func (q *Maxheap3) Heapify(varr []Maxheap3Element) {
	q.buf = append(q.buf, varr...)
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}

func (q *Maxheap3) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if !Maxheap3Less(parent, newitem) {
			break
		}
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}

func (q *Maxheap3) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !Maxheap3Less(q.buf[rightpos], q.buf[childpos]) {
			childpos = rightpos
		}
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}
