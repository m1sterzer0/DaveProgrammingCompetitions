package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	S := gs()
	N := len(S)
	d := make([]int,26)
	fact,factinv := makefact(5010,MOD)
	for _,c := range S { d[int(byte(c)-'a')]++ }
	dp := make([]int,N+1)
	ndp := make([]int,N+1)
	dp[0]=1
	for i:=0;i<26;i++ {
		for j:=0;j<=N;j++ { ndp[j] = 0 }
		for j:=0;j<=d[i];j++ {
			for k:=0;k+j<=N;k++ {
				comb := fact[k+j] * factinv[k] % MOD * factinv[j] % MOD
				ndp[k+j] += dp[k] * comb % MOD
			}
		}
		for j:=0;j<=N;j++ { ndp[j] %= MOD }
		dp,ndp = ndp,dp
	}
	ans := 0
	for i:=1;i<=N;i++ { ans += dp[i] }
	ans %= MOD
	fmt.Println(ans)
}
