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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); B := gis(N)
	l0lj,l0rj,r0lj,r0rj := A[0],1<<60,1<<60,0
	for i:=1;i<N;i++ {
		nl0lj := min(l0lj+B[i-1]+A[i],l0rj+A[i])
		nl0rj := min(l0lj,l0rj+B[i-1])
		nr0lj := min(r0lj+B[i-1]+A[i],r0rj+A[i])
		nr0rj := min(r0lj,r0rj+B[i-1])
		l0lj,l0rj,r0lj,r0rj = nl0lj,nl0rj,nr0lj,nr0rj
	}
	l0lj += B[N-1]; r0rj += B[N-1]
	ans := min(min(l0lj,l0rj),min(r0lj,r0rj))
	fmt.Println(ans)
}

