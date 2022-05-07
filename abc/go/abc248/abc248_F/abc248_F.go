package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,P := gi(),gi()
	var dp  [2][3010]int
	var ndp [2][3010]int
	dp[0][2] = 3; dp[1][1] = 4; dp[1][0] = 1
	for n:=3;n<=N;n++ {
		for i:=0;i<N-2;i++ { ndp[0][i+2] += 2*dp[1][i] }
		for i:=0;i<N-1;i++ { ndp[0][i+1] += dp[0][i]; ndp[1][i+1] += 3*dp[1][i] }
		for i:=0;i<N;i++   { ndp[1][i] += dp[0][i] + dp[1][i] }
		for i:=0;i<N;i++   { dp[0][i] = ndp[0][i] % P; ndp[0][i] = 0; dp[1][i] = ndp[1][i] % P; ndp[1][i] = 0 }
	}
	ansarr := make([]int,0,N); for i:=1;i<=N-1;i++ { ansarr = append(ansarr,dp[1][i]) }
	ansstr := vecintstring(ansarr)
	fmt.Println(ansstr)
}
