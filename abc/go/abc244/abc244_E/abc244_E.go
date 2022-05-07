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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,K,S,T,X := gi(),gi(),gi(),gi(),gi(),gi()
	U,V := fill2(M)
	dp := twodi(2,N+1,0); ndp := twodi(2,N+1,0)
	dp[0][S] = 1
	for i:=0;i<K;i++ {
		for j:=0;j<M;j++ {
			u,v := U[j],V[j]
			if v == X { 
				ndp[1][v] += dp[0][u]; ndp[0][v] += dp[1][u] 
			} else { 
				ndp[0][v] += dp[0][u]; ndp[1][v] += dp[1][u] 
			}
			if u == X {
				ndp[1][u] += dp[0][v]; ndp[0][u] += dp[1][v] 
			} else {
				ndp[0][u] += dp[0][v]; ndp[1][u] += dp[1][v] 
			}
		}
		for j:=1;j<=N;j++ { dp[0][j] = ndp[0][j] % MOD; dp[1][j] = ndp[1][j] % MOD; ndp[0][j] = 0; ndp[1][j] = 0 }
	}
	ans := dp[0][T]; fmt.Println(ans)
}

