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
type letter struct {idx int; c byte}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S := gs(); N := len(S); ll := make([]letter,N); for i,c := range S { ll[i] = letter{i,byte(c)} }
		sort.Slice(ll,func(i,j int) bool {return ll[i].c < ll[j].c} )
		A := make([]byte,N); for i,l := range ll { A[i] = l.c }
		B := make([]byte,N); copy(B,A); for i,j := 0,len(S)-1; i<j; i,j = i+1,j-1 { B[i],B[j] = B[j],B[i] }
		ptr := 0; good := true
		for i:=0;i<N;i++ {
			if A[i] != B[i] { continue }
			for ptr < N && (A[ptr] == A[i] || B[ptr] == A[i]) { ptr++ }
			if ptr == N { good = false; break }
			B[i],B[ptr] = B[ptr],B[i]; ptr++
		}
		ansstr := ""
		if !good { 
			ansstr = "IMPOSSIBLE"
		} else {
			ansarr := make([]byte,N)
			for i,l := range ll { ansarr[l.idx] = B[i]}
			ansstr = string(ansarr)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
	}
}
