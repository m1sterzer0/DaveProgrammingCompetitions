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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); S := gs()
		ep := []int{}
		for i:=0;i<N-1;i++ { if S[i] != S[i+1] { ep = append(ep,i) } }; ep = append(ep,N-1)
		larr := []int{}; larr = append(larr,ep[0]+1); for i:=1;i<len(ep);i++ { larr = append(larr,ep[i]-ep[i-1]) }
		odds := []int{}; for i:=0;i<len(larr);i++ { if larr[i] % 2 == 1 { odds = append(odds,i) } }
		ans := 0;  for i:=0;i<len(odds);i+=2 {	ans += odds[i+1]-odds[i] }
		fmt.Fprintln(wrtr,ans)
	}
}

