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
	N,X := gi(),gi(); A,B := fill2(N)
	dp := make([]bool,X+1)
	ndp := make([]bool,X+1)
	dp[0] = true
	for i:=0;i<N;i++ {
		a,b := A[i],B[i]
		for j:=0;j<=X;j++ { ndp[j] = false }
		for _,d := range []int{a,b} {
			for j:=0;j+d<=X;j++ { if dp[j] { ndp[j+d] = true } }
		}
		dp,ndp = ndp,dp
	}
	if dp[X] { fmt.Println("Yes") } else { fmt.Println("No") }
}

