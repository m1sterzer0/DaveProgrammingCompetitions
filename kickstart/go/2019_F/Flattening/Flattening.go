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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	dp := twodi(103,103,inf)
	olddp := twodi(103,103,inf)
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// dp[i][j][k] = minimum number of changes necessary in first i squares to
		//               have k iso-height segments with j as last unchanged block (index N denotes all prev blocks have been flattened)
		N,K := gi2(); A := gis(N)
		K++ // Found it easier to think about the total number of segments instead of the total number of changes
		for k:=0;k<=K;k++ {	for j:=0;j<=N;j++ {	dp[j][k] = inf } }
		dp[N][0] = 0 //Base case
		for i:=0;i<N;i++ {
			olddp,dp = dp,olddp
			for k:=0;k<=K;k++ {	for j:=0;j<=N;j++ { dp[j][k] = inf } }
			for k:=0;k<=K;k++ {	
				for j:=0;j<=N;j++ {
					if j < N &&  A[j] == A[i] {  dp[i][k]   = min(dp[i][k],olddp[j][k]) }    // Keep when I match
					if k < K                  {  dp[i][k+1] = min(dp[i][k+1],olddp[j][k]) }  // Keep when I don't match
					dp[j][k] = min(dp[j][k],olddp[j][k]+1)                                   // Throw away case
				}
			}
		}
		//for k:=0;k<=K;k++ { for j:=0;j<=N;j++ { fmt.Printf("DBG dp[%v][%v]=%v\n",j,k,dp[j][k])}}
		best := inf
		for k:=0;k<=K;k++ {	for j:=0;j<=N;j++ {	best = min(best,dp[j][k] ) } }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

