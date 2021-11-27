package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,S,T,A,B := gi(),gi(),gi(),gi(),gi()
	// We surmise that there is some region near the goal where we want to just use +1, but outside of that, we should be rerolling.
	// We need to solve for the "best" reroll region.  Ignore the endpoint constraint at first.
	// For a given interval [T-i+1,T-i+2,... T], the EV is B * (N/i) + A * (i-1)/2.
	// Setting derivative equal to zero, this is minimized at i == sqrt(2NB/A)
	ans := 0.0
	if S == T {
		ans = 0.0 
	} else if T == 1 {
		ans = float64(B)*float64(N)
	} else {
		ii := int(math.Sqrt(2.0*float64(N)*float64(B)/float64(A)))
		bestev,besti := 1e99,-1
		for i:=ii-10000;i<=ii+10000;i++ {
			if i < 0 { continue }
			cand := float64(B) * float64(N)/float64(i) + float64(A) * 0.5 * float64(i-1)
			if cand < bestev { bestev = cand; besti = i } 
		}
		if besti > T { besti = T; bestev = float64(B) * float64(N)/float64(besti) + float64(A) * 0.5 * float64(besti-1) }
		if S > T { 
			ans = bestev
		} else {
			ans = float64(A) * float64(T-S)
			if ans > bestev { ans = bestev }
		}
	}
	fmt.Println(ans)
}

