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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N:=gi(); gi()
		X,Y := fill2(N)
		a,b := 0,0
		best := X[0]
		for i:=0;i<N;i++ {
			newb := b + X[i]*Y[i]
			if b + X[i] > 0 && newb < 0 { // local maximum case
				firstterm := b+X[i]
				numterms := b/(-X[i])
				lastterm := b + numterms*X[i]
				xx := (firstterm+lastterm)*numterms / 2
				best = max(best,a+xx)
			}
			a += Y[i] * b + Y[i] * (Y[i]+1) / 2 * X[i]
			best = max(best,a)
			b = newb
		}
		fmt.Fprintln(wrtr,best)
	}
}

