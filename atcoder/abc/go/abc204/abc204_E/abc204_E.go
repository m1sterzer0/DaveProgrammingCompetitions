package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type mhnode struct {n2,d int}
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
type edge struct { n,c,d int }
func isqrt(x int) int {
    if x == 0 { return 0 }
    s := int(math.Sqrt(float64(x)))
    s = (s + x/s) >> 1
	if s*s > x { return s-1 } else { return s }
}
func solvemin(t,c,d int) int {
	tapproxopt := isqrt(d)
	if t > tapproxopt + 5 { return t + c + d / (t+1) }
	minv,maxv := max(t,tapproxopt-5),tapproxopt+5
	best := 1_000_000_000_000_000_000
	for i:=minv; i<=maxv; i++ {
		best = min(best,i + c + d / (i+1) )
	}
	return best
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); A,B,C,D := fill4(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { 
		a,b,c,d := A[i],B[i],C[i],D[i]
		gr[a] = append(gr[a],edge{b,c,d})
		gr[b] = append(gr[b],edge{a,c,d})
	}
	mh := Newminheap(func(a,b mhnode)bool{return a.d < b.d})
	mh.Push(mhnode{0,0})
	myinf := 1_000_000_000_000_000_000
	dist := iai(N,myinf)
	for !mh.IsEmpty() {
		xx := mh.Pop()
		if dist[xx.n2] != myinf { continue }
		dist[xx.n2] = xx.d
		for _,c := range gr[xx.n2] {
			dnew := solvemin(xx.d,c.c,c.d)
			mh.Push(mhnode{n2:c.n,d:dnew})
		}
	}
	ans := -1; if dist[N-1] < myinf { ans = dist[N-1] }; fmt.Println(ans)
}



