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

func funnysqrt(x int) int {
	l,r := 0,100001
	for r-l > 1 {
		m := (r+l)>>1
		if m*m <= x { l = m } else { r = m }
	}
	if (2*l+1)*(2*l+1) >= 4*x { return l } else { return r }
}

func countFilled(R int) int {
	cnt := 4*R+1; y:=R
	for x:=1;x<=R;x++ {
		for funnysqrt(x*x+y*y) > R { y-- }
		cnt += 4*y
	}
	return cnt
}

// You have to see that
// a) countFilled is a superset of countPerim
// b) the countPerim radii never overlap
// c) you never mirror from below the 45 degree line to above the 45 degree line
// All of this seems difficult, but the top solvers seem to have seen this nearly right away.
func countPerim(R int) int {
	cnt := 4*R+1; x:=0
	for rr:=1;rr<=R;rr++ {
		for funnysqrt(rr*rr-(x+1)*(x+1)) >= (x+1) { x++ }
		if funnysqrt(rr*rr-x*x) > x { cnt += 8*x } else { cnt += 8*x - 4 } 
	}
	return cnt
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
		// The two observations that you need
		// a) the area count is a superset of the perimeter count
		// b) for fixed x, the perimeter count emits a unique y value for each x coordinate
		R := gi()
		cnt1 := countFilled(R)
		cnt2 := countPerim(R)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,cnt1-cnt2)
	}
}

