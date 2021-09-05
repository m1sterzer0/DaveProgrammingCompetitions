package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil { fmt.Println(e.Error()); panic(e) }
		buf = append(buf, l...)
		if !p { break }
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }
func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func gtf() float64 { var a float64; fmt.Fscan(rdr,&a); return a }
func gts() string { var a string; fmt.Fscan(rdr,&a); return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

type Minheap3Element struct { t,n,c int }
func Minheap3Less(a, b Minheap3Element) bool { return a.t < b.t }
type Minheap3 struct { buf []Minheap3Element }
func NewMinheap3() *Minheap3 { buf := make([]Minheap3Element, 0); return &Minheap3{buf} }
func (q *Minheap3) Empty() bool { return len(q.buf) == 0 }
func (q *Minheap3) Clear() { q.buf = q.buf[:0] }
func (q *Minheap3) Len() int { return len(q.buf) }
func (q *Minheap3) Push(v Minheap3Element) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *Minheap3) Head() Minheap3Element {	return q.buf[0] }
func (q *Minheap3) Pop() Minheap3Element {
	v1 := q.buf[0]
	if len(q.buf) == 1 {
		q.buf = q.buf[:0]
	} else {
		l := len(q.buf); q.buf[0] = q.buf[l-1];  q.buf = q.buf[:l-1]; q.siftup(0)
	}
	return v1
}
func (q *Minheap3) Heapify(varr []Minheap3Element) {
	q.buf = append(q.buf, varr...); n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {	q.siftup(i)	}
}
func (q *Minheap3) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		parentpos := (pos - 1) >> 1
		parent := q.buf[parentpos]
		if !Minheap3Less(newitem, parent) {	break }
		q.buf[pos], pos = parent, parentpos
	}
	q.buf[pos] = newitem
}
func (q *Minheap3) siftup(pos int) {
	endpos, startpos, newitem, childpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for childpos < endpos {
		rightpos := childpos + 1
		if rightpos < endpos && !Minheap3Less(q.buf[childpos], q.buf[rightpos]) { childpos = rightpos }
		q.buf[pos], pos = q.buf[childpos], childpos
		childpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

type edge struct { n2,a,b int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil { panic(e) }
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}
	
    // NON-BOILERPLATE STARTS HERE
	N := gti(); M := gti(); S := gti()
	U := make([]int,M)
	V := make([]int,M)
	A := make([]int,M)
	B := make([]int,M)
	for i:=0;i<M;i++ { U[i] = gti()-1; V[i] = gti()-1; A[i] = gti(); B[i] = gti() }
	C := make([]int,N)
	D := make([]int,N)
	for i:=0;i<N;i++ { C[i] = gti(); D[i] = gti() }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { 
		gr[U[i]] = append(gr[U[i]],edge{V[i],A[i],B[i]})
		gr[V[i]] = append(gr[V[i]],edge{U[i],A[i],B[i]})
	}
	myinf := 1_000_000_000_000_000_000
	maxcoin := make([]int,N)
	dist := make([]int,N); for i:=0;i<N;i++ { dist[i] = myinf }
	moneycap := maxarr(A) * (N-1)
	S = min(S,moneycap)
	mh := NewMinheap3()
	mh.Push(Minheap3Element{0,0,S})
	numtovisit := N
	for !mh.Empty() {
		xx := mh.Pop()
		t := xx.t; n := xx.n; c := xx.c
		if dist[n] <= t && maxcoin[n] >= c { continue }
		if dist[n] == myinf { 
			dist[n] = t; numtovisit -= 1
			if numtovisit == 0 { break } 
		}
		if maxcoin[n] < c { maxcoin[n] = c }
		if c < moneycap { 
			newc := min(moneycap,c + C[n]); newt := t + D[n]
			mh.Push(Minheap3Element{newt,n,newc})
		}
		for _,ee := range gr[n] {
			if c == moneycap {
				mh.Push(Minheap3Element{t+ee.b,ee.n2,moneycap})
			} else if c >= ee.a {
				mh.Push(Minheap3Element{t+ee.b,ee.n2,c-ee.a})
			}
		}
	}
	for i:=1;i<N;i++ { fmt.Fprintln(wrtr,dist[i]) }
	wrtr.Flush()
}



