package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); A := gis(N); ans := 0
		if N <= 3 {
			ans = N
		} else {
			// dp[i] is the longest arithmetic sequence that starts at index i
			dp := ia(N); dp[N-1] = 1; dp[N-2] = 2
			for i:=N-3;i>=0;i-- {
				if A[i+1] - A[i] == A[i+2]-A[i+1] { dp[i] = 1 + dp[i+1] } else { dp[i] = 2 }
			}
			ans = 3
			for i:=0;i<N-2;i++ { // loop over new starting point
				iend := i + dp[i]
				// Change ourselves to be compatible with the one in front of us
				ans = max(ans,dp[i+1]+1)
				// Change the second one to be compatible with the arithmetic sequence starting 2 ahead of us.
				if dp[i] == 2 && i+3 < N && A[i+2]-A[i] == 2*(A[i+3]-A[i+2]) { ans = max(ans,dp[i+2]+2) } 
				// Change the one just past our sequence and make it compatible with ourself
				if iend < N { ans = max(ans,dp[i]+1) }
				// Change the one just past our sequence and see if we pick up 2
				if (iend+1) < N && A[iend+1] - A[iend-1] == 2 * (A[i+1]-A[i]) { ans = max(ans,dp[i]+2) }
				// Change the one just past our sequence and see if we pick up the next sequence
				if (iend+2) < N && A[iend+1] - A[iend-1] == 2 * (A[i+1]-A[i]) && A[iend+2]-A[iend+1] == A[i+1]-A[i] { ans = max(ans,dp[i]+1+dp[iend+1]) }
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

