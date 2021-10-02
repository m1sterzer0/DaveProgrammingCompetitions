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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
type event struct { t,typ,x,y int }
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
	N,M,Q := gi3(); A,B,S,T := fill4(M); X,Y,Z := fill3(Q)
	for i:=0;i<M;i++ { A[i]--; B[i]--; }; for i:=0;i<Q;i++ { Y[i]-- }
	// 4 types of events
	// 1: add a person with ID x to city y
    // 2: query where is person x
	// 3: unload people from bus N+x to city y
	// 4: load people at city x onto bus N+y
	loc := iai(Q,-1)
	person := iai(N+M,-1)
	events := make([]event,0)
	for i:=0;i<M;i++ {
		events = append(events,event{S[i],4,A[i],i})
		events = append(events,event{T[i],3,i,B[i]})
	} 
	for i:=0;i<Q;i++ {
		events = append(events,event{X[i],1,i,Y[i]})
		events = append(events,event{Z[i],2,i,0})
	} 
	sort.Slice(events,func(i,j int)bool{return events[i].t < events[j].t || events[i].t == events[j].t && events[i].typ < events[j].typ} )
	uf := NewDsu(Q); ans := make([]string,Q)
	for _,ee := range events {
		if ee.typ == 1 {
			pid := ee.x
			if person[ee.y] >= 0 { x2 := person[ee.y]; uf.Merge(x2,ee.x); pid = uf.Leader(x2) }
			loc[pid] = ee.y; person[ee.y] = pid
		} else if ee.typ == 2 {
			pid := uf.Leader(ee.x); l := loc[pid]
			if l < N { ans[ee.x] = strconv.Itoa(l+1) } else { ans[ee.x] = fmt.Sprintf("%v %v",A[l-N]+1,B[l-N]+1) }
		} else if ee.typ == 3 {
			pid := person[N+ee.x]
			if pid == -1 { continue }
			pid2 := person[ee.y]; person[N+ee.x] = -1; person[ee.y] = -1
			if pid2 >= 0 { uf.Merge(pid,pid2); pid = uf.Leader(pid) }
			loc[pid] = ee.y; person[ee.y] = pid
		} else {
			pid := person[ee.x]
			if pid == -1 { continue }
			person[N+ee.y] = pid; person[ee.x] = -1; loc[pid] = N+ee.y
		}
	}
	for _,aa := range ans { fmt.Fprintln(wrtr,aa) }
}



