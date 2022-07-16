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

	myle := func(a,b string) bool {
		if len(a) < len(b) { return true }
		if len(a) == len(b) && a <= b { return true }
		return false
	}

	makestr := func(a,m int) string {
		s := ""
		for i:=a;i<a+m;i++ { s = s + strconv.Itoa(i) }
		return s
	}

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// Looks like strings are the way to go
		// Need to search from 2 to 14 consecutive numbers
		Y := gs()
		best := "12345678910111213141516"
		for clen:=2;clen<=14;clen++ {
			l,u := 0,10000000000
			for u-l > 1 {
				m := (u+l)>>1
				s := makestr(m,clen)
				if myle(s,Y) { 
					l = m 
				} else {
					if myle(s,best) { best = s } 
					u = m 
				}
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

