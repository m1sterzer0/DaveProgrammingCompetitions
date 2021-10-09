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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi()
	m := 2*N; u:=ia(N+10); u[0] = 1; inv := iai(N+10,1)
	// For a tree to be beautiful, we need 
	// Root and leaves to have 1's written on them
	// Sum of all nodes is N
	// Not having two consecutive 0s -- two consecutive zeros means it takes more than N-1 steps to aggregate the 1s to the root.
	// TODO -- UNDERSTAND THE SOLUTION -- right now, i'm just translating the solution in the editorial to go
	u[1] = u[0]*3*(m+1-1) % MOD * inv[1] % MOD 
	inv[2] = MOD - inv[MOD % 2] * (MOD / 2) % MOD
	for k:=2;k<N;k++ {
		u[k] = (u[k-1]*3*(m+1-k) + u[k-2]*(2*m+2-k)) % MOD * inv[k] % MOD
		inv[k+1] = MOD - inv[MOD % (k+1)] * (MOD / (k+1)) % MOD
	}
	fmt.Println(u[N-1] * inv[N] % MOD)
}

