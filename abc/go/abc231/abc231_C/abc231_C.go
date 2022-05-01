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
type query struct { idx,x int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); A := gis(N); X := gis(Q)
	qq := make([]query,Q); for i:=0;i<Q;i++ { qq[i] = query{i,X[i]} }
	ansarr := make([]int,Q)
	sort.Slice(A,func(i,j int) bool { return A[i] < A[j] } )
	sort.Slice(qq,func(i,j int) bool { return qq[i].x < qq[j].x } )
	aptr := 0
	for _,q := range qq {
		for aptr < N && A[aptr] < q.x { aptr++ }
		ansarr[q.idx] = N-aptr
	}
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}


