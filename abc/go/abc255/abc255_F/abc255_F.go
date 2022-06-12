package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); P := gis(N); I := gis(N); for i:=0;i<N;i++ { P[i]--; I[i]-- }
	Irev := ia(N); for i,n := range I { Irev[n] = i }
	good := P[0] == 0; l := iai(N,-1); r := iai(N,-1)
	var solveTree func(pstart,pend,istart,iend int)
	solveTree = func(pstart,pend,istart,iend int) {
		head := P[pstart]
		ihead := Irev[head]
		if ihead < istart || ihead > iend { good = false; return }
		lsize := ihead-istart
		rsize := iend-ihead
		if lsize > 0 { left := P[pstart+1]; l[head] = left; solveTree(pstart+1,pstart+lsize,istart,ihead-1) }
		if rsize > 0 { right := P[pstart+1+lsize]; r[head] = right; solveTree(pstart+1+lsize,pend,ihead+1,iend) }
	}
	solveTree(0,N-1,0,N-1)
	if !good {
		fmt.Fprintln(wrtr,-1)
	} else {
		for i:=0;i<N;i++ {
			fmt.Fprintf(wrtr,"%v %v\n",l[i]+1,r[i]+1)
		}
	}
}

