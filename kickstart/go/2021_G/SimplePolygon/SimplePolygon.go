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
func gi2() (int,int) { return gi(),gi() }
type pt2 struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,A := gi2()
		if A < N-2 { 
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"POSSIBLE")
			ansarr := make([]pt2,N+1)
			if N == 3 {
				ansarr[0] = pt2{0,0}
				ansarr[N] = pt2{0,0}
				ansarr[1] = pt2{0,A}
				ansarr[2] = pt2{1,0}
			} else {
				ansarr[0] = pt2{0,1}
				ansarr[N] = pt2{0,1}
				ansarr[1] = pt2{0,4+A-N}
				i,j := 2,N-1
				for i<=j {
					ansarr[i].x = ansarr[i-1].x+1
					ansarr[j].x = ansarr[i-1].x+1
					if i != j {
						if i%2 == 0 { ansarr[i].y = 1; ansarr[j].y = 0 } else { ansarr[i].y = 2; ansarr[j].y = 1 }
					} else {
						if i%2 == 0 { ansarr[i].y = 0 } else { ansarr[i].y = 2 }
					}
					i++;j--
				}
			}
			for i:=0;i<N;i++ { fmt.Fprintf(wrtr,"%v %v\n",ansarr[i].x,ansarr[i].y) }
		}
	}
}
