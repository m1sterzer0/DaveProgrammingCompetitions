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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	const inf = 1000000000
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); A := gis(N)
		cntneg := 0; for _,a := range A { if a < 0 { cntneg++ } }
		// Negatives clearly go on the left end of the array
		for i:=0;i<N;i++ { if i < cntneg { A[i] = -abs(A[i]) } else { A[i] = abs(A[i]) } }
		ans := "YES"
		for i:=0;i<N-1;i++ { if A[i] > A[i+1] { ans = "NO"; break } }
		fmt.Fprintln(wrtr,ans)
	}
}

