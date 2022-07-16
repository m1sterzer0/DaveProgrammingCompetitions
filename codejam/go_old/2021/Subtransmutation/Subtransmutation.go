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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		/////////////////////////////////////////////////////////////////////////////////////
		// a) Notice that if both A and B are even, and I have both even and odd
		//    numbers in U, then we are screwed.
		// b) More generally, since we only use steps of A,B == ng,mg where g is gcd(A,B),
		//    We must have (as a necessary condition) that start - kg == end -->
		//    start % g == end % g.
		// c) Is this condition sufficient?  Here we turn to offline experiments.  There
		//    are only so many values of A & B, so we can fairly exhaustively cover things
		//    without the time limit.  The routine below suggests 402 is the Maximum
		//    I need.  Thus, we can just set a simulation maximum of 500, and if we
		//    don't succeed, we print IMPOSSIBLE.
		/////////////////////////////////////////////////////////////////////////////////////
		N,A,B := gi3(); U := gis(N)

		tryit := func(v int) bool {
			inv := iai(max(N,v+1),0); inv[v] = 1
			for i:=len(inv)-1;i>=0;i-- {
				if i < N {
					if inv[i] < U[i] { return false }
					inv[i] -= U[i]
				}
				if i - A >= 0 { inv[i-A] += inv[i] }
				if i - B >= 0 { inv[i-B] += inv[i] }
				inv[i] = 0 // Just bookkeeping -- not strictly needed
			}
			return true
		}
		ans := -1
		for i:=0;i<=500;i++ {
			if tryit(i) { ans = i+1; break }
		}
		if ans == -1 { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
    }
}

