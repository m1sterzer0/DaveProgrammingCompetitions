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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); B := gis(N)
	dp := ia(3001); ndp := ia(3001); cdp := iai(3001,1); dp[0] = 1; 
	for i:=0;i<N;i++ {
		for j:=0;j<=3000;j++ { ndp[j] = 0 }
		for j:=A[i]; j<=B[i]; j++ {	ndp[j] = cdp[j] % MOD }
		dp,ndp = ndp,dp
		for j:=0;j<=3000;j++ { if j == 0 { cdp[j] = dp[j] } else { cdp[j] = cdp[j-1] + dp[j] } }
	}
	fmt.Println(cdp[3000] % MOD)
}

