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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9
	// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + .. 90
	// 1 + 2 + 3 + ......................   900
	N := gi(); 
	twoinv := powmod(2,MOD-2,MOD); ans := 0; first := 1; next := 10
	tri := func(n int) int { a := n % MOD; b := (n+1) % MOD; return a * b % MOD * twoinv % MOD }
	for {
		if N < next { ans += tri(N-first+1); break }
		ans += tri(next-first); first*=10; next*=10
	}
	ans %= MOD
	fmt.Println(ans)
}

