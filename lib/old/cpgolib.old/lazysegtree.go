package cpgolib

type LazySegTreeElem struct {
	x int
}

type LazySegTreeFuncIdx struct {
	x int
}

// This version good for Point updates and Range queries
type LazySegTree struct {
	n           int
	size        int
	log         int
	op          func(a, b LazySegTreeElem) LazySegTreeElem
	mapping     func(a LazySegTreeFuncIdx, x LazySegTreeElem) LazySegTreeElem
	composition func(f LazySegTreeFuncIdx, g LazySegTreeFuncIdx) LazySegTreeFuncIdx
	e           LazySegTreeElem
	id          LazySegTreeFuncIdx
	d           []LazySegTreeElem
	lz          []LazySegTreeFuncIdx
}

func NewLazySegTree(n int, op func(a, b LazySegTreeElem) LazySegTreeElem,
	mapping func(a LazySegTreeFuncIdx, x LazySegTreeElem) LazySegTreeElem,
	composition func(f LazySegTreeFuncIdx, g LazySegTreeFuncIdx) LazySegTreeFuncIdx,
	e LazySegTreeElem, id LazySegTreeFuncIdx) *LazySegTree {
	v := make([]LazySegTreeElem, n)
	for i := 0; i < n; i++ {
		v[i] = e
	}
	return NewLazySegTreeVec(v, op, mapping, composition, e, id)
}

func NewLazySegTreeVec(v []LazySegTreeElem,
	op func(a, b LazySegTreeElem) LazySegTreeElem,
	mapping func(a LazySegTreeFuncIdx, x LazySegTreeElem) LazySegTreeElem,
	composition func(f LazySegTreeFuncIdx, g LazySegTreeFuncIdx) LazySegTreeFuncIdx,
	e LazySegTreeElem, id LazySegTreeFuncIdx) *LazySegTree {
	n := len(v)
	sz, log := 1, 0
	for n < sz {
		sz <<= 1
		log += 1
	}
	d := make([]LazySegTreeElem, 2*sz)
	lz := make([]LazySegTreeFuncIdx, sz)
	d[0] = e
	for i := 0; i < n; i++ {
		d[sz+i] = v[i]
		lz[i] = id
	}
	st := &LazySegTree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}

func (q *LazySegTree) Set(p int, v LazySegTreeElem) {
	p += q.size
	for i := q.log; i >= 1; i-- {
		q.push(p >> i)
	}
	q.d[p] = v
	for i := 1; i <= q.log; i++ {
		q.update(p >> i)
	}
}

func (q *LazySegTree) Get(p int) LazySegTreeElem {
	p += q.size
	for i := q.log; i >= 1; i-- {
		q.push(p >> i)
	}
	return q.d[p]
}

func (q *LazySegTree) Prod(l int, r int) LazySegTreeElem {
	// We add 1 to right vs. atcoder, as we want to get all the points from l->r inclusive
	if r < l {
		return q.e
	}
	r += 1
	l += q.size
	r += q.size
	// A bit of overkill, but still O(log(n)) per operation
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			q.push(l >> i)
		}
		if ((r >> i) << i) != r {
			q.push((r - 1) >> i)
		}
	}
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

func (q *LazySegTree) Allprod() LazySegTreeElem {
	return q.d[1]
}

func (q *LazySegTree) Apply(p int, f LazySegTreeFuncIdx) {
	p += q.size
	for i := q.log; i >= 1; i-- {
		q.push(p >> i)
	}
	q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ {
		q.update(p >> i)
	}
}

func (q *LazySegTree) ApplyRange(l int, r int, f LazySegTreeFuncIdx) {
	// Add one, as our versioin applies from l to r inclusive
	if r < l {
		return
	}
	r += 1
	l += q.size
	r += q.size
	// A bit of overkill, but still O(log(n)) per operation
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			q.push(l >> i)
		}
		if ((r >> i) << i) != r {
			q.push((r - 1) >> i)
		}
	}
	l2, r2 := l, r
	for l < r {
		if l&1 != 0 {
			q.allApply(l, f)
			l += 1
		}
		if r&1 != 0 {
			r -= 1
			q.allApply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			q.update(l >> i)
		}
		if ((r >> i) << i) != r {
			q.update((r - 1) >> i)
		}
	}
}

// TODO MaxRight
// TODO MinLeft

func (q *LazySegTree) update(k int) {
	q.d[k] = q.op(q.d[2*k], q.d[2*k+1])
}

func (q *LazySegTree) allApply(k int, f LazySegTreeFuncIdx) {
	q.d[k] = q.mapping(f, q.d[k])
	if k < q.size {
		q.lz[k] = q.composition(f, q.lz[k])
	}
}

func (q *LazySegTree) push(k int) {
	q.allApply(2*k, q.lz[k])
	q.allApply(2*k+1, q.lz[k])
	q.lz[k] = q.id
}
