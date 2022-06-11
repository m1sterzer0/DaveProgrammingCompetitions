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
func min(a,b int) int { if a > b { return b }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := gs()
	numblocks,sumblocks := 0,0
	for i:=1;i<N-1;i++ {
		if S[i-1] != 'A' || S[i] != 'R' || S[i+1] != 'C' { continue }
		j,k := i-2,i+2; d := 1
		for j >= 0 && k < N && S[j] == 'A' && S[k] == 'C' { d++; j--; k++ }
		numblocks++; sumblocks += d
	}
	ans := min(2*numblocks,sumblocks)
	fmt.Println(ans)
}

