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
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi();
	A,B := fill2(N);
	A2 := append([]int{},A...)
	B2 := append([]int{},B...)
	rev(A2); rev(B2)
	totlen := sumarr(A)
	l,r := float64(0),float64(totlen)

	timeit := func (A []int, B[]int, m float64) float64 {
		rs := float64(0); t := float64(0)
		for i:=0;i<N;i++ {
			if rs + float64(A[i]) < m {
				t += float64(A[i]) / float64(B[i])
				rs += float64(A[i])
			} else {
				return t + (m - rs) / float64(B[i])
			}
		}
		return 0.0 // Shouldn't get here
	}
	for r-l > 1e-6 {
		m := 0.5*(l+r)
		lt := timeit(A,B,m); rt := timeit(A2,B2,float64(totlen)-m)
		if lt < rt { l = m } else { r = m }
	}
	m := 0.5*(l+r)
	fmt.Println(m)
}

