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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N);  for i,a := range A { if a != -1 { A[i]-- } }
	pown := make([]int,N+1); pown[0] = 1; for i:=1;i<=N;i++ { pown[i] = pown[i-1] * N % MOD }
	fact := make([]int,N+1); fact[0] = 1; for i:=1;i<=N;i++ { fact[i] = fact[i-1] * i % MOD }
	uf := NewDsu(N); dp := make([]int,N+5); dp[0] = 1; num := 0; ans := 0
	for i,a := range A { if a == -1 { num++ } else { uf.Merge(a,i) } }
	for i:=0;i<N;i++ {
		if uf.Leader(i) != i { continue }
		open := false
		for j:=0;j<N;j++ { if uf.Leader(j) == i && A[j] == -1 { open = true; break } }
		if !open { ans += pown[num]; ans %= MOD }
		if open { for x:=num;x>=0;x-- { dp[x+1] += dp[x] * uf.Size(i) % MOD; dp[x+1] %= MOD } }
	}
	for k:=1;k<=num;k++ { ans += fact[k-1] * dp[k] % MOD * pown[num-k] % MOD; ans %= MOD }
	fmt.Println(ans)
}

