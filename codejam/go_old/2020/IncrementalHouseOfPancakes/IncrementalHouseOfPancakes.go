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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

	solvetri := func(x int) int {
		l,u := 0,1500000000
		for u-l > 1 { m := (l+u)>>1; res := m * (m+1) / 2; if res <= x { l = m } else { u = m }}
		return l
	}

	solve2 := func(x,st int) int {
		l,u := 0,min(x/st+1,1000000001)
		for u-l > 1 { m := (l+u)>>1; res := st*m + m*(m-1); if res <= x { l = m} else { u = m} }
		return l
	}

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		L,R := gi(),gi(); swapped := false; if L < R { L,R = R,L; swapped = true }
		n1 := solvetri(L-R); L -= n1*(n1+1)/2
		if L == R { swapped = false }
		n2 := solve2(L,n1+1); n3 := solve2(R,n1+2)
		L -= n1*n2+n2*n2; R -= n1*n3+n3*n3+n3
		if swapped { L,R = R,L }
        fmt.Fprintf(wrtr,"Case #%v: %v %v %v\n",tt,n1+n2+n3,L,R)	
    }
}

