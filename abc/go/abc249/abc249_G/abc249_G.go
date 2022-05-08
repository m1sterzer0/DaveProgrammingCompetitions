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
func max(a,b int) int { if a > b { return a }; return b }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// This is a basis problem mixed with a digit DP problem
	// Key is to find the basis (or largest sub-basis) of (Ai << 30) | Bi
	// TODO: figure out how to check for the fake 0 vs. real 0 case
	N,K := gi(),gi(); A,B := fill2(N); best := 0
	var basis [60]int
	addBasis := func (x int) {
		for i:=59;i>=0;i-- { if x & (1<<i) != 0 { x = x ^ basis[i] } }
		if x == 0 { return }
		idx := 59; for x & (1<<idx) == 0 { idx-- }; basis[idx] = x
	}
	var basis2 [30]int
	addBasis2 := func (x int) {
		for i:=29;i>=0;i-- { if x & (1<<i) != 0 { x = x ^ basis2[i] } }
		if x == 0 { return }
		idx := 29; for x & (1<<idx) == 0 { idx-- }; basis2[idx] = x
	}
	maxFromBasis2 := func (x int) int {
		for i:=29;i>=0;i-- {
			if x & (1<<i) == 0 { x = x ^ basis2[i] }
		}
		return x
	}
	solveit := func(x,idx int) int {
		x = x & 0x3fffffff
		for i:=0;i<30;i++ { basis2[i] = 0 }
		for i:=idx;i>=0;i-- { addBasis2(basis[i] & 0x3fffffff) }
		return maxFromBasis2(x)
	}
	checkPossible := func() bool {
		// For this to be possible, we must have
		// a) Some row-reduced basis vector for the Ai's <= K
		// b) Have fewer than N basis vectors
		for i:=0;i<30;i++ { basis2[i] = 0 }
		for _,a := range A { addBasis2(a) }
		good := false; cnt := 0
		for i:=0;i<=29;i++ { if basis2[i] != 0 { cnt++; if  basis2[i] <= K { good = true } } }
		if cnt < N { good = true }
		return good
	}
	if !checkPossible() {
		best = -1
	} else if K == (1<<30)-1 { // Special case
		for i:=0;i<30;i++ { basis2[i] = 0 }
		for _,b := range B { addBasis2(b) }
		best = maxFromBasis2(0)
	} else {
		K += 1
		for i:=0;i<N;i++ { x := (A[i] << 30) | B[i]; addBasis(x) }
		bound := K << 30
		x := 0;
		for i:=59;i>=30;i-- {
			if bound & (1<<i) != 0 { // If the bound is a one, we have the ability to set the bit to zero and search on lower order bits freely
				if x & (1<<i) == 0 { cand := solveit(x,i-1); best = max(best,cand) }
				if x & (1<<i) != 0 && basis[i] != 0 { cand := solveit(x ^ basis[i],i-1); best = max(best,cand) }
			}
			if x & (1<<i) != bound & (1<<i) {
				if basis[i] == 0 { break }
				x = x ^ basis[i]
			}
		}
	}
	fmt.Println(best)
}

