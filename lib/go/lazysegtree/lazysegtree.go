package lazysegtree

type DATATYPE int
type FUNCTYPE int

// START HERE
type LAZYSEGTREE struct {
	n, size, log int
	op           func(DATATYPE, DATATYPE) DATATYPE
	mapping      func(FUNCTYPE, DATATYPE) DATATYPE
	composition  func(FUNCTYPE, FUNCTYPE) FUNCTYPE
	e            DATATYPE
	id           FUNCTYPE
	d            []DATATYPE
	lz           []FUNCTYPE
}

func NewLAZYSEGTREE(n int, op func(DATATYPE, DATATYPE) DATATYPE, mapping func(FUNCTYPE, DATATYPE) DATATYPE, composition func(FUNCTYPE, FUNCTYPE) FUNCTYPE, e DATATYPE, id FUNCTYPE) *LAZYSEGTREE {
	v := make([]DATATYPE, n)
	for i := 0; i < n; i++ {
		v[i] = e
	}
	return NewLAZYSEGTREEVec(v, op, mapping, composition, e, id)
}
func NewLAZYSEGTREEVec(v []DATATYPE, op func(DATATYPE, DATATYPE) DATATYPE, mapping func(FUNCTYPE, DATATYPE) DATATYPE, composition func(FUNCTYPE, FUNCTYPE) FUNCTYPE, e DATATYPE, id FUNCTYPE) *LAZYSEGTREE {
	n, sz, log := len(v), 1, 0
	for n < sz {
		sz <<= 1
		log += 1
	}
	d := make([]DATATYPE, 2*sz)
	lz := make([]FUNCTYPE, sz)
	d[0] = e
	for i := 0; i < n; i++ {
		d[sz+i] = v[i]
		lz[i] = id
	}
	st := &LAZYSEGTREE{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}
func (q *LAZYSEGTREE) Set(p int, v DATATYPE) {
	p += q.size
	for i := q.log; i >= 1; i-- {
		q.push(p >> i)
	}
	q.d[p] = v
	for i := 1; i <= q.log; i++ {
		q.update(p >> i)
	}
}
func (q *LAZYSEGTREE) Get(p int) DATATYPE {
	p += q.size
	for i := q.log; i >= 1; i-- {
		q.push(p >> i)
	}
	return q.d[p]
}

// Gets the product from l to r incluseive
func (q *LAZYSEGTREE) Prod(l int, r int) DATATYPE {
	if r < l {
		return q.e
	}
	l += q.size
	r += q.size
	r += 1 // r+1 for close right end interval
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
func (q *LAZYSEGTREE) Allprod() DATATYPE { return q.d[1] }
func (q *LAZYSEGTREE) Apply(p int, f FUNCTYPE) {
	p += q.size
	for i := q.log; i >= 1; i-- {
		q.push(p >> i)
	}
	q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ {
		q.update(p >> i)
	}
}
func (q *LAZYSEGTREE) ApplyRange(l int, r int, f FUNCTYPE) {
	// Add one, as our versioin applies from l to r inclusive
	if r < l {
		return
	}
	r += 1
	l += q.size
	r += q.size
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

// Given monotone f, finds maximum r such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *LAZYSEGTREE) MaxRight(l int, f func(DATATYPE) bool) int {
	if l == q.n {
		return q.n - 1
	}
	l += q.size
	for i := q.log; i >= 1; i-- {
		q.push(l >> i)
	}
	sm := q.e
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(q.op(sm, q.d[l])) {
			for l < q.size {
				q.push(l)
				l *= 2
				if f(q.op(sm, q.d[l])) {
					sm = q.op(sm, q.d[l])
					l++
				}
			}
			return l - q.size - 1
		}
		sm = q.op(sm, q.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return q.n - 1
}

// Given monotone f, finds minimum l such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *LAZYSEGTREE) MinLeft(r int, f func(DATATYPE) bool) int {
	if r < 0 {
		return 0
	}
	r += q.size
	r++
	for i := q.log; i >= 1; i-- {
		q.push((r - 1) >> i)
	}
	sm := q.e //r++ for the fully closed vs. half open
	for {
		r--
		for r > 1 && r%2 == 1 {
			r >>= 1
		}
		if !f(q.op(q.d[r], sm)) {
			for r < q.size {
				q.push(r)
				r = 2*r + 1
				if f(q.op(q.d[r], sm)) {
					sm = q.op(q.d[r], sm)
					r--
				}
			}
			return r + 1 - q.size
		}
		sm = q.op(q.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}
func (q *LAZYSEGTREE) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *LAZYSEGTREE) allApply(k int, f FUNCTYPE) {
	q.d[k] = q.mapping(f, q.d[k])
	if k < q.size {
		q.lz[k] = q.composition(f, q.lz[k])
	}
}
func (q *LAZYSEGTREE) push(k int) {
	q.allApply(2*k, q.lz[k])
	q.allApply(2*k+1, q.lz[k])
	q.lz[k] = q.id
}
