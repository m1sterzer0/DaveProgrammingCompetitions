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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type PI struct{ x, y int }
func Kosaraju(n int, diredges []PI) (int, []int) {
	g, grev, visited, visitedInv, scc, s, counter := make([][]int, n), make([][]int, n), make([]bool, n), make([]bool, n), make([]int, n), make([]int, 0, n), 0
	var dfs1, dfs2 func(int)
	for _, xx := range diredges { x, y := xx.x, xx.y; g[x] = append(g[x], y); grev[y] = append(grev[y], x) }
	dfs1 = func(u int) { if !visited[u] { visited[u] = true; for _, c := range g[u] { dfs1(c) }; s = append(s, u) } }
	for i := 0; i < n; i++ { dfs1(i) }
	dfs2 = func(u int) {
		if !visitedInv[u] { visitedInv[u] = true; for _, c := range grev[u] { dfs2(c) }; scc[u] = counter }
	}
	for i := n - 1; i >= 0; i-- { nn := s[i]; if !visitedInv[nn] { dfs2(nn); counter += 1 } }; return counter, scc
}
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
	inf := 1_000_000_000_000_000_000; res := 0; h := make([]int, q.n); prv_v := make([]int, q.n)
	prv_e := make([]int, q.n); f := 0; dist := make([]int, q.n); for i := 0; i < q.n; i++ { dist[i] = inf }
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
	N,M,K := gi3(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	X := gis(N)
	edges := make([]PI,M); for i:=0;i<M;i++ { edges[i] = PI{A[i],B[i]} }
	numscc,scc := Kosaraju(N,edges)
	value := ia(numscc)
	for i:=0;i<N;i++ { value[scc[i]] += X[i] }
	newedges := []PI{}
	leafedges := make([]bool,numscc); for i:=0;i<numscc;i++ { leafedges[i] = true }
	for i:=0; i<M; i++ {
		a,b := A[i],B[i]
		scca,sccb := scc[a],scc[b]
		if scca == sccb { continue }
		newedges = append(newedges,PI{scca,sccb})
	}
	sort.Slice(newedges, func(i,j int) bool { return newedges[i].x < newedges[j].x || newedges[i].x == newedges[j].x && newedges[i].y < newedges[j].y})
	cumval := ia(numscc)
	cumval[0] = value[0]
	for i:=1;i<numscc;i++ { cumval[i] = cumval[i-1] + value[i] }
	mcf := NewMinCostFlow(2*numscc+2)
	vv := 0; if scc[0] != 0 { vv = cumval[scc[0]-1] }
	mcf.AddEdge(2*numscc,2*scc[0],K,vv)
	inf := 1_000_000_000_000_000_000
	for i:=0;i<numscc;i++ {
		mcf.AddEdge(2*i,2*i+1,1,0)
		mcf.AddEdge(2*i,2*i+1,inf,value[i])
	}
	for i:=0;i<numscc;i++ { 
		if leafedges[i] { mcf.AddEdge(2*i+1,2*numscc+1,inf,cumval[len(cumval)-1]-cumval[i]) }
	}
	for i:=0; i<len(newedges); i++ {
		if i > 0 && newedges[i] == newedges[i-1] { continue }
		a,b := newedges[i].x,newedges[i].y
		mcf.AddEdge(2*a+1,2*b,inf,cumval[b]-cumval[a]-value[b])
	}
	_,mincost := mcf.Flowssp(2*numscc,2*numscc+1)
	ans := K * cumval[len(cumval)-1] - mincost
	fmt.Println(ans)
}

