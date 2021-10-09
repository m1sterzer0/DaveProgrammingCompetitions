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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func min(a,b int) int { if a > b { return b }; return a }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W,C := gi3(); A := make([][]int,H)
	for i:=0;i<H;i++ { A[i] = gis(W) }
	inf := powint(10,18); dp1,dp2 := twodi(H,W,inf),twodi(H,W,inf)
	// Pass 1, looking up and to the left
	ans := inf
	for i:=0;i<H;i++ {
		for j:=0;j<W;j++ {
			best := inf
			if i > 0 { best = min(best,C + min(dp1[i-1][j],A[i-1][j])) }
			if j > 0 { best = min(best,C + min(dp1[i][j-1],A[i][j-1])) }
			dp1[i][j] = best
			ans = min(ans,A[i][j]+best)
		}
		for j:=W-1;j>=0;j-- {
			best := inf
			if i > 0   { best = min(best,C + min(dp2[i-1][j],A[i-1][j])) }
			if j < W-1 { best = min(best,C + min(dp2[i][j+1],A[i][j+1])) }
			dp2[i][j] = best
			ans = min(ans,A[i][j]+best)
		}
	}
	fmt.Println(ans)
}

