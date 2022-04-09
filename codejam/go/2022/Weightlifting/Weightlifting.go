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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	var commonByWeight [100][100][100]int
	var solutionsByInterval [100][100]int
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		E,W := gi2()
		bd := make([][]int,E); for i:=0;i<E;i++ { bd[i] = gis(W) }

		// Calculate the common weights between the intervals
		for delta:=0;delta<E;delta++ {
			for i:=0;i+delta<E;i++ {
				j := i+delta
				ltot := 0
				for w:=0;w<W;w++ {
					if i == j { 
						commonByWeight[i][j][w] = bd[i][w]
					} else {
						commonByWeight[i][j][w] = min(commonByWeight[i][j-1][w],commonByWeight[i+1][j][w])
					}
					ltot += commonByWeight[i][j][w]
				}
				solutionsByInterval[i][j] = inf
				if i == j {
					solutionsByInterval[i][j] = 2 * ltot
				} else {
					best := inf
					for k:=i;k<j;k++ { best = min(best,solutionsByInterval[i][k]+solutionsByInterval[k+1][j]) }
					solutionsByInterval[i][j] = best - 2*ltot
				}
			}
		}
		ans := solutionsByInterval[0][E-1]
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

