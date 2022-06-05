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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
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
const inf = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); A,B := fill2(M); Q := gi(); X,K := fill2(Q)
	gr := make([][]int,N+1);
	for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a); }
	darr := iai(N+1,inf); q := Newqueue(); varr := make([]int,0)
	for i:=0;i<Q;i++ {
		x,k := X[i],K[i];
		varr = varr[:0]
		varr = append(varr,x); darr[x] = 0; if 0 < k { q.Push(x) }
		for !q.IsEmpty() {
			xx := q.Pop()
			for _,c := range gr[xx] { 
				if darr[c] != inf { continue }
				darr[c] = darr[xx]+1; varr = append(varr,c); if darr[c] < k { q.Push(c) }
			}
		}
		ans := sumarr(varr)
		for _,v := range varr { darr[v] = inf }
		fmt.Fprintln(wrtr,ans)
	}
}

