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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func min(a,b int) int { if a > b { return b }; return a }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}

const MOD int = 998244353

// Adapted from https://codeforces.com/blog/entry/45233
func sumoversubsets(N int, A []int, inplace bool) []int {
	F := A; if !inplace { F = make([]int,len(A)); copy(F,A) }
	for i:=0;i<N;i++ {
		for mask:=0;mask<1<<N;mask++ {
			if mask & (1<<i) != 0 { F[mask] += F[mask^(1<<i)] }
		}
	}
	return F
}

func sumoversupersets(N int, A []int, inplace bool) []int {
	F := A; if !inplace { F = make([]int,len(A)); copy(F,A) }
	rev(F); sumoversubsets(N,F,true); rev(F); return F
}

// AKA Mobius Transorm https://codeforces.com/blog/entry/72488
func inversesumoversubsets(N int, A []int, inplace bool) []int {
	F := A; if !inplace { F = make([]int,len(A)); copy(F,A) }
	for i:=0;i<N;i++ {
		for mask:=0;mask<1<<N;mask++ {
			if mask & (1<<i) != 0 { 
				F[mask] += MOD - F[mask^(1<<i)]
				F[mask] %= MOD
			}
		}
	}
	return F
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); A := gis(N); B := gis(M)
	C := make([][]int,N); for i:=0;i<N;i++ { C[i] = gis(M) }
	supply := ia(1<<N)
	for bm:=0;bm<1<<N;bm++ {
		for n:=0;n<N;n++ {
			if bm & (1<<n) != 0 { supply[bm] += A[n] }
		}
	}
	demand := ia(1<<N)
	for m:=0;m<M;m++ {
		bm := 0
		for n:=0;n<N;n++ { if C[n][m] != 0 { bm |= (1<<n) } }
		demand[bm] += B[m]
	}
	sumoversubsets(N,demand,true)
	inf := 1_000_000_000_000_000_000
	minexcess := inf
	for x:=0;x<1<<N;x++ { if demand[x] > 0 { minexcess = min(minexcess,supply[x]-demand[x])}}
	if minexcess < 0 { fmt.Println("0 1"); return }
	snuke := minexcess+1
	feasible := ia(1<<N)
	for x:=0;x<1<<N;x++ { if demand[x] > 0 && supply[x]-demand[x] == minexcess { feasible[x] = 1 } }
	sumoversupersets(N,feasible,true)
	maxsupply := maxarr(supply)
	fact,factinv := makefact(maxsupply,MOD)
	binom := ia(maxsupply+1)
	for x:=snuke;x<=maxsupply;x++ { binom[x] = fact[x] * factinv[snuke] % MOD * factinv[x-snuke] % MOD }
	uniqueways := ia(1<<N)
	for x:=0;x<1<<N;x++ { if feasible[x] > 0 { uniqueways[x] = binom[supply[x]] } }
	inversesumoversubsets(N,uniqueways,true)
	snukeways := 0
	for x:=0;x<1<<N;x++ { if feasible[x] > 0 { snukeways += uniqueways[x] } }
	snukeways %= MOD
	fmt.Printf("%v %v\n",snuke,snukeways)
}

