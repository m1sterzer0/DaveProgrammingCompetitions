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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs(); }
	ans := inf
	for targ:=0;targ<=9;targ++ {
		cnts := make([]int,10)
		for _,s := range S {
			for i,c := range s { if int(c-'0') == targ { cnts[i]++} }
		}
		mytime := 0
		for i:=0;i<=9;i++ {
			if cnts[i] == 0 { continue }
			mytime = max(mytime,i+10*(cnts[i]-1))
		}
		ans = min(ans,mytime)
	}
	fmt.Println(ans)
}

