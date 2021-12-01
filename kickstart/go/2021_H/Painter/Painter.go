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
func ia(m int) []int { return make([]int,m) }

func countsegs(r []int) int {
	ret,last := 0,0
	for _,c := range r { 
		if last == 0 && c == 1 { ret++ }
		last = c
	}
	return ret
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
		N := gi(); P := gs()
		r,y,b := ia(N),ia(N),ia(N)
		for i,c := range P {
			if c == 'R' || c == 'O' || c == 'P' || c == 'A' { r[i] = 1 }
			if c == 'Y' || c == 'O' || c == 'G' || c == 'A' { y[i] = 1 }
			if c == 'B' || c == 'G' || c == 'P' || c == 'A' { b[i] = 1 }
		}
		ans := 0; ans += countsegs(r); ans += countsegs(y); ans += countsegs(b)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

