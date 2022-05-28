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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); P := gis(N)
	var dp [510][510]int
	for del:=0;del<N-1;del++ {
		for i:=1;i+del<N;i++ {
			j:=i+del
			if i == j { dp[i][j] = 1; continue }
			dp[i][j] = dp[i+1][j] // Case when we have no peers
			for k:=i+1;k<=j;k++ { 
				if P[k] < P[i] { continue }
				if k == i+1 { dp[i][j] += dp[k][j] } else { dp[i][j] += dp[i+1][k-1]*dp[k][j] % MOD } // Searching for the next peer
			} 
			dp[i][j] %= MOD
		}
	}
	ans := dp[1][N-1]
	fmt.Println(ans)
}

