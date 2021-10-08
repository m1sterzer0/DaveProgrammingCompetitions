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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	G := [3][3]int{}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		G[0][0],G[0][1],G[0][2] = gi3()
		G[1][0],        G[1][2] = gi2()
		G[2][0],G[2][1],G[2][2] = gi3()
		ans := 0; dmap := make(map[int]int)
		for _,i := range []int{0,2} { 
			if G[i][1]-G[i][0] == G[i][2]-G[i][1] { ans++ }
			if G[1][i]-G[0][i] == G[2][i]-G[1][i] { ans++ }
		}
		v := G[0][0] + G[2][2]; if v & 1 == 0 { dmap[v/2]++ }
		v =  G[2][0] + G[0][2]; if v & 1 == 0 { dmap[v/2]++ }
		v =  G[0][1] + G[2][1]; if v & 1 == 0 { dmap[v/2]++ }
		v =  G[1][0] + G[1][2]; if v & 1 == 0 { dmap[v/2]++ }
		ccontrib := 0
		for _,v := range dmap { ccontrib = max(ccontrib,v) }
		ans += ccontrib
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

