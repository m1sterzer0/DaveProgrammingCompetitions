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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	dp := [1501]int{}; ndp := [1501]int{}; plates := [50][30]int{}
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K,P := gi3()
		for i:=0;i<N;i++ { for j:=0;j<K;j++ { plates[i][j] = gi()} }
		for i:=0;i<=P;i++ { dp[i] = 0; ndp[i] = 0 }
		for i:=0;i<N;i++ {
			sumb := 0
			for j:=1;j<=K;j++ {
				sumb += plates[i][j-1]
				for k:=P-j;k>=0;k-- { 
					ndp[k+j] = max(ndp[k+j],dp[k]+sumb)
				}
			}
			for j:=0;j<=P;j++ { dp[j] = ndp[j] }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,dp[P])
    }
}

