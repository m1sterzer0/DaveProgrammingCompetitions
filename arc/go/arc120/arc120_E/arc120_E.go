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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N)
	try := func(k int) bool {
		dplr,dprl := A[0]+k,A[0]+k
		for i,a := range A {
			if i == 0 { continue }
			ndplr,ndprl,pa := -1,-1,A[i-1]
			// Case 1, prev guy goes rl, and we go lr
			if dprl >= (a+pa)/2 { ndplr = max(ndplr,pa+k) }
			// Case 2, prev guy goes lr, and we go rl
			if dplr >= 0 {
				ldist := (pa+k-dplr)/2
				if ldist + (a-pa)/2 <= k { ndplr = max(ndplr,pa+k-2*ldist) }
			}
			// Case 3, prev guy goes lr, and we go rl
			if dplr >= 0 && a-dplr <= k { ndprl = max(ndprl,a+(k-(a-dplr))/2) }
			// Case 4, prev guy goes rl, and we go rl
			if dprl >= (a+pa)/2 { ndprl = max(ndprl,a+(dprl-pa-(a-dprl))/2) }
			// Check
			if ndprl < 0 && ndplr < 0 { return false }
			dprl,dplr = ndprl,ndplr
		}
		return true
	}
	l,u := 0,1000000001
	for u-l > 1 {
		m := (l+u)>>1
		if try(m) { u = m } else { l = m }
	}
	fmt.Println(u)
}

