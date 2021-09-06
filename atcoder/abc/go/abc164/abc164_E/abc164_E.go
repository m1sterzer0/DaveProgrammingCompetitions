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
func min(a,b int) int { if a > b { return b }; return a }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }

type mystate struct {t,n,c int}
type minheap struct { buf []mystate; less func(mystate,mystate)bool }
func Newminheap(f func(mystate,mystate)bool) *minheap { buf := make([]mystate, 0); return &minheap{buf,f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v mystate) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() mystate { return q.buf[0] }
func (q *minheap) Pop() mystate {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }
	return v1
}
func (q *minheap) Heapify(pri []mystate) { q.buf=append(q.buf,pri...); n:=len(q.buf); for i:=n/2-1;i>=0;i-- { q.siftup(i) } }
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos { ppos:=(pos-1)>>1; p:=q.buf[ppos]; if !q.less(newitem,p) { break } ;q.buf[pos], pos = p, ppos }
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos+1; if rtpos < endpos && !q.less(q.buf[chpos],q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos],pos = q.buf[chpos],chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}
type edge struct {n2,a,b int}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,S := gi(),gi(),gi()
	U,V,A,B := ia(M),ia(M),ia(M),ia(M)
	for i:=0;i<M;i++ { U[i],V[i],A[i],B[i] = gi()-1,gi()-1,gi(),gi() }
	C,D := ia(N),ia(N)
	for i:=0;i<N;i++ { C[i],D[i] = gi(),gi() }
	gr := make([][]edge,N)
	for i:=0;i<M;i++ {
		gr[U[i]] = append(gr[U[i]],edge{V[i],A[i],B[i]})
		gr[V[i]] = append(gr[V[i]],edge{U[i],A[i],B[i]})
	}
	myinf := 1_000_000_000_000_000_000
	maxcoin := ia(N)
	dist := iai(N,myinf)
	moneycap := maxarr(A) * (N-1)
	S = min(S,moneycap)
	mh := Newminheap(func(a,b mystate) bool {return a.t < b.t})
	mh.Push(mystate{0,0,S})
	numtovisit := N
	for !mh.IsEmpty() {
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
			mh.Push(mystate{newt,n,newc})
		}
		for _,ee := range gr[n] {
			if c == moneycap {
				mh.Push(mystate{t+ee.b,ee.n2,moneycap})
			} else if c >= ee.a {
				mh.Push(mystate{t+ee.b,ee.n2,c-ee.a})
			}
		}
	}
	for i:=1;i<N;i++ { fmt.Fprintln(wrtr,dist[i]) }
}

