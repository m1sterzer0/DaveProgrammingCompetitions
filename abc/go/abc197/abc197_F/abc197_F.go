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
func gi2() (int,int) { return gi(),gi() }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
type mhnode struct {n1,n2,d int}
type minheap struct { buf []mhnode; less func(mhnode, mhnode) bool }
func Newminheap(f func(mhnode, mhnode) bool) *minheap { buf := make([]mhnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v mhnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() mhnode { return q.buf[0] }
func (q *minheap) Pop() mhnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []mhnode) {
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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); A := make([]int,M); B := make([]int,M); C := make([]byte,M)
	for i:=0;i<M;i++ { xx := gss(3); A[i],_ = strconv.Atoi(xx[0]); B[i],_ = strconv.Atoi(xx[1]); C[i] = xx[2][0] }
	for i:=0;i<M;i++ { A[i]--; B[i]-- }

	// One graph that just shows the paths from one node to another
	gr := make([][]int,N)
	for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }

	// Second graph maps pairs of nodes to pairs of other nodes
	gr2 := make([][]int,N*N)
	for i:=0;i<M;i++ { 
		a1,b1,c1 := A[i],B[i],C[i]
		for j:=0;j<M;j++ {
			if i == j { continue }
			a2,b2,c2 := A[j],B[j],C[j]
			if c1 != c2 { continue }
			gr2[N*a1+a2] = append(gr2[N*a1+a2],N*b1+b2)
			gr2[N*a1+b2] = append(gr2[N*a1+b2],N*b1+a2)
			gr2[N*b1+a2] = append(gr2[N*b1+a2],N*a1+b2)
			gr2[N*b1+b2] = append(gr2[N*b1+b2],N*a1+a2)
		}
	}

    ans := -1; myinf := 1_000_000_000_000_000_000; dist := iai(N*N,myinf); mh := Newminheap(func(a,b mhnode)bool{return a.d < b.d})
	mh.Push(mhnode{0,N-1,0})
	for !mh.IsEmpty() {
		xx := mh.Pop(); if dist[N*xx.n1+xx.n2] != myinf { continue}; dist[N*xx.n1+xx.n2] = xx.d
		if xx.n1 == xx.n2 { ans = xx.d; break }
		// Now we look for any ways to get from n1/n2 to a common node in one edge
		for _,e := range gr[xx.n1] { if e == xx.n2 { mh.Push(mhnode{xx.n2,xx.n2,xx.d+1}) } }
		// Now we look for edge pair transitions
		for _,e := range gr2[N*xx.n1+xx.n2] { n1b,n2b := e/N,e%N; mh.Push(mhnode{n1b,n2b,xx.d+2}) }
	}
	fmt.Println(ans)
}



