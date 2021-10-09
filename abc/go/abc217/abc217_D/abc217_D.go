package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
type query struct {c,x int}
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
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	L,Q := gi2(); C,X := fill2(Q)
	q := make([]query,Q); for i:=0;i<Q;i++ { q[i] = query{C[i],X[i]} }
	sort.Slice(q,func(i,j int) bool { return q[i].x < q[j].x || q[i].x == q[j].x && q[i].c < q[j].c })
	q = append(q,query{1,L})
	mark2segid := make(map[int]int); seglen := ia(Q+2)
	curseg,lastcut := 0,0
	for i,xx := range q {
		c,x := xx.c,xx.x
		if i > 0 && x == q[i-1].x { continue }
		mark2segid[x] = curseg
		if c == 1 { seglen[curseg] = x-lastcut; lastcut = x; curseg++ }
	}
	uf := NewDsu(curseg+10)
	ans := []int{}
	for i:=Q-1;i>=0;i-- {
		c,x := C[i],X[i]
		s := mark2segid[x]
		if c == 2 { ans = append(ans,seglen[uf.Leader(s)]); continue }
		l1,l2 := seglen[uf.Leader(s)],seglen[uf.Leader(s+1)]
		uf.Merge(s,s+1)
		seglen[uf.Leader(s)] = l1+l2
	}
	rev(ans)
	for _,a := range ans { fmt.Fprintln(wrtr,a) }
}

