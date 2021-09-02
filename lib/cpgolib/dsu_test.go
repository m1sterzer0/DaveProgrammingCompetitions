package cpgolib

import (
	"fmt"
	"math/rand"
	"testing"
)

type dsubrute1 struct {
	n int
	l []int
}

func newdsubrute1(n int) *dsubrute1 {
	buf := make([]int, n)
	for i := 0; i < n; i++ {
		buf[i] = -1
	}
	return &dsubrute1{n, buf}
}

func (q *dsubrute1) leader(a int) int {
	if q.l[a] < 0 {
		return a
	}
	return q.l[a]
}

func (q *dsubrute1) merge(a int, b int) int {
	l1 := q.leader(a)
	l2 := q.leader(b)
	if l1 != l2 {
		s1 := 1
		s2 := 1
		for _, v := range q.l {
			if v == l1 {
				s1++
			} else if v == l2 {
				s2++
			}
		}
		if s1 < s2 {
			l1, l2, s1, s2 = l2, l1, s2, s1
		}
		for i := 0; i < q.n; i++ {
			if q.l[i] == l2 {
				q.l[i] = l1
			}
		}
		q.l[l2] = l1
		q.l[l1] = -(s1 + s2)
	}
	return l1
}

func (q *dsubrute1) same(a int, b int) bool {
	return q.leader(a) == q.leader(b)
}

func (q *dsubrute1) size(a int) int {
	l := q.leader(a)
	return -q.l[l]
}

func TestDsuRand(t *testing.T) {
	r := rand.New(rand.NewSource(8675309))
	for iter := 0; iter < 10; iter++ {
		q1 := NewDsu(1000)
		q2 := NewDsu2()
		q3 := newdsubrute1(1000)
		for i := 0; i < 1000; i++ {
			q2.Add(i)
		}
		for iter2 := 0; iter2 < 1000; iter2++ {
			idx := r.Intn(1000)
			idx2 := r.Intn(1000)
			q1.Merge(idx, idx2)
			q2.Merge(idx, idx2)
			q3.merge(idx, idx2)
			idx3 := r.Intn(1000)
			v1 := q1.Leader(idx3)
			v2 := q2.Leader(idx3)
			v3 := q3.leader(idx3)
			if v1 != v2 || v1 != v3 {
				s := fmt.Sprintf("Leaders do not match idx3:%v v1:%v v2:%v v3:%v\n", idx3, v1, v2, v3)
				t.Error(s)
			}
			idx4 := r.Intn(1000)
			v1 = q1.Size(idx4)
			v2 = q2.Size(idx4)
			v3 = q3.size(idx4)
			if v1 != v2 || v1 != v3 {
				s := fmt.Sprintf("Sizes do not match idx4:%v v1:%v v2:%v v3:%v\n", idx4, v1, v2, v3)
				t.Error(s)
			}
			idx5 := r.Intn(1000)
			idx6 := r.Intn(1000)
			b1 := q1.Same(idx5, idx6)
			b2 := q2.Same(idx5, idx6)
			b3 := q3.same(idx5, idx6)
			if b1 != b2 || b1 != b3 {
				s := fmt.Sprintf("Same call outputs do not match idx5:%v idx6:%v b1:%v b2:%v b3:%v\n", idx5, idx6, v1, v2, v3)
				t.Error(s)
			}
		}
	}
}
