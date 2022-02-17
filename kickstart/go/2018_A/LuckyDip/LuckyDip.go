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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2(); V := gis(N)
		sort.Slice(V,func(i,j int) bool { return V[i] > V[j] } )
		cs := ia(N); cs[0] = V[0]; for i:=1;i<N;i++ { cs[i] = cs[i-1] + V[i] }
		ev := 0.000; idx := N-1 // idx is the last index you keep
		for i:=0; i<=K; i++ {
			for idx >= 0 && float64(V[idx]) < ev { idx-- }
			s := 0.000
			if idx == -1 { s = ev * float64(N) } else { s = float64(cs[idx]) + ev * float64(N-idx-1) }
			ev = s / float64(N)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ev)
    }
}

