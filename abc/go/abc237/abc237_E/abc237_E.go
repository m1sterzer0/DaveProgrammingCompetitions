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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
type edge struct {n2,d int}
type minheap struct { buf []edge; less func(edge, edge) bool }
func Newminheap(f func(edge, edge) bool) *minheap { buf := make([]edge, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v edge) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() edge { return q.buf[0] }
func (q *minheap) Pop() edge {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []edge) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

const inf = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); H := gis(N); U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { 
		u,v := U[i],V[i]
		hu,hv := H[u],H[v]
		cuv,cvu := tern(hv<=hu,0,hv-hu),tern(hu<=hv,0,hu-hv)
		gr[u] = append(gr[u],edge{v,cuv})
		gr[v] = append(gr[v],edge{u,cvu})
	}
	darr := make([]int,N); for i:=0;i<N;i++ { darr[i] = inf }; darr[0] = 0
	mh := Newminheap(func (a,b edge) bool { return a.d < b.d })
	for _,e := range gr[0] { mh.Push(e) }
	for !mh.IsEmpty() {
		e := mh.Pop()
		if darr[e.n2] < inf { continue }
		darr[e.n2] = e.d
		for _,ee := range gr[e.n2] { mh.Push(edge{ee.n2,ee.d+e.d}) }
	}
	ans := 0
	for i:=0;i<N;i++ { cand := H[0]-H[i]-darr[i]; if cand > ans { ans = cand } }
	fmt.Println(ans)
}

