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
	N,M := gi2()
	ways := ia(N+1)
	biggestset := 1 + (N-1) / M
	smallestset := biggestset-1
	numbiggest := N - smallestset * M
	numsmallest := M - numbiggest
	fact,factinv := makefact(5000,MOD)
	// Ignore the non-empty requirement for ways
	for i:=1;i<=N;i++ {
		if biggestset > i { ways[i] = 0; continue }
		waysperbig := fact[i] * factinv[i-biggestset] % MOD
		waysbig := powmod(waysperbig,numbiggest,MOD)
		wayspersmall := fact[i] * factinv[i-smallestset] % MOD
		wayssmall := powmod(wayspersmall,numsmallest,MOD)
		ways[i] = waysbig * wayssmall % MOD
	}
	for i:=1;i<=N;i++ {
		lans,s := ways[i],-1
		for j:=i-1;j>0;j-- {
			adder := fact[i] * factinv[j] % MOD * factinv[i-j] % MOD * ways[j] % MOD * s
			lans = (lans + MOD + adder) % MOD
			s *= -1 
		}
		lans = lans * factinv[i] % MOD
		fmt.Fprintln(wrtr,lans)
	}
}

