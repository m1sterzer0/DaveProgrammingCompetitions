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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func abs(a int) int { if a < 0 { return -a }; return a }
func min(a,b int) int { if a > b { return b }; return a }
const inf = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,X,Y := gi(),gi(),gi(); A := gis(N); B := gis(N)
	// We will do a subset DP on the minimum cost of matching the first xx members of B with the particular subset of A
	dp := make([]int,1<<N)
	for bm := uint(1); bm<1<<N; bm++ {
		k := bits.OnesCount(bm)
		bidx := k-1
		best:=inf; numbefore := 0
		for i:=0;i<N;i++ {
			if bm & (1<<uint(i)) == 0 { numbefore++; continue }
			aidx := k-1+numbefore
			cand := dp[bm ^ (1<<uint(i))] + Y * (aidx-bidx) + X * abs(A[i]-B[bidx])
			best = min(best,cand)
			//fmt.Printf("DBG: bm:%04b i:%v aidx:%v bidx:%v cand:%v best:%v X:%v Y:%v\n",bm,i,aidx,bidx,cand,best,X,Y)
		}
		dp[bm] = best
	}
	ans := dp[1<<N-1]
	fmt.Println(ans)
}

