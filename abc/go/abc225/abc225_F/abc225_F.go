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
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi2(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
	sort.Slice(S,func(i,j int) bool { return S[i] + S[j] < S[j] + S[i] } )
	dp := make([]string,N+1); ndp := make([]string,N+1)
	for i:=0;i<=N;i++ { dp[i] = "" }
	for k:=1;k<=K;k++ {
		ndp[N-k] = S[N-k] + dp[N-k+1]
		for i:=N-k-1;i>=0;i-- { ndp[i] = ndp[i+1]; cand := S[i] + dp[i+1]; if cand < ndp[i] { ndp[i] = cand } }
		dp,ndp = ndp,dp
	}
	fmt.Println(dp[0])
}
