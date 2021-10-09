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
func ia(m int) []int { return make([]int,m) }
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
	N := gi(); C := gis(N)
	fact,factinv := makefact(N,MOD)
	d := make(map[int]int)
	for _,c := range C { d[c]++ }
	sb := ia(N+1); for _,v := range d { sb[v] += 1 }
	ans := ia(N+1)
	combnk := ia(N+1); combnk[0] = 1
	combnkinv := ia(N+1); combnkinv[0] = 1
	for k:=1;k<=N;k++ { combnk[k]    = fact[N] * factinv[k] % MOD * factinv[N-k] % MOD }
	for k:=1;k<=N;k++ { combnkinv[k] = fact[k] * fact[N-k] % MOD * factinv[N] % MOD }
	for sz:=1;sz<=N;sz++ {
		if sb[sz] == 0 { continue }
		cnt := sb[sz]
		for k:=1;k<=N;k++ {
			term := combnk[k]
			if N-sz >= k { 
				xx := fact[N-sz] * factinv[k] % MOD * factinv[N-sz-k] % MOD
				term = (term+MOD-xx) % MOD
			}
			ans[k] += cnt * term ; ans[k] %= MOD
		}
	}
	for k:=1;k<=N;k++ {
		ans[k] *= combnkinv[k]; ans[k] %= MOD 
	}
	for i:=1;i<=N;i++ { fmt.Fprintln(wrtr,ans[i]) }
}

