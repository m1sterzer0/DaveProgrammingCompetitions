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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a > b { return b }; return a }
func abs(a int) int { if a < 0 { return -a }; return a }
const inf int = 2000000000000000000

type minheap struct { buf []int; less func(int, int) bool }
func Newminheap(f func(int, int) bool) *minheap { buf := make([]int, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v int) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() int { return q.buf[0] }
func (q *minheap) Pop() int {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []int) {
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

type pt struct {x,c int}
func solve(K,N int, XX,CC []int) int {
	// Sort the points first
	pts := make([]pt,N)
	for i:=0;i<N;i++ { pts[i] = pt{XX[i],CC[i]} }
	sort.Slice(pts,func(i,j int) bool { return pts[i].x < pts[j].x} )
	C := make([]int,N); X := make([]int,N)
	for i:=0;i<N;i++ { C[i] = pts[i].c; X[i] = pts[i].x }

	nleft,nright := K/2,K-K/2
	minw,maxw := nleft,N-1-nright
	cost := iai(N,inf); for i:=minw;i<=maxw;i++ { cost[i] = C[i] }

	mh1 := Newminheap(func(a,b int) bool { return a > b }); penalty := 0; cursum := 0

	if nleft > 0 {
		for i:=0;i<minw;i++ {
			cursum += C[i] + X[minw] - X[i]
			mh1.Push(C[i] + X[minw] - X[i])
		}
		cost[minw] += cursum
		for i:=minw+1;i<=maxw;i++ {
			penalty += X[i]-X[i-1]
			cand := C[i-1] + X[i]-X[i-1] - penalty
			if cand < mh1.Head() { cursum -= mh1.Head(); mh1.Pop(); cursum += cand; mh1.Push(cand) }
			cost[i] += cursum + nleft*penalty
		}
	}

	mh1.Clear(); penalty = 0; cursum = 0

	if nright > 0 { 
		for i:=N-1;i>maxw;i-- {
			cursum += C[i] + X[i] - X[maxw]
			mh1.Push(C[i] + X[i] - X[maxw])
		}
		cost[maxw] += cursum
		for i:=maxw-1;i>=minw;i-- {
			penalty += X[i+1]-X[i]
			cand := C[i+1] + X[i+1]-X[i] - penalty
			if cand < mh1.Head() { cursum -= mh1.Head(); mh1.Pop(); cursum += cand; mh1.Push(cand) }
			cost[i] += cursum + nright*penalty
		}
	}

	best := inf
	for i:=minw;i<=maxw;i++ {
		best = min(best,cost[i])
	}
	return best
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
		K,N := gi2(); X := gis(N); C := gis(N)
		ans := solve(K,N,X,C)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

