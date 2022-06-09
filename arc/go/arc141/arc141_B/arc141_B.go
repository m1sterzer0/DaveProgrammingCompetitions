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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi()
	// Key observation is that each term of A must increase the position of the MSB
	if N > 60 { fmt.Println(0); return }
	dp := [61][61]int{}
	for i:=0;i<=59;i++ {
		lo := 1<<uint(i); hi := (1<<uint(i+1))-1;
		if lo > M { continue }
		if hi <= M { dp[1][i] = (hi-lo+1) % MOD } else { dp[1][i] = (M-lo+1) % MOD }
	}
	for i:=2;i<=N;i++ {
		for j:=0;j<=59;j++ {
			for k:=j+1;k<=59;k++ {
				dp[i][j] += dp[i-1][k] 
			}
			dp[i][j] %= MOD; dp[i][j] *= dp[1][j]; dp[i][j] %= MOD
		}
	}
	ans := 0
	for j:=0;j<=59;j++ { ans += dp[N][j] }
	ans %= MOD
	fmt.Println(ans)
}

