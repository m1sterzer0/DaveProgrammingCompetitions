package cpgolib

type SegTreeElem struct {
	x int
}

// This version good for Point updates and Range queries
type SegTree struct {
	n    int
	size int
	log  int
	op   func(a, b SegTreeElem) SegTreeElem
	e    SegTreeElem
	d    []SegTreeElem
}

func NewSegTree(n int, op func(a, b SegTreeElem) SegTreeElem, e SegTreeElem) *SegTree {
	v := make([]SegTreeElem, n)
	for i := 0; i < n; i++ {
		v[i] = e
	}
	return NewSegTreeVec(v, op, e)
}

func NewSegTreeVec(v []SegTreeElem, op func(a, b SegTreeElem) SegTreeElem, e SegTreeElem) *SegTree {
	n := len(v)
	sz, log := 1, 0
	for n < sz {
		sz <<= 1
		log += 1
	}
	d := make([]SegTreeElem, 2*sz)
	d[0] = e
	for i := 0; i < n; i++ {
		d[sz+i] = v[i]
	}
	st := &SegTree{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}

func (q *SegTree) Set(p int, v SegTreeElem) {
	p += q.size
	q.d[p] = v
	for i := 1; i <= q.log; i++ {
		q.update(p >> i)
	}
}

func (q *SegTree) Get(p int) SegTreeElem {
	return q.d[p+q.size]
}

func (q *SegTree) Prod(l int, r int) SegTreeElem {
	// We add 1 to right vs. atcoder, as we want to get all the points from l->r inclusive
	if r < l {
		return q.e
	}
	r += 1
	sml, smr := q.e, q.e
	l += q.size
	r += q.size
	for l < r {
		if l&1 != 0 {
			sml = q.op(sml, q.d[l])
			l++
		}
		if r&1 != 0 {
			r--
			smr = q.op(q.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return q.op(sml, smr)
}

func (q *SegTree) Allprod() SegTreeElem {
	return q.d[1]
}

// TODO MaxRight
// TODO MinLeft

func (q *SegTree) update(k int) {
	q.d[k] = q.op(q.d[2*k], q.d[2*k+1])
}
