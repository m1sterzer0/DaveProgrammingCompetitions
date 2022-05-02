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
func next_permutation(a []int) bool {
	la := len(a); var i,j int
	for i=la-2;i>=0;i-- { if a[i] < a[i+1] { break } }
	if i<0 { i,j = 0,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- } ; return false }
	for j=la-1;j>=0;j-- { if a[i] < a[j] { break } }
	a[i],a[j] = a[j],a[i]
	i,j = i+1,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- }
	return true
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); A,B := fill2(M); C,D := fill2(M)
	for i:=0;i<M;i++ { A[i]--; B[i]--; C[i]--; D[i]-- }
	ans := "No"
	adj := make([][]bool,N); for i:=0;i<N;i++ { adj[i] = make([]bool,N) }
	P := make([]int,0); for i:=0;i<N;i++ { P = append(P,i) }
	for i:=0;i<M;i++ { adj[A[i]][B[i]] = true; adj[B[i]][A[i]] = true }
	for {
		good := true
		for i:=0;i<M;i++ { c,d := P[C[i]],P[D[i]]; if !adj[c][d] { good = false; break } }
		if good { ans = "Yes"; break }
		if !next_permutation(P) { break }
	}
	fmt.Println(ans)
}
