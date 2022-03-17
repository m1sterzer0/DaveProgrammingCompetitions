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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y := fill2(N); for i:=0;i<N;i++ { X[i]--; Y[i]-- }

		// Find the edge that completes the cycle
		uf := NewDsu(N); gr := make([][]int,N); x := -1; y := -1
		for i:=0;i<N;i++ {
			xx,yy := X[i],Y[i]
			l1,l2 := uf.Leader(xx),uf.Leader(yy)
			if l1 == l2 { x,y = xx,yy } else { uf.Merge(l1,l2); gr[xx] = append(gr[xx],yy); gr[yy] = append(gr[yy],xx) }
		}

		// Now we have a tree.  Assume x is a root and find y via dfs
		dd := iai(N,-1)
		var dfs func(n,p int) bool
		dfs = func(n,p int) bool {
			if n == y { dd[n] = 0; return true }
			for _,c := range gr[n] {
				if c == p { continue }
				if dfs(c,n) { dd[n] = 0; return true }
			}
			return false
		}
		dfs(x,-1)

		// Finally, do a bfs to find all of the distances
		q := ia(0)
		for i:=0;i<N;i++ { if dd[i] == 0 { q = append(q,i) } }
		for len(q) > 0 {
			n := q[0]; q = q[1:]
			for _,c := range gr[n] {
				if dd[c] >= 0 { continue }
				dd[c] = dd[n]+1
				q = append(q,c)
			}
		}

		ansstr := vecintstring(dd)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

