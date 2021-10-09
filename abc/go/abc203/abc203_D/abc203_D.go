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
func check(m int, N int, K int, A [][]int ) bool {
	dp := twodi(N,N,0);
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			dp[i][j] = 0; if A[i][j] <= m { dp[i][j] = 1 }
			if i > 0 { dp[i][j] += dp[i-1][j] }
			if j > 0 { dp[i][j] += dp[i][j-1] }
			if i > 0 && j > 0 { dp[i][j] -= dp[i-1][j-1] }
		}
	}
	for i:=K-1;i<N;i++ {
		for j:=K-1;j<N;j++ {
			cnt := dp[i][j]
			if i-K >= 0 { cnt -= dp[i-K][j] }
			if j-K >= 0 { cnt -= dp[i][j-K] }
			if i-K >= 0 && j-K >= 0 { cnt += dp[i-K][j-K] }
			if 2*cnt >= K*K { return true }
		}
	}
	return false
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,K := gi2(); A := make([][]int,N); for i:=0;i<N;i++ { A[i] = gis(N) }
	l,u := -1,1_000_000_001
	for u-l > 1 { m := (u+l)/2; if check(m,N,K,A) { u = m } else {l = m } }
	fmt.Println(u)
}



