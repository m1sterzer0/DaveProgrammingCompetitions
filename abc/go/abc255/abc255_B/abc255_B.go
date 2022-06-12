package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); A := gis(K); for i:=0;i<K;i++ { A[i]--; }; X,Y := fill2(N)
	darr := make([]int,N)
	for i:=0;i<N;i++ { darr[i] = 1000000000000000000 }
	for _,a := range A {
		x1,y1 := X[a],Y[a]
		for i:=0;i<N;i++ {
			x2,y2 := X[i],Y[i]
			cand := (x2-x1)*(x2-x1)+(y2-y1)*(y2-y1)
			if cand < darr[i] { darr[i] = cand }
		}
	}
	m := maxarr(darr)
	ans := math.Sqrt(float64(m))
	fmt.Println(ans)
}

