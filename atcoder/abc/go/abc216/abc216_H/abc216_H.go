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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// LGV problem -- need to do math offline and reduce to this form
	K,N := gi2(); X := gis(K)
	f,finv := makefact(N,MOD)
	comb := ia(N+1); for i:=0;i<=N;i++ { comb[i] = f[N] * finv[i] % MOD * finv[N-i] % MOD }
	dp := ia(1<<K); ndp := ia(1<<K); dp[0] = 1
	lb,ub := minarr(X),maxarr(X)+N
	for i:=lb;i<=ub;i++ {
		copy(ndp,dp)
		for j:=(1<<K)-1;j>=0;j-- {
			sgn := 1
			for k:=K-1;k>=0;k-- {
				if j & (1<<k) == 0 { continue }
				if i - X[k] < 0 || i - X[k] > N { continue }
				ndp[j] += dp[j ^ (1<<k)] * comb[i-X[k]] * sgn % MOD + MOD
				ndp[j] %= MOD
				sgn *= -1
			}
		}
		dp,ndp = ndp,dp
	}
	twopow := powmod(2,N*K,MOD)
	twopowinv := powmod(twopow,MOD-2,MOD)
	ans := dp[(1<<K)-1] * twopowinv % MOD
	fmt.Println(ans)
}

