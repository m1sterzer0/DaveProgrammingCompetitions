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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	adjm := make([][]bool,2*N); for i:=0;i<2*N;i++ { adjm[i] = make([]bool,2*N) }
	for i:=0;i<M;i++ { adjm[A[i]][B[i]] = true; adjm[B[i]][A[i]] = true }
	dp := twodi(2*N,2*N,0)
	comb := twodi(2*N+1,2*N+1,0)
	for i:=0;i<=2*N;i++ { comb[0][i] = 0; comb[i][0] = 1; comb[i][i] = 1 }
	for i:=1;i<=2*N;i++ {
		for j:=1;j<i;j++ { 
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % MOD
		}
	}
	for sz:=2;sz<=2*N;sz+=2 {
		for i:=0;i<2*N;i++ {
			j := i+sz-1
			if j >= 2*N { break }
			if sz == 2 {
				if adjm[i][j] { dp[i][j] = 1 }
			} else {
				for k:=i+1;k<=j;k+=2 {
					if !adjm[i][k] { continue }
					w1,w2 := 1,1
					if k-i > 1 { w1 = dp[i+1][k-1] }
					if k != j { w2 = dp[k+1][j] }
					w3 := comb[sz/2][(k-i+1)/2]
					dp[i][j] += w1 * w2 % MOD * w3 % MOD
				}
				dp[i][j] %= MOD
			}
		}
	}
	fmt.Println(dp[0][2*N-1])
}

