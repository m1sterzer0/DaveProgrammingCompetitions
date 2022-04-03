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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		F := gis(N)
		P := gis(N); for i:=0;i<N;i++ { P[i]-- }
		gr := make([][]int,N)
		roots := make([]int,0)
		for i,p := range P {
			if p == -1 { roots = append(roots,i) } else { gr[p] = append(gr[p],i) }
		}
		// DFS1 -- roll up minimum value of chain reaction ending at node i
		ans := 0
		var dfs1 func(n int) int
		dfs1 = func(n int) int {
			if len(gr[n]) == 0 {
				return F[n]
			} else {
				sum := 0
				minchild := inf
				for _,c := range gr[n] {
					cval := dfs1(c); sum += cval; minchild = min(minchild,cval)
				}
				ans += sum-minchild
				return max(F[n],minchild)
			}
		}
		for _,r := range roots { ans += dfs1(r) }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

