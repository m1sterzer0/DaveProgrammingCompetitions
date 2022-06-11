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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func min(a,b int) int { if a > b { return b }; return a }
const inf  = 1000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); A := gis(N);
	psum := make([]int,N); psum[0] = A[0]; for i:=1;i<N;i++ { psum[i] = psum[i-1] + A[i] }
	parr := make([]int,0,M); for i:=0;i<N;i++ { for j:=0;j<A[i];j++ { parr = append(parr,i) } }
	parrcum := make([]int,M+1); parrcum[0] = 0; for i:=0;i<M;i++ { parrcum[i+1] = parrcum[i] + parr[i] }
	// dp[i][j][k] = cheapest way to fix prefix through i with last term at least j pancakes and prefix sum == k
	dp := twodi(M+1,M+1,inf)
	ndp := twodi(M+1,M+1,inf)
	for j:=0;j<=M;j++ { if psum[0] >= j { dp[j][j] = psum[0]-j } else { dp[j][j] = parrcum[j] } }
	for j:=M-1;j>=0;j-- { for k:=j;k<=M;k++ { dp[j][k] = min(dp[j][k],dp[j+1][k]) } } // fixes the "at least" part
	for i:=1;i<N;i++ {
		for j:=0;j<=M;j++ { for k:=0;k<=M;k++ { ndp[j][k] = inf } }
		for j:=0;j<=M;j++ {
			for k:=j; k<=M; k++ {
				add := 0
				if psum[i] >= k { 
					add = psum[i] - k
				} else {
					lend := min(j,k-psum[i])
					add = parrcum[k] - parrcum[k-lend] - i * lend
				}
				ndp[j][k] = dp[j][k-j] + add
			}
		}
		for j:=M-1;j>=0;j-- { for k:=j;k<=M;k++ { ndp[j][k] = min(ndp[j][k],ndp[j+1][k]) } }
		dp,ndp = ndp,dp
	}
	ans := dp[0][M]
	fmt.Println(ans)
}
