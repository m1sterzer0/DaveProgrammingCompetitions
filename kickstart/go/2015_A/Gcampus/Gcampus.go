package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
const inf int = 2000000000000000000
type edge struct { c,n2,idx int }
type dnode struct { d,n2,idx int }
type minheap struct { buf []dnode; less func(dnode, dnode) bool }
func Newminheap(f func(dnode, dnode) bool) *minheap { buf := make([]dnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v dnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() dnode { return q.buf[0] }
func (q *minheap) Pop() dnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []dnode) {
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
		N,M := gi(),gi(); U,V,C := fill3(M)
		// Just to a repeated modified dijkstra
		gr := make([][]edge,N)
		for i:=0;i<M;i++ {
			u,v,c := U[i],V[i],C[i]
			gr[u] = append(gr[u],edge{c,v,i})
			gr[v] = append(gr[v],edge{c,u,i})
		}
		sb := make([]bool,M)
		darr := make([]int,N)
		h := Newminheap(func (a,b dnode) bool { return a.d < b.d })
		doDijkstra := func(s int) {
			for i:=0;i<N;i++ { darr[i] = inf }
			darr[s] = 0; for _,e := range gr[s] { h.Push(dnode{e.c,e.n2,e.idx}) }
			for !h.IsEmpty() {
				dn := h.Pop()
				if darr[dn.n2] < dn.d { continue }
				sb[dn.idx] = true
				if darr[dn.n2] == inf {
					darr[dn.n2] = dn.d
					for _,e := range gr[dn.n2] { h.Push(dnode{dn.d+e.c,e.n2,e.idx}) }
				}
			}
		}
		for i:=0;i<N;i++ { doDijkstra(i) }
		ansarr := make([]int,0)
		for i:=0;i<M;i++ { if !sb[i] { ansarr = append(ansarr,i) } }
        fmt.Fprintf(wrtr,"Case #%v:\n",tt)
		for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
    }
}

