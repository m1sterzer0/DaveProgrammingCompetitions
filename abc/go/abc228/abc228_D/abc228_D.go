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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }

type segtree struct { n, size, log int; op func(int, int) int; e int; d []int }
func Newsegtree(n int, op func(int, int) int, e int) *segtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }; return NewsegtreeVec(v, op, e)
}
func NewsegtreeVec(v []int, op func(int, int) int, e int) *segtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz); d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i] }; st := &segtree{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *segtree) Set(p int, v int) {
	p += q.size; q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *segtree) Get(p int) int { return q.d[p+q.size] }
func (q *segtree) Prod(l int, r int) int {
	if r < l { return q.e }; r += 1; sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *segtree) Allprod() int { return q.d[1] }
func (q *segtree) MaxRight(l int, f func(int) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; sm := q.e
	for {
		for l%2 == 0 { l >>= 1 }
		if !f(q.op(sm, q.d[l])) {
			for l < q.size { l *= 2; if f(q.op(sm, q.d[l])) { sm = q.op(sm, q.d[l]); l++ } }; return l - q.size - 1
		}
		sm = q.op(sm, q.d[l]); l++; if l&-l == l { break }
	}
	return q.n - 1
}
func (q *segtree) MinLeft(r int, f func(int) bool) int {
	if r < 0 { return 0 }; r += q.size; sm := q.e; r++ 
	for {
		r--; for r > 1 && r%2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r], sm)) {
			for r < q.size { r = 2*r + 1; if f(q.op(q.d[r], sm)) { sm = q.op(q.d[r], sm); r-- } }; return r + 1 - q.size
		}
		sm = q.op(q.d[r], sm); if r&-r == r { break }
	}
	return 0
}
func (q *segtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := powint(2,20)
	Q := gi(); T,X := fill2(Q)
	a := iai(N,-1)
	a2 := iai(N,1)
	sumi := func(a,b int) int { return a+b }
	s := NewsegtreeVec(a2,sumi,0)
	for i:=0;i<Q;i++ {
		x := X[i] % N
		if T[i] == 2 { 
			fmt.Fprintln(wrtr,a[x])
		} else {
			if s.Get(x) == 1 {
				a[x] = X[i]
				s.Set(x,0)
			} else if s.Prod(x,N-1) > 0 {
				l,u := x,N-1
				for u-l > 1 {
					m := (u+l)/2
					if s.Prod(x,m) > 0 { u = m } else { l = m }
				}
				a[u] = X[i]
				s.Set(u,0)
			} else if s.Get(0) == 1 {
				a[0] = X[i]
				s.Set(0,0)
			} else {
				l,u := 0,x-1
				for u-l > 1 {
					m := (u+l)/2
					if s.Prod(0,m) > 0 { u = m } else { l = m }
				}
				a[u] = X[i]
				s.Set(u,0)
			}
		}
	}
}

