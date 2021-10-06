package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
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
type edge struct {n1,n2,c int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		A := make([][]int,N); for i:=0;i<N;i++ { A[i] = gis(N) }
		B := make([][]int,N); for i:=0;i<N;i++ { B[i] = gis(N) }
		gis(N); gis(N)  // Don't care about the values
		// Can't have a cycle, so we want a maximum spanning tree for the spaces we can solve with checksums
		// Nodes 0 to N-1 are rows, and nodes N to 2N-1 are cols
		// Lets see if Kruskal is fast enough
		totcost := 0
		edges := make([]edge,0,N*N)
		for i:=0;i<N;i++ { 
			for j:=0;j<N;j++ { 
				if A[i][j] == -1 { 
					totcost += B[i][j]
					edges = append(edges,edge{i,N+j,B[i][j]}) 
				}
			}
		}
		sort.Slice(edges,func(i,j int) bool { return edges[i].c > edges[j].c })
		uf := NewDsu(2*N)
		savedcost := 0
		for _,e := range edges {
			if uf.Same(e.n1,e.n2) { continue }
			savedcost += e.c
			uf.Merge(e.n1,e.n2)
		}
		ans := totcost - savedcost
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}
