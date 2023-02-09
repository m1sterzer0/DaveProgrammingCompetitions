package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)

func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
func sortUniq(a []int) []int {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}

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
	//f1, _ := os.Create("cpu.prof"); penv mprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	// This is a max flow with min capacities problem.
	// Algorithm for determining if max flow with min capacity exists
	// * if (u->v, cap=r, mincap=l), replace with (s'->v, cap=l),(u->t', cap=l), (u->v, cap=r-l)
	// * add flows in the following order: s'->t', s->t', s'->t, s->t
	// * then check to see that all s' edges are fully booked and all t' edges are fully booked (think you can skip the last step)
	H,W := gi(),gi(); bd := make([]string,H); for i:=0;i<H;i++ { bd[i] = gs() }
	// Node numbers 0 to H*W-1 are board nodes
	// Node HW is s
	// Node HW+1 is t
	// Node HW+2 is s'
	// Node HW+3 is t'
	mf := NewMfgraph(H*W+4)
	s,t,ss,tt := H*W,H*W+1,H*W+2,H*W+3
	for i:=0;i<H;i++ {
		for j:=0;j<W;j++ {
			n := W*i+j
			if (i+j)%2 == 0 {
				if bd[i][j] == '1' { continue }
				if bd[i][j] == '?' { mf.Addedge(s,n,1)  }
				if bd[i][j] == '2' { mf.Addedge(ss,n,1) } // Omit s->tt since s has infinite capacity
				if i+1 <  H && bd[i+1][j] != '1' { mf.Addedge(n,n+W,1) }
				if i-1 >= 0 && bd[i-1][j] != '1' { mf.Addedge(n,n-W,1) }
				if j+1 <  W && bd[i][j+1] != '1' { mf.Addedge(n,n+1,1) }
				if j-1 >= 0 && bd[i][j-1] != '1' { mf.Addedge(n,n-1,1) }
			} else {
				if bd[i][j] == '1' { continue }
				if bd[i][j] == '?' { mf.Addedge(n,t,1)  }
				if bd[i][j] == '2' { mf.Addedge(n,tt,1) } // Omit ss->t since ss has infinite capacity
			}
		}
	}
	mf.Flow(ss,tt); mf.Flow(s,tt); mf.Flow(ss,t) // Prioritize flow to use ss & tt instead of s & t
	ans := "Yes"
	edges := mf.Edges()
	for _,e := range edges {
		if e.from == ss && e.flow != 1 { ans = "No" }
		if e.to == tt && e.flow != 1 { ans = "No" }
	}
	fmt.Println(ans)
}

