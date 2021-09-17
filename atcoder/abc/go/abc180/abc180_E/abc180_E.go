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
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func abs(a int) int { if a >= 0 { return a }; return -a }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); X,Y,Z := fill3(N)
	inf := 2_000_000_000_000_000_000
	bmmax := powint(2,N)
	dp := twodi(N,bmmax,inf)
	dp[0][1] = 0
	for m := 3; m < bmmax; m += 2 {
		for i:=1;i<N;i++ {
			if (1<<i) & m == 0 { continue }
			m2 := m ^ (1 << i)
			for j:=0;j<N;j++ {
				if (1 << j) & m2 == 0 { continue }
				d := dp[j][m2]
				if d == inf { continue }
				dp[i][m] = min(dp[i][m],d+abs(X[i]-X[j]) + abs(Y[i]-Y[j]) + max(0,Z[i]-Z[j]))
			}
		}
	}
	ans := inf
	for i:=1;i<N;i++ { ans = min(ans,dp[i][bmmax-1] + abs(X[0]-X[i]) + abs(Y[0]-Y[i]) + max(0,Z[0]-Z[i])) }
	fmt.Println(ans)
}



