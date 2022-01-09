package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000

type query struct { id, amt int }
type event struct { id, qid, typ int }
type stsearch struct {id,l,r int }

type Bitset struct { m int; c []uint64 }
func NewBitset(cap int) *Bitset { return &Bitset{0, make([]uint64, 0, cap)} }
func (q *Bitset) Copy() *Bitset {
	c2 := make([]uint64, len(q.c)); for i, x := range q.c { c2[i] = x }; return &Bitset{q.m, c2}
}
func (q *Bitset) Ins(n int) { for q.m <= n { q.c = append(q.c, 0); q.m += 64 }; q.c[n/64] |= 1 << uint(n % 64) }
func (q *Bitset) Del(n int) { if q.m > n { q.c[n/64] &= 0xffffffffffffffff ^ (1 << uint(n % 64)) } }
func (q *Bitset) Flip(n int) { for q.m <= n { q.c = append(q.c, 0); q.m += 64 }; q.c[n/64] ^= 1 << uint(n % 64) }
func (q *Bitset) Size() int { return q.m }
func (q *Bitset) Any() bool { for _, cc := range q.c { if cc != 0 { return true } }; return false }
func (q *Bitset) None() bool { for _, cc := range q.c { if cc != 0 { return false } }; return true }
func (q *Bitset) Count() int { ans := 0; for _, cc := range q.c { ans += bits.OnesCount64(cc) }; return ans }
func (q *Bitset) PadTo(a *Bitset) { for q.m < a.m { q.c = append(q.c, 0); q.m += 64 } }
func (q *Bitset) And(a *Bitset) { if a.m < q.m { q.shrinkTo(a.m) }; lq := len(q.c); for i := 0; i < lq; i++ { q.c[i] &= a.c[i] } }
func (q *Bitset) Or(a *Bitset) { q.PadTo(a); la := len(a.c); for i := 0; i < la; i++ { q.c[i] |= a.c[i] } }
func (q *Bitset) Xor(a *Bitset) { q.PadTo(a); la := len(a.c); for i := 0; i < la; i++ { q.c[i] ^= a.c[i] } }
func (q *Bitset) Cap(n int) { q.shrinkTo(n) }
func (q *Bitset) Not() { lc := len(q.c); for i := 0; i < lc; i++ { q.c[i] = ^q.c[i] } }
func (q *Bitset) Shl(a int) {
	q.shrink(); if q.m == 0 { return }; mm := q.max() + 1; newmm := mm + a
	for q.m < newmm { q.c = append(q.c, 0); q.m += 64 }; g, b := a/64, a%64
	for i := len(q.c) - 1; i >= 0; i-- {
		if i-g < 0 {
			q.c[i] = 0
		} else {
			q.c[i] = q.c[i-g] << uint(b); if i-g-1 >= 0 && b != 0 { q.c[i] |= q.c[i-g-1] >> uint(64 - b) }
		}
	}
}
func (q *Bitset) Shr(a int) {
	g, b, lc := a/64, a%64, len(q.c)
	for i := 0; i < lc; i++ {
		if i+g >= lc {
			q.c[i] = 0
		} else {
			q.c[i] = q.c[i+g] >> uint(b); if i+g+1 < lc && b != 0 { q.c[i] |= q.c[i+g+1] << uint(64 - b) }
		}
	}
	q.shrink()
}
func (q *Bitset) GetBits() []int {
	base := 0; ans := []int{}
	for _, c := range q.c {
		for c != 0 { offset := bits.TrailingZeros64(c); ans = append(ans, base+offset); c ^= 1 << uint(offset) }; base += 64
	}
	return ans
}
func (q *Bitset) shrink() { for i := len(q.c) - 1; i >= 0 && q.c[i] == 0; i-- { q.c = q.c[:i]; q.m -= 64 } }
func (q *Bitset) shrinkTo(a int) { i := len(q.c) - 1; for q.m-64 > a { q.c = q.c[:i]; q.m -= 64 } }
func (q *Bitset) max() int {
	lc := len(q.c); if q.c[lc-1] == 0 { q.shrink(); lc = len(q.c) }; if lc == 0 { return -1 }
	return 64*lc - 1 - bits.LeadingZeros64(q.c[lc-1])
}
func BitsetAnd(a, b *Bitset) *Bitset { c := a.Copy(); c.And(b); return c }
func BitsetOr(a, b *Bitset) *Bitset { c := a.Copy(); c.Or(b); return c }
func BitsetXor(a, b *Bitset) *Bitset { c := a.Copy(); c.Xor(b); return c }
func BitsetShl(a *Bitset, n int) *Bitset { c := a.Copy(); c.Shl(n); return c }
func BitsetShr(a *Bitset, n int) *Bitset { c := a.Copy(); c.Shr(n); return c }

type minheap struct { buf []query; less func(query, query) bool }
func Newminheap(f func(query, query) bool) *minheap { buf := make([]query, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v query) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() query { return q.buf[0] }
func (q *minheap) Pop() query {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []query) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

type lazysegtree struct {
	n, size, log int; op func(int, int) int; mapping func(int, int) int
	composition func(int, int) int; e int; id int; d []int; lz []int;
	// New for this problem
	qq []stsearch; res []int 
}
func Newlazysegtree(n int, op func(int, int) int, mapping func(int, int) int, composition func(int, int) int, e int, id int) *lazysegtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []int, op func(int, int) int, mapping func(int, int) int, composition func(int, int) int, e int, id int) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz)
	lz := make([]int, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	qq := make([]stsearch,0); res := make([]int,0)
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz, qq, res}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) int { p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p] }
func (q *lazysegtree) Prod(l int, r int) int {
	if r < l { return q.e }; l += q.size; r += q.size; r += 1 
	for i := q.log; i >= 1; i-- {
		if ((l >> uint(i)) << uint(i)) != l { q.push(l >> uint(i)) }; if ((r >> uint(i)) << uint(i)) != r { q.push((r - 1) >> uint(i)) }
	}
	sml, smr := q.e, q.e
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *lazysegtree) Allprod() int { return q.d[1] }
func (q *lazysegtree) Apply(p int, f int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f int) {
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := q.log; i >= 1; i-- {
		if ((l >> uint(i)) << uint(i)) != l { q.push(l >> uint(i)) }; if ((r >> uint(i)) << uint(i)) != r { q.push((r - 1) >> uint(i)) }
	}
	l2, r2 := l, r
	for l < r { if l&1 != 0 { q.allApply(l, f); l += 1 }; if r&1 != 0 { r -= 1; q.allApply(r, f) }; l >>= 1; r >>= 1 }
	l, r = l2, r2
	for i := 1; i <= q.log; i++ {
		if ((l >> uint(i)) << uint(i)) != l { q.update(l >> uint(i)) }; if ((r >> uint(i)) << uint(i)) != r { q.update((r - 1) >> uint(i)) }
	}
}
func (q *lazysegtree) MaxRight(l int, f func(int) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; for i := q.log; i >= 1; i-- { q.push(l >> uint(i)) }; sm := q.e
	for {
		for l%2 == 0 { l >>= 1 }
		if !f(q.op(sm, q.d[l])) {
			for l < q.size { q.push(l); l *= 2; if f(q.op(sm, q.d[l])) { sm = q.op(sm, q.d[l]); l++ } }
			return l - q.size - 1
		}
		sm = q.op(sm, q.d[l]); l++; if l&-l == l { break }
	}
	return q.n - 1
}
func (q *lazysegtree) MinLeft(r int, f func(int) bool) int {
	if r < 0 { return 0 }; r += q.size; r++; for i := q.log; i >= 1; i-- { q.push((r - 1) >> uint(i)) }; sm := q.e 
	for {
		r--; for r > 1 && r%2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r], sm)) {
			for r < q.size { q.push(r); r = 2*r + 1; if f(q.op(q.d[r], sm)) { sm = q.op(q.d[r], sm); r-- } }
			return r + 1 - q.size
		}
		sm = q.op(q.d[r], sm); if r&-r == r { break }
	}
	return 0
}
func (q *lazysegtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *lazysegtree) allApply(k int, f int) {
	q.d[k] = q.mapping(f, q.d[k]); if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) }
}
func (q *lazysegtree) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }

// Custom code for this problem
func (q *lazysegtree) findones(l,r int) []int {
	q.qq = q.qq[:0]; q.res = q.res[:0]
	q.qq = append(q.qq,stsearch{1,0,q.size-1})
	ptr := 0
	for ptr < len(q.qq) {
		xx := q.qq[ptr]; ptr++
		nid,ll,rr := xx.id,xx.l,xx.r
		if rr < l || r < ll { continue }
		if q.d[nid] != 1 { continue }
		if ll == rr { 
			q.res = append(q.res,ll)
		} else {
			if q.lz[nid] != q.id { q.push(nid) }
			m := (ll+rr)>>1
			q.qq = append(q.qq,stsearch{2*nid,ll,m})
			q.qq = append(q.qq,stsearch{2*nid+1,m+1,rr})
		}
	}
	return q.res
}

func solve(N,Q int, L,R []int) int {
	cmap := make(map[int]bool)
	for _,l := range L { cmap[l] = true }
	for _,r := range R { cmap[r+1] = true }
	C := make([]int,0,len(cmap))
	for c := range cmap { C = append(C,c) }
	sort.Slice(C,func(i,j int) bool { return C[i] < C[j] })
	cold2cnew := make(map[int]int)
	for i,c := range C { cold2cnew[c] = i }

	// Create the bitmasks and the covercount for each interval with sweepline
	events := make([]event,0)
	for i:=0;i<Q;i++ { 
		id1 := cold2cnew[L[i]]
		id2 := cold2cnew[R[i]+1]
		events = append(events,event{id1,i,1} )
		events = append(events,event{id2,i,-1} )
	}
	sort.Slice(events,func(i,j int) bool { return events[i].id < events[j].id} )
	bm := NewBitset(1+Q>>6); ptr := 0; n := 0; bms := make([]*Bitset,0); sb := ia(Q); ticketIntervals := make([][]int,Q)
	v := ia(len(C)-1)
	for i:=0;i<len(C)-1;i++ {
		ll := C[i+1]-C[i]
		for events[ptr].id == i {
			if events[ptr].typ == 1 { n++; bm.Ins(events[ptr].qid) }
			if events[ptr].typ == -1 { n--; bm.Del(events[ptr].qid) }
			ptr++
		}
		if n == 0 {
			v[i] = inf
		} else if n == 1 {
			id := bm.GetBits()[0]; sb[id] += ll; ticketIntervals[id] = append(ticketIntervals[id],i); v[i] = n
		} else {
			v[i] = n
		}
		bms = append(bms,bm.Copy())
	}

	// Now we set up the maxheap
	// Create a maxheap with a scoreboard.
	mh := Newminheap(func (a,b query) bool { return a.amt > b.amt })
	for i:=0;i<Q;i++ {
		mh.Push(query{i,sb[i]})
	}

	// Create the segtree for the minimum cover of a 
	lst := NewlazysegtreeVec(v,func(a,b int) int { if a < b { return a }; return b }, func(a,b int) int { return a+b }, func(a,b int) int { return a+b }, inf, 0)

	best := inf
	bsleft := NewBitset(1+Q>>6)
	for i:=0;i<Q;i++ { bsleft.Ins(i) }
	for !mh.IsEmpty() {
		q := mh.Pop()
		if sb[q.id] == -1 { continue }
		best = min(best,sb[q.id])
		sb[q.id] = -1
		bsleft.Del(q.id)
		if best == 0 { break }
		l,r := cold2cnew[L[q.id]],cold2cnew[R[q.id]+1]-1
		lst.ApplyRange(l,r,-1)
		for _,i := range ticketIntervals[q.id] { lst.Set(i,inf) }

		// Now for the unique code to find the new leaf ones in the seg tree efficiently
		// The search look inefficient, but we do a maximum Qlog(Q) + (2Q)*log(2Q) nodes in total, which should fit.
		newSingletons := lst.findones(l,r)
		for _,id := range newSingletons {
			ll := C[id+1]-C[id]
			qid := BitsetAnd(bms[id],bsleft).GetBits()[0]
			ticketIntervals[qid] = append(ticketIntervals[qid],id)
			sb[qid] += ll
			mh.Push(query{qid,sb[qid]})
		}
	}
	return best
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2()
		L,R := fill2(Q)
		ans := solve(N,Q,L,R)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}
