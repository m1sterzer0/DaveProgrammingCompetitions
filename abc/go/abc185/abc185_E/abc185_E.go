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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gi2() (int,int) { return gi(),gi() }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func min(a,b int) int { if a > b { return b }; return a }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); A := gis(N); B := gis(M); inf := 1_000_000_000_000_000_000
	// Standard edit distance DP
	dp := twodi(N+1,M+1,inf) 
	for i:=0;i<=N;i++ { dp[i][0] = i }
	for j:=0;j<=M;j++ { dp[0][j] = j }
	for i:=1;i<=N;i++ {
		for j:=1;j<=M;j++ {
			cand := inf
			if A[i-1] == B[j-1] { cand = min(cand,dp[i-1][j-1]) } else { cand = min(cand,1+dp[i-1][j-1]) } //keep
			cand = min(cand,1+dp[i-1][j])
			cand = min(cand,1+dp[i][j-1])
			dp[i][j] = cand
		}
	}
	ans := dp[N][M]; fmt.Println(ans)
}
