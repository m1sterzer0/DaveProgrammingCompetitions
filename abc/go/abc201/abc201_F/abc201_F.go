package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func min(a,b int) int { if a > b { return b }; return a }
type segtree struct { n, size, log int; op func(int, int) int; e int; d []int }
func Newsegtree(n int, op func(int, int) int, e int) *segtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }; return NewsegtreeVec(v, op, e)
}
func NewsegtreeVec(v []int, op func(int, int) int, e int) *segtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz); d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i] }; st := &segtree{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *segtree) Set(p int, v int) { p += q.size; q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> i) } }
func (q *segtree) Prod(l int, r int) int {
	if r < l { return q.e }; r += 1; sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *segtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); P := gis(N); A,B,C := fill3(N)
	pos := iai(N,0); for i:=0;i<N;i++ { pos[P[i]-1] = i }
	lfcost := iai(N,0); for i:=0;i<N;i++ { lfcost[i] = min(A[i],B[i]) }
	rtcost := iai(N,0); for i:=0;i<N;i++ { rtcost[i] = min(A[i],C[i]) }
	cumlf := iai(N,0); s := 0; for i:=0;i<N;i++ { s += lfcost[i]; cumlf[i] = s }
	cumrt := iai(N,0); s = 0; for i:=0;i<N;i++ { s += rtcost[i]; cumrt[i] = s }
	cuma  := iai(N,0); s = 0; for i:=0;i<N;i++ { s += A[i]; cuma[i] = s }
	inf := 1_000_000_000_000_000_000
	st := Newsegtree(N,min,inf)
	dp := iai(N,inf)    // dp[i] = cost of sorting all elements <= i
	for i:=0;i<N;i++ {  // Assume i is the rightmost fixed element
		v1 := 0; if i > 0 { v1 = cumlf[i-1] }  // This is the case where i is the only fixed element
		v2 := inf; if i > 0 && pos[i] > 0 { v2 = st.Prod(0,pos[i]-1) + cuma[i-1] }  // This is the case where there at least one more fixed element to the left of us
		dp[i] = min(v1,v2)
		st.Set(pos[i],dp[i]-cuma[i])
	}
	best := inf
	for i:=0;i<N;i++ { cand := dp[i] + cumrt[N-1] - cumrt[i]; best = min(best,cand) }
	fmt.Println(best)
}



