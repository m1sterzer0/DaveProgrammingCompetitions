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
func twodi(n int,m int,v int) [][]int { 
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	R,C,K := gi(),gi(),gi()
	RR,CC,VV := fill3(K); for i:=0;i<K;i++ { RR[i]--; CC[i]-- }
	bd := twodi(R,C,0)
	for i:=0;i<K;i++ { bd[RR[i]][CC[i]] = VV[i] }
	dp,ndp := twodi(C,4,0),twodi(C,4,0)
	// Letting you come in anywhere along the top row isn't going to hurt
	for i:=0;i<R;i++ {
		for j:=0;j<C;j++ {
			for k:=0;k<=3;k++ { ndp[j][k] = 0 }
			if j > 0 { for k:=0;k<=3;k++ { ndp[j][k] = ndp[j-1][k] } }
			bestAbove := max(max(dp[j][0],dp[j][1]),max(dp[j][2],dp[j][3]))
			ndp[j][0] = max(ndp[j][0],bestAbove)
			if bd[i][j] > 0 {
				itemval := bd[i][j]
				ndp[j][1] = max(ndp[j][1],bestAbove+itemval)
				if j > 0 { for k:=0;k<=2;k++ { ndp[j][k+1] = max(ndp[j][k+1],ndp[j-1][k]+itemval) } }
			}
		}
		dp,ndp = ndp,dp
	}
	ans := max(max(dp[C-1][0],dp[C-1][1]),max(dp[C-1][2],dp[C-1][3]))
	fmt.Println(ans)
}



