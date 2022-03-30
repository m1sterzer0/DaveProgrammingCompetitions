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
func gi3() (int,int,int) { return gi(),gi(),gi() }
type rec struct { s string; v int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,P := gi3()
		prefarr := make([]string,N); for i:=0;i<N;i++ { prefarr[i] = gs() }
		badarr  := make([]string,M); for i:=0;i<M;i++ { badarr[i] = gs()  }
		last := make([]rec,0); last = append(last,rec{"",0})
		next := make([]rec,0)
		badset := make(map[string]int)
		for i:=1;i<=P;i++ {
			badset = make(map[string]int)
			for _,b := range badarr { badset[b[0:i]] = 1 }
			next = next[:0]
			idx := i - 1; z,o := 0,0
			for _,p := range prefarr { 
				if p[idx] == '0' { o++ } else { z++ }
			}
			for _,rr := range last {
				next = append(next,rec{rr.s + "0", rr.v + z})
				next = append(next,rec{rr.s + "1", rr.v + o})
			}
			sort.Slice(next,func(i,j int) bool { return next[i].v < next[j].v })
			for i,rr := range next {
				if badset[rr.s] == 0 { next = next[:i+1]; break }
			}
			last,next = next,last
		}
		ans := 0
		for _,rr := range last {
			if badset[rr.s] == 0 { ans = rr.v; break }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

