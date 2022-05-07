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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); A := gis(N); B := gis(N)
	dpa,dpb := true,true
	for i:=1;i<N;i++ {
		ndpa := dpa && abs(A[i]-A[i-1]) <=K || dpb && abs(A[i]-B[i-1]) <= K
		ndpb := dpa && abs(B[i]-A[i-1]) <=K || dpb && abs(B[i]-B[i-1]) <= K
		dpa,dpb = ndpa,ndpb
	}
	ans := "No"; if dpa || dpb { ans = "Yes" }; fmt.Println(ans)
}

