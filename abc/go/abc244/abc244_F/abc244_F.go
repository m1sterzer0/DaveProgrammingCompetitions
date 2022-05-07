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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }
const inf = 2000000000000000000
type st struct { bm,n int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
	darr := twodi(1<<uint(N),N,-1)
	adj := make([][]bool,N); for i:=0;i<N;i++ { adj[i] = make([]bool,N) }
	for i:=0;i<M;i++ { u,v := U[i],V[i]; adj[u][v] = true; adj[v][u] = true }
	darr[0][0] = 0; q := make([]st,0,2300000)
	for i:=0;i<N;i++ { bm := 1<<uint(i); darr[bm][i] = 1; q = append(q,st{bm,i}) }
	for len(q) > 0 {
		s := q[0]; q = q[1:]; bm,n := s.bm,s.n
		for i:=0;i<N;i++ {
			if !adj[n][i] { continue }
			nbm := bm ^ (1<<uint(i))
			if darr[nbm][i] >= 0 { continue }
			darr[nbm][i] = darr[bm][n] + 1
			q = append(q,st{nbm,i})
		}
	}
	ans := 0
	for bm:=0;bm<1<<uint(N);bm++ {
		best := inf
		for n:=0;n<N;n++ {
			if darr[bm][n] == -1 { continue }
			best = min(best,darr[bm][n])
		}
		if best == inf { fmt.Fprintf(os.Stderr,"ERROR: bm:%017b\n",bm) }
		ans += best
	}
	fmt.Println(ans)
}

