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

func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

type nodeparent struct { n,p int}

type queue struct {	buf []nodeparent; head,tail,sz,bm,l int }
func Newqueue() *queue { buf := make([]nodeparent,8); return &queue{buf,0,0,8,7,0} }
func (q *queue) IsEmpty() bool { return q.l == 0 }
func (q *queue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *queue) Len() int { return q.l }
func (q *queue) Push(x nodeparent) {
	if q.l == q.sz { q.sizeup() }
	if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *queue) Pop() nodeparent {
	if q.l == 0 { panic("Empty queue Pop()") }
	v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }
	return v
}
func (q *queue) Head() nodeparent {if q.l == 0 { panic("Empty queue Head()") }; return q.buf[q.head] }
func (q *queue) Tail() nodeparent {if q.l == 0 { panic("Empty queue Tail()") }; return q.buf[q.tail] }
func (q *queue) sizeup() {
	buf := make([]nodeparent, 2*q.sz)
	for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }
	q.buf = buf; q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi(),gi()
	A,B := fill2(M)
	gr := make([][]int,N)
	for i:=0;i<M;i++ {
		a,b := A[i]-1,B[i]-1 
		gr[a] = append(gr[a],b)
		gr[b] = append(gr[b],a)
	}
	par := iai(N,-1)
	q := Newqueue()
	q.Push(nodeparent{0,0})
	for !q.IsEmpty() {
		xx := q.Pop(); n,p := xx.n,xx.p
		if par[n] >= 0 { continue }
		par[n] = p
		for _,c := range gr[n] {
			if c == p { continue }
			q.Push(nodeparent{c,n})
		}
	}
	good := true
	for i:=0;i<N;i++ { if par[i] == -1 { good = false; break} }
	if !good {
		fmt.Fprintln(wrtr,"No")
	} else {
		fmt.Fprintln(wrtr,"Yes")
		for _,p := range par[1:] { fmt.Fprintln(wrtr,p+1) }
	}
}




