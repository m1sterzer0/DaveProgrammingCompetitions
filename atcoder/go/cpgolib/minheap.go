package cpgolib

////////////////////////////////////////////////////////////////////////////////
// Minheap 1 -- Works with an integer key for priority value
////////////////////////////////////////////////////////////////////////////////
type Minheapelement1 struct {
	pri int
	val interface{}
}

type Minheap1 struct {
	buf []Minheapelement1
}

func NewMinheap1() *Minheap1 {
	buf := make([]Minheapelement1, 0)
	return &Minheap1{buf}
}

func (q *Minheap1) Empty() bool {
	return len(q.buf) == 0
}

func (q *Minheap1) Clear() {
	q.buf = q.buf[:0]
}

func (q *Minheap1) Len() int {
	return len(q.buf)
}

func (q *Minheap1) Push(pri int, val interface{}) {
	q.buf = append(q.buf, Minheapelement1{pri, val})
	q.siftdown(0, len(q.buf)-1)
}

func (q *Minheap1) Head() (int, interface{}) {
	return q.buf[0].pri, q.buf[0].val
}

func (q *Minheap1) Pop() (int, interface{}) {
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

func (q *Minheap1) Heapify(pri []int, val []interface{}) {
	if len(pri) != len(val) {
		panic("In Heapify, need the length of pri and val arguments to match")
	}
	for i, p := range pri {
		v := val[i]
		q.buf = append(q.buf, Minheapelement1{p, v})
	}
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}

func (q *Minheap1) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if newitem.pri >= parent.pri {
			break
		}
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}

func (q *Minheap1) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !(q.buf[childpos].pri < q.buf[rightpos].pri) {
			childpos = rightpos
		}
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

////////////////////////////////////////////////////////////////////////////////
// Minheap 2 -- Works with all ints
////////////////////////////////////////////////////////////////////////////////

type Minheap2 struct {
	buf []int
}

func NewMinheap2() *Minheap2 {
	buf := make([]int, 0)
	return &Minheap2{buf}
}

func (q *Minheap2) Empty() bool {
	return len(q.buf) == 0
}

func (q *Minheap2) Clear() {
	q.buf = q.buf[:0]
}

func (q *Minheap2) Len() int {
	return len(q.buf)
}

func (q *Minheap2) Push(v int) {
	q.buf = append(q.buf, v)
	q.siftdown(0, len(q.buf)-1)
}

func (q *Minheap2) Head() int {
	return q.buf[0]
}

func (q *Minheap2) Pop() int {
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

func (q *Minheap2) Heapify(pri []int) {
	q.buf = append(q.buf, pri...)
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}

func (q *Minheap2) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if newitem >= parent {
			break
		}
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}

func (q *Minheap2) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !(q.buf[childpos] < q.buf[rightpos]) {
			childpos = rightpos
		}
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

////////////////////////////////////////////////////////////////////////////////
// Minheap 3 -- Create a custom class with a Less function
////////////////////////////////////////////////////////////////////////////////

type Minheap3Element struct {
	x int
	y int
}

func Minheap3Less(a, b Minheap3Element) bool {
	return a.x < b.x || a.x == b.x && a.y < b.y
}

type Minheap3 struct {
	buf []Minheap3Element
}

func NewMinheap3() *Minheap3 {
	buf := make([]Minheap3Element, 0)
	return &Minheap3{buf}
}

func (q *Minheap3) Empty() bool {
	return len(q.buf) == 0
}

func (q *Minheap3) Clear() {
	q.buf = q.buf[:0]
}

func (q *Minheap3) Len() int {
	return len(q.buf)
}

func (q *Minheap3) Push(v Minheap3Element) {
	q.buf = append(q.buf, v)
	q.siftdown(0, len(q.buf)-1)
}

func (q *Minheap3) Head() Minheap3Element {
	return q.buf[0]
}

func (q *Minheap3) Pop() Minheap3Element {
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

func (q *Minheap3) Heapify(varr []Minheap3Element) {
	q.buf = append(q.buf, varr...)
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}

func (q *Minheap3) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if !Minheap3Less(newitem, parent) {
			break
		}
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}

func (q *Minheap3) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !Minheap3Less(q.buf[childpos], q.buf[rightpos]) {
			childpos = rightpos
		}
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}
