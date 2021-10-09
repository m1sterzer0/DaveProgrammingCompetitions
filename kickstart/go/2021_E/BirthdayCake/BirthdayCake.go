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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	//experiment()
	T := gi()
    for tt:=1;tt<=T;tt++ {	
		// PROGRAM STARTS HERE
		// 1) Perimeter, just brute force the possibilities from left edge, right edge, top edge, and bot edge
		// 2) Internals, observations
		//    a) once the cake gets down to LxW s.t. L <= K and W <= K, every cut will chop the cake into one more piece, so we will have L*W-1 cuts
		//    b) once the cake gets down so that the short dimension is <= K, also every cut will chop the cake into one more piece, so we will have L*W-1 cuts
		//    c) It believe that if we view a large cake as an array of K*K cakes with some leftovers on the edges, we can just chop the board into the
		//       K x K cakes, and then deal with each K*K cake as per (a) and (b) above.
		//    d) Taking one step further, we say that if the length is larger than l, we can just chop K off of the length with a transverse cut and 
		//       iterate on the two pieces.
		//    e) We can validate our intuition with a brute force that does a DP on all possible ways to divide the rectange into 2 (I can't immediately see
		//       why you would want to do anything more complex than a full vertical cut across the rectangle).
		solvePerimeter := func(R,C,r1,c1,r2,c2,k int) int {
			l := r2-r1+1; w := c2-c1+1
			// All 4 sides are free
			if r1 == 1 && r2 == R && c1 == 1 && c2 == C { return 0 }
			// 3 sides are free
			if (r1== 1 || r2 == R) && c1 ==1 && c2 == C { return (w+k-1)/k }
			if r1== 1 && r2 == R && (c1 ==1 || c2 == C) { return (l+k-1)/k }
			// 2 parallel sides are free
			if c1 ==1 && c2 == C { return (w+k-1)/k * 2 }
			if r1 ==1 && r2 == R { return (l+k-1)/k * 2 }
			// 2 perpendicular sides are free
			if (r1== 1 || r2 == R) && (c1 ==1 || c2 == C) { return (w+k-1)/k + (l+k-1)/k }
			// 1 side is free
			if c1 == 1 || c2 == C { return (w+k-1)/k * 2 + (l+k-1)/k }
			if r1 == 1 || r2 == R { return (w+k-1)/k     + (l+k-1)/k * 2}
			// Now we have to cut in from the side
			ans :=          (c2+k-1)/k     + (w+k-1)/k     + (l+k-1)/k * 2
			ans = min(ans,  (C-c1+1+k-1)/k + (w+k-1)/k     + (l+k-1)/k * 2)
			ans = min(ans,  (r2+k-1)/k     + (w+k-1)/k * 2 + (l+k-1)/k)
			ans = min(ans,  (R-r1+1+k-1)/k + (w+k-1)/k * 2 + (l+k-1)/k)
			return ans
		}
		solveInternal := func(l,w,k int) int {
			if l % k == 0 { return (l/k - 1) * ((w+k-1)/k) + (l/k) * (w*k-1) }
			return l/k * ((w+k-1)/k + (w*k-1)) + (l%k) * w - 1
		}
		R,C,K := gi3(); r1,c1,r2,c2 := gi4()
		ans := 0
		ans += solvePerimeter(R,C,r1,c1,r2,c2,K)
		ans += solveInternal(r2-r1+1,c2-c1+1,K)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

