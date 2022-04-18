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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); bd := make([][]int,N); for i:=0;i<N;i++ { bd[i] = make([]int,N) }
		for i:=0;i<N;i++ { for j:=0;j<N;j++ { bd[i][j] = gi() } }
		tr,badr,badc := 0,0,0
		sb := make([]int,N+1)
		for i:=0;i<N;i++ { tr += bd[i][i] }
		for i:=0;i<N;i++ {
			for j:=1;j<=N;j++ { sb[j] = 0 }
			for j:=0;j<N;j++ { sb[bd[i][j]]++ }
			for j:=1;j<=N;j++ { if sb[j] > 1 { badr++; break } }
		}
		for j:=0;j<N;j++ {
			for i:=1;i<=N;i++ { sb[i] = 0 }
			for i:=0;i<N;i++ { sb[bd[i][j]]++ }
			for i:=1;i<=N;i++ { if sb[i] > 1 { badc++; break } }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v %v %v\n",tt,tr,badr,badc)
    }
}

