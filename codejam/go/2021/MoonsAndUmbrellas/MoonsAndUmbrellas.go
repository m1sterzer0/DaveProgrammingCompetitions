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
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	var dp [2][1000]int
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		X,Y,S := gi(),gi(),gs()
		N := len(S)
		// Init DP
		for i:=0;i<N;i++ { dp[0][i] = inf; dp[1][i] = inf }
		// Base case
		if S[0] != 'J' { dp[0][0] = 0 }
		if S[0] != 'C' { dp[1][0] = 0 }
		// Inductive step
		for i:=1;i<N;i++ {
			if S[i] != 'J' { dp[0][i] = min(dp[0][i],min(dp[0][i-1],dp[1][i-1]+Y)) }
			if S[i] != 'C' { dp[1][i] = min(dp[1][i],min(dp[1][i-1],dp[0][i-1]+X)) }
		}
		ans := min(dp[0][N-1],dp[1][N-1])
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

