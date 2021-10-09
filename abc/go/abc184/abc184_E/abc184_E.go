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
type PI struct { x,y int }

type queue struct { buf []PI; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]PI, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x PI) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() PI {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() PI { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() PI { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]PI, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2(); bd := make([]string,H); for i:=0;i<H;i++ { bd[i] = gs() }
	st,en := PI{0,0},PI{0,0}; tele := make([][]PI,26); inf := 1_000_000_000_000_000_000; sb:= twodi(H,W,inf)
	for i:=0;i<H;i++ { 
		for j:=0;j<W;j++ {
			if bd[i][j] == 'S' {
				st = PI{i,j}
			} else if bd[i][j] == 'G' {
				en = PI{i,j}
			} else if bd[i][j] == '#' || bd[i][j] == '.' {
				_ = true
			} else {
				idx := int(bd[i][j])-int('a')
				tele[idx] = append(tele[idx],PI{i,j}) 
			}
		}
	}
	// Regular BFS
	q := Newqueue()
	q.Push(st); sb[st.x][st.y] = 0
	for !q.IsEmpty() {
		xx := q.Pop()
		if xx == en { break }
		d := sb[xx.x][xx.y]
		for _,p := range [4]PI{{xx.x-1,xx.y},{xx.x+1,xx.y},{xx.x,xx.y-1},{xx.x,xx.y+1}} {
			if p.x < 0 || p.x >= H || p.y < 0 || p.y >= W { continue }
			if bd[p.x][p.y] == '#' { continue }
			if sb[p.x][p.y] != inf { continue }
			sb[p.x][p.y] = d+1
			q.Push(p)
		}
		cidx := int(bd[xx.x][xx.y])-int('a')
		if cidx >= 0 && cidx < 26 {
			for _,a := range tele[cidx] {
				if sb[a.x][a.y] != inf { continue }
				sb[a.x][a.y] = d+1
				q.Push(a)
			}
			tele[cidx] = tele[cidx][:0] // Prevent more searches on this teleporter
		}
	}
	ans := -1
	if sb[en.x][en.y] < inf { ans = sb[en.x][en.y] }
	fmt.Println(ans)
}



