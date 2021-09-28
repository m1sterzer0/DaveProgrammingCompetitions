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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
const MOD int = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); ans := 1 // Only one way for one full segment
	cumsum := iai(N,0); s := 0; for i,a := range A { s += a; cumsum[i] = s }
	dp,ndp,ways := iai(N,1),iai(N,0),iai(N+1,0)
	for i:=2;i<=N;i++ { // Iterate over total number of segments
		for j:=0;j<i;j++ { ways[j] = 0 }
		for j:=0;j<N;j++ {
		    // ways[j] is the number of ways to partition some proper prefix the sequence A[0:j]
			//         into i-1 chunks such that the sum of the elements in that prefix % i is j
			idx := cumsum[j] % i
			ndp[j] = ways[idx]
			ways[idx] += dp[j]; ways[idx] %= MOD
		}
		dp,ndp = ndp,dp
		ans += dp[N-1]; ans %= MOD
	}
	fmt.Println(ans)
}

