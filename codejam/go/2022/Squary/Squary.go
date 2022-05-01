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
		N,K := gi(),gi(); E := gis(N)
		// (a1^2+a2^2+...+an^2) = (a1^2+a2^2+...+an^2)+2*sum_over_i<j(ai * aj)
		// for those to be equal, we need sum_over_i<j(ai*aj) = 0
		cumsum := 0; cumprodsum := 0
		for _,e := range E { cumprodsum += e*cumsum; cumsum += e }
		ansstr := "IMPOSSIBLE"
		if cumprodsum == 0 {
			ansstr = "0"
		} else if K == 1 && cumsum != 0 && cumprodsum % cumsum == 0 {
			ansstr = fmt.Sprintf("%v",-cumprodsum/cumsum)
		} else if K > 1 {
			// Use one number to get the cumsum to 1
			// Use the second number to get the modified cumprodsum to 0
			a := 1-cumsum; cumprodsum += a*cumsum; cumsum += a
			b := -cumprodsum
			ansstr = fmt.Sprintf("%v %v",a,b)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

