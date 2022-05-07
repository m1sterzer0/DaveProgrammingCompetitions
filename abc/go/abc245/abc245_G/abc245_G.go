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
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type dsnode struct { n, l, d int}
type minheap struct { buf []dsnode; less func(dsnode, dsnode) bool }
func Newminheap(f func(dsnode, dsnode) bool) *minheap { buf := make([]dsnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v dsnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() dsnode { return q.buf[0] }
func (q *minheap) Pop() dsnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []dsnode) {
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
type edge struct {n2,d int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Tricky Dijkstra problem
	// Need the concept of "2nd best" to prevent this from becoming an N^2 problem
	N := gi(); M := gi(); gi(); L := gi()
	A := gis(N); B := gis(L); U,V,C := fill3(M)
	for i:=0;i<L;i++ { B[i]-- }; for i:=0;i<M;i++ { U[i]--; V[i]-- }
	d1arr := make([]int,N); for i:=0;i<N;i++ { d1arr[i] = inf }
	d2arr := make([]int,N); for i:=0;i<N;i++ { d2arr[i] = inf }
	carr  := make([]int,N); for i:=0;i<N;i++ { carr[i] = -1 }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { 
		u,v,c := U[i],V[i],C[i]
		gr[u] = append(gr[u],edge{v,c})
		gr[v] = append(gr[v],edge{u,c})
	}
	mh := Newminheap(func(a,b dsnode) bool { return a.d < b.d })
	for _,b := range B { l := A[b]; mh.Push(dsnode{b,l,0}) }
	for !mh.IsEmpty() {
		e := mh.Pop(); n,l,d := e.n,e.l,e.d
		if d2arr[n] < inf || carr[n] == l { continue }
		if d1arr[n] == inf { d1arr[n] = d; carr[n] = l } else { d2arr[n] = d }
		for _,ee := range gr[n] { mh.Push(dsnode{ee.n2,l,ee.d+d}) }
	}
	ansarr := make([]int,N)
	for i:=0;i<N;i++ { 
		d := d1arr[i]; if A[i] == carr[i] { d = d2arr[i] }
		if d == inf { ansarr[i] = -1 } else { ansarr[i] = d }
	}
	ans := vecintstring(ansarr)
	fmt.Println(ans)
}


