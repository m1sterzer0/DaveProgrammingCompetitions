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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W,K := gi3()
	A := make([][]int,H)
	for i:=0;i<H;i++ { A[i] = gis(W) }
	m := make(map[int]bool); for i:=0;i<H;i++ { for j:=0;j<W;j++ { m[A[i][j]] = true } }
	vals := []int{}; for k := range m { vals = append(vals,k) }
	dp := [61][30][30]int{}
	myinf := 1_000_000_000_000_000_000
	ans := myinf
	for _,v := range vals {
		for i:=0;i<61;i++ { for j:=0;j<30;j++ { for k:=0;k<30;k++ { dp[i][j][k] = myinf } } }
		for i:=0;i<=K;i++ {
			for j:=0;j<H;j++ {
				for k:=0;k<W;k++ {
					if j == 0 && k == 0 {
						if i == 0 && A[j][k] <= v { dp[i][j][k] = 0 }
						if i == 1 && A[j][k] >= v { dp[i][j][k] = A[j][k] }
					} else {
						if i > 0 && j > 0 && A[j][k] >= v { dp[i][j][k] = min(dp[i][j][k], dp[i-1][j-1][k] + A[j][k]) }
						if i > 0 && k > 0 && A[j][k] >= v { dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k-1] + A[j][k]) }
						if j > 0 && A[j][k] <= v { dp[i][j][k] = min(dp[i][j][k],dp[i][j-1][k]) }
						if k > 0 && A[j][k] <= v { dp[i][j][k] = min(dp[i][j][k],dp[i][j][k-1]) }
					}
				}
			}
		}
		ans = min(ans,dp[K][H-1][W-1])
	} 
	fmt.Println(ans)
}

