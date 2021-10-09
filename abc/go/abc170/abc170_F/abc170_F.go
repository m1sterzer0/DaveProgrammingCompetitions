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

func min(a,b int) int { if a > b { return b }; return a }

type dstate struct { t,x,y,dir int}

type minheap struct { buf []dstate; less func(dstate,dstate)bool }
func Newminheap(f func(dstate,dstate)bool) *minheap { buf := make([]dstate, 0,24_000_000); return &minheap{buf,f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v dstate) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() dstate { return q.buf[0] }
func (q *minheap) Pop() dstate {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }
	return v1
}
func (q *minheap) Heapify(pri []dstate) { q.buf=append(q.buf,pri...); n:=len(q.buf); for i:=n/2-1;i>=0;i-- { q.siftup(i) } }
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W,K := gi(),gi(),gi()
	// Inverting y and x, since it is unnatural
	y1,x1,y2,x2 := gi()-1,gi()-1,gi()-1,gi()-1
	C := make([]string,H)
	for i:=0;i<H;i++ { C[i] = gs() }
	dist := make([][][]int,4)
	for dir:=0;dir<4;dir++ {
		a := make([][]int,H)
		for i:=0;i<H;i++ { a[i] = make([]int,W); }
		dist[dir] = a
	}
	myinf := 1_000_000_000_000_000_000
	for dir:=0;dir<4;dir++ {
		for y:=0;y<H;y++ {
			for x:=0;x<W;x++ {
				dist[dir][y][x] = myinf
			}
		}
	}
	mh := Newminheap(func(p,q dstate)bool { return p.t < q.t })
	mh.Push(dstate{0,x1,y1,0})
	mh.Push(dstate{0,x1,y1,1})
	mh.Push(dstate{0,x1,y1,2})
	mh.Push(dstate{0,x1,y1,3})
	for !mh.IsEmpty() {
		xx := mh.Pop(); t,x,y,d := xx.t,xx.x,xx.y,xx.dir
		if dist[d][y][x] < myinf { continue }
		dist[d][y][x] = t
		// Edges in the current direction
		if d == 0 && y != 0 && C[y-1][x] != '@'   { mh.Push(dstate{t+1,x,y-1,d}) }
		if d == 2 && y != H-1 && C[y+1][x] != '@' { mh.Push(dstate{t+1,x,y+1,d}) }
		if d == 1 && x != W-1 && C[y][x+1] != '@' { mh.Push(dstate{t+1,x+1,y,d}) }
		if d == 3 && x != 0 && C[y][x-1] != '@'   { mh.Push(dstate{t+1,x-1,y,d}) }
		// Edges involving a 90 degree rotation (trick)
		t2 := (t+K-1)/K*K
		mh.Push(dstate{t2,x,y,(d+1)%4})
		mh.Push(dstate{t2,x,y,(d+3)%4})
	}
	ans := myinf
	for d:=0;d<4;d++ { ans = min(ans,dist[d][y2][x2]) }
	if ans == myinf { fmt.Println(-1) } else { fmt.Println((ans+K-1)/K) }
}






