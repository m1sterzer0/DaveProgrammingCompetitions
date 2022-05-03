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
	N,A,B,C := gi(),gi(),gi(),gi()
	fact,factinv := makefact(N+10,MOD)
	binom := func(n,k int) int { 
		if n < 0 || k < 0 || k > n { return 0 }
		return fact[n] * factinv[k] % MOD * factinv[n-k] % MOD
	}
	ans,a,b,c := 0,1,1,1
	sgn := 1; if N % 2 == 1 { sgn = -1 }
	for i:=0;i<=N;i++ {
		cur := binom(N,i) * a % MOD * b % MOD * c % MOD
		ans += cur * sgn
		a = (a + a - binom(i,A)) % MOD
		b = (b + b - binom(i,B)) % MOD
		c = (c + c - binom(i,C)) % MOD
		sgn *= -1
	}
	ans %= MOD; ans += MOD; ans %= MOD
	fmt.Println(ans)
}

