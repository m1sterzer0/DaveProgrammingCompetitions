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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); B := make([][]int,N); for i:=0;i<N;i++ { B[i] = gis(M) }
	for i:=0;i<N;i++ { for j:=0;j<M;j++ { B[i][j]-- } }
	// 3 checks
	// a) First row is subset of a legal row
	// b) All numbers not in first column are one more than their left neighbor
	// c) All numbers not in the first row are one more than their upper neighbor
	ans := "Yes"
	if B[0][0] / 7 != B[0][M-1] / 7 { ans = "No" }
	for i:=0;i<N;i++ { for j:=1;j<M;j++ { if B[i][j] != B[i][j-1] + 1 { ans = "No" } } }
	for i:=1;i<N;i++ { for j:=0;j<M;j++ { if B[i][j] != B[i-1][j] + 7 { ans = "No" } } }
	fmt.Println(ans)
}