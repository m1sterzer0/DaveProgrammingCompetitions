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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
type edge struct {s,t int}
type queue struct { buf []int; head, tail, sz, bm, l int }
func Newqueue() *queue { buf := make([]int, 8); return &queue{buf, 0, 0, 8, 7, 0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() int {
	if q.l == 0 { panic("Empty queue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *queue) Head() int { if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() int { if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf
	q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
func findPath(N int, gr [][]int, sbad,tbad int) []int {
	par := iai(N,-1)
	par[0] = 0; q := Newqueue(); q.Push(0)
	for !q.IsEmpty() {
		n := q.Pop()
		for _,n2 := range gr[n] {
			if par[n2] >= 0 || n == sbad && n2 == tbad { continue }
			par[n2] = n; q.Push(n2)
		}
	}
	if par[N-1] == -1 { return []int{} }
	path := ia(0); n := N-1
	for { path = append(path,n); p := par[n]; if n == p { break }; n = p }
	rev(path); return path
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); S,T := fill2(M); for i:=0;i<M;i++ { S[i]--; T[i]-- }
	gr := make([][]int,N)
	for i:=0;i<M;i++ { s,t := S[i],T[i]; gr[s] = append(gr[s],t) }
	p := findPath(N,gr,-1,-1); bestDist := len(p)-1; bestEdges := make(map[edge]bool)
	for i:=0;i<len(p)-1;i++ { bestEdges[edge{p[i],p[i+1]}] = true }
	for i:=0;i<M;i++ {
		s,t := S[i],T[i]
		if !bestEdges[edge{s,t}] { fmt.Fprintln(wrtr,bestDist) } else { p := findPath(N,gr,s,t); fmt.Fprintln(wrtr,len(p)-1) }
	}
}

