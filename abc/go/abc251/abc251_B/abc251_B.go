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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,W := gi(),gi(); A := gis(N)
	sb := make([]bool,W+1)
	for i:=0;i<N;i++ {
		w := A[i]; if w <= W { sb[w] = true }
		for j:=i+1;j<N;j++ {
			w = A[i]+A[j]; if w <= W { sb[w] = true }
			for k:=j+1;k<N;k++ {
				w = A[i]+A[j]+A[k]; if w <= W { sb[w] = true }
			}
		}
	}
	cnt := 0; for i:=1;i<=W;i++ { if sb[i] { cnt++ } }
	fmt.Println(cnt)
}

