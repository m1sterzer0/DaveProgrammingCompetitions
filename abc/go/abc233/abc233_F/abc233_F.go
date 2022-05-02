package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}
type edge struct { idx,n2 int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); P := gis(N); for i:=0;i<N;i++ { P[i]-- }
	M := gi(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	// Use Union-Find to see if i and P[i] are connected
	uf := NewDsu(N); for i:=0;i<M;i++ { a,b := A[i],B[i]; uf.Merge(a,b) }
	good := true
	for i:=0;i<N;i++ { if !uf.Same(i,P[i]) { good = false; break } }
	if !good { 
		fmt.Fprintln(wrtr,-1)
	} else {
		gr := make([][]edge,N)
		for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],edge{i,b}); gr[b] = append(gr[b],edge{i,a}) }
		// Now we need to order the nodes and reduce the graph to O(N) with a spanning tree
		gr2 := make([][]edge,N)
		order := make([]int,0)
		visited := make([]bool,N)
		var dfs1 func(n int)
		dfs1 = func(n int) {
			if visited[n] { return }
			visited[n] = true
			for _,e := range gr[n] {
				if visited[e.n2] { continue }
				gr2[n] = append(gr2[n],edge{e.idx,e.n2})
				gr2[e.n2] = append(gr2[e.n2],edge{e.idx,n}) 
				dfs1(e.n2)
			}
			order = append(order,n)
		}
		for i:=0;i<N;i++ { dfs1(i) }
		res := make([]int,0)
		var dfs2 func(n,targ int) bool
		dfs2 = func(n,targ int) bool {
			if visited[n] { return false }
			visited[n] = true
			if P[n] == targ { return true }
			for _,e := range gr2[n] { if dfs2(e.n2,targ) { P[n],P[e.n2] = P[e.n2],P[n]; res = append(res,e.idx+1) } }
			if P[n] == targ { return true }
			return false
		}
		for _,n := range order {
			for i:=0;i<N;i++ { visited[i] = false }
			dfs2(n,n)
		}
		fmt.Fprintln(wrtr,len(res))
		ansstr := vecintstring(res)
		fmt.Fprintln(wrtr,ansstr)
	}
}

