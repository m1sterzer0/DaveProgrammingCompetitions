package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
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
		W,N := gi2()
		X := gis(W); for i:=0;i<W;i++ { X[i]-- }
		sort.Slice(X,func(i,j int) bool { return X[i] < X[j] } )
		// Special case when all of the values are equal -- helps prevent an infinite wraparound
		ans := 0
		if X[0] != X[W-1] { 
			// Figure out where the calipers go
			minrval := inf; ridx := -1; minlval := inf; lidx := -1; dec := 0; inc := 0; cursum := 0
			for i,v := range X {
				cursum += min(v,N-v)
				if N % 2 == 0 {
					if v < N/2 { dec++ } else { inc++ }
					rval := v; lval := v-N/2; if lval < 0 { lval += N }
					if rval < minrval { minrval = rval;  ridx = i }
					if lval < minlval { minlval = lval;  lidx = i }
				} else {
					if v <= N/2 { dec++ } else { inc++ }
					rval := v; lval := v-(N+1)/2; if lval < 0 { lval += N }
					if rval < minrval { minrval = rval;  ridx = i }
					if lval < minlval { minlval = lval;  lidx = i }
				}
			}
			best := cursum; curs1,curs2 := 0,N/2; if N % 2 == 1 { curs2++ }
			for ridx < W && X[ridx] == curs1 { dec--; inc++; ridx++; if ridx >= W { ridx -= W } }
			for X[lidx] == curs2 { dec++; inc--; if N % 2 == 1 { cursum++ }; lidx++; if lidx >= W { lidx -= W } }
			for curs1 < N-1 {
				dv1 := (X[ridx] + N - curs1) % N
				dv2 := (X[lidx] + N - curs2) % N
				dv3 := N-1-curs1
				dv := min(min(dv1,dv2),dv3)
				curs1 += dv; curs2 += dv
				if curs2 >= N { curs2 -= N }
				cursum += dv*(inc-dec)
				best = min(best,cursum)
				for ridx < W && X[ridx] == curs1 { dec--; inc++; ridx++; if ridx >= W { ridx -= W } }
				for X[lidx] == curs2 { dec++; inc--; if N % 2 == 1 { cursum++ }; lidx++; if lidx >= W { lidx -= W } }
			}
			ans = best
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

