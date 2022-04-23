package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

	guess := func(x,y int) int { 
		fmt.Fprintf(wrtr,"%v %v\n",x,y); wrtr.Flush(); res := gs()
		if res == "HIT" { return 1 } else if res == "MISS" { return -1 } else if res == "CENTER" { return 0 }
		fmt.Fprintf(os.Stderr,"SOMETHING BAD HAPPENED\n"); os.Exit(1); return -2
	}

	solveCase := func() {
		cmin,cmax := -1000000000,1000000000
		res,x1,y1 := -1,0,0
		// Phase 1: Find a point in the circle
		for res != 1 {
			x1 = cmin + rand.Intn(cmax-cmin+1)
			y1 = cmin + rand.Intn(cmax-cmin+1)
			res = guess(x1,y1); if res == 0 { return }
		}
		// Phase 2: Trace out the cross
		xl,xr,yl,yr := x1,x1,y1,y1
		for a := 1<<40;a>=1;a = a>>1 {
			if xr + a <= cmax { res = guess(xr+a,y1); if res == 0 { return }; if res == 1 { xr += a } }
			if xl - a >= cmin { res = guess(xl-a,y1); if res == 0 { return }; if res == 1 { xl -= a } }
		}
		for a := 1<<40;a>=1;a = a>>1 {
			if yr + a <= cmax { res = guess(x1,yr+a); if res == 0 { return }; if res == 1 { yr += a } }
			if yl - a >= cmin { res = guess(x1,yl-a); if res == 0 { return }; if res == 1 { yl -= a } }
		}
		// Phase 3: Guess the center
		res = guess((xr+xl)/2,(yr+yl)/2)
		if res != 0 { fmt.Fprintf(os.Stderr,"We did not win when we thought we should. Exiting\n"); os.Exit(1) }
	}

    T := gi(); gi(); gi()
	rand.Seed(8675309)
	for tt:=1;tt<=T;tt++ {
		solveCase()
	}
}
