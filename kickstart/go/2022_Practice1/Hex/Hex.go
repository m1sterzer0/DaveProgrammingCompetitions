package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gbs() []byte { return []byte(gs()) }

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
func (q *Mfgraph) Flow(s, t int) int { return q.FlowCapped(s, t, 1000000000000000000) }
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
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE

		// For a valid non-win, we must have
		// * Red is not connected from top to bottom
		// * Blue is not connected from left to right
		// * Red and blue total moves differ by no more than 1

		// For a valid Red win, we must have
		// * Red is connected from top to bottom
		// * Blue is NOT connected from left to right
		// * # Blue pieces <= # Red Pieces <= # Blue pieces + 1
		// * There must exist a red piece that we can 'pluck out' and find that red is not connected from top to bottom

		// (Similar for blue win -- just transpose the board)
		// Can use max-flow/min-cut theorem to just check if the max flow from top to bottom is 0,1,>=2.

		N := gi()
		bd := make([][]byte,N)
		for i:=0;i<N;i++ { bd[i] = gbs() }
		bd2 := make([][]byte,N)
		for i:=0;i<N;i++ { bd2[i] = make([]byte,N) }
		for i:=0;i<N;i++ { for j:=0;j<N;j++ { bd2[i][j] = bd[j][i] } }

		// Count pieces
		nr,nb := 0,0
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				if bd[i][j] == 'R' { nr++ }
				if bd[i][j] == 'B' { nb++ }
			}
		}

		// Node 0 is start
		// Node 1 to N*N for node inputs
		// Node N*N+1 to 2*N*N for node outputs
		// Node 2*N*N+1 is sink
		gennid := func(i,j int) int { return 1 + j + N*i }
		getflow := func(bb [][]byte, c byte) int {
			qq := NewMfgraph(2*N*N+2)
			for i:=0;i<N;i++ {
				for j:=0;j<N;j++ {
					if bb[i][j] != c { continue }
					nid := gennid(i,j)
					qq.Addedge(nid,nid+N*N,1)
					if i == 0 { qq.Addedge(0,nid,1) }
					if i == N-1 { qq.Addedge(N*N+nid,2*N*N+1,1) }
					if i-1 >= 0 && bb[i-1][j] == c { qq.Addedge(N*N+nid,gennid(i-1,j),1) }
					if i+1 < N  && bb[i+1][j] == c { qq.Addedge(N*N+nid,gennid(i+1,j),1) }
					if j-1 >= 0 && bb[i][j-1] == c { qq.Addedge(N*N+nid,gennid(i,j-1),1) }
					if j+1 < N  && bb[i][j+1] == c { qq.Addedge(N*N+nid,gennid(i,j+1),1) }
					if i-1 >= 0 && j+1 < N && bb[i-1][j+1] == c { qq.Addedge(N*N+nid,gennid(i-1,j+1),1) }
					if i+1 < N && j-1 >= 0 && bb[i+1][j-1] == c { qq.Addedge(N*N+nid,gennid(i+1,j-1),1) }
				}
			}
			return qq.FlowCapped(0,2*N*N+1,2)
		}

		fr := getflow(bd,'R')
		fb := getflow(bd2,'B')
		ans := "Nobody wins"
		if fr >= 2 || fb >= 2 || (fr == 1 && fb == 1) || (fr == 1 && nb > nr) || (fb ==1 && nr > nb) || nr - nb >= 2 || nb - nr >= 2 {
			ans = "Impossible"
		} else if fr == 1 {
			ans = "Red wins"
		} else if fb == 1 {
			ans = "Blue wins"
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

