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
func max(a,b int) int { if a > b { return a }; return b }
func maxf(a,b float64) float64 { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N)
	var dp1 [2][100010]int
	var dp2 [2][100010]float64
	// Do the average code
	lavg,ravg := 0.00,1000000000.1
	for ravg-lavg > 0.0001 {
		mavg := 0.5*(lavg+ravg)
		dp2[0][0] = 0.00; dp2[1][0] = float64(A[0])-mavg
		for i:=1;i<N;i++ {
			a := float64(A[i])
			dp2[0][i] = dp2[1][i-1]; dp2[1][i] = (a-mavg) + maxf(dp2[0][i-1],dp2[1][i-1])
		}
		score := maxf(dp2[0][N-1],dp2[1][N-1])
		if score >= 0 { lavg = mavg } else { ravg = mavg }
	}
	fmt.Println(0.5*(lavg+ravg))
	// Do the median code
	lmed,rmed := 0,1000000001
	for rmed-lmed > 1 {
		m := (lmed+rmed)>>1
		dp1[0][0] = 0; dp1[1][0] = 0; if A[0] < m { dp1[1][0] = -1 } else { dp1[1][0] = 1}
		for i:=1;i<N;i++ { 
			a := A[i]; adder := -1; if a >= m { adder = 1 } 
			dp1[0][i] = dp1[1][i-1]; dp1[1][i] = adder + max(dp1[0][i-1],dp1[1][i-1])
		}
		score := max(dp1[0][N-1],dp1[1][N-1])
		if score <= 0 { rmed = m } else { lmed = m }
	}
	fmt.Println(lmed)
}

