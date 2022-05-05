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
func abs(a int) int { if a < 0 { return -a }; return a }
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
	// Assume X >= Y >= Z >= 0
	// One-dimensional : If N<X or N-X is odd, then 0, otherwise comb(N,(N-X)/2).
	// Two-dimensional : Do (x,y) --> (x+y,x-y).  This makes choices independent, so assuming parity and minimums work, ans is f1(N,X+Y) * f1(N,X-y)
	// Three-dimensional : Simple linear sum.
	// The two dimension part is the hard part to see.
	fact,factinv := makefact(10000010,MOD)
	binom := func(n,r int) int { if n < 0 || r < 0 || r > n { return 0 }; return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD }
	f1 := func(n,x int) int { if x > n || (n & 1 != x & 1) { return 0 }; return binom(n,(n-x)/2) }
	f2 := func(n,x,y int) int { return f1(n,x+y) * f1(n,abs(x-y)) % MOD }
	N,X,Y,Z := gi(),gi(),gi(),gi()
	X = abs(X); Y = abs(Y); Z = abs(Z)
	ans := 0
	if N >= X+Y+Z && (N-(X+Y+Z)) & 1 == 0 {
		for z := Z; X+Y+z <= N; z += 2 { ans += binom(N,z) * f1(z,Z) % MOD * f2(N-z,X,Y) % MOD }
	}
	ans %= MOD; fmt.Println(ans)
}

