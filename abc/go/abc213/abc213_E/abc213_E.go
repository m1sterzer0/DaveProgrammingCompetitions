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
func gi2() (int,int) { return gi(),gi() }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func abs(a int) int { if a < 0 { return -a }; return a }
type xxx struct { x,y,d int }
type deque struct { buf []xxx; head, tail, sz, bm, l int }
func Newdeque() *deque { buf := make([]xxx, 8); return &deque{buf, 0, 0, 8, 7, 0} }
func (q *deque) IsEmpty() bool { return q.l == 0 }
func (q *deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *deque) PushFront(x xxx) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *deque) PushBack(x xxx) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *deque) PopFront() xxx {
	if q.l == 0 { panic("Empty deque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) PopBack() xxx {
	if q.l == 0 { panic("Empty deque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) Len() int { return q.l }
func (q *deque) Head() xxx { if q.l == 0 { panic("Empty deque Head()") }; return q.buf[q.head] }
func (q *deque) Tail() xxx { if q.l == 0 { panic("Empty deque Tail()") }; return q.buf[q.tail] }
func (q *deque) sizeup() {
	buf := make([]xxx, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W := gi2(); gr := make([]string,H); for i:=0;i<H;i++ { gr[i] = gs() } 
	inf := 1_000_000_000_000_000_000
	dist := twodi(H,W,inf); deq := Newdeque(); deq.PushBack(xxx{0,0,0})
	for !deq.IsEmpty() {
		xx := deq.PopFront(); x,y,d := xx.x,xx.y,xx.d
		if dist[x][y] < inf { continue }
		dist[x][y] = d
		for dx:=-2;dx<=2;dx++ {
			for dy:=-2;dy<=2;dy++ {
				mh := abs(dx)+abs(dy)
				if mh == 0 || mh == 4 { continue }
				if x+dx < 0 || x+dx >= H || y+dy < 0 || y+dy >= W { continue }
				if mh == 1 && gr[x+dx][y+dy] == '.' { deq.PushFront(xxx{x+dx,y+dy,d}) }
				deq.PushBack(xxx{x+dx,y+dy,d+1})
			}
		}
	}
	fmt.Println(dist[H-1][W-1])
}


