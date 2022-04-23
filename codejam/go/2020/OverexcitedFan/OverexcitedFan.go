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
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		X,Y := gi(),gi(); M := gs()
		// Strategy.  Assume we are at (0,0) and Peppurr is at (X,Y).  We simulate peppers movements, and
		// at each step, we just look at manhattan distance to see if we can get there in time.  O(N) solution.
		x,y,ans := X,Y,-1
		for i:=0;i<len(M);i++ {
			if M[i] == 'N' { y++ } else if M[i] == 'S' { y-- } else if M[i] == 'E' { x++ } else { x-- }
			if abs(x)+abs(y) <= (i+1) { ans = i+1; break }
		}
		if ans == -1 {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
    }
}

