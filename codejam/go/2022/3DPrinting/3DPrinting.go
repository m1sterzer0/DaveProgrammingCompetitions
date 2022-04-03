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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		C1,M1,Y1,K1 := gi4()
		C2,M2,Y2,K2 := gi4()
		C3,M3,Y3,K3 := gi4()
		cmax := min(C1,min(C2,C3))
		mmax := min(M1,min(M2,M3))
		ymax := min(Y1,min(Y2,Y3))
		kmax := min(K1,min(K2,K3))
		left := 1000000
		c := min(left,cmax); left -= c
		m := min(left,mmax); left -= m
		y := min(left,ymax); left -= y
		k := min(left,kmax); left -= k
		if left > 0 {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v %v %v %v\n",tt,c,m,y,k)
		}
    }
}

