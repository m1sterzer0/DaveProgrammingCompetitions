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
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }

type mfpreedge struct{ to, rev, cap int }
type mfedge struct{ from, to, cap, flow int }
type mfpos struct{ x, y int }
type Mfgraph struct { n int; pos []mfpos; g [][]mfpreedge }
func NewMfgraph(n int) *Mfgraph { g := make([][]mfpreedge, n); pos := make([]mfpos, 0); return &Mfgraph{n, pos, g} }
func (q *Mfgraph) Addedge(from, to, cap int) int {
	m := len(q.pos); fromid := len(q.g[from]); toid := len(q.g[to]); q.pos = append(q.pos, mfpos{from, fromid})
	if from == to { toid++ }; q.g[from] = append(q.g[from], mfpreedge{to, toid, cap})
	q.g[to] = append(q.g[to], mfpreedge{from, fromid, 0}); return m
}
func (q *Mfgraph) Getedge(i int) mfedge {
	e := q.g[q.pos[i].x][q.pos[i].y]; re := q.g[e.to][e.rev]; return mfedge{q.pos[i].x, e.to, e.cap + re.cap, re.cap}
}
func (q *Mfgraph) Edges() []mfedge {
	m := len(q.pos); res := make([]mfedge, 0); for i := 0; i < m; i++ { res = append(res, q.Getedge(i)) }; return res
}
func (q *Mfgraph) Changeedge(i int, newcap int, newflow int) {
	e := &(q.g[q.pos[i].x][q.pos[i].y]); re := &(q.g[e.to][e.rev]); e.cap = newcap - newflow; re.cap = newflow
}
func (q *Mfgraph) Flow(s, t int) int { return q.FlowCapped(s, t, 1_000_000_000_000_000_000) }
func (q *Mfgraph) FlowCapped(s int, t int, flowlimit int) int {
	level := make([]int, q.n); iter := make([]int, q.n)
	bfs := func() {
		for i := 0; i < q.n; i++ { level[i] = -1 }; level[s] = 0; que := make([]int, 0, q.n); que = append(que, s)
		for len(que) > 0 {
			v := que[0]; que = que[1:]
			for _, e := range q.g[v] {
				if e.cap == 0 || level[e.to] >= 0 { continue }; level[e.to] = level[v] + 1; if e.to == t { return }
				que = append(que, e.to)
			}
		}
	}
	var dfs func(int, int) int
	dfs = func(v int, up int) int {
		if v == s { return up }; res := 0; level_v := level[v]
		for i := iter[v]; i < len(q.g[v]); i++ {
			e := q.g[v][i]; cap := q.g[e.to][e.rev].cap; if level_v <= level[e.to] || cap == 0 { continue }
			newup := up - res; if cap < up-res { newup = cap }; d := dfs(e.to, newup); if d <= 0 { continue }
			q.g[v][i].cap += d; q.g[e.to][e.rev].cap -= d; res += d; if res == up { return res }
		}
		level[v] = q.n; return res
	}
	flow := 0
	for flow < flowlimit {
		bfs(); if level[t] == -1 { break }; for i := 0; i < q.n; i++ { iter[i] = 0 }; f := dfs(t, flowlimit-flow)
		if f == 0 { break }; flow += f
	}
	return flow
}
func (q *Mfgraph) Mincut(s int) []bool {
	visited := make([]bool, q.n); que := make([]int, 0, q.n); que = append(que, s)
	for len(que) > 0 {
		p := que[0]; que = que[1:]; visited[p] = true
		for _, e := range q.g[p] { if e.cap > 0 && !visited[e.to] { visited[e.to] = true; que = append(que, e.to) } }
	}
	return visited
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }
	myinf := powint(10,18)
	mf := NewMfgraph(N*N+2); source := N*N; sink := N*N+1 
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			id := N*i+j
			if i!=0 { mf.Addedge(id,id-N,1) }
			if i!=N-1 { mf.Addedge(id,id+N,1) }
			if j!=0 { mf.Addedge(id,id-1,1) }
			if j!=N-1 { mf.Addedge(id,id+1,1)}
			if bd[i][j] == 'B' && (i+j) % 2 == 0 || bd[i][j] == 'W' && (i+j) % 2 == 1 { mf.Addedge(source,id,myinf) }
			if bd[i][j] == 'W' && (i+j) % 2 == 0 || bd[i][j] == 'B' && (i+j) % 2 == 1 { mf.Addedge(id,sink,myinf) }
		}
	}
	cutsize := mf.Flow(source,sink)
	ans := 2*N*(N-1) - cutsize
	fmt.Println(ans)
}