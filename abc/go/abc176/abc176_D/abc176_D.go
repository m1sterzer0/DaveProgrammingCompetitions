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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}

type loc struct {d,i,j int}
type deque struct { buf []loc; head, tail, sz, bm, l int }
func Newdeque() *deque { buf := make([]loc, 8); return &deque{buf, 0, 0, 8, 7, 0} }
func (q *deque) IsEmpty() bool { return q.l == 0 }
func (q *deque) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *deque) PushFront(x loc) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *deque) PushBack(x loc) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *deque) PopFront() loc {
	if q.l == 0 { panic("Empty deque PopFront()") }; v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) PopBack() loc {
	if q.l == 0 { panic("Empty deque PopBack()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *deque) Len() int { return q.l }
func (q *deque) Head() loc { if q.l == 0 { panic("Empty deque Head()") }; return q.buf[q.head] }
func (q *deque) Tail() loc { if q.l == 0 { panic("Empty deque Tail()") }; return q.buf[q.tail] }
func (q *deque) sizeup() {
	buf := make([]loc, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2()
	Ch,Cw := gi2(); Ch--; Cw--
	Dh,Dw := gi2(); Dh--; Dw--
	bd := make([]string,H)
	for i:=0;i<H;i++ { bd[i] = gs() }
	myinf := 1_000_000_000
	d := twodi(H,W,myinf)
	q := Newdeque()
	q.PushFront(loc{0,Ch,Cw})

	for !q.IsEmpty() {
		xx := q.PopFront()
		if d[xx.i][xx.j] != myinf { continue }
		d[xx.i][xx.j] = xx.d
		if xx.i == Dh && xx.j == Dw { break }
		vx,vy := 1,0
		for i:=0;i<4;i++ { 
			x2,y2 := xx.i+vx,xx.j+vy
			if x2 >= 0 && x2 < H && y2 >= 0 && y2 < W && bd[x2][y2] != '#' {q.PushFront(loc{xx.d,x2,y2}) }
			vx,vy = -vy,vx
		}
		for i:=xx.i-2;i<=xx.i+2;i++ {
			for j:=xx.j-2;j<=xx.j+2;j++ {
				if i >= 0 && i < H && j >= 0 && j < W && bd[i][j] != '#' { q.PushBack(loc{xx.d+1,i,j}) }
			}
		}
	}
	ans := d[Dh][Dw]
	if ans == myinf { ans = -1 }
	fmt.Println(ans)
}



