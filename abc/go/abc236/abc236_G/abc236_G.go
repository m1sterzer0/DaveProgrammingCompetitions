package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func vecintstring(a []int) string { 
	astr := make([]string,len(a));
	for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ")
}
const inf = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Let dp[i][u][v] be the minimum time needed to get from 1 to v in exactly i steps
	// If we try to compute this linearly we have
	// dp[i+1][u][v] = min_over_w [ max (dp[i][u][w],dp[1][w][v]) ]
	// This takes O(N^3) to calculate, but we can create the powers of two and then use the binary
	// representation of L to assemble (log L) of them into the final solution.
	N,T,L := gi(),gi(),gi(); U,V := fill2(T); for i:=0;i<T;i++ { U[i]--; V[i]-- }
	dp := make(map[int][][]int)

	makearr := func(a,b [][]int) [][]int {
		c := make([][]int,N); for i:=0;i<N;i++ { c[i] = make([]int,N) }
		for i:=0;i<N;i++ { for j:=0;j<N;j++ { c[i][j] = inf } }
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				for k:=0;k<N;k++ {
					v := max(a[i][k],b[k][j])
					if v < c[i][j] { c[i][j] = v }
				}
			}
		}
		return c
	}
	for i:=1;i<=L;i*=2 {
		if i > 1 {
			dp[i] = makearr(dp[i>>1],dp[i>>1])
		} else {
			vv := make([][]int,N); for i:=0;i<N;i++ { vv[i] = make([]int,N) }
			for i:=0;i<N;i++ { for j:=0;j<N;j++ { vv[i][j] = inf} }
			for i:=0;i<T;i++ {
				u,v,t := U[i],V[i],i+1
				if t < vv[u][v] { vv[u][v] = t }
			}
			dp[1] = vv
		}
	}
	state := make([][]int,N); for i:=0;i<N;i++ { state[i] = make([]int,N) }
	for i:=0;i<N;i++ { for j:=0;j<N;j++ { state[i][j] = inf } }
	state[0][0] = 0
	for i,l:=1,L;l>0;i,l=i*2,l>>1 {
		if l & 1 == 1 { state = makearr(state,dp[i]) }
	}
	ansarr := make([]int,N)
	for i:=0;i<N;i++ { 
		if state[0][i] == inf { ansarr[i] = -1 } else { ansarr[i] = state[0][i] }
	}
	ansstr := vecintstring(ansarr)
	fmt.Println(ansstr)
}
