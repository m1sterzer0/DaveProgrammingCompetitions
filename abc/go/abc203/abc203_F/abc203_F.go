package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	N,K := gi2(); A := gis(N)
	sort.Slice(A,func(i,j int)bool{ return A[i] > A[j] } )
	inf := 1_000_000_000_000_000_000
	dp := twodi(N+1,33,inf)  // dp[i][j] = number of weeds aoki needs to pull to clear out first i weeds in no more than j moves from Takahashi
	dp[0][0] = 0
	ep := 0
	for i,a := range A {
		// Option 1 -- have Aoiki pull weed a
		for j:=0;j<33;j++ { dp[i+1][j] = min(dp[i+1][j],dp[i][j]+1)}
		// Option 2 -- Takahashi can do a cut starting with a as the tallest week
		for ep < N && 2*A[ep] > a { ep++ }
		for j:=0;j<32;j++ { dp[ep][j+1] = min(dp[ep][j+1],dp[i][j]) }
	}
	ansn,ansk := 0,0
	for j:=0;j<33;j++ {	if dp[N][j] <= K { ansn = j; ansk = dp[N][j]; break } }
	fmt.Printf("%v %v\n",ansn,ansk)
}



