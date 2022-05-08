package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
	var dp[1<<15][26]int; ans := 0
	for bm:=uint(1);bm<1<<N;bm++ {
		if bits.OnesCount(bm) == 1 {
			idx := bits.TrailingZeros(bm)
			for _,c := range S[idx] { dp[bm][c-'a']++ }
		} else {
			idx := bits.TrailingZeros(bm)
			bm1 := uint(1) << uint(idx)
			bm2 := bm ^ bm1
			for i:=0;i<26;i++ { dp[bm][i] += dp[bm1][i] + dp[bm2][i] }
		}
		cand := 0; for i:=0;i<26;i++ { if dp[bm][i] == K { cand++ } }
		ans = max(ans,cand)
	}
	fmt.Println(ans)
}
