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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
const MOD int = 1_000_000_007
func convolve(a,b []int) []int {
	c := iai(len(a)+len(b)-1,0)
	for i,a := range a { for j,b := range b { c[i+j] += a*b % MOD } }
	for i:=0;i<len(c);i++ { c[i] %= MOD }
	return c
}
func rightshift(a []int) []int { return append([]int{0},a[:len(a)-1]...) }
func arrsum2(a,b []int) []int { c := iai(len(a),0); for i:=0;i<len(a);i++ { c[i] = (a[i]+b[i]) % MOD }; return c }
func arrsum3(a,b,c []int) []int { d := iai(len(a),0); for i:=0;i<len(a);i++ { d[i] = (a[i]+b[i]+c[i]) % MOD }; return d }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	var dfs func(n,p int) ([]int,[]int,[]int)
	
	dfs = func(n,p int) ([]int,[]int,[]int) {
		// dp0 are cases where I don't have a takahashi, and none of my children have a takahashi
		// dp1 are cases where I don't have a takahashi, but at least one of my children has a takahashi
		// dp2 are cases where I have a takahashi
		// Transitions
		// -- lets say a child has dp0c,dp1c,dp2c
		// -- dp0new = convolve ( dp0old , dp0c + dp1c )
		// -- dp1new = rightshift ( convolve (dp0old, dp2c) ) + convolve ( dp1old, dp0c + dp1c + dp2c )
		// -- dp2new = convolve(dp2old, rightshift (dp0c) + dp1c + dp2c )
		dp0,dp1,dp2 := []int{1,0},[]int{0,0},[]int{0,1}
		for _,c := range gr[n] {
			if c == p { continue }
			dp0c,dp1c,dp2c := dfs(c,n)
			newdp0 := convolve(dp0,arrsum2(dp0c,dp1c))
			newdp1a := rightshift(convolve(dp0,dp2c))
			newdp1b := convolve(dp1,arrsum3(dp0c,dp1c,dp2c))
			newdp1 := arrsum2(newdp1a,newdp1b)
			newdp2 := convolve(dp2,arrsum3(rightshift(dp0c),dp1c,dp2c))
			dp0,dp1,dp2 = newdp0,newdp1,newdp2
		}
		return dp0,dp1,dp2
	} 
	dp0,dp1,dp2 := dfs(0,-1)
	for i:=0;i<=N;i++ {	fmt.Fprintln(wrtr, (dp0[i]+dp1[i]+dp2[i]) % MOD) }
}

