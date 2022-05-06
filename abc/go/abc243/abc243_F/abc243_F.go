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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
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
	// Let (c0,c1,c2,...,cn-1) be the counts of the the prizes we get after K pulls
	// then the solution is
	// sum_over_XX (K! / prod(ci!) * prod(pi^ci) )
	// where XX if the collections of tuples (c0,c1,...cn-1) which sum to K which have exactly M non-zero terms
	// We can alter this to K! * sum_over_xx ( prod(pi^ci)/prod(ci!) )
	// This can be computed term-by-term with a DP
	N,M,K := gi(),gi(),gi(); W := gis(N)
	sumw := sumarr(W); sumwinv := powmod(sumw,MOD-2,MOD)
	P := make([]int,N); for i:=0;i<N;i++ { P[i] = W[i] * sumwinv % MOD }
	fact,factinv := makefact(100,MOD)
	dp := twodi(M+1,K+1,0)
	ndp := twodi(M+1,K+1,0)
	dp[0][0] = 1; ndp[0][0] = 1
	for i:=0;i<N;i++ {
		pp := 1
		for cnt:=1;cnt<=K;cnt++ {
			pp *= P[i]; pp %= MOD
			f := pp * factinv[cnt] % MOD
			for oldcnt:=0;oldcnt+cnt<=K;oldcnt++ {
				for numuniq:=0;numuniq<M;numuniq++ {
					ndp[numuniq+1][oldcnt+cnt] += dp[numuniq][oldcnt] * f % MOD
				}
			}
		}
		for i:=0;i<=M;i++ { for j:=0;j<=K;j++ { ndp[i][j] %= MOD; dp[i][j] = ndp[i][j] } }
	}
	ans := dp[M][K] * fact[K] % MOD
	fmt.Println(ans)
}

