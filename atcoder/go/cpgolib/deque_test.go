package cpgolib

import (
	"fmt"
	"math/rand"
	"testing"
)

type dequebrute struct {
	buf []interface{}
}

func newdequebrute() *dequebrute {
	buf := make([]interface{}, 0)
	return &dequebrute{buf}
}

func (q *dequebrute) empty() bool { return len(q.buf) == 0 }

func (q *dequebrute) clear() {
	if len(q.buf) > 0 {
		q.buf = q.buf[:0]
	}
}

func (q *dequebrute) pushFront(x interface{}) {
	m := []interface{}{x}
	q.buf = append(m, q.buf...)
}

func (q *dequebrute) pushBack(x interface{}) {
	q.buf = append(q.buf, x)
}

func (q *dequebrute) popFront() interface{} {
	if len(q.buf) == 0 {
		q.errorPopWhenEmpty()
	}
	v := q.buf[0]
	q.buf = q.buf[1:]
	return v
}

func (q *dequebrute) popBack() interface{} {
	if len(q.buf) == 0 {
		q.errorPopWhenEmpty()
	}
	n := len(q.buf)
	v := q.buf[n-1]
	q.buf = q.buf[0 : n-1]
	return v
}

func (q *dequebrute) len() int {
	return len(q.buf)
}

func (q *dequebrute) head() interface{} {
	if len(q.buf) == 0 {
		q.errorEmptyAccess()
	}
	return q.buf[0]
}

func (q *dequebrute) tail() interface{} {
	if len(q.buf) == 0 {
		q.errorEmptyAccess()
	}
	return q.buf[len(q.buf)-1]
}

func (q *dequebrute) errorPopWhenEmpty() {
	panic("Tried to pop from an empty dequebrute. Panicking...")
}

func (q *dequebrute) errorEmptyAccess() {
	panic("Tried to access element from an empty deque. Panicking...")
}

func TestFIFOSimple(t *testing.T) {
	q1 := NewDeque()
	q2 := NewDequeint()
	q3 := newdequebrute()
	for j := 0; j < 10; j++ {
		for i := 0; i < 16; i++ {
			q1.Push(i)
			q2.Push(i)
			q3.pushBack(i)
		}
		for i := 15; i >= 0; i-- {
			v1 := q1.Pop()
			v2 := q2.Pop()
			v3 := q3.popBack()
			if v1 != i {
				s := fmt.Sprintf("Wrong value popped from Deque. v1:%v v2:%v v3:%v i:%v", v1, v2, v3, i)
				t.Error(s)
			}
			if v2 != i {
				s := fmt.Sprintf("Wrong value popped from Dequeint. v1:%v v2:%v v3:%v i:%v", v1, v2, v3, i)
				t.Error(s)
			}
			if v3 != i {
				s := fmt.Sprintf("Wrong value popped from dequebrute. v1:%v v2:%v v3:%v i:%v", v1, v2, v3, i)
				t.Error(s)
			}

		}
	}
}

func TestDequeRand(t *testing.T) {
	r := rand.New(rand.NewSource(8675309))
	q1 := NewDeque()
	q2 := NewDequeint()
	q3 := newdequebrute()
	for iter := 0; iter < 10; iter++ {
		q1.Clear()
		q2.Clear()
		q3.clear()
		l := 0
		for phase := 0; phase < 20; phase++ {
			phaselen := 250 + r.Intn(2000)
			backprob := r.Float64()
			writeprob := r.Float64()
			popprob := r.Float64()
			for iter := 0; iter < phaselen; iter++ {
				s1, s2, s3, s4 := r.Float64(), r.Float64(), r.Float64(), r.Float64()
				if s1 < writeprob {
					v := r.Intn(1_000_000_000_000_000_000)
					if s2 < backprob {
						if s4 < 0.333 {
							q1.PushBack(v)
							q2.PushBack(v)
						} else if s4 < 0.666 {
							q1.Append(v)
							q2.Append(v)
						} else {
							q1.Push(v)
							q2.Push(v)
						}
						q3.pushBack(v)
					} else {
						if s4 < 0.500 {
							q1.PushFront(v)
							q2.PushFront(v)
						} else {
							q1.AppendLeft(v)
							q2.AppendLeft(v)
						}
						q3.pushFront(v)
					}
					l += 1
				} else if l > 0 {
					v1, v2, v3 := 0, 0, 0
					if s2 < backprob {
						v1 = q1.Tail().(int)
						v2 = q2.Tail()
						v3 = q3.tail().(int)
					} else {
						v1 = q1.Head().(int)
						v2 = q2.Head()
						v3 = q3.head().(int)
					}
					if v1 != v2 && v2 == v3 {
						t.Error("Read failure in Deque")
					}
					if v1 != v3 && v1 == v2 {
						t.Error("Read failure in Dequebrute")
					}
					if v1 != v2 && v1 == v3 {
						t.Error("Read failure in Dequeint")
					}
					if s3 < popprob {
						if s2 < backprob {
							if s4 < 0.500 {
								v1 = q1.PopBack().(int)
								v2 = q2.PopBack()
							} else {
								v1 = q1.Pop().(int)
								v2 = q2.Pop()
							}
							v3 = q3.popBack().(int)
						} else {
							if s4 < 0.500 {
								v1 = q1.PopFront().(int)
								v2 = q2.PopFront()
							} else {
								v1 = q1.PopLeft().(int)
								v2 = q2.PopLeft()
							}
							v3 = q3.popFront().(int)
						}
						if v1 != v2 && v2 == v3 {
							t.Error("Pop failure in Deque")
						}
						if v1 != v3 && v1 == v2 {
							t.Error("Pop failure in Dequebrute")
						}
						if v1 != v2 && v1 == v3 {
							t.Error("Pop failure in Dequeint")
						}
						l -= 1
					}
				}
				l1 := q1.Len()
				l2 := q2.Len()
				l3 := q3.len()
				e1 := q1.Empty()
				e2 := q2.Empty()
				e3 := q3.empty()
				refe := l == 0
				if l1 != l || l2 != l || l3 != l {
					t.Error("Length error")
				}
				if refe != e1 || refe != e2 || refe != e3 {
					t.Error("Empty error")
				}

			}
		}
	}
}
