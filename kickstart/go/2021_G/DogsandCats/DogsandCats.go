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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		_,D,C,M := gi4(); S := gs()
		totdogs := 0; for _,s := range S { if s == 'D' { totdogs++ } }
		dogsfed := 0
		for _,s := range S {
			if s == 'C' { if C == 0 { break }; C-- }
			if s == 'D' { if D == 0 { break }; D--; dogsfed++; C += M }
		}
		ans := "NO"; if dogsfed == totdogs { ans = "YES" }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

