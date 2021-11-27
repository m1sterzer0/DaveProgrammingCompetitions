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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi2()
	// Find all of the primes between 2 and 1_000_000
	darr := ia(K); for i:=0;i<K;i++ { darr[i] = i+1 }
	narr := ia(K); for i:=0;i<K;i++ { narr[i] = N-K+1+i }
	sieve := make([]bool,1_000_001)
	for i:=0;i<=1_000_000;i++ { sieve[i] = true }
	sieve[0] = false; sieve[1] = false
	sieve[2] = true; for k:=4;k<=1_000_000;k+=2 { sieve[k] = false }
	for i:=3;i<=1000;i+=2 {
		if !sieve[i] { continue }
		for k:= i*i; k<=1_000_000; k += 2*i { sieve[k] = false }
	}
	ans := 1
	for p:=0;p<=1_000_000;p++ {
		if !sieve[p] { continue }
		// denominator factors
		dcnt := 1
		n1 := (1+p-1) / p * p
		for n1 <= K { 
			idx := n1-1;
			for darr[idx] % p == 0 { dcnt++; darr[idx] /= p }
			n1 += p
		}

		// numerator factors
		ncnt := 1
		n2 := (N-K+1+p-1) / p * p
		for n2 <= N {
			idx := n2 - (N-K+1)
			for narr[idx] % p == 0 { ncnt++; narr[idx] /= p }
			n2 += p
		}

		ans *= (ncnt-dcnt+1); ans %= MOD
	}

	// Now for the primes that remain in the numerator
	for i:=0;i<K;i++ {
		if narr[i] != 1 { ans *= 2; ans %= MOD }  // Since the primes here are greater than 1_000_000, only one copy of each type can show up in numerator
	}
	fmt.Println(ans)
}

