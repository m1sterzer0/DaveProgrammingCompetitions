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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func min(a,b int) int { if a > b { return b }; return a }
type lazysegtree struct {
	n, size, log int; op func(int, int) int; mapping func(int, int) int
	composition func(int, int) int; e int; id int; d []int; lz []int
}
func Newlazysegtree(n int, op func(int, int) int, mapping func(int, int) int, composition func(int, int) int, e int, id int) *lazysegtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []int, op func(int, int) int, mapping func(int, int) int, composition func(int, int) int, e int, id int) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz)
	lz := make([]int, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *lazysegtree) Get(p int) int { p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; return q.d[p] }
func (q *lazysegtree) Prod(l int, r int) int {
	if r < l { return q.e }; l += q.size; r += q.size; r += 1 
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l { q.push(l >> i) }; if ((r >> i) << i) != r { q.push((r - 1) >> i) }
	}
	sml, smr := q.e, q.e
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *lazysegtree) Allprod() int { return q.d[1] }
func (q *lazysegtree) Apply(p int, f int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f int) {
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l { q.push(l >> i) }; if ((r >> i) << i) != r { q.push((r - 1) >> i) }
	}
	l2, r2 := l, r
	for l < r { if l&1 != 0 { q.allApply(l, f); l += 1 }; if r&1 != 0 { r -= 1; q.allApply(r, f) }; l >>= 1; r >>= 1 }
	l, r = l2, r2
	for i := 1; i <= q.log; i++ {
		if ((l >> i) << i) != l { q.update(l >> i) }; if ((r >> i) << i) != r { q.update((r - 1) >> i) }
	}
}
func (q *lazysegtree) MaxRight(l int, f func(int) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; for i := q.log; i >= 1; i-- { q.push(l >> i) }; sm := q.e
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
	if r < 0 { return 0 }; r += q.size; r++; for i := q.log; i >= 1; i-- { q.push((r - 1) >> i) }; sm := q.e 
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
func lstsum(a,b int) int { return a + b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi2(); S := gs(); T,L,R := fill3(Q); for i:=0;i<Q;i++ { L[i]--; R[i]-- }
	ansarr := []string{}
	working := []byte(S)
	running := 0
	lst := Newlazysegtree(N,min,lstsum,lstsum,1_000_000_000,0)
	for i,s := range(S) { if s == '(' { running += 1 } else { running -= 1}; lst.Set(i,running) }
	for i:=0;i<Q;i++ {
		t,l,r := T[i],L[i],R[i]
		if t == 1 {
			if working[l] == working[r] { continue }
			amt := 2; if working[l] == '(' { amt = -2 }
			lst.ApplyRange(l,r-1,amt)
			working[l],working[r] = working[r],working[l] 
		} else {
			base := 0; if l > 0 { base = lst.Get(l-1) }
			rval := lst.Get(r)
			mval := lst.Prod(l,r)
			qval := "No"
			if rval == base && mval >= base { qval = "Yes" }
			ansarr = append(ansarr,qval)
		}
	}
	for _,s := range ansarr { fmt.Fprintln(wrtr,s) }
}

