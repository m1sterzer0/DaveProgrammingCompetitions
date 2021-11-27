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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
type piece struct { idx,r,c,a int }
type queue struct { buf []piece; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]piece, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x piece) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() piece {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() piece { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() piece { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]piece, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W,N := gi3()
	R,C,A := fill3(N); for i:=0;i<N;i++ { R[i]--; C[i]-- }
	besth,bestw := ia(H),ia(W)
	p := make([]piece,N)
	for i:=0;i<N;i++ { p[i] = piece{i,R[i],C[i],A[i]} }
	sort.Slice(p,func(i,j int) bool { return p[i].a > p[j].a } )
	ansarr := ia(N)
	q := Newqueue()
	pidx := 0
	for pidx < N {
		v := p[pidx].a
		for pidx < N && p[pidx].a == v {
			pp := p[pidx]; idx,r,c := pp.idx,pp.r,pp.c
			myv := max(besth[r],bestw[c])
			ansarr[idx] = myv
			q.Push(piece{idx,r,c,myv+1})
			pidx++
		}
		for !q.IsEmpty() {
			pp := q.Pop(); r,c,a := pp.r,pp.c,pp.a
			besth[r] = max(besth[r],a)
			bestw[c] = max(bestw[c],a)
		}
	}
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}

