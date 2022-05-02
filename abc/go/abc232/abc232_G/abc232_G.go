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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
const inf = 2000000000000000000
type edge struct {n2,w int}
type dnode struct {n,d int}

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
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Too many edges for a dijkstra search.
	// Consider mapping to the nodes of a clock
	// we turn i --> j to i --> clock node ---> series of clock jumps --> j
	// i --> clock is (-Ai) % M
	// clock --> j is Bj % M
	// Note that the distance is then Ai+Bj
	// We can prune the graph to only the clock nodes we need which makes it around 3N which is small enough for dijkstra.
	N,M := gi(),gi(); A := gis(N); B := gis(N)
	s := make(map[int]bool)
	for i:=0;i<N;i++ { a,b := (M-A[i])%M,B[i]; s[a] = true; s[b] = true }
	clocknodes := make([]int,0)
	for k := range s { clocknodes = append(clocknodes,k) }
	sort.Slice(clocknodes,func(i,j int) bool { return clocknodes[i] < clocknodes[j] })
	lookup := make(map[int]int)
	for i,v := range clocknodes { lookup[v] = N+i }
	N2 := len(clocknodes)
	gr := make([][]edge,N+N2)
	for i:=0;i<N;i++ {
		a := (M-A[i])%M; n2 := lookup[a]; gr[i] = append(gr[i],edge{n2,0})
		b := B[i];       n3 := lookup[b]; gr[n3] = append(gr[n3],edge{i,0})
	}
	for i:=0;i<N2-1;i++ { gr[N+i] = append(gr[N+i],edge{N+i+1,clocknodes[i+1]-clocknodes[i]}) }
	gr[N+N2-1] = append(gr[N+N2-1],edge{N,M+clocknodes[0]-clocknodes[N2-1]})
	mh := Newminheap(func(a,b dnode) bool { return a.d < b.d })
	darr := make([]int,N+N2); for i:=0;i<N+N2;i++ { darr[i] = inf }; mh.Push(dnode{0,0})
	for !mh.IsEmpty() {
		x := mh.Pop()
		if darr[x.n] < inf { continue }
		darr[x.n] = x.d
		for _,e := range gr[x.n] { if darr[e.n2] == inf { mh.Push(dnode{e.n2,x.d+e.w}) } }
	}
	ans := darr[N-1]
	fmt.Println(ans)
}


