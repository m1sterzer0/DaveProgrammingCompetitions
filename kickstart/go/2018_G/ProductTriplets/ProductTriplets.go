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
		N := gi(); A := gis(N)
		sb := make([]int,200001)
		for _,a := range A { sb[a]++ }
		ans := 0
		// All three zero
		if sb[0] >= 3 { ans += sb[0] * (sb[0]-1) * (sb[0]-2) / 6 }
		// Two zeros and one nonzero
		if sb[0] >= 2 { ans += sb[0] * (sb[0]-1) / 2 * (N - sb[0]) }
		// Three ones
		if sb[1] >= 3 { ans += sb[1] * (sb[1]-1) * (sb[1]-2) / 6 }
		// Process pairs.  Two cases: 1*x=x and x*x=y
		for i:=2;i<=200000;i++ {
			if sb[i] < 2 { continue }
			numpair := sb[i] * (sb[i]-1) / 2
			ans += numpair * sb[1]
			if i*i <= 200000 { ans += numpair * sb[i*i] }
		}
		// All three different numbers all greater than 1
		for i:=2;i<=1000;i++ {
			jmax := 200000/i
			for j:=i+1;j<=jmax;j++ {
				ans += sb[i] * sb[j] * sb[i*j]
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

