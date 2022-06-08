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
		K := gi(); i:=1; for 1<<uint(i)-1 < K { i++ }
		inv := false
		for;i>=2;i-- {
			l := (1<<uint(i)-1)
			c := (l+1)/2
			if K < c { continue }
			if K == c { break } 
			inv = !inv; K = c - (K-c)
		}
		ans := 0; if inv { ans = 1 }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

