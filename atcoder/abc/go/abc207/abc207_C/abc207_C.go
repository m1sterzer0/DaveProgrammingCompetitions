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
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); T,L,R := fill3(N); ans := 0
	for i:=0;i<N;i++ { L[i] *= 2; R[i] *= 2 }
	for i:=0;i<N;i++ { if T[i] == 2 { R[i]-- } else if T[i] == 3 { L[i]++ } else if T[i] == 4 { L[i]++; R[i]-- } }
	for i:=0;i<N;i++ {
		for j:=i+1;j<N;j++ {
			if R[i] < L[j] || R[j] < L[i] { continue }
			ans++
		}
	}
	fmt.Println(ans)
}

