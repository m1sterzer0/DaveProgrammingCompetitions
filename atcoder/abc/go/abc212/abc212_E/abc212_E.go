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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,K := gi3(); U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
	dp := iai(N,0); ndp := iai(N,0); dp[0] = 1
	for k:=0;k<K;k++ {
		sdp := 0; for _,d := range dp { sdp += d }
		for i:=0;i<N;i++ { ndp[i] = sdp - dp[i] }
		for i:=0;i<M;i++ {
			u,v := U[i],V[i]
			ndp[u] -= dp[v]
			ndp[v] -= dp[u]
		}
		for i:=0;i<N;i++ { ndp[i] %= MOD }
		dp,ndp = ndp,dp
	}
	fmt.Println(dp[0])
}

