package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func max(a,b int) int { if a > b { return a }; return b }

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
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) int {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtree) Prod(l int, r int) int {
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
func (q *lazysegtree) Allprod() int { return q.d[1] }
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,S := gi2(); A := gis(N); for i:=0;i<N;i++ { A[i]-- }
		hist := make([][]int,100000)
		lstsum := func(a,b int) int { return a + b }
		lst := Newlazysegtree(N,max,lstsum,lstsum,0,0)
		best := 0
		for i,a := range(A) {
			hist[a] = append(hist[a],i)
			if len(hist[a]) <= S { 
				lst.ApplyRange(0,i,1)
			} else if len(hist[a]) == S+1 {
				lst.ApplyRange(0,hist[a][0],-S)
				lst.ApplyRange(hist[a][0]+1,i,1)
			} else {
				n := len(hist[a])
				lst.ApplyRange(hist[a][n-S-2]+1,hist[a][n-S-1],-S)
				lst.ApplyRange(hist[a][n-S-1]+1,i,1)
			}
			best = max(best,lst.Allprod())
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
	}
}

