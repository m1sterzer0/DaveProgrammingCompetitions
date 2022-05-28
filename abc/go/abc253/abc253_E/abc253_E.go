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
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,K := gi(),gi(),gi()
	dp := ia(M+1); ndp := ia(M+1); cumdp := ia(M+1)
	for i:=1;i<=M;i++ { dp[i] = 1 }
	for n:=2;n<=N;n++ {
		cumdp[0] = 0; for i:=1;i<=M;i++ { cumdp[i] = (dp[i] + cumdp[i-1]) % MOD }
		for i:=1;i<=M;i++ {
			ndp[i] = 0
			if K == 0 { ndp[i] = cumdp[M]; continue }
			if i-K >= 1 { ndp[i] += cumdp[i-K] }
			if i+K <= M { ndp[i] += MOD + cumdp[M] - cumdp[i+K-1] }
			ndp[i] %= MOD
		}
		dp,ndp = ndp,dp
	}
	ans := sumarr(dp); ans %= MOD
	fmt.Println(ans)
}

