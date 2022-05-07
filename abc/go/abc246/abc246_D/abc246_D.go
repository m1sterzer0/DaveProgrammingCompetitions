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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Term is (a+b)*(a^2+b^2) -- result is between 1/2(a+b)^3 and (a+b)^3
	// Instead, say we fix the larger element: a, and solve for the smaller element b
	// If b <= a, then  a^3 = (a+0) * (a^2+0^2) <= (a+b) * (a^2+b^2) <= (2a) * (2a^2) = 4a^3
	// a is no bigger than 10^6, so we can do this with binary search
	N := gi(); best := 1000000000000000000 // (this is realizable with (10^6,0))
	if N == 0 { best = 0 }
	for a:=1;a<=1000000;a++ {
		ac := a*a*a
		if 4*ac < N { continue }
		if ac >= N { best = min(best,ac); break }
		l,u := 0,a
		for u-l > 1 { m := (u+l)>>1; cand := (a+m)*(a*a+m*m); if cand < N { l = m } else { u = m }}
		best = min(best,(a+u)*(a*a+u*u))
	}
	fmt.Println(best)
}

