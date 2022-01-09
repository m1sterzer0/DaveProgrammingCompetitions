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
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	var brd [1000][1000]int
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ { 
				brd[i][j] = gi()
			}
		}
		best := 0
		for si:=0;si<N;si++ {
			i,j := si,0; cand := brd[i][j]
			for i+1 < N && j+1 < N { i++; j++ ; cand += brd[i][j] }
			best = max(best,cand)
		}
		for sj:=0;sj<N;sj++ {
			i,j := 0,sj; cand := brd[i][j]
			for i+1 < N && j+1 < N { i++; j++; cand += brd[i][j] }
			best = max(best,cand)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

