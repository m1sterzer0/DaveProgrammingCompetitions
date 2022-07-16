package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,C := gi2()
		ans := "IMPOSSIBLE"
		// N + N-1 + N-2 + .... + 2 = (N)*(N+1)/2 - 1
		if C >= N-1 && C <= N*(N+1)/2-1 {
			C -= N-1
			A := ia(N); for i:=0;i<N;i++ { A[i] = i+1 }
			doRev := func(i,j int) {
				for i<j { A[i],A[j] = A[j],A[i]; i++; j-- }
			}
			for i:=N-2;i>=0;i-- {
				if C == 0 { break }
				if C >= N-1-i {
					doRev(i,N-1)
					C -= N-1-i
				} else {
					doRev(i,i+C)
					C = 0
				}
			}
			ans = vecintstring(A)
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

