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
func ia(m int) []int { return make([]int,m) }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
func lcm(a,b int) int { return a / gcd(a,b) * b }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	fact,factinv := makefact(50,MOD)
	comb := func(n,r int) int { return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD }
	// PROGRAM STARTS HERE
	N,K := gi2();
	ans := 0
	p := ia(N+1); p[N] = 1
	for {
		left := N
		cnt := 1
		s := 1
		for i:=N;i>=2;i-- {
			if p[i] == 0 { continue }
			s = lcm(s,i)
			// Counting is ways to choose the groups * ways to assign to subgroups * ways to arrange the cycle in each subgroup
			cnt *= factinv[p[i]]; cnt %= MOD
			for j:=0;j<p[i];j++ { cnt *= comb(left,i); cnt %= MOD; left-=i }
			cnt *= powmod(fact[i-1],p[i],MOD); cnt %= MOD
		}
		s2 := powmod(s,K,MOD)
		inc := cnt * s2 % MOD
		ans += inc

		// Next-partition code
		if p[1] == N { break }
		i := 2; for p[i] == 0 { i++ }
		v := p[1] + 1; p[1] = 0; p[i]--; i--; p[i]++
		p[i] += v/i; v %= i; if v > 0 { p[v]++ }
	}
	ans %= MOD
	fmt.Println(ans)
}
