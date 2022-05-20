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
		N,D := gi(),gi()
		minbuckets := N; maxbuckets := N; ans := 0
		for f:=D;f<=N;f+=D {
			// Let n be the number of buckets
			// we need n >= 1 + (N-D)/(D+2),  we also need n <= N/D
			for maxbuckets*f > N { maxbuckets-- }
			for (minbuckets-1)*(f+2) > N-f { minbuckets-- }
			minbuckets = max(1,minbuckets)
			for b:=minbuckets; b<=maxbuckets;b++ {
				r := N - b*f
				bm1 := b-1
				if bm1 == 0 && r == 0 { 
					ans++ 
				} else if bm1 > 0 && r <= 2*bm1 {
					max2s := r/2
					min2s := max(0,r-bm1)
					ans += (max2s-min2s+1)
				}
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

