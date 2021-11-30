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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a > b { return b }; return a }
const myinf int = 1<<60
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2(); B := gis(N)
		dp := iai(K+1,myinf); dp[0] = 0
		cumsum := ia(N); rs := 0; for i:=0;i<N;i++ { rs += B[i]; cumsum[i] = rs }
		best := myinf
		for en:=N-1;en>=0;en-- {
			for i:=en;i>=0;i-- {
				v := cumsum[en]
				if i != 0 { v -= cumsum[i-1] }
				if v <= K { cand := dp[K-v] + (en-i+1); best = min(best,cand) }
			}
			st := en
			for i:=st;i<N;i++ {
				v := cumsum[i]
				if st != 0 { v -= cumsum[st-1] }
				if v <= K { cand := i-st+1; dp[v] = min(dp[v],cand) }
			}
		}
		ans := -1; if best < myinf { ans = best }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}
