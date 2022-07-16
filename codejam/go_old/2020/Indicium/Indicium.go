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

func print2dIntArr(sb [][]int) {
	for i:=0;i<len(sb);i++ {
		n := len(sb[i])
		ansarr := make([]string,n)
		for j:=0;j<n;j++ { ansarr[j] = strconv.Itoa(sb[i][j]) }
		ansstr := strings.Join(ansarr," ")
		fmt.Fprintf(wrtr,"%v\n",ansstr)
	}
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
		N,K := gi(),gi()
		if K == N+1 || K == N*N-1 || (N == 3 && (K==5 || K==7)) {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			A,B,C,done := 0,0,0,false
			for a:=1;a<=N;a++ {
				for b:=1;b<=N;b++ {
					for c:=1;c<=N;c++ {
						if (N-2)*a+b+c != K { continue }
						if b != c && (a == b || a == c) { continue }
						A,B,C,done = a,b,c,true
						break
					}
					if done { break }
				}
				if done { break}
			}

			sb := make([][]int,N); for i:=0;i<N;i++ { sb[i] = make([]int,N) }
			sb[0][0] = B; sb[1][1] = C; for i:=2;i<N;i++ { sb[i][i] = A }
			for r:=0;r<N;r++ {
				// Node 0 to N-1 are columns
				// Node N is source
				// Node N+1 to N+N are numbers
				// Node 2N+1 is sink
				taken := make(map[int]bool)
				for i:=0;i<N;i++ { if sb[r][i] != 0 { taken[sb[r][i]] = true } }
				mf := NewMfgraph(2*N+2)

				// Add the edges
				for i:=0;i<N;i++ { mf.Addedge(N,i,1) }
				for i:=N+1;i<=2*N;i++ { mf.Addedge(i,2*N+1,1) }
				for j:=0;j<N;j++ {
					if sb[r][j] != 0 {
						mf.Addedge(j,N+sb[r][j],1)
					} else {
						loctaken := make(map[int]bool)
						for i:=0;i<N;i++ { if sb[i][j] != 0 { loctaken[sb[i][j]] = true } }
						for n:=1;n<=N;n++ {
							if loctaken[n] { continue }
							mf.Addedge(j,N+n,1)
						}
					}
				}

				// Run the flow and reap the rewards
				f := mf.Flow(N,2*N+1)
				if f != N { fmt.Printf("SOMETHING BAD HAPPENED. EXITING\n"); os.Exit(1) }
				ee := mf.Edges()
				for _,e := range ee {
					if e.flow == 1 && e.from < N && e.to >= N && e.to <= 2*N { sb[r][e.from] = e.to-N }
				}
			}
			// Print the table
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"POSSIBLE")
			print2dIntArr(sb)
		}
	}
}

