package main

import (
	"bufio"
	"fmt"
	"os"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }

const MOD = 998244353

func powmod(a, e int) int {
	res := 1; m := a
	for e > 0 { if e & 1 != 0 { res = res * m % MOD}; e >>= 1; m = m * m % MOD }
	return res
}

func makefact(n int) ([]int,[]int) {
	fact := make([]int,n+1); fact[0] = 1
	factinv := make([]int,n+1)
	for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % MOD }
	factinv[n] = powmod(fact[n],MOD-2)
	for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % MOD }
	return fact,factinv
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewReaderSize(f, BUFSIZE) }
	
    // NON-BOILERPLATE STARTS HERE
	// Ways to paint with zero pairs is M * (M-1)^(N-1)
	// Ways to paint with one pair is comb(N-1,1) * M * (M-1)^(N-1-1)
	// Ways to paint with two pairs is comb(N-1,2) * M * (M-1)^(N-1-2)
	// Special case M==1 just to avoid headaches.
	N,M,K := gti(),gti(),gti()
	fact,factinv := makefact(N)
	ans := 0
	if M == 1 {
		ans = tern(K==N-1,1,0)
	} else {
		for k:=0;k<=K;k++ {
			loc := fact[N-1] * factinv[k] % MOD * factinv[N-1-k] % MOD * M % MOD * powmod(M-1,N-1-k) % MOD
			ans = (ans + loc) % MOD
		}
	}
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



