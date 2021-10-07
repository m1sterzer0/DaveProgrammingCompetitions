package segtree

type DATATYPE int

// START HERE
type SEGTREE struct {
	n, size, log int
	op           func(DATATYPE, DATATYPE) DATATYPE
	e            DATATYPE
	d            []DATATYPE
}

func NewSEGTREE(n int, op func(DATATYPE, DATATYPE) DATATYPE, e DATATYPE) *SEGTREE {
	v := make([]DATATYPE, n)
	for i := 0; i < n; i++ {
		v[i] = e
	}
	return NewSEGTREEVec(v, op, e)
}
func NewSEGTREEVec(v []DATATYPE, op func(DATATYPE, DATATYPE) DATATYPE, e DATATYPE) *SEGTREE {
	n, sz, log := len(v), 1, 0
	for sz < n {
		sz <<= 1
		log += 1
	}
	d := make([]DATATYPE, 2*sz)
	d[0] = e
	for i := 0; i < n; i++ {
		d[sz+i] = v[i]
	}
	st := &SEGTREE{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}
func (q *SEGTREE) Set(p int, v DATATYPE) {
	p += q.size
	q.d[p] = v
	for i := 1; i <= q.log; i++ {
		q.update(p >> uint(i))
	}
}
func (q *SEGTREE) Get(p int) DATATYPE { return q.d[p+q.size] }

// Gives product from l to r inclusive
func (q *SEGTREE) Prod(l int, r int) DATATYPE {
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
func (q *SEGTREE) Allprod() DATATYPE { return q.d[1] }

// Given monotone f, finds maximum r such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *SEGTREE) MaxRight(l int, f func(DATATYPE) bool) int {
	if l == q.n {
		return q.n - 1
	}
	l += q.size
	sm := q.e
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(q.op(sm, q.d[l])) {
			for l < q.size {
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
func (q *SEGTREE) MinLeft(r int, f func(DATATYPE) bool) int {
	if r < 0 {
		return 0
	}
	r += q.size
	sm := q.e
	r++ //r++ for the fully closed vs. half open
	for {
		r--
		for r > 1 && r%2 == 1 {
			r >>= 1
		}
		if !f(q.op(q.d[r], sm)) {
			for r < q.size {
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
func (q *SEGTREE) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
