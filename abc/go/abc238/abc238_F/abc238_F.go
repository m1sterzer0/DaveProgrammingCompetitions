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
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const MOD = 998244353
type Z struct {p,q int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); P := gis(N); Q := gis(N)
	scores := make([]Z,0)
	for i:=0;i<N;i++ { p,q := P[i],Q[i]; scores = append(scores,Z{p-1,q-1}) }
	sort.Slice(scores,func(i,j int) bool { return scores[i].p < scores[j].p || scores[i].p == scores[j].p && scores[i].q < scores[j].q })

	//Sort items from worst to best on the first test
	dp := twodi(K+1,N+1,0)
	ndp := twodi(K+1,N+1,0)
	//dp[i][j][k] = number of ways to select j from among the first i folks, where 
	//the lowest q value skipped so far is k
	dp[0][N] = 1
	for i:=0;i<N;i++ {
		for j:=0;j<=K;j++ { for k:=0;k<=N;k++ { ndp[j][k] = 0 } }
		for j:=0;j<=K;j++ {
			for k:=0;k<=N;k++ {
				// Option 1 -- add the current element to the set
				// Requirements: Need to have room, and need to make sure that I have not skipped a strictly better element
				if j < K && scores[i].q < k {
					ndp[j+1][k] += dp[j][k]; ndp[j+1][k] %= MOD
				}
				k2 := min(k,scores[i].q)
				ndp[j][k2] += dp[j][k]; ndp[j][k2] %= MOD
			}
		}
		dp,ndp = ndp,dp
	}
	ans := 0
	for i:=0;i<=N;i++ { ans += dp[K][i] }
	ans %= MOD
	fmt.Println(ans)
}
