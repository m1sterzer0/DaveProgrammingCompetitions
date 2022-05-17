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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func min(a,b int) int { if a > b { return b }; return a }
type pt struct { i,m int }
const inf int = 2000000000000000000
const MOD int = 1000000007

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,L := gi(),gi(),gi()
		A,B,P := fill3(N)
		dp,ndp := iai(L+1,inf),iai(L+1,inf); dp[0] = 0
		q := make([]pt,0)
		for i:=0;i<N;i++ {
			a,b,p := A[i],B[i],P[i]
			q = q[:0]; lq := -1
			for i,d := range dp { ndp[i] = d }
			for j:=a;j<=L;j++ {
				l,r := j-b,j-a
				// Sliding range minimum query using a stack
				for lq >= 0 && q[lq].m >= dp[r] { q = q[:lq]; lq-- }
				q = append(q,pt{r,dp[r]}); lq++
				for q[0].i < l { q = q[1:]; lq-- }
				ndp[j] = min(ndp[j],q[0].m+p)
			}
			dp,ndp = ndp,dp
		}
		if dp[L] > M { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,dp[L])
		}
    }
}

