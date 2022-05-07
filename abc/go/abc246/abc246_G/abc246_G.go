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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := make([]int,N); for i:=1;i<N;i++ { A[i] = gi() }; U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	// Binsearch on score and instead solve the 0-1 problem
	// Keep invariant of "how many sqaures do I need to eliminate BEFORE reaching this tree to make this tree non-winning"
	// Check answer for root
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	l,u,m := 0,1000000001,0
	var dfs func(n,p int) int
	dfs = func(n,p int) int {
		res := 0
		for _,c := range gr[n] { if c != p { res += dfs(c,n) } } // Must make all children non-winning to make position non-winning
		if res > 0 { res-- }                                     // This if for the one move I get once I get to this tree
		if A[n] >= m { res++ }                                   // If current square is good, we need to eliminate it before we get here
		return res
	}
	for u-l > 1 { m = (u+l)>>1; r := dfs(0,-1); if r == 0 { u = m } else { l = m } }
	fmt.Println(l)
}

