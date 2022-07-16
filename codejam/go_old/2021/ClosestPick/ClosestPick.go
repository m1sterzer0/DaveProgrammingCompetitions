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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi(),gi(); P := gis(N);
		sort.Slice(P,func(i,j int) bool { return P[i] < P[j]} )
		singles := []int{0,0}
		doubles := []int{0}
		// Do the endpoints first
		if P[0] > 1 { singles = append(singles,P[0]-1) }
		if P[N-1] < K { singles = append(singles,K-P[N-1])}
		// Now for the gaps
		for i:=1;i<N;i++ {
			gapsize := P[i]-P[i-1]-1
			if gapsize == 1 { singles = append(singles,1) }
			if gapsize > 1  { singles = append(singles,(gapsize+1)/2); doubles = append(doubles,gapsize) }
		}
		sort.Slice(singles,func(i,j int) bool { return singles[i] > singles[j]} )
		sort.Slice(doubles,func(i,j int) bool { return doubles[i] > doubles[j]} )
		good := max(singles[0]+singles[1],doubles[0])
		ans := float64(good)/float64(K)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

