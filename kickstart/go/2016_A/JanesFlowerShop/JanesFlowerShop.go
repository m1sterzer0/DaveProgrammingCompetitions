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
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		M := gi(); C := gis(M+1);
		C[0] *= -1
		for C[len(C)-1] == 0 { C = C[:len(C)-1] } // Chop off any zeros at the end
		// Now we should have a negative number in the first slot and a positive number in the last slot
		calcsum := func(m float64) float64 {
			f := 1.00; s := 0.00
			for i:=len(C)-1; i>=0; i-- { s += float64(C[i])*f; f *= (1.0+m) }
			return s
		}
		l,r := -1.000,1.000
		for r-l > 1e-8 {
			m := 0.5*(r+l)
			c := calcsum(m)
			if c > 0 { l = m } else { r = m }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,0.5*(l+r))
	
    }
}

