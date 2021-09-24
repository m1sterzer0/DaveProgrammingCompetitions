package convolver

import "math/bits"

// START HERE
var docstr = "(998244353,3) works"
func CONVOLVERpowmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
type CONVOLVER struct{ mod, primroot, rank2 int; root,iroot,rate2,irate2,rate3,irate3 []int }
func NewCONVOLVER(mod, primroot int) *CONVOLVER {
	rank2 := bits.TrailingZeros(uint(mod-1))
	if rank2 < 3 { panic("Hard wired to work for a significantly large power of 2 in the modulus") }
	root   := make([]int,rank2+1)
	iroot  := make([]int,rank2+1)
	rate2  := make([]int,rank2-2+1)
	irate2 := make([]int,rank2-2+1)
	rate3  := make([]int,rank2-3+1)
	irate3 := make([]int,rank2-3+1)
	root[rank2]  = CONVOLVERpowmod(primroot,(mod-1)>>rank2,mod)
	iroot[rank2] = CONVOLVERpowmod(root[rank2],mod-2,mod)
	for i:=rank2-1;i>=0;i-- {
		root[i] = root[i+1]*root[i+1] % mod
		iroot[i] = iroot[i+1]*iroot[i+1] % mod
	}
	prod,iprod := 1,1
	for i:=0;i<=rank2-2;i++ {
		rate2[i] = root[i+2] * prod % mod
		irate2[i] = iroot[i+2] * iprod % mod
		prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod,iprod = 1,1
	for i:=0;i<=rank2-3;i++ {
		rate3[i] = root[i+3] * prod % mod
		irate3[i] = iroot[i+3] * iprod % mod
		prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &CONVOLVER{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}

func (q *CONVOLVER) butterfly(a []int) {
	mod := q.mod
	n := len(a)
	h := 0; for (1<<h) < n { h++ }
	ll := 0
	for ll < h {
		if (h - ll == 1) {
			p := 1 << (h-ll-1)
			rot := 1
			for s:=0; s < (1 << ll); s++ {
				offset := s << (h - ll)
				for i:=0;i<p;i++ {
					l := a[i+offset]
					r := a[i+offset+p] * rot % mod
					u := l + r; if u >= mod { u -= mod }
					v := l - r; if v < 0 { v += mod }
					a[i+offset] = u
					a[i+offset+p] = v
				}
				if s + 1 != (1 << ll) {
					rot = rot * q.rate2[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll++ 
		} else {
			p := 1 << (h-ll-2)
			rot := 1; imag := q.root[2]
			for s:=0; s < (1 << ll); s++ {
				rot2 := rot * rot % mod
				rot3 := rot2 * rot % mod
				offset := s << (h - ll)
				for i:=0;i<p;i++ {
					mod2 := mod * mod
					a0 := a[i+offset]
					a1 := a[i+offset+p] * rot
					a2 := a[i+offset+2*p] * rot2
					a3 := a[i+offset+3*p] * rot3
					a1na3imag := (a1+mod2-a3) % mod * imag
					na2 := mod2 - a2
					a[i+offset]   = (a0 + a2 + a1 + a3) % mod
					a[i+offset+p] = (a0 + a2 + (2 * mod2 - a1 - a3)) % mod
					a[i+offset+2*p] = (a0 + na2 + a1na3imag) % mod
					a[i+offset+3*p] = (a0 + na2 + (mod2-a1na3imag)) % mod
				}
				if s + 1 != (1 << ll) {
					rot = rot * q.rate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll += 2
		}
	}
}

func (q *CONVOLVER) butterflyinv(a []int) {
	mod := q.mod
	n := len(a)
	h := 0; for (1<<h) < n { h++ }
	ll := h
	for ll > 0 {
		if (ll == 1) {
			p := 1 << (h-ll)
			irot := 1
			for s:=0; s < (1 << (ll-1)); s++ {
				offset := s << (h - ll + 1)
				for i:=0;i<p;i++ {
					l := a[i+offset]
					r := a[i+offset+p]
					u := l + r; if u >= mod { u -= mod }
					v := (mod+l-r) * irot % mod
					a[i+offset] = u
					a[i+offset+p] = v
				}
				if s + 1 != (1 << (ll-1)) {
					irot = irot * q.irate2[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll-- 
		} else {
			p := 1 << (h-ll)
			irot := 1; iimag := q.iroot[2]
			for s:=0; s < (1 << (ll-2)); s++ {
				irot2 := irot * irot % mod
				irot3 := irot2 * irot % mod
				offset := s << (h - ll + 2)
				for i:=0;i<p;i++ {
					a0 := a[i+offset]
					a1 := a[i+offset+p]
					a2 := a[i+offset+2*p]
					a3 := a[i+offset+3*p]
					a2na3iimag := (mod + a2 - a3) * iimag % mod
					a[i+offset]   = (a0 + a1 + a2 + a3) % mod
					a[i+offset+p] = (a0 + (mod-a1) + a2na3iimag) * irot % mod
					a[i+offset+2*p] = (a0 + a1 + (mod-a2) + (mod-a3)) * irot2 % mod
					a[i+offset+3*p] = (a0 + (mod-a1) + (mod - a2na3iimag)) * irot3 % mod 
				}
				if s + 1 != (1 << (ll-2)) {
					irot = irot * q.irate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll -= 2
		}
	}
	iz := CONVOLVERpowmod(n,mod-2,mod)
	for i:=0;i<n;i++ { a[i] = a[i] * iz % mod }
}

func (q *CONVOLVER) convolvefft(a []int, b []int) []int {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	z := 1; for z < finalsz { z *= 2 }
	lena, lenb := len(a), len(b)
	la := make([]int, z); lb := make([]int, z)
	for i := 0; i < lena; i++ { la[i] = a[i] }
	for i := 0; i < lenb; i++ { lb[i] = b[i] }
	q.butterfly(la)
	q.butterfly(lb)
	for i := 0; i < z; i++ { la[i] *= lb[i]; la[i] %= mod }
	q.butterflyinv(la)
	return la[:finalsz]
}

func (q *CONVOLVER) convolvenaive(a []int, b []int) []int {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	ans := make([]int, finalsz)
	for i,a := range a {
		for j,b := range b {
			ans[i+j] += a * b; ans[i+j] %= mod
		}
	}
	return ans
}

func (q *CONVOLVER) Convolve(a []int, b []int) []int {
	lmin := len(a); if len(b) < lmin { lmin = len(b) }
	if lmin <= 60 { return q.convolvenaive(a,b) } else { return q.convolvefft(a,b) } 
}







//func NewCONVOLVER(mod, root, rootdepth int) *CONVOLVER {
//	rootinv, e, m := 1, mod-2, root
//	for e > 0 {
//		if e&1 != 0 {
//			rootinv = rootinv * m % mod
//		}
//		m = m * m % mod
//		e >>= 1
//	}
//	return &CONVOLVER{mod, root, rootinv, 1 << rootdepth}
//}
//
//// Not as fancy as atcoder version which takes double-steps to save modulus operations. Instead modelled after cpalgorithms version
//func (q *CONVOLVER) NTT(a []int, invert bool) {
//	mod := q.mod
//	n := len(a)
//	for i, j := 1, 0; i < n; i++ {
//		bit := n >> 1
//		for ; j&bit != 0; bit >>= 1 {
//			j ^= bit
//		}
//		j ^= bit
//		if i < j {
//			a[i], a[j] = a[j], a[i]
//		}
//	}
//	for ll := 2; ll <= n; ll <<= 1 {
//		wlen := q.root
//		if invert {
//			wlen = q.rootinv
//		}
//		for i := ll; i < q.rootpw; i <<= 1 {
//			wlen = wlen * wlen % mod
//		}
//		for i := 0; i < n; i += ll {
//			w := 1
//			lover2 := ll >> 1
//			for j := 0; j < lover2; j++ {
//				idx1 := i + j
//				idx2 := idx1 + lover2
//				u := a[idx1]
//				v := a[idx2] * w % mod
//				v1 := u + v
//				v2 := u - v
//				if v1 >= mod {
//					v1 -= mod
//				}
//				if v2 < 0 {
//					v2 += mod
//				}
//				a[idx1], a[idx2] = v1, v2
//				w = w * wlen % mod
//			}
//		}
//	}
//	if invert {
//		ninv, e, m := 1, mod-2, n
//		for e > 0 {
//			if e&1 != 0 {
//				ninv = ninv * m % mod
//			}
//			m = m * m % mod
//			e >>= 1
//		}
//		for i := 0; i < n; i++ {
//			a[i] *= ninv
//			a[i] %= mod
//		}
//	}
//}
//func (q *CONVOLVER) Convolve(a []int, b []int) []int {
//	mod := q.mod
//	finalsz := len(a) + len(b) - 1
//	z := 1
//	for z < finalsz {
//		z *= 2
//	}
//	lena, lenb := len(a), len(b)
//	la := make([]int, z)
//	lb := make([]int, z)
//	for i := 0; i < lena; i++ {
//		la[i] = a[i]
//	}
//	for i := 0; i < lenb; i++ {
//		lb[i] = b[i]
//	}
//	q.NTT(la, false)
//	q.NTT(lb, false)
//	for i := 0; i < z; i++ {
//		la[i] *= lb[i]
//		la[i] %= mod
//	}
//	q.NTT(la, true)
//	return la[:finalsz]
//}
