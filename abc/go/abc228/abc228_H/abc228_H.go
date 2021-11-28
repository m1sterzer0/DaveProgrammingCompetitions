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
// Linear version rules
// a) Query points are given in strictly increasing order
// b) Slopes are negative and monotonically non-increasing
// c) Line is added before first query
type chtline struct { m,b int }
type chtdeque struct { buf []chtline; head, tail, sz, bm, l int }
func Newchtdeque() *chtdeque { buf := make([]chtline, 8); return &chtdeque{buf, 0, 0, 8, 7, 0} }
func (q *chtdeque) IsEmpty() bool { return q.l == 0 }
func (q *chtdeque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *chtdeque) PushFront(x chtline) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *chtdeque) PushBack(x chtline) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *chtdeque) PopFront() chtline {
	if q.l == 0 { panic("Empty chtdeque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *chtdeque) PopBack() chtline {
	if q.l == 0 { panic("Empty chtdeque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *chtdeque) Len() int { return q.l }
func (q *chtdeque) Head() chtline { if q.l == 0 { panic("Empty chtdeque Head()") }; return q.buf[q.head] }
func (q *chtdeque) Tail() chtline { if q.l == 0 { panic("Empty chtdeque Tail()") }; return q.buf[q.tail] }
func (q *chtdeque) sizeup() {
	buf := make([]chtline, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
type chtmininc struct { dq *chtdeque }
func Newchtmininc() *chtmininc { dq := Newchtdeque(); return &chtmininc{dq} }
func (q *chtmininc) check(l1,l2,l3 chtline) bool { 
	// Let x0 be intersection of l1 and l3
	// check return true if y value of l2 at x0 is < y value at x0 of l1/l3
	// Let l1 = m1 * x + b1, l2 = m2 * x + b2, l3 = m3 * x + b3
	// Then l1 and l3 intersect at x0 = (b3-b1)/(m1-m3).  Note that m1-m3 is guaranteed to be positive.
	// The y value of l1 at x0 is (b3-b1)/(m1-m3)*m1 + b1.  Similarly, the y value of l2 at x0 is (b3-b1)/(m1-m3)*m2+b2
	// Then we return (b3-b1)/(m1-m3)*m2 + b2 < (b3-b1)/(m1-m3)*m1 + b1
	//    iff (b3-b1)/(m1-m3)*(m2-m1) < b1-b2
	//    iff (b3-b1) * (m2-m1) < (b1-b2) * (m1-m3)
	//    iff (b1-b3) * (m1-m2) < (b1-b2) * (m1-m3)
	if l1.m <= l2.m || l2.m <= l3.m { fmt.Println("ERROR: slopes not strictly decreasing in check"); os.Exit(1) }
	return (l1.b-l3.b) * (l1.m-l2.m) < (l1.b-l2.b) * (l1.m-l3.m) // Need to watch for overflow
}

func (q *chtmininc) Add(m,b int) {
	if q.dq.Len() == 0 { q.dq.PushBack(chtline{m,b}); return }
	for !q.dq.IsEmpty() && q.dq.Tail().m == m {
		if b >= q.dq.Tail().b { return }
		q.dq.PopBack()
	}
	for q.dq.Len() >= 2 {
		l1 := q.dq.PopBack()
		if q.check(q.dq.Tail(),l1,chtline{m,b}) { q.dq.PushBack(l1); break }
	}
	q.dq.PushBack(chtline{m,b})
}
func (q *chtmininc) Query(x int) int {
	l := q.dq.PopFront(); v1 := l.m*x+l.b
	for !q.dq.IsEmpty() {
		l2 := q.dq.Head(); v2 := l2.m*x+l2.b
		if v2 > v1 { break }
		l,v1 = l2,v2; q.dq.PopFront()
	}
	q.dq.PushFront(l); return v1
}
type pair struct {a,c int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Answer transcription from maspy -- need to code cht
	N,X := gi2(); A,C := fill2(N)
	ac := make([]pair,N); for i:=0;i<N;i++ { ac[i] = pair{A[i],C[i]} }
	sort.Slice(ac,func (i,j int) bool { return ac[i].a < ac[j].a } )
	for i:=0;i<N;i++ { A[i] = ac[i].a; C[i] = ac[i].c }
	D := ia(N); for i:=0;i<N;i++ { D[i] = A[i]*C[i] }
	cumC := ia(N+1); cumC[0] = 0; for i:=0;i<N;i++ { cumC[i+1] = cumC[i] + C[i] }
	cumD := ia(N+1); cumD[0] = 0; for i:=0;i<N;i++ { cumD[i+1] = cumD[i] + D[i] }
	cht := Newchtmininc(); cht.Add(0,0); ans := 0
	for i:=1;i<=N;i++ {
		if i == N {
			y := cht.Query(A[i-1]) + X
			ans = y + A[i-1] * cumC[i] - cumD[i]
		} else {
			y := cht.Query(A[i-1]) + X
			dp := y + A[i-1] * cumC[i] - cumD[i]
			cht.Add(-cumC[i],cumD[i]+dp)
		}
	}
	fmt.Println(ans)
}

