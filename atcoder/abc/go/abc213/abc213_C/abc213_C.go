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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func copyarr(a []int) []int { a2 := make([]int,len(a)); for i:=0;i<len(a);i++ { a2[i] = a[i] }; return a2 }
func uniqSorted(a []int) []int {
	a2 := []int{}
	for i:=0;i<len(a);i++ { if i == 0 || a[i] != a[i-1] { a2 = append(a2,a[i]) } }
	return a2
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	_,_,N := gi3()
	A,B := fill2(N)
	A2 := copyarr(A); sort.Slice(A2,func(i,j int)bool{return A2[i] < A2[j]}); A3 := uniqSorted(A2)
	B2 := copyarr(B); sort.Slice(B2,func(i,j int)bool{return B2[i] < B2[j]}); B3 := uniqSorted(B2)
	ridx,cidx := 1,1; rmap := make(map[int]int); cmap := make(map[int]int)
	for _,r := range A3 { rmap[r] = ridx; ridx++ }
	for _,c := range B3 { cmap[c] = cidx; cidx++ }
	for i:=0;i<N;i++ {
		a,b := A[i],B[i]
		fmt.Fprintf(wrtr,"%v %v\n",rmap[a],cmap[b])
	}
}

