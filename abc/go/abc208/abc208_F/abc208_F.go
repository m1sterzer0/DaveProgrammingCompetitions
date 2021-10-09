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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD int = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,K := gi3()
	d := M+K+10 // Can err on the side of a few extra terms
	dp := ia(d); w,nw := iai(31,0),iai(31,0); dp[0] = 0
	for i:=1;i<d;i++ {
		nw[0] = powmod(i,K,MOD)
		for i:=1;i<=M;i++ { v := nw[i-1] + w[i]; if v >= MOD { v -= MOD }; nw[i] = v }
		dp[i] = nw[M]; w,nw = nw,w
	}
	// Now we have M+K+1 function values so we need to do Lagrange interpolation
	// Given strict limits, we have to be careful about the modular inverses
	// The trick here is that given the points are 0,1,2,...,d-1, the denominator looks like the product of
	// two factorials times a sign correction term.
	if N < d { fmt.Println(dp[N]); return }
	nmod := N % MOD
	n1 := iai(d,1); for i:=1;i<d;i++    { f := nmod-(i-1); if f < 0 { f += MOD }; n1[i] = n1[i-1] * f % MOD }
	n2 := iai(d,1); for i:=d-2;i>=0;i-- { f := nmod-(i+1); if f < 0 { f += MOD }; n2[i] = n2[i+1] * f % MOD }
	_,factinv := makefact(d,MOD)
	ans := 0
	for i:=0;i<d;i++ {
		adder := n1[i] * n2[i] % MOD * dp[i] % MOD * factinv[i] % MOD * factinv[d-1-i] % MOD
		if (d-1-i) & 1 == 1 { adder = MOD - adder }
		ans += adder
	}
	ans %= MOD
	fmt.Println(ans)
}

