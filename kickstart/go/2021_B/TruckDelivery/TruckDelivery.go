package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func vecintstring(a []int) string {	astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type segtree struct { n, size, log int; op func(int, int) int; e int; d []int }
func Newsegtree(n int, op func(int, int) int, e int) *segtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }; return NewsegtreeVec(v, op, e)
}
func NewsegtreeVec(v []int, op func(int, int) int, e int) *segtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz); d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i] }; st := &segtree{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *segtree) Set(p int, v int) { p += q.size; q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) } }
func (q *segtree) Get(p int) int { return q.d[p+q.size] }
func (q *segtree) Prod(l int, r int) int {
	if r < l { return q.e }; r += 1; sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *segtree) Allprod() int { return q.d[1] }
func (q *segtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
type edge struct {n2,l,a int}
type query struct { i,w int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	stfunc := func(a,b int) int { if a == 0 { return b }; if b == 0 { return a }; return gcd(a,b) }
	st := Newsegtree(200010,stfunc,0)
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2(); X,Y,L,A := fill4(N-1); C,W := fill2(Q)
		for i:=0;i<N-1;i++ { X[i]--; Y[i]-- }; for i:=0;i<Q;i++ { C[i]-- }
		gr := make([][]edge,N)
		for i:=0;i<N-1;i++ {
			x,y,l,a := X[i],Y[i],L[i],A[i]
			gr[x] = append(gr[x],edge{y,l,a})
			gr[y] = append(gr[y],edge{x,l,a})
		}
		queries := make([][]query,N)
		for i:=0;i<Q;i++ {
			c,w := C[i],W[i]
			queries[c] = append(queries[c],query{i,w})
		}
		ans := ia(Q)
		var dfs func(n,p int)
		dfs = func(n,p int) {
			for _,q := range queries[n] { ans[q.i] = st.Prod(1,q.w) }
			for _,e := range gr[n] {
				if e.n2 == p { continue }
				st.Set(e.l,e.a)
				dfs(e.n2,n)
				st.Set(e.l,0)
			}
		}
		dfs(0,-1)
		ansstr := vecintstring(ans)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}