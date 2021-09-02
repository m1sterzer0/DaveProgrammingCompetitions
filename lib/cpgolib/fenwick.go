package cpgolib

type Fenwick struct {
	n   int
	tot int
	bit []int
}

func NewFenwick(n int) *Fenwick {
	buf := make([]int, n+1)
	return &Fenwick{n, 0, buf}
}

func (q *Fenwick) Clear() {
	for i := 0; i <= q.n; i++ {
		q.bit[i] = 0
	}
	q.tot = 0
}

func (q *Fenwick) Inc(idx int, val int) {
	for idx <= q.n {
		q.bit[idx] += val
		idx += idx & (-idx)
	}
	q.tot += val
}

func (q *Fenwick) Dec(idx int, val int) {
	q.Inc(idx, -val)
}

func (q *Fenwick) IncDec(left int, right int, val int) {
	q.Inc(left, val)
	q.Dec(right, val)
}

func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 {
		return 0
	}
	ans := 0
	for idx > 0 {
		ans += q.bit[idx]
		idx -= idx & (-idx)
	}
	return ans
}

func (q *Fenwick) Suffixsum(idx int) int {
	return q.tot - q.Prefixsum(idx-1)
}

func (q *Fenwick) Rangesum(left int, right int) int {
	if right < left {
		return 0
	}
	return q.Prefixsum(right) - q.Prefixsum(left-1)
}
