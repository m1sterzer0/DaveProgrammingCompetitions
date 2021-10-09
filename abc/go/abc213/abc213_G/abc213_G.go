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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	subg := ia(1<<N)
	for m:=0;m<1<<N;m++ {
		nume := 0
		for i:=0;i<M;i++ { 
			a,b := A[i],B[i]
			if m & (1<<a) != 0 && m & (1<<b) != 0 { nume++ }
		}
		subg[m] = powmod(2,nume,MOD)
	}
	conn := ia(1<<N)
	for m:=1;m<1<<N;m+=2 {
		val := subg[m]
		submask := (m-1) & m
		for submask > 0 {
			if submask & 1 != 0 {
				antimask := m ^ submask
				ways := subg[antimask] * conn[submask] % MOD
				val = (MOD + val - ways) % MOD
			}
			submask = (submask-1) & m
		}
		conn[m] = val
	}

	fullmask := (1<<N)-1
	ans := ia(N)
	for m:=1;m<1<<N;m+=2 {
		antimask := fullmask ^ m
		ways := conn[m] * subg[antimask] % MOD
		for i:=1;i<N;i++ {
			if m & (1<<i) == 0 { continue }
			ans[i] += ways; ans[i] %= MOD
		}
	}
	for i:=1;i<N;i++ { fmt.Fprintln(wrtr,ans[i]) }
}

