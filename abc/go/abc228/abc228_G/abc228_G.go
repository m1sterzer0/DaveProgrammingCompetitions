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
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W,N := gi3(); C := make([]string,H); for i:=0;i<H;i++ { C[i] = gs() }
	dp := ia(1<<H); ndp := ia(1<<W); dp[(1<<H)-1] = 1
	// For each set of rows/cols, precalculate the set of cols/rows you can end up at if you want to add a given number
	rowset2colset := twodi(1<<H,10,0)
	for rbm:=1;rbm < 1<<H;rbm++ {
		for i:=0;i<H;i++ {
			if rbm & (1<<i) == 0 { continue }
			for j:=0;j<W;j++ {
				c := C[i][j] - '0'
				rowset2colset[rbm][c] |= (1<<j)
			}
		}
	}
	colset2rowset := twodi(1<<W,10,0)
	for cbm:=1;cbm < 1<<W;cbm++ {
		for j:=0;j<W;j++ {
			if cbm & (1<<j) == 0 { continue }
			for i:=0;i<H;i++ {
				c := C[i][j] - '0'
				colset2rowset[cbm][c] |= (1<<i)
			}
		}
	}

	for i:=0;i<2*N;i++ {
		if i % 2 == 0 {
			for i:=0;i<(1<<W);i++ { ndp[i] = 0 }
			for i:=1;i<(1<<H);i++ {
				dp[i] %= MOD
				for c:=1;c<=9;c++ {
					colset := rowset2colset[i][c]
					if colset == 0 { continue }
					ndp[colset] += dp[i]
				}
			}
		} else {
			for i:=0;i<(1<<H);i++ { ndp[i] = 0 }
			for i:=0;i<(1<<W);i++ {
				dp[i] %= MOD
				for c:=1;c<=9;c++ {
					rowset := colset2rowset[i][c]
					if rowset == 0 { continue }
					ndp[rowset] += dp[i]
				}
			}
		}
		dp,ndp = ndp,dp
	}
	ans := 0
	for i:=1;i<(1<<H);i++ { 
		dp[i] %= MOD
		ans += dp[i]
	}
	ans %= MOD
	fmt.Println(ans)
}

