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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type deque struct { buf []int; head, tail, sz, bm, l int }
func Newdeque() *deque { buf := make([]int, 8); return &deque{buf, 0, 0, 8, 7, 0} }
func (q *deque) IsEmpty() bool { return q.l == 0 }
func (q *deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *deque) PushFront(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *deque) PushBack(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *deque) PopFront() int {
	if q.l == 0 { panic("Empty deque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) PopBack() int {
	if q.l == 0 { panic("Empty deque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) Len() int { return q.l }
func (q *deque) Head() int { if q.l == 0 { panic("Empty deque Head()") }; return q.buf[q.head] }
func (q *deque) Tail() int { if q.l == 0 { panic("Empty deque Tail()") }; return q.buf[q.tail] }
func (q *deque) sizeup() {
	buf := make([]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Solution transcription
	N,K := gi2(); A := gis(K); A = append([]int{0},A...)
	nxt := iai(K+1,K+1); zan := ia(N+1)
	for i:=0;i<K;i++ { nxt[zan[A[i+1]]] = i+1; zan[A[i+1]] = i+1 }
	mn,mx := iai(K+1,1),iai(K+1,K)
	for i:=0;i<K;i++ { mn[i+1] = max(mn[i],min(K,nxt[i+1])) }
	for i:=K-1;i>=0;i-- { mx[i] = min(mx[i+1],nxt[i+1]-1) }
	ggo := iai(K+1,-1)
	deq := Newdeque(); deq.PushBack(K)
	for i:=K-1;i>=0;i-- {
		for j:=mn[i+1]-1;j>=mn[i];j-- {	if ggo[j] != -1 { deq.PushFront(j) } }
		for !deq.IsEmpty() && deq.Tail() > mx[i] { deq.PopBack() }
		if !deq.IsEmpty() { ggo[i] = deq.Tail() }
	}
	if ggo[0] == -1 {
		fmt.Println(-1)
	} else {
		val := iai(N+1,K+1)
		zan,t := 0,1
		for zan < K {
			for i:=zan+1;i<=ggo[zan];i++ { val[A[i]] = min(val[A[i]],t)}
			zan = ggo[zan]; t++
		}
		ansarr := []int{}
		for i:=0;i<N;i++ { ansarr = append(ansarr,min(val[i+1],t-1)) }
		ans := vecintstring(ansarr)
		fmt.Println(ans)
	}
} 
