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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func abs(a int) int { if a < 0 { return -a }; return a }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N)
	cross := func(x1,y1,x2,y2 int) int { return x1*y2-y1*x2 }
	tri := func(i0,i1,i2 int) int {
		x0 := X[i0 % N]; y0 := Y[i0 % N]
		x1 := X[i1 % N]; y1 := Y[i1 % N]
		x2 := X[i2 % N]; y2 := Y[i2 % N]
		return abs(cross(x1-x0,y1-y0,x2-x0,y2-y0))
	}
	// First calculate 2*area of ngon
	A := 0
	for i:=2;i<N;i++ { A += tri(0,i-1,i) }  // x1,y1,x2,y2 := X[i-1]-X[0],Y[i-1]-Y[0],X[i]-X[0],Y[i]-Y[0]; A += x1*y2-y1*x2 }
	best,i0,i1,a := A,0,1,0
	for {
		for 4*a < A  { i1++; a += tri(i0,i1-1,i1); best = min(best,abs(A-4*a)) }
		for 4*a >= A { i0++; a -= tri(i0-1,i0,i1); best = min(best,abs(A-4*a)) } 
		if i0 >= N   { break }
	}
	fmt.Println(best)
}
