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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func abs(a int) int { if a < 0 { return -a }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type state struct {pidx,st int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,P := gi(),gi(); X := make([][]int,N); for i:=0;i<N;i++ { X[i] = gis(P) }
		pmax,pmin := make([]int,N),make([]int,N)
		for i:=0;i<N;i++ {
			pmax[i],pmin[i] = X[i][0],X[i][0]
			for j:=1;j<P;j++ { pmax[i] = max(pmax[i],X[i][j]); pmin[i] = min(pmin[i],X[i][j]) }
		}
		cache := make(map[state]int)
		var solvecase func(pidx,st int) int
		solvecase = func(pidx,st int) int {
			v,ok := cache[state{pidx,st}]
			if !ok {
				if pidx == N {
					v = 0
				} else {
					mymax,mymin := pmax[pidx],pmin[pidx]
					v1 := abs(mymin-st) + abs(mymax-mymin) + solvecase(pidx+1,mymax)
					v2 := abs(mymax-st) + abs(mymax-mymin) + solvecase(pidx+1,mymin)
					v = min(v1,v2)
				}
				cache[state{pidx,st}] = v
			}
			return v
		}
		ans := solvecase(0,0)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

