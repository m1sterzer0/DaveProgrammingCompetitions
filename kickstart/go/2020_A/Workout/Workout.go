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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2(); M := gis(N)
		check := func(m int) bool {
			cnt := 0
			for i:=0;i<N-1;i++ {
				if M[i+1] - M[i] <= m  { continue }
				if m == 0 { return false }
				cnt += (M[i+1]-M[i]-1) / m
			}
			return cnt <= K
		}
		l,u := -1,1000000001
		for u-l > 1 { m := (u+l)>>1; if check(m) { u = m } else { l = m } }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,u)
    }
}

