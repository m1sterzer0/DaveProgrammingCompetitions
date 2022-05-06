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
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi()
	dp := []int{0,1,1,1,1,1,1,1,1,1}
	ndp := []int{0,0,0,0,0,0,0,0,0,0}
	for i:=2;i<=N;i++ {
		ndp[1] = (dp[1]+dp[2]) % MOD
		ndp[9] = (dp[8]+dp[9]) % MOD
		for i:=2;i<=8;i++ { ndp[i] = (dp[i-1]+dp[i]+dp[i+1]) % MOD }
		dp,ndp = ndp,dp
	}
	ans := sumarr(dp); ans %= MOD; fmt.Println(ans)
}

