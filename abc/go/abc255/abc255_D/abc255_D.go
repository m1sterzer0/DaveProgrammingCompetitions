package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func abs(a int) int { if a < 0 { return -a }; return a }
func bisect_left(arr []int, targ int) int {
	l, u := -1, len(arr); for u-l > 1 { m := (u + l) >> 1; if arr[m] < targ { l = m } else { u = m } }; return u
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); A := gis(N); X := gis(Q)
	sort.Slice(A,func(i,j int) bool { return A[i]<A[j] })
	Amin := A[0]; Amax := A[N-1]
	prefixSum := make([]int,N); prefixSum[0] = A[0]; for i:=1;i<N;i++ { prefixSum[i] = prefixSum[i-1] + A[i] }
	for _,x := range X {
		ans := 0
		if x <= Amin || x >= Amax { 
			ans = abs(x * N - prefixSum[N-1])
		} else {
			idx := bisect_left(A,x)
			ans = abs(x * (idx) - prefixSum[idx-1]) + abs(x*(N-idx) - (prefixSum[N-1]-prefixSum[idx-1]))
		}
		fmt.Fprintln(wrtr,ans)
	}
}

