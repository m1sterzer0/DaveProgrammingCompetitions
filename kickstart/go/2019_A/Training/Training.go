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
func min(a,b int) int { if a > b { return b }; return a }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,P := gi2()
		S := gis(N)
		sort.Slice(S,func(i,j int) bool { return S[i] < S[j]})
		segsum := sumarr(S[0:P])
		best := P * S[P-1] - segsum
		for i:=P;i<N;i++ {
			segsum += S[i]-S[i-P]
			cand := P*S[i] - segsum
			best = min(cand,best)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

