package cpgolib

import (
	"math/rand"
	"testing"
)

type minheapbrute struct {
	buf []int
}

func newminheapbrute() *minheapbrute {
	buf := make([]int, 0)
	return &minheapbrute{buf}
}

func (q *minheapbrute) empty() bool {
	return len(q.buf) == 0
}

func (q *minheapbrute) clear() {
	q.buf = q.buf[:0]
}

func (q *minheapbrute) len() int {
	return len(q.buf)
}

func (q *minheapbrute) push(v int) {
	q.buf = append(q.buf, v)
}

func (q *minheapbrute) head() int {
	m := q.buf[0]
	for j := 0; j < len(q.buf); j++ {
		if q.buf[j] < m {
			m = q.buf[j]
		}
	}
	return m
}

func (q *minheapbrute) pop() int {
	i, m := 0, q.buf[0]
	for j := 0; j < len(q.buf); j++ {
		if q.buf[j] < m {
			i = j
			m = q.buf[j]
		}
	}
	q.buf = append(q.buf[:i], q.buf[i+1:]...)
	return m
}

//func (q *minheapbrute) heapify(pri []int) {
//	q.buf = append(q.buf, pri...)
//}

func TestMinheapRand(t *testing.T) {
	r := rand.New(rand.NewSource(8675309))
	q1 := NewMinheap1()
	q2 := NewMinheap2()
	q3 := NewMinheap3()
	q4 := newminheapbrute()
	for iter := 0; iter < 10; iter++ {
		q1.Clear()
		q2.Clear()
		q3.Clear()
		q4.clear()
		l := 0
		for phase := 0; phase < 20; phase++ {
			phaselen := 250 + r.Intn(200)
			writeprob := r.Float64()
			for iter := 0; iter < phaselen; iter++ {
				s1 := r.Float64()
				if s1 < writeprob {
					v := r.Intn(1_000_000_000_000_000_000)
					q1.Push(v, v)
					q2.Push(v)
					q3.Push(Minheap3Element{v, v})
					q4.push(v)
					l += 1
				} else if l > 0 {
					_, prev1 := q1.Pop()
					v1 := prev1.(int)
					v2 := q2.Pop()
					v3 := q3.Pop().x
					v4 := q4.pop()
					if v1 != v2 || v1 != v3 || v1 != v4 {
						t.Error("Popped values mismatch")
					}
					l -= 1
				}
				if l > 0 {
					_, prev1 := q1.Head()
					v1 := prev1.(int)
					v2 := q2.Head()
					v3 := q3.Head().x
					v4 := q4.head()
					if v1 != v2 || v1 != v3 || v1 != v4 {
						t.Error("Head values mismatch")
					}
				}

				l1 := q1.Len()
				l2 := q2.Len()
				l3 := q3.Len()
				l4 := q4.len()
				if l1 != l2 || l1 != l3 || l1 != l4 {
					t.Error("Lengths mismatch")
				}

				b1 := q1.Empty()
				b2 := q2.Empty()
				b3 := q3.Empty()
				b4 := q4.empty()
				b5 := l == 0
				if b1 != b2 || b1 != b3 || b1 != b4 || b1 != b5 {
					t.Error("Empty mismatch")
				}
			}
		}
	}
}
