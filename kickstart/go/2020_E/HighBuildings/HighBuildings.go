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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
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
		N,A,B,C := gi4()
		A -= C; B -= C
		ansarr := iai(N,1)
		var ansstr string
		if A + B + C > N {
			ansstr = "IMPOSSIBLE"
		} else if N == 1 {
			if (A > 0 || B > 0 || C != N) {
				ansstr = "IMPOSSIBLE"
			} else {
				ansstr = vecintstring(ansarr)
			}
		} else if N == 2 {
			if C == 2 { 
				ansstr = vecintstring(ansarr)
			} else if A == 1 {
				ansarr[1] = 2; ansstr = vecintstring(ansarr)
			} else if B == 1 {
				ansarr[0] = 2; ansstr = vecintstring(ansarr)
			} else {
				ansstr = "IMPOSSIBLE" //Shouldn't get here
			}
		} else if C == 1 {
			D := N - A - B - C
			if A == 0 && B == 0 {
				ansstr = "IMPOSSIBLE"
			} else if A > 0 {
				for i:=0;i<A;i++ { ansarr[i] = 2}
				for i:=A+D;i<A+D+C;i++ { ansarr[i] = 3 }
				ansstr = vecintstring(ansarr)
			} else {
				for i:=A;i<A+C;i++ { ansarr[i] = 3}
				for i:=A+C+D;i<N;i++ { ansarr[i] = 2 }
				ansstr = vecintstring(ansarr)
			}
		} else {
			D := N - A - B - C
			for i:=A;i<A+1;i++ { ansarr[i] = 2}
			for i:=A+D+1;i<A+D+C;i++ {ansarr[i] = 2}
			ansstr = vecintstring(ansarr)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

