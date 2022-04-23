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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func abs(a int) int { if a < 0 { return -a }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
type pair struct {a,b int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	reduce := func(a,b int) pair {
		if a < 0 { a = -a; b = -b }
		if a == 0 { return pair{0,1} }
		if b == 0 { return pair{1,0} }
		g := gcd(a,abs(b))
		return pair{a/g,b/g}
	}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y := fill2(N)
		// Get the list of reduced coefficients (a,b) such that ax+by=0 for the given pairs
		pairset := make(map[pair]bool)
		for i:=0;i<N;i++ {
			for j:=i+1;j<N;j++ {
				aa := Y[j]-Y[i]; bb := X[i]-X[j]
				p := reduce(aa,bb)
				pairset[p] = true
			}
		}
		best := 1  // Can always hit one hole
		for p := range pairset {
			cnts := make(map[int]int)
			for i:=0;i<N;i++ { c := p.a*X[i]+p.b*Y[i]; cnts[c]++ }
			singles,twos,threes := 0,0,0
			for _,v := range cnts { 
				if v == 1 { singles++ } else if v % 2 == 0 { twos += v/2 } else { threes++; twos += (v-3)/2 }
			}
			// Singles    Twos     Threes     Description
            //    >=2     ----      Even      2 + All twos + All 3s
			//    1       ----      Even      1 + All twos + All 3s
			//    0       ----      Even      All Twos + All Threes  (Either split a 2 into 2 singles, or split a pair of 3s to 2+2+1+1)
			//  Therefore, if we have an even number of threes, we always coutn all threes, all twos, and up to 2 singles
			//   >=2      ----      Odd       1 + All Twos + All threes
			//   1        ----      Odd       1 + All Twos + All threes
			//   0        ----      Odd       All Twos + All threes
			//  Therefore, if we have an odd number of threes, we always coutn all threes, all twos, and up to 1 singles
			if threes % 2 == 1 { singles = min(1,singles) } else { singles = min(2,singles) }
			best = max(best,singles+2*twos+3*threes)
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
	}
}

