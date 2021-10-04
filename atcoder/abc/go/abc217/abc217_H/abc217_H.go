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
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
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
// Slope trick problem
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); T,D,X := fill3(N)
	iless := func(a,b int) bool { return a < b }
	igreater := func(a,b int) bool { return a > b }
	tmp := ia(N+5)
	mhleft := Newminheap(igreater); mhleft.Heapify(tmp)
	mhright := Newminheap(iless); mhright.Heapify(tmp)
	minval,addL,addR,mytime := 0,0,0,0
	pushleft := func(x int) { mhleft.Push(x-addL) }
	pushright := func(x int) { mhright.Push(x-addR) }
	topleft := func() int { return mhleft.Head()+addL }
	topright := func() int { return mhright.Head()+addR }
	popleft := func() int { v := topleft(); mhleft.Pop(); return v }
	popright := func() int { v := topright(); mhright.Pop(); return v }
	addRightDamage := func(x,minval int) int { 
		minval += max(0,topleft()-x)
		pushleft(x); pushright(popleft())
		return minval
	}
	addLeftDamage := func(x,minval int) int {
		minval += max(0,x-topright())
		pushright(x); pushleft(popright())
		return minval
	}
	for i:=0;i<N;i++ {
		t,d,x := T[i],D[i],X[i]
		addL -= t-mytime; addR += (t-mytime); mytime = t
		if d == 0 { 
			minval = addLeftDamage(x,minval)
		} else { 
			minval = addRightDamage(x,minval) 
		}
	}
	fmt.Println(minval)
}

