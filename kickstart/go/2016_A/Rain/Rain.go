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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
type sq struct {i,j int}
type queue struct { buf []sq; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]sq, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x sq) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() sq {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() sq { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() sq { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]sq, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
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
		R,C := gi(),gi(); H := make([][]int,R); for i:=0;i<R;i++ { H[i] = gis(C) }
		cc := make([][]bool,R); for i:=0;i<R;i++ { cc[i] = make([]bool,C) }
		cccnt := 0; sumht := 0
		q := Newqueue()
		ds := []sq{{-1,0},{1,0},{0,-1},{0,1}}
		for  {
			// Empty the array
			for i:=0;i<R;i++ { for j:=0;j<C;j++ { cc[i][j] = false } }
			// Reset the count
			cccnt = 0 
			// Stick edge cells on the array
			for i:=0;i<R;i++ { for j:=0;j<C;j++ { if i == 0 || i == R-1 || j == 0 || j == C-1 { cc[i][j] = true; cccnt++; q.Push(sq{i,j})}} }
			// Do bfs
			for !q.IsEmpty() { 
				s := q.Pop()
				for _,d := range ds { 
					ci,cj := s.i+d.i,s.j+d.j
					if ci >= 0 && ci < R && cj >= 0 && cj < C && !cc[ci][cj] && H[ci][cj] >= H[s.i][s.j] {
						cc[ci][cj] = true; cccnt++; q.Push(sq{ci,cj})
					}
				}
			}
			// If cccnt == R*H { break }
			if cccnt == R*C { break }
			// Find the lowest nodes in the unreached sea, and raise them each up by 1
			m := inf
			for i:=0;i<R;i++ { for j:=0;j<C;j++ { if !cc[i][j] { m = min(m,H[i][j]) } } }
			for i:=0;i<R;i++ { for j:=0;j<C;j++ { if !cc[i][j] && H[i][j] == m { sumht++; H[i][j]++ } } } 
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,sumht)
    }
}

