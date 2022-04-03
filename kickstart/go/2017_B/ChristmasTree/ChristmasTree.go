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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,K := gi3(); bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }
		for i,j := 0,N-1; i<j; i,j=i+1,j-1 { bd[i],bd[j] = bd[j],bd[i] } //Flip board upside-down for convenience

		// Pass 1 : figure out the largest triangle we can make with each vertex as apex
		dp1 := twodi(N,M,0)
		for i:=0;i<N;i++ {
			for j:=0;j<M;j++ {
				v := 0
				if bd[i][j] == '#' { 
					v = 1 
					if i > 0 && j > 0 && j < M-1 {
						v = 1 + min(min(dp1[i-1][j-1],dp1[i-1][j]),dp1[i-1][j+1])
					}
				}
				dp1[i][j] = v
			}
		}

		// Pass 2 : Do the big DP for the answer
		H := min(N,(M+1)/2)
		dp := make([]int,M*(K+1)*(H+1))
		olddp := make([]int,M*(K+1)*(H+1))
		dp2 := make([]int,M*(K+1))
		olddp2 := make([]int,M*(K+1))
		ldp := len(dp); ldp2 := len(dp2)
		ans := 0
		for i:=0;i<N;i++ {
			dp,olddp = olddp,dp
			dp2,olddp2 = olddp2,dp2
			for i:=0; i<ldp; i++ { dp[i] = 0 }
			for i:=0; i<ldp2; i++ { dp2[i] = 0 }
			for j:=0;j<M;j++ {
				//fmt.Printf("DBG i:%v j:%v\n",i,j)
				for k:=1;k<=K;k++ {
					for h:=1;h<=dp1[i][j];h++ {
						v := 0
						if h == 1 && k == 1 { 
							v = 1
						} else if h == 1 {
							if olddp2[j*(K+1)+(k-1)] > 0 { v = 1 + olddp2[j*(K+1)+(k-1)] }
						} else {
							v1 := olddp[(j-1)*(K+1)*(H+1)+k*(H+1)+(h-1)]
							v2 := olddp[  (j)*(K+1)*(H+1)+k*(H+1)+(h-1)]
							v3 := olddp[(j+1)*(K+1)*(H+1)+k*(H+1)+(h-1)]
							if v1 > 0 || v2 > 0 || v3 > 0 {
								v = max(max(v1,v2),v3) + 2*h - 1
							}
						}
						if k == K && v > ans { ans = v }
						if v > dp2[j*(K+1)+k] { dp2[j*(K+1)+k] = v } 
						dp[j*(K+1)*(H+1)+k*(H+1)+h] = v
					}
				}
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

