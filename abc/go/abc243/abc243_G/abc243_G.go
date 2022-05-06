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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func isqrt(x int) int {
    if x == 0 { return 0 }
    s := int(math.Sqrt(float64(x)))
    s = (s + x/s) >> 1
        if s*s > x { return s-1 } else { return s }
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Prework, find the number of suffixes starting with x for all x up to 100000
	cntarr := make([]int,100001)
	cumarr := make([]int,100001)
	cntarr[1] = 1; cumarr[1] = 1
	for i:=2;i<=100000;i++ {
		j := isqrt(i)
		cntarr[i] = cumarr[j]
		cumarr[i] = cumarr[i-1] + cntarr[i]
	}
	T := gi()
	for tt:=1;tt<=T;tt++ {
		X := gi()
		y := isqrt(X)
		ans := 0
		for i:=1;i<=100000;i++ { // Looping on candidates for maximum 3rd number
			l,r := i*i,min(i*i+2*i,y)
			if l > r { break }
			w := r-l+1           // Calculating how many choices for 2nd number have the same set of choices for 3rd number
			ans += w * cumarr[i]
		}
		fmt.Fprintln(wrtr,ans)
	}
}
