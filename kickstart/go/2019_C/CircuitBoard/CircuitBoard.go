package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func max(a,b int) int { if a > b { return a }; return b }

type deque struct { buf []entry; head, tail, sz, bm, l int }
func Newdeque() *deque { buf := make([]entry, 8); return &deque{buf, 0, 0, 8, 7, 0} }
func (q *deque) IsEmpty() bool { return q.l == 0 }
func (q *deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *deque) PushFront(x entry) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *deque) PushBack(x entry) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *deque) PopFront() entry {
	if q.l == 0 { panic("Empty deque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) PopBack() entry {
	if q.l == 0 { panic("Empty deque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) Len() int { return q.l }
func (q *deque) Head() entry { if q.l == 0 { panic("Empty deque Head()") }; return q.buf[q.head] }
func (q *deque) Tail() entry { if q.l == 0 { panic("Empty deque Tail()") }; return q.buf[q.tail] }
func (q *deque) sizeup() {
	buf := make([]entry, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

type entry struct {pos,val int}

func maxAreaHist(a []int) int {
	left := ia(len(a))
	right := ia(len(a))
	q := Newdeque()
	for i,x := range a {
		for !q.IsEmpty() && q.Tail().val >= x { q.PopBack() }
		if q.IsEmpty() { left[i] = 0 } else { left[i] = q.Tail().pos+1 }
		q.PushBack(entry{i,x})
	}
	q.Clear()
	for i:=len(a)-1;i>=0;i-- {
		x := a[i]
		for !q.IsEmpty() && q.Tail().val >= x { q.PopBack() }
		if q.IsEmpty() { right[i] = len(a)-1 } else { right[i] = q.Tail().pos-1 }
		q.PushBack(entry{i,x})
	}
	best := 0
	for i:=0;i<len(a);i++ {
		w := right[i]-left[i]+1; h := a[i]
		best = max(best,w*h)
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
		R,C,K := gi3()
		V := make([][]int,R)
		for i:=0;i<R;i++ { V[i] = gis(C) }
		mindq,maxdq := Newdeque(),Newdeque()
		// Now we calculate the length of the longest segment within row i starting at i,j and moving right.  We do this with 
		// a deque for the min and a deque for the max
		P := twodi(R,C,0)
		for i:=0;i<R;i++ {
			mindq.Clear(); maxdq.Clear(); ptr := -1
			for j:=0;j<C;j++ {
				if !mindq.IsEmpty() && mindq.Head().pos < j { mindq.PopFront() }
				if !maxdq.IsEmpty() && maxdq.Head().pos < j { maxdq.PopFront() }
				if ptr < j { ptr++; mindq.PushBack(entry{0,V[i][ptr]});  maxdq.PushBack(entry{0,V[i][ptr]}) }
				for ptr+1 < C && V[i][ptr+1] - mindq.Head().val <= K && maxdq.Head().val - V[i][ptr+1] <= K {
					ptr++
					for !mindq.IsEmpty() && V[i][ptr] <= mindq.Tail().val { mindq.PopBack() }
					mindq.PushBack(entry{ptr,V[i][ptr]})
					for !maxdq.IsEmpty() && V[i][ptr] >= maxdq.Tail().val { maxdq.PopBack() }
					maxdq.PushBack(entry{ptr,V[i][ptr]})
				}
				P[i][j] = ptr-j+1
			}
		}
		best := 0; hist := ia(R)
		for j:=0;j<C;j++ {
			for i:=0;i<R;i++ { hist[i] = P[i][j] }
			cand := maxAreaHist(hist)
			best = max(best,cand)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

