package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}

const MOD = 1_000_000_007

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	K := gi()
	S := gs()
	ls := len(S)
	inv25 := powmod(25,MOD-2,MOD)
	fact,factinv := makefact(K+ls+3,MOD)
	ans := 0; t1 := 1; t2 := powmod(25,K,MOD)
	for suffixLen:=0;suffixLen<=K;suffixLen++ {
		nonsuf := K - suffixLen
		// need to divide nonsuffix into ls distinct bins
		// by stars and bars, this is Comb(nonsuffix+ls-1,ls-1)
		t3 := fact[nonsuf+ls-1] * factinv[ls-1] % MOD * factinv[nonsuf] % MOD
		lans := t1 * t2 % MOD * t3 % MOD
		ans = (ans + lans) % MOD
		t1 = t1 * 26 % MOD // SUFFIX TERM
		t2 = t2 * inv25 % MOD // PREFIX TERM
	}
	fmt.Println(ans)
}



