package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func solveDp(N int) float64 {
	dp := make([]float64,N+1)
	dp[N] = 0; cumsum := 0.00
	for i:=N-1;i>=0;i-- { dp[i] = 1 + cumsum/float64(N-i); cumsum += dp[i] }
	return dp[0]
}
func solve(N int) float64 {
	// Answer is just 1 + 1/2 + 1/3 + 1/4 + 1/5 + ... + 1/N (https://en.wikipedia.org/wiki/Harmonic_number)
	// Harmonic numbers are approximated by ln(N) + gamma + 1/(2*n) - 1/(12*n^2) + 1/(120*n^4) - 1/(252*n^6) + 1/(240*n^8) - 5/(660*n^10) + 691/(32760*n^12)
	if N <= 10 { return solveDp(N) }
	fn := float64(N)
	gamma := 0.57721566490153286060651209 // https://en.wikipedia.org/wiki/Euler%27s_constant
	ans := math.Log(fn) + gamma
	pn := fn;      ans += 1.0/2.0/pn
	pn *= fn;      ans -= 1.0/12.0/pn
	pn *= fn * fn; ans += 1.0/120.0/pn
	pn *= fn * fn; ans -= 1.0/252.0/pn
	pn *= fn * fn; ans += 1.0/240.0/pn
	pn *= fn * fn; ans -= 5.0/660.0/pn
	pn *= fn * fn; ans += 691.0/32760.0/pn
	return ans
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		ans := solve(N)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

