package convolver

// START HERE
type CONVOLVER struct{ mod, root, rootinv, rootpw int }

// (998244353,31,23) works
func NewCONVOLVER(mod, root, rootdepth int) *CONVOLVER {
	rootinv, e, m := 1, mod-2, root
	for e > 0 {
		if e&1 != 0 {
			rootinv = rootinv * m % mod
		}
		m = m * m % mod
		e >>= 1
	}
	return &CONVOLVER{mod, root, rootinv, 1 << rootdepth}
}

// Not as fancy as atcoder version which takes double-steps to save modulus operations. Instead modelled after cpalgorithms version
func (q *CONVOLVER) NTT(a []int, invert bool) {
	mod := q.mod
	n := len(a)
	for i, j := 1, 0; i < n; i++ {
		bit := n >> 1
		for ; j&bit != 0; bit >>= 1 {
			j ^= bit
		}
		j ^= bit
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	for ll := 2; ll <= n; ll <<= 1 {
		wlen := q.root
		if invert {
			wlen = q.rootinv
		}
		for i := ll; i < q.rootpw; i <<= 1 {
			wlen = wlen * wlen % mod
		}
		for i := 0; i < n; i += ll {
			w := 1
			lover2 := ll >> 1
			for j := 0; j < lover2; j++ {
				idx1 := i + j
				idx2 := idx1 + lover2
				u := a[idx1]
				v := a[idx2] * w % mod
				v1 := u + v
				v2 := u - v
				if v1 >= mod {
					v1 -= mod
				}
				if v2 < 0 {
					v2 += mod
				}
				a[idx1], a[idx2] = v1, v2
				w = w * wlen % mod
			}
		}
	}
	if invert {
		ninv, e, m := 1, mod-2, n
		for e > 0 {
			if e&1 != 0 {
				ninv = ninv * m % mod
			}
			m = m * m % mod
			e >>= 1
		}
		for i := 0; i < n; i++ {
			a[i] *= ninv
			a[i] %= mod
		}
	}
}
func (q *CONVOLVER) Convolve(a []int, b []int) []int {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	z := 1
	for z < finalsz {
		z *= 2
	}
	lena, lenb := len(a), len(b)
	la := make([]int, z)
	lb := make([]int, z)
	for i := 0; i < lena; i++ {
		la[i] = a[i]
	}
	for i := 0; i < lenb; i++ {
		lb[i] = b[i]
	}
	q.NTT(la, false)
	q.NTT(lb, false)
	for i := 0; i < z; i++ {
		la[i] *= lb[i]
		la[i] %= mod
	}
	q.NTT(la, true)
	return la[:finalsz]
}
