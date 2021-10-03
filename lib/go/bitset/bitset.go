package bitset

import "math/bits"

// START HERE

type Bitset struct {
	m int
	c []uint64
}

func NewBitset(cap int) *Bitset { return &Bitset{0, make([]uint64, 0, cap)} }
func (q *Bitset) Copy() *Bitset {
	c2 := make([]uint64, len(q.c))
	for i, x := range q.c {
		c2[i] = x
	}
	return &Bitset{q.m, c2}
}
func (q *Bitset) Ins(n int) {
	for q.m <= n {
		q.c = append(q.c, 0)
		q.m += 64
	}
	q.c[n/64] |= 1 << (n % 64)
}
func (q *Bitset) Del(n int) {
	if q.m > n {
		q.c[n/64] &= 0xffffffffffffffff ^ (1 << (n % 64))
	}
}
func (q *Bitset) Flip(n int) {
	for q.m <= n {
		q.c = append(q.c, 0)
		q.m += 64
	}
	q.c[n/64] ^= 1 << (n % 64)
}
func (q *Bitset) Size() int { return q.m }
func (q *Bitset) Any() bool {
	for _, cc := range q.c {
		if cc != 0 {
			return true
		}
	}
	return false
}
func (q *Bitset) None() bool {
	for _, cc := range q.c {
		if cc != 0 {
			return false
		}
	}
	return true
}
func (q *Bitset) Count() int {
	ans := 0
	for _, cc := range q.c {
		ans += bits.OnesCount64(cc)
	}
	return ans
}
func (q *Bitset) PadTo(a *Bitset) {
	for q.m < a.m {
		q.c = append(q.c, 0)
		q.m += 64
	}
}
func (q *Bitset) And(a *Bitset) {
	q.shrinkTo(q.m)
	la := len(a.c)
	for i := 0; i < la; i++ {
		q.c[i] &= a.c[i]
	}
}
func (q *Bitset) Or(a *Bitset) {
	q.PadTo(a)
	la := len(a.c)
	for i := 0; i < la; i++ {
		q.c[i] |= a.c[i]
	}
}
func (q *Bitset) Xor(a *Bitset) {
	q.PadTo(a)
	la := len(a.c)
	for i := 0; i < la; i++ {
		q.c[i] ^= a.c[i]
	}
}
func (q *Bitset) Cap(n int) { q.shrinkTo(n) }
func (q *Bitset) Not() {
	lc := len(q.c)
	for i := 0; i < lc; i++ {
		q.c[i] = ^q.c[i]
	}
}
func (q *Bitset) Shl(a int) {
	q.shrink()
	if q.m == 0 {
		return
	}
	mm := q.max() + 1
	newmm := mm + a
	for q.m < newmm {
		q.c = append(q.c, 0)
		q.m += 64
	}
	g, b := a/64, a%64
	for i := len(q.c) - 1; i >= 0; i-- {
		if i-g < 0 {
			q.c[i] = 0
		} else {
			q.c[i] = q.c[i-g] << b
			if i-g-1 >= 0 && b != 0 {
				q.c[i] |= q.c[i-g-1] >> (64 - b)
			}
		}
	}
}
func (q *Bitset) Shr(a int) {
	g, b, lc := a/64, a%64, len(q.c)
	for i := 0; i < lc; i++ {
		if i+g >= lc {
			q.c[i] = 0
		} else {
			q.c[i] = q.c[i+g] >> b
			if i+g+1 < lc && b != 0 {
				q.c[i] |= q.c[i+g+1] << (64 - b)
			}
		}
	}
	q.shrink()
}
func (q *Bitset) GetBits() []int {
	base := 0
	ans := []int{}
	for _, c := range q.c {
		for c != 0 {
			offset := bits.TrailingZeros64(c)
			ans = append(ans, base+offset)
			c ^= 1 << offset
		}
		base += 64
	}
	return ans
}
func (q *Bitset) shrink() {
	for i := len(q.c) - 1; i >= 0 && q.c[i] == 0; i-- {
		q.c = q.c[:i]
		q.m -= 64
	}
}
func (q *Bitset) shrinkTo(a int) {
	i := len(q.c) - 1
	for q.m-64 > a {
		q.c = q.c[:i]
		q.m -= 64
	}
}
func (q *Bitset) max() int {
	lc := len(q.c)
	if q.c[lc-1] == 0 {
		q.shrink()
		lc = len(q.c)
	}
	if lc == 0 {
		return -1
	}
	return 64*lc - 1 - bits.LeadingZeros64(q.c[lc-1])
}
func BitsetAnd(a, b *Bitset) *Bitset     { c := a.Copy(); c.And(b); return c }
func BitsetOr(a, b *Bitset) *Bitset      { c := a.Copy(); c.Or(b); return c }
func BitsetXor(a, b *Bitset) *Bitset     { c := a.Copy(); c.Xor(b); return c }
func BitsetShl(a *Bitset, n int) *Bitset { c := a.Copy(); c.Shl(n); return c }
func BitsetShr(a *Bitset, n int) *Bitset { c := a.Copy(); c.Shr(n); return c }
