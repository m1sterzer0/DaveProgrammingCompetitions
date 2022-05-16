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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
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

type edge struct { n1,n2 int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)

	// PROGRAM STARTS HERE
	N,M := gi(),gi(); U,V := fill2(M)
	gr := make([][]int,N+1)
	for i:=0;i<M;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	// Lets make T1
	uf1 := NewDsu(N+1)
	T1 := make([]edge,0,N-1)
	var dfs func(n,p int)
	dfs = func(n,p int) {
		for _,c := range gr[n] {
			if uf1.Same(c,n) { continue }
			T1 = append(T1,edge{n,c})
			uf1.Merge(c,n)
			dfs(c,n)
		}
	}
	dfs(1,-1)
	uf2 := NewDsu(N+1)
	T2 := make([]edge,0,N-1)
	q := make([]int,0,N+1); q = append(q,1)
	for len(q) > 0 {
		n := q[0]; q = q[1:]
		for _,c := range gr[n] {
			if uf2.Same(c,n) { continue }
			T2 = append(T2,edge{n,c})
			uf2.Merge(c,n)
			q = append(q,c)
		}
	}
	for _,x := range T1 { fmt.Fprintf(wrtr,"%v %v\n",x.n1,x.n2)}
	for _,x := range T2 { fmt.Fprintf(wrtr,"%v %v\n",x.n1,x.n2)}
}

