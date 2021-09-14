package cpgolib

type Deque struct {
	buf  []interface{}
	head int
	tail int
	sz   int
	bm   int
	l    int
}

func NewDeque() *Deque {
	buf := make([]interface{}, 8)
	return &Deque{buf, 0, 0, 8, 7, 0}
}

func (q *Deque) Empty() bool { return q.l == 0 }

func (q *Deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }

func (q *Deque) PushFront(x interface{}) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.head = (q.head - 1) & q.bm
	}
	q.l++
	q.buf[q.head] = x
}

func (q *Deque) PushBack(x interface{}) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.tail = (q.tail + 1) & q.bm
	}
	q.l++
	q.buf[q.tail] = x
}

func (q *Deque) PopFront() interface{} {
	if q.l == 0 {
		q.errorPopWhenEmpty()
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

func (q *Deque) PopBack() interface{} {
	if q.l == 0 {
		q.errorPopWhenEmpty()
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

func (q *Deque) Len() int { return q.l }

func (q *Deque) Head() interface{} {
	if q.l == 0 {
		q.errorEmptyAccess()
	}
	return q.buf[q.head]
}

func (q *Deque) Tail() interface{} {
	if q.l == 0 {
		q.errorEmptyAccess()
	}
	return q.buf[q.tail]
}

func (q *Deque) sizeup() {
	buf := make([]interface{}, 2*q.sz)
	for i := 0; i < q.l; i++ {
		buf[i] = q.buf[(q.head+i)&q.bm]
	}
	q.buf = buf
	q.head = 0
	q.tail = q.sz - 1
	q.sz = 2 * q.sz
	q.bm = q.sz - 1
}

func (q *Deque) errorPopWhenEmpty() {
	panic("Tried to pop from an empty deque. Panicking...")
}

func (q *Deque) errorEmptyAccess() {
	panic("Tried to access element from an empty deque. Panicking...")
}

func (q *Deque) Append(x interface{}) { q.PushBack(x) }

func (q *Deque) AppendLeft(x interface{}) { q.PushFront(x) }

func (q *Deque) Push(x interface{}) { q.PushBack(x) }

func (q *Deque) Pop() interface{} { return q.PopBack() }

func (q *Deque) PopLeft() interface{} { return q.PopFront() }

type Dequeint struct {
	buf                   []int
	head, tail, sz, bm, l int
}

func NewDequeint() *Dequeint    { buf := make([]int, 8); return &Dequeint{buf, 0, 0, 8, 7, 0} }
func (q *Dequeint) Empty() bool { return q.l == 0 }
func (q *Dequeint) Clear()      { q.head = 0; q.tail = 0; q.l = 0 }
func (q *Dequeint) PushFront(x int) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.head = (q.head - 1) & q.bm
	}
	q.l++
	q.buf[q.head] = x
}
func (q *Dequeint) PushBack(x int) {
	if q.l == q.sz {
		q.sizeup()
	}
	if q.l > 0 {
		q.tail = (q.tail + 1) & q.bm
	}
	q.l++
	q.buf[q.tail] = x
}
func (q *Dequeint) Len() int { return q.l }
func (q *Dequeint) Head() int {
	if q.l == 0 {
		q.errorEmptyAccess()
	}
	return q.buf[q.head]
}
func (q *Dequeint) Tail() int {
	if q.l == 0 {
		q.errorEmptyAccess()
	}
	return q.buf[q.tail]
}
func (q *Dequeint) PopFront() int {
	if q.l == 0 {
		q.errorPopWhenEmpty()
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
func (q *Dequeint) PopBack() int {
	if q.l == 0 {
		q.errorPopWhenEmpty()
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
func (q *Dequeint) sizeup() {
	buf := make([]int, 2*q.sz)
	for i := 0; i < q.l; i++ {
		buf[i] = q.buf[(q.head+i)&q.bm]
	}
	q.buf = buf
	q.head = 0
	q.tail = q.sz - 1
	q.sz = 2 * q.sz
	q.bm = q.sz - 1
}
func (q *Dequeint) errorPopWhenEmpty() { panic("Tried to pop from an empty deque. Panicking...") }
func (q *Dequeint) errorEmptyAccess() {
	panic("Tried to access element from an empty deque. Panicking...")
}
func (q *Dequeint) Append(x int)     { q.PushBack(x) }
func (q *Dequeint) AppendLeft(x int) { q.PushFront(x) }
func (q *Dequeint) Push(x int)       { q.PushBack(x) }
func (q *Dequeint) Pop() int         { return q.PopBack() }
func (q *Dequeint) PopLeft() int     { return q.PopFront() }

//type Dequeint struct {
//	buf                   []int
//	head, tail, sz, bm, l int
//}
//
//func NewDequeint() *Dequeint    { buf := make([]int, 8); return &Dequeint{buf, 0, 0, 8, 7, 0} }
//func (q *Dequeint) Empty() bool { return q.l == 0 }
//func (q *Dequeint) Clear()      { q.head = 0; q.tail = 0; q.l = 0 }
//func (q *Dequeint) PushFront(x int) {
//	if q.l == q.sz {
//		q.sizeup()
//	}
//	if q.l > 0 {
//		q.head = (q.head - 1) & q.bm
//	}
//	q.l++
//	q.buf[q.head] = x
//}
//func (q *Dequeint) PushBack(x int) {
//	if q.l == q.sz {
//		q.sizeup()
//	}
//	if q.l > 0 {
//		q.tail = (q.tail + 1) & q.bm
//	}
//	q.l++
//	q.buf[q.tail] = x
//}
//func (q *Dequeint) Len() int { return q.l }
//func (q *Dequeint) Head() int {
//	if q.l == 0 {
//		q.errorEmptyAccess()
//	}
//	return q.buf[q.head]
//}
//func (q *Dequeint) Tail() int {
//	if q.l == 0 {
//		q.errorEmptyAccess()
//	}
//	return q.buf[q.tail]
//}
//func (q *Dequeint) PopFront() int {
//	if q.l == 0 {
//		q.errorPopWhenEmpty()
//	}
//	v := q.buf[q.head]
//	q.l--
//	if q.l > 0 {
//		q.head = (q.head + 1) & q.bm
//	} else {
//		q.Clear()
//	}
//	return v
//}
//func (q *Dequeint) PopBack() int {
//	if q.l == 0 {
//		q.errorPopWhenEmpty()
//	}
//	v := q.buf[q.tail]
//	q.l--
//	if q.l > 0 {
//		q.tail = (q.tail - 1) & q.bm
//	} else {
//		q.Clear()
//	}
//	return v
//}
//func (q *Dequeint) sizeup() {
//	buf := make([]int, 2*q.sz)
//	for i := 0; i < q.l; i++ {
//		buf[i] = q.buf[(q.head+i)&q.bm]
//	}
//	q.buf = buf
//	q.head = 0
//	q.tail = q.sz - 1
//	q.sz = 2 * q.sz
//	q.bm = q.sz - 1
//}
//func (q *Dequeint) errorPopWhenEmpty() { panic("Tried to pop from an empty deque. Panicking...") }
//func (q *Dequeint) errorEmptyAccess() {
//	panic("Tried to access element from an empty deque. Panicking...")
//}
//func (q *Dequeint) Append(x int)     { q.PushBack(x) }
//func (q *Dequeint) AppendLeft(x int) { q.PushFront(x) }
//func (q *Dequeint) Push(x int)       { q.PushBack(x) }
//func (q *Dequeint) Pop() int         { return q.PopBack() }
//func (q *Dequeint) PopLeft() int     { return q.PopFront() }
//