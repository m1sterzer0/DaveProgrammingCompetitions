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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func min(a,b int) int { if a > b { return b }; return a }
const inf = 2000000000000000000
type edge struct { idx,n2,d int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); A,B,C := fill3(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	darr := twodi(N,N,inf)
	for i:=0;i<N;i++ { darr[i][i] = 0 }
	for i:=0;i<M;i++ { a,b,c := A[i],B[i],C[i]; darr[a][b] = min(darr[a][b],c); darr[b][a] = min(darr[b][a],c) }
	for k:=0;k<N;k++ {
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				darr[i][j] = min(darr[i][j],darr[i][k]+darr[k][j])
			}
		}
	}
	// * We can remove an edge iff it doesn't affect the distance matrix
	// * If an edge has longer length than the shortest path between a,b, then it can be safely removed
	// * If we can find a path a --> k --> b in the distance matrix for the same distance, then the edge is also redundant
	ans := 0
	for i:=0;i<M;i++ {
		a,b,c := A[i],B[i],C[i]
		redundant := false
		for k:=0;k<N;k++ {
			if k == a || k == b { continue }
			if darr[a][k]+darr[k][b] <= c { redundant = true }
		}
		if redundant { ans++ }
	}
	fmt.Println(ans)
}
