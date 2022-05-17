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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// For 1 rows, the best pattern is 
		//   ..x..x..x..x..x 
		// For 2 rows , the best pattern is also two rows of
		//   ..x..x..x..x..x
		//   ..x..x..x..x..x
		// For 3 (or more row), we want to do the following repeating pattern
		//   ..x..x..x..x..x..x..x..x..x..x..x..x..x
		//   .x..x..x..x..x..x..x..x..x..x..x..x..x.
		//   x..x..x..x..x..x..x..x..x..x..x..x..x..
		R,C := gi(),gi()
		if R > C { R,C = C,R }
		ans := 0
		if R <= 2 { 
			ans += R*(C/3*2+C%3)
		} else { 
			ans += R/3*2*C
			if R%3 > 0 { ans += C/3*2+C%3 }
			if R%3 > 1 { ans += C/3*2; if C%3 != 0 { ans++ } }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

