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
func ia(m int) []int { return make([]int,m) }
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
		N,M := gi2(); A := gis(N)
		zcnt,ocnt := ia(62),ia(62)
		for i:=0;i<62;i++ {
			for _,a := range A { if a & (1<<uint(i)) == 0 { zcnt[i]++ } else { ocnt[i]++ } }
		}
		ans := 0; budget := M
		for i:=0;i<62;i++ {
			pv := 1 << uint(i)
			if min(zcnt[i],ocnt[i]) > budget / pv { ans = -1; break }
			if ocnt[i] >= zcnt[i] { ans |= 1 << uint(i) }
			budget -= min(zcnt[i],ocnt[i]) * pv
		}
		// Now greedily add in ones where we can
		if ans != -1 {
			for i:=61;i>=0;i-- {
				pv := 1 << uint(i)
				if zcnt[i] > ocnt[i] && (zcnt[i]-ocnt[i]) <= budget / pv { ans |= 1 << uint(i); budget -= pv * (zcnt[i]-ocnt[i]) }
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

