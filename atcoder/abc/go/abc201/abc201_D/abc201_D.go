package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2(); A := make([]string,H); for i:=0;i<H;i++ { A[i] = gs() }
	dp := [2000][2000]int{}
	for i:=H-1;i>=0;i-- {
		for j:=W-1;j>=0;j-- {
			if i==H-1 && j==W-1 { dp[i][j] = 0; continue }
			sgn := 1; if (i+j) % 2 == 1 { sgn = -1 }
			delta1,delta2 := sgn,sgn
			if j == W-1 {
				if A[i+1][j] == '-' { delta1 *= -1 }
				dp[i][j] = dp[i+1][j] + delta1
			} else if i == H-1 {
				if A[i][j+1] == '-' { delta2 *= -1 }
				dp[i][j] = dp[i][j+1] + delta2
			} else {
				if A[i+1][j] == '-' { delta1 *= -1 }
				cand1 := dp[i+1][j] + delta1
				if A[i][j+1] == '-' { delta2 *= -1 }
				cand2 := dp[i][j+1] + delta2
				if sgn == 1 { dp[i][j] = max(cand1,cand2) } else {dp[i][j] = min(cand1,cand2) }
			}
		}
	}
	ans := "Takahashi"; if dp[0][0] == 0 { ans = "Draw" }; if dp[0][0] < 0 { ans = "Aoki" }
	fmt.Println(ans)
}



