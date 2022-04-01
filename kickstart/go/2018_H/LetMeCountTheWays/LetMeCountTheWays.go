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
func gi2() (int,int) { return gi(),gi() }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	fact,factinv := makefact(200010,MOD)
	comb := func(n,r int) int { return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD }
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M := gi2()
		twon := 2 * N
		ans := fact[twon]
		for i:=1;i<=M;i++ {
			waysToChooseCouples := comb(M,i)
			waysToOrderPairWithinCouples := powmod(2,i,MOD)
			waysToOrderRemainingPeople := fact[twon-i]
			adder := waysToChooseCouples * waysToOrderPairWithinCouples % MOD * waysToOrderRemainingPeople % MOD
			if i % 2 == 1 { ans += MOD - adder } else { ans += adder }
			ans %= MOD
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

