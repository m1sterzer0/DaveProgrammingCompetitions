package cpgolib

import (
	"math/rand"
	"testing"
)

type fenwickbrute struct {
	n   int
	buf []int
}

func newfenwickbrute(n int) *fenwickbrute {
	buf := make([]int, n+1)
	return &fenwickbrute{n, buf}
}

func (q *fenwickbrute) clear() {
	for i := 0; i <= q.n; i++ {
		q.buf[i] = 0
	}
}

func (q *fenwickbrute) inc(idx int, val int) {
	q.buf[idx] += val
}

func (q *fenwickbrute) dec(idx int, val int) {
	q.buf[idx] -= val
}

func (q *fenwickbrute) incdec(left int, right int, val int) {
	q.buf[left] += val
	q.buf[right] -= val
}

func (q *fenwickbrute) prefixsum(idx int) int {
	ans := 0
	for i := 0; i <= idx; i++ {
		ans += q.buf[i]
	}
	return ans
}

func (q *fenwickbrute) suffixsum(idx int) int {
	ans := 0
	for i := idx; i <= q.n; i++ {
		ans += q.buf[i]
	}
	return ans
}

func (q *fenwickbrute) rangesum(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		ans += q.buf[i]
	}
	return ans
}

func TestFenwickRand(t *testing.T) {
	r := rand.New(rand.NewSource(8675309))
	q1 := NewFenwick(1000)
	q2 := newfenwickbrute(1000)
	for iter := 0; iter < 10; iter++ {
		q1.Clear()
		q2.clear()
		for iter2 := 0; iter2 < 10_000; iter2++ {
			idx := 1 + r.Intn(1000)
			if iter2%4 == 0 {
				inc := -2000 + r.Intn(4001)
				q1.Inc(idx, inc)
				q2.inc(idx, inc)
			} else if iter2%4 == 1 {
				dec := -2000 + r.Intn(4001)
				q1.Dec(idx, dec)
				q2.dec(idx, dec)
			} else {
				idx2 := 1 + r.Intn(1000)
				if idx2 < idx {
					idx, idx2 = idx2, idx
				}
				val := -2000 + r.Intn(4001)
				q1.IncDec(idx, idx2, val)
				q2.incdec(idx, idx2, val)
			}

			i1, i2, i3, i4 := 1+r.Intn(100), 1+r.Intn(1000), 1+r.Intn(1000), 1+r.Intn(1000)
			v1a := q1.Prefixsum(i1)
			v1b := q2.prefixsum(i1)
			if v1a != v1b {
				t.Error("prefix sum mismatch")
			}
			v2a := q1.Suffixsum(i2)
			v2b := q2.suffixsum(i2)
			if v2a != v2b {
				t.Error("suffix sum mismatch")
			}
			if i3 < i4 {
				v3a := q1.Rangesum(i3, i4)
				v3b := q2.rangesum(i3, i4)
				if v3a != v3b {
					t.Error("range sum mismatch")
				}
			}
		}
	}
}
