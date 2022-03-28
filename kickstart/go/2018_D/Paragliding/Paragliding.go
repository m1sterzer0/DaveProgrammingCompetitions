package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
type obj struct {idx,x,y int; balloon bool}

type minheap struct { buf []obj; less func(obj, obj) bool }
func Newminheap(f func(obj, obj) bool) *minheap { buf := make([]obj, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v obj) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() obj { return q.buf[0] }
func (q *minheap) Pop() obj {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []obj) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
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
		N,K := gi2()
		p1,p2,A1,B1,C1,M1 := gi(),gi(),gi(),gi(),gi(),gi()
		h1,h2,A2,B2,C2,M2 := gi(),gi(),gi(),gi(),gi(),gi()
		x1,x2,A3,B3,C3,M3 := gi(),gi(),gi(),gi(),gi(),gi()
		y1,y2,A4,B4,C4,M4 := gi(),gi(),gi(),gi(),gi(),gi()
		P := ia(N); H := ia(N); X := ia(K); Y := ia(K)
		P[0] = p1; P[1] = p2; H[0] = h1; H[1] = h2; X[0] = x1; X[1] = x2; Y[0] = y1; Y[1] = y2
		for i:=2;i<N;i++ { P[i] = (A1 * P[i-1] + B1 * P[i-2] + C1) % M1 + 1 }
		for i:=2;i<N;i++ { H[i] = (A2 * H[i-1] + B2 * H[i-2] + C2) % M2 + 1 }
		for i:=2;i<K;i++ { X[i] = (A3 * X[i-1] + B3 * X[i-2] + C3) % M3 + 1 }
		for i:=2;i<K;i++ { Y[i] = (A4 * Y[i-1] + B4 * Y[i-2] + C4) % M4 + 1 }
		objs := make([]obj,0)
		sb := make([]bool,K)
		doit := func(objs []obj) {
			sort.Slice(objs,func(i,j int) bool { return objs[i].x < objs[j].x || objs[i].x == objs[j].x && objs[i].balloon && !objs[j].balloon} )
			h := Newminheap(func(a,b obj) bool { return a.y-a.x < b.y-b.x } )
			for _,oo := range objs {
				if oo.balloon {
					h.Push(oo)
				} else {
					for !h.IsEmpty() && h.Head().y - h.Head().x <= oo.y - oo.x {
						sb[h.Head().idx] = true
						h.Pop()
					}
				}
			}
		}
		for i:=0;i<K;i++ { objs = append(objs,obj{i,X[i],Y[i],true})  }
		for i:=0;i<N;i++ { objs = append(objs,obj{i,P[i],H[i],false}) }
		doit(objs) // Pass with towers on right
		for i:=0;i<N+K;i++ { objs[i].x = 1000000001 - objs[i].x }
		doit(objs) // Pass with towers on left
		ans := 0
		for i:=0;i<K;i++ { if sb[i] { ans++ } }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

