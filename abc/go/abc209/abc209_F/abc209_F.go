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
const MOD int = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); H := gis(N)
	dp,cdp,ndp,ncdp := ia(N),ia(N),ia(N),ia(N)
	dp[0],cdp[0] = 1,1
	for i:=1;i<N;i++ {
		cumsum := 0
		for j:=0;j<=i;j++ {
			if H[i] == H[i-1] { 
				ndp[j] = cdp[i-1]
			} else if H[i] < H[i-1]  {
				ndp[j] = cdp[i-1]
				if j > 0 { ndp[j] -= cdp[j-1]; if ndp[j] < 0 { ndp[j] += MOD } }
			} else {
				ndp[j] = 0
				if j > 0 { ndp[j] += cdp[j-1] }
			}
			cumsum += ndp[j]; if cumsum > MOD { cumsum -= MOD }
			ncdp[j] = cumsum
		}
		dp,cdp,ndp,ncdp = ndp,ncdp,dp,cdp
	}
	fmt.Println(cdp[N-1])
}

