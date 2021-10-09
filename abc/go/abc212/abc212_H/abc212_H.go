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
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 998244353
func hadamard(n int, a []int, inv bool) []int {
	A := make([]int,n)
	for i:=0;i<n;i++ { A[i] = a[i] }
	h := 2
	for h <= n {
		hf := h/2
		for i:=0;i<n;i+=h {
			for j:=0;j<hf;j++ {
				u,v := A[i+j],A[i+j+hf]
				A[i+j],A[i+j+hf] = u+v,u-v
			}
		}
		h <<= 1
	}
	for i:=0;i<n;i++ { A[i] = (MOD + A[i] % MOD) % MOD }
	if inv {
		xx := powmod(n,MOD-2,MOD)
		for i:=0;i<n;i++ { A[i] *= xx; A[i] %= MOD }
	}
	return A
}
func sumpow(x,n int) int { //returns x + x^2 + x^3 + ... + x^n}
	if x == 0 || n == 0 { return 0 }
	if n == 1 { return x % MOD }
	if x == 1 { return n % MOD }
	num := (x * (powmod(x,n,MOD) - 1) % MOD + MOD) % MOD
	return num * powmod(x-1,MOD-2,MOD) % MOD
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi2(); A := gis(K)
	aa := ia(1<<16); for _,a := range A { aa[a]++ }
	ha := hadamard(1<<16,aa,false)
	ha2 := ia(1<<16); for i,a := range ha { ha2[i] = sumpow(a,N) }
	aa2 := hadamard(1<<16,ha2,true)
	ans := (sumpow(K,N) + MOD - aa2[0]) % MOD
	fmt.Println(ans)
}

