package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
type stnode struct { sz,numlt,numeq,numgt int}
type lazysegtree struct {
	n, size, log int; op func(stnode, stnode) stnode; mapping func(int, stnode) stnode
	composition func(int, int) int; e stnode; id int; d []stnode; lz []int
}
func Newlazysegtree(n int, op func(stnode, stnode) stnode, mapping func(int, stnode) stnode, composition func(int, int) int, e stnode, id int) *lazysegtree {
	v := make([]stnode, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []stnode, op func(stnode, stnode) stnode, mapping func(int, stnode) stnode, composition func(int, int) int, e stnode, id int) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]stnode, 2*sz)
	lz := make([]int, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v stnode) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) stnode {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtree) Prod(l int, r int) stnode {
	if r < l { return q.e }; l += q.size; r += q.size; r += 1 
	for i := q.log; i >= 1; i-- {
		if ((l >> uint(i)) << uint(i)) != l { q.push(l >> uint(i)) }
		if ((r >> uint(i)) << uint(i)) != r { q.push((r - 1) >> uint(i)) }
	}
	sml, smr := q.e, q.e
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *lazysegtree) Allprod() stnode { return q.d[1] }
func (q *lazysegtree) Apply(p int, f int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f int) {
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := q.log; i >= 1; i-- {
		if ((l >> uint(i)) << uint(i)) != l { q.push(l >> uint(i)) }
		if ((r >> uint(i)) << uint(i)) != r { q.push((r - 1) >> uint(i)) }
	}
	l2, r2 := l, r
	for l < r { if l&1 != 0 { q.allApply(l, f); l += 1 }; if r&1 != 0 { r -= 1; q.allApply(r, f) }; l >>= 1; r >>= 1 }
	l, r = l2, r2
	for i := 1; i <= q.log; i++ {
		if ((l >> uint(i)) << uint(i)) != l { q.update(l >> uint(i)) }
		if ((r >> uint(i)) << uint(i)) != r { q.update((r - 1) >> uint(i)) }
	}
}
func (q *lazysegtree) MaxRight(l int, f func(stnode) bool) int {
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
func (q *lazysegtree) MinLeft(r int, f func(stnode) bool) int {
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q,X := gi(),gi(),gi(); P := gis(N); C,L,R := fill3(Q)
	// This is the classic game of "follow the shell"
	// We can classify the numbers into three categories
	// Less than target, Equal to targer, Greater than target
	op := func(a,b stnode) stnode { return stnode{a.sz+b.sz,a.numlt+b.numlt,a.numeq+b.numeq,a.numgt+b.numgt} }
	mapping := func(a int, b stnode) stnode { 
		if a == -1 { return stnode{b.sz,b.sz,0,0} } 
		if a == 0  { return stnode{b.sz,0,b.sz,0} } 
		if a == 1  { return stnode{b.sz,0,0,b.sz} }
		return b
	}
	comp := func(a,b int) int { 
		if a == -1 || a == 0 || a == 1 { return a } else { return b }
	}
	e := stnode{0,0,0,0}
	id := -2
	vv := make([]stnode,N+1); vv[0] = e
	for i:=1;i<=N;i++ { 
		v := e; v.sz=1
		if P[i-1] < X { v.numlt = 1 } else if P[i-1] > X { v.numgt = 1} else { v.numeq = 1 }
		vv[i] = v
	}
	lst := NewlazysegtreeVec(vv,op,mapping,comp,e,id)
	for i:=0;i<Q;i++ {
		c,l,r := C[i],L[i],R[i]; w := r-l+1
		p := lst.Prod(l,r)
		if l==r || p.numlt == w || p.numgt == w { continue }
		curs := l
		if c == 1 {
			if p.numlt > 0 { lst.ApplyRange(curs,curs+p.numlt-1,-1); curs += p.numlt }
			if p.numeq > 0 { lst.ApplyRange(curs,curs+p.numeq-1,0); curs += p.numeq }
			if p.numgt > 0 { lst.ApplyRange(curs,curs+p.numgt-1,1); curs += p.numgt }
		} else {
			if p.numgt > 0 { lst.ApplyRange(curs,curs+p.numgt-1,1); curs += p.numgt }
			if p.numeq > 0 { lst.ApplyRange(curs,curs+p.numeq-1,0); curs += p.numeq }
			if p.numlt > 0 { lst.ApplyRange(curs,curs+p.numlt-1,-1); curs += p.numlt }
		}
	}
	ans := 0
	for i:=1;i<=N;i++ { if lst.Get(i).numeq == 1 { ans = i } }
	fmt.Println(ans)
}
