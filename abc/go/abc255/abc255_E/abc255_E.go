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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); S := gis(N-1); X := gis(M)
	// Need to calculate what I need the starting number to be in order for A_i = X_j
	// Special case for A0
	// A1 = S0 - A0
	// A2 = (S1-S0) + A0
	// A3 = (S2-S1+S0) - A0
	// A4 = (S3-S2+S1-S0) + A0
	dp := twodi(N,M,-1)
	// Do the evens first
	ss := 0
	for i:=0;i<N;i+=2 {
		for j,x := range X { dp[i][j] = x-ss }
		if i+2 < N { ss += S[i+1]-S[i]}
	}
	// Now the odds
	ss = S[0]
	for i:=1;i<N;i+=2 {
		for j,x := range X { dp[i][j] =  ss-x }
		if i+2 < N { ss += S[i+1]-S[i]}
	}
	// Now count the most popular starting number
	sb  := make(map[int]int)
	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			sb[dp[i][j]]++
		}
	}
	ans := 0
	for _,v := range sb { ans = max(ans,v) }
	fmt.Println(ans)
} 
