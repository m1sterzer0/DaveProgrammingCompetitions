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
func gi() int  { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); S := gs(); ans := N
	sb := [1000][26]int{} 
	for i:=N-1;i>=1;i-- {
		if N % i != 0 { continue }
		for j:=0;j<i;j++ { for k:=0;k<26;k++ { sb[j][k] = 0 } }
		for j,c := range S { sb[j%i][c-'a']++ }
		t := 0
		for j:=0;j<i;j++ {
			sv,mv := 0,0
			for k:=0;k<26;k++ { sv += sb[j][k]; mv = max(mv,sb[j][k]) }
			t += sv-mv
		}
		if t <= K { ans = i }
	}
	fmt.Println(ans)
}

