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
func twodf(n int,m int,v float64) [][]float64 {
	r := make([][]float64,n); for i:=0;i<n;i++ { x := make([]float64,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
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
		N,M := gi(),gi()
		dp := twodf(N+1,M+1,0.00)
		for i:=0;i<=N;i++ {
			for j:=0;j<=M;j++ {
				if i == 0 && j == 0 { dp[i][j] = 1.00; continue }
				if i<=j { continue }
				n := N-i; m := M-j
				if i > 0 { dp[i][j] += dp[i-1][j] * float64(n+1)/float64(n+m+1) }
				if j > 0 { dp[i][j] += dp[i][j-1] * float64(m+1)/float64(n+m+1) }
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,dp[N][M])
    }
}

