package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type interval struct { s,e int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2(); S,E := fill2(N)
		ivals := make([]interval,N)
		for i:=0;i<N;i++ { ivals[i] = interval{S[i],E[i]}}
		sort.Slice(ivals,func(i,j int) bool { return ivals[i].s < ivals[j].s} )
		kend := -1; ans := 0
		for _,ii := range ivals {
			if kend >= ii.e { continue }
			if kend >= ii.s { ii.s = kend }
			numdep := (ii.e-ii.s+K-1)/K
			kend = ii.s + K * numdep
			ans += numdep
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

