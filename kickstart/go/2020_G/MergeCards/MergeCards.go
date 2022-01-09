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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

	// Precalculate all of the coefficients through N*2
	coeff := make([][]float64,5001)
	coeff[2] = append(coeff[2],1.0,1.0)
	for i:=3; i<=5000; i++ {
		dfact := 1.0 / float64(i-1)
		for j:=0; j<i; j++ {
			v := 0.00 
			if j == 0 || j == i-1 { v = dfact * (1.0 + float64(i-1)*coeff[i-1][0] ) }
			if j > 0 && j < i-1   { v = dfact * (2.0 + float64(j) * coeff[i-1][j-1] + float64(i-1-j) * coeff[i-1][j] ) }
			coeff[i] = append(coeff[i],v)
		}
	}

	// Now solution is a simple dot product
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); A := gis(N)
		ans := 0.0
		for i:=0;i<N;i++ { ans += float64(A[i]) * coeff[N][i] }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

