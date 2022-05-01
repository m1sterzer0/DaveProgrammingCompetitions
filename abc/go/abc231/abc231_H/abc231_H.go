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
type MinCostFlowPI struct{ c, v int }
type MinHeapMinCostFlow struct { buf []MinCostFlowPI; less func(MinCostFlowPI, MinCostFlowPI) bool }
func NewMinHeapMinCostFlow(f func(MinCostFlowPI, MinCostFlowPI) bool) *MinHeapMinCostFlow {
	buf := make([]MinCostFlowPI, 0); return &MinHeapMinCostFlow{buf, f}
}
func (q *MinHeapMinCostFlow) IsEmpty() bool { return len(q.buf) == 0 }
func (q *MinHeapMinCostFlow) Push(v MinCostFlowPI) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *MinHeapMinCostFlow) Pop() MinCostFlowPI {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *MinHeapMinCostFlow) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *MinHeapMinCostFlow) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}
type MinCostFlow struct { n, numedges int; g [][]int; to, cap, cost []int }
func NewMinCostFlow(n int) *MinCostFlow {
	g := make([][]int, n); to := make([]int, 0); cap := make([]int, 0); cost := make([]int, 0)
	return &MinCostFlow{n, 0, g, to, cap, cost}
}
func (q *MinCostFlow) AddEdge(fr, to, cap, cost int) {
	q.to = append(q.to, to); q.to = append(q.to, fr); q.cap = append(q.cap, cap); q.cap = append(q.cap, 0)
	q.cost = append(q.cost, cost); q.cost = append(q.cost, -cost); q.g[fr] = append(q.g[fr], q.numedges)
	q.g[to] = append(q.g[to], q.numedges+1); q.numedges += 2
}
func (q *MinCostFlow) Flowssp(s, t int) (int, int) {
	inf := 1000000000000000000; res := 0; h := make([]int, q.n); prv_v := make([]int, q.n); prv_e := make([]int, q.n)
	f := 0; dist := make([]int, q.n); for i := 0; i < q.n; i++ { dist[i] = inf }
	for {
		for i := 0; i < q.n; i++ { dist[i] = inf }; dist[s] = 0
		que := NewMinHeapMinCostFlow(func(a, b MinCostFlowPI) bool { return a.c < b.c }); que.Push(MinCostFlowPI{0, s})
		for !que.IsEmpty() {
			xx := que.Pop(); c, v := xx.c, xx.v; if dist[v] < c { continue }; r0 := dist[v] + h[v]
			for _, e := range q.g[v] {
				w, cap, cost := q.to[e], q.cap[e], q.cost[e]
				if cap > 0 && r0+cost-h[w] < dist[w] {
					r := r0 + cost - h[w]; dist[w] = r; prv_v[w] = v; prv_e[w] = e; que.Push(MinCostFlowPI{r, w})
				}
			}
		}
		if dist[t] == inf { return f, res }; for i := 0; i < q.n; i++ { h[i] += dist[i] }; d := inf; v := t
		for v != s { dcand := q.cap[prv_e[v]]; if dcand < d { d = dcand }; v = prv_v[v] }; f += d; res += d * h[t]
		v = t; for v != s { e := prv_e[v]; e2 := e ^ 1; q.cap[e] -= d; q.cap[e2] += d; v = prv_v[v] }
	}
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Main idea -- traditional min cost flow is too hard.
	// Instead, we turn them all to black and try to save as much as we can with removals
	// The matching of "removal budget" between row and column works better than dealing with the thorny issue
	// of sometimes covering a (row,col) pair with a black piece and a singular row/col.
	H,W,N := gi(),gi(),gi(); A,B,C := fill3(N); for i:=0;i<N;i++ { A[i]--; B[i]-- }
	rowcnt := make([]int,H); colcnt := make([]int,W)
	for i:=0;i<N;i++ { rowcnt[A[i]]++; colcnt[B[i]]++ }
	totcost := 0; for i:=0;i<N;i++ { totcost += C[i] }
	// Nodes
	// 0 to H-1 are rows
	// H to H+W-1 are cols
	// H+W is source
	// H+W+1 is governor
	// H+W+2 is sink
	mcf := NewMinCostFlow(H+W+3)
	mcf.AddEdge(H+W,H+W+1,0,0) // This is the edge we are going to keep playing with
	for i:=0;i<H;i++ { mcf.AddEdge(H+W+1,i,rowcnt[i]-1,0) }
	for i:=0;i<W;i++ { mcf.AddEdge(H+i,H+W+2,colcnt[i]-1,0) }
	for i:=0;i<N;i++ { mcf.AddEdge(A[i],H+B[i],1,-C[i]) }
	best := totcost
	for {
		mcf.cap[0] = 1
		f,res := mcf.Flowssp(H+W,H+W+2)
		if f == 0 || res >= 0 { break }
		best += res
	}
	fmt.Println(best)
}


