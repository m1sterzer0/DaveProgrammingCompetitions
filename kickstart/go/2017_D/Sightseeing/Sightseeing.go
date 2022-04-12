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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Ts,Tf := gi3(); S,F,D := fill3(N-1)

		getTime := func(s,f,d,tstart int) int {
			if s >= tstart { return s + d }
			n := (tstart-s+f-1)/f
			return s + n*f + d
		}

		// Now do the base case
		dp,ldp := iai(N,inf),iai(N,inf); dp[0] = 0
		for i:=1;i<N;i++ { dp[i] = getTime(S[i-1],F[i-1],D[i-1],dp[i-1]) }
		if dp[N-1] > Tf {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			// Here is the real DP
			ans := 0
			for i:=1;i<=N-1;i++ {
				ldp,dp = dp,ldp
				for j:=0;j<i;j++ { dp[j] = inf }
				for j:=i;j<N;j++ {
					departTime := min(dp[j-1],ldp[j-1]+Ts)
					dp[j] = getTime(S[j-1],F[j-1],D[j-1],departTime)
				}
				if dp[N-1] > Tf { break }
				ans++
			}
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
	}
}

