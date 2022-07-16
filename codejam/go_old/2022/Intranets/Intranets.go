package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 1000000007
var fact [1000010]int
var factinv [1000010]int
// Still don't know why this works, but multiple people submitted something similar to this, so 
// I need to dig in to understand the inclusion/exclusion here.
func solve(M,K int) int {
	twoinv := powmod(2,MOD-2,MOD)
	temp := make([]int,M/2+5)
	temp[0] = 1
	sum := 0
	for i:=1;i<=M/2;i++ {
		sum += 2*M - 4*i + 1; sum %= MOD
		temp[i] = temp[i-1] * (M-2*i+2) % MOD * (M-2*i+1) % MOD * twoinv % MOD * powmod(sum,MOD-2,MOD) % MOD
	}
	ans := 0
	for i:=K;i<=M/2;i++ {
		sgn := -1; if i&1==K&1 { sgn = 1 }
		ans += sgn * temp[i] * fact[i] % MOD * factinv[K] % MOD * factinv[i-K] % MOD
		ans += MOD; ans %= MOD
	}
	return ans
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	nmax := 1000000
	fact[0] = 1; for i:=1;i<=nmax;i++ { fact[i] = fact[i-1] * i % MOD }
	factinv[nmax] = powmod(fact[nmax],MOD-2,MOD); for i:=nmax-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % MOD }
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		M,K := gi(),gi()
		ans := solve(M,K)
		fmt.Printf("Case #%v: %v\n",tt,ans)
    }
}

