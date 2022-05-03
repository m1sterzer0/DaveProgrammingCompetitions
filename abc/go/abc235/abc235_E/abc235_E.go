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
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
type edge struct {idx,n2,cost int}
const inf = 2000000000000000000
type minheap struct { buf []edge; less func(edge, edge) bool }
func Newminheap(f func(edge, edge) bool) *minheap { buf := make([]edge, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v edge) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() edge { return q.buf[0] }
func (q *minheap) Pop() edge {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []edge) {
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
type segtree struct { n, size, log int; op func(int, int) int; e int; d []int }
func Newsegtree(n int, op func(int, int) int, e int) *segtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }; return NewsegtreeVec(v, op, e)
}
func NewsegtreeVec(v []int, op func(int, int) int, e int) *segtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz); d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i] }; st := &segtree{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *segtree) Set(p int, v int) {
	p += q.size; q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *segtree) Get(p int) int { return q.d[p+q.size] }
func (q *segtree) Prod(l int, r int) int {
	if r < l { return q.e }; r += 1; sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *segtree) Allprod() int { return q.d[1] }
func (q *segtree) MaxRight(l int, f func(int) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; sm := q.e
	for {
		for l%2 == 0 { l >>= 1 }
		if !f(q.op(sm, q.d[l])) {
			for l < q.size { l *= 2; if f(q.op(sm, q.d[l])) { sm = q.op(sm, q.d[l]); l++ } }; return l - q.size - 1
		}
		sm = q.op(sm, q.d[l]); l++; if l&-l == l { break }
	}
	return q.n - 1
}
func (q *segtree) MinLeft(r int, f func(int) bool) int {
	if r < 0 { return 0 }; r += q.size; sm := q.e; r++ 
	for {
		r--; for r > 1 && r%2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r], sm)) {
			for r < q.size { r = 2*r + 1; if f(q.op(q.d[r], sm)) { sm = q.op(q.d[r], sm); r-- } }; return r + 1 - q.size
		}
		sm = q.op(q.d[r], sm); if r&-r == r { break }
	}
	return 0
}
func (q *segtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,Q := gi(),gi(),gi(); A,B,C := fill3(M); U,V,W := fill3(Q)
	// General strategy:
	// * Build a minimum spanning tree of the original graph using Prim
	// * Record the "time" in which each node was added to the graph, and record the weight of the edge added at that "time"
	// * Build a segemnt tree where we can calculate a range maximum query (will resort to fancier if we need linear)
	// * For each query, we find the times for the two nodes, and we determine if any edges after the first node was inserted
	//   and up to the last node inserted had larger cost than the new proposed node.
	gr := make([][]edge,N+1)
	for i:=0;i<M;i++ { a,b,c := A[i],B[i],C[i]; gr[a] = append(gr[a],edge{i,b,c}); gr[b] = append(gr[b],edge{i,a,c}) }
	visitedEdge := make([]bool,M)
	visitedNode := make([]bool,N+1)
	insertTime := make([]int,N+1); for i:=0;i<=N;i++ { insertTime[i] = inf }
	edgeCostTime := make([]int,N+1)
	curtime := 0; insertTime[1] = 0; edgeCostTime[0] = 0; visitedNode[1] = true
	mh := Newminheap(func(e1,e2 edge) bool { return e1.cost < e2.cost })
	for _,ee := range gr[1] { if !visitedEdge[ee.idx] { visitedEdge[ee.idx] = true; mh.Push(ee) } }
	for !mh.IsEmpty() {
		e := mh.Pop()
		if visitedNode[e.n2] { continue }
		curtime++; insertTime[e.n2] = curtime; edgeCostTime[curtime] = e.cost; visitedNode[e.n2] = true
		for _,ee := range gr[e.n2] { if !visitedEdge[ee.idx] { visitedEdge[ee.idx] = true; mh.Push(ee) } }
	}
	st := NewsegtreeVec(edgeCostTime,max,0)
	for i:=0;i<Q;i++ {
		u,v,w := U[i],V[i],W[i]
		t1,t2 := insertTime[u],insertTime[v]
		if t1>t2 { t1,t2 = t2,t1 }
		w2 := st.Prod(t1+1,t2)
		if w2 < w { fmt.Fprintln(wrtr,"No") } else { fmt.Fprintln(wrtr,"Yes") }
	}
}

