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
		X,Y := gi(),gi()
		ansarr := make([]byte,0)
		var tryit func(x,y int) bool
		tryit = func(x,y int) bool {
			if x == 0 && y == 0 { return true }
			if (x & 1) == (y & 1) { return false }
			if x == 1 && y == 0 { ansarr = append(ansarr,'E'); return true }
			if x == -1 && y == 0 { ansarr = append(ansarr,'W'); return true }
			if x == 0 && y == 1 { ansarr = append(ansarr,'N'); return true }
			if x == 0 && y == -1 { ansarr = append(ansarr,'S'); return true }
			if x & 1 == 1 {
				ansarr = append(ansarr,'W')
				if tryit((x+1)/2,y/2) { return true }
				ansarr = ansarr[:len(ansarr)-1]
				ansarr = append(ansarr,'E')
				if tryit((x-1)/2,y/2) { return true }
				ansarr = ansarr[:len(ansarr)-1]
			} else {
				ansarr = append(ansarr,'S')
				if tryit(x/2,(y+1)/2) { return true }
				ansarr = ansarr[:len(ansarr)-1]
				ansarr = append(ansarr,'N')
				if tryit(x/2,(y-1)/2) { return true }
				ansarr = ansarr[:len(ansarr)-1]
			}
			return false
		}
		res := tryit(X,Y)
		if res { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,string(ansarr))
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		}
    }
}

