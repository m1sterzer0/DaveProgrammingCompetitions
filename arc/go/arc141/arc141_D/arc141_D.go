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
	N,M := gi(),gi()
	A := gis(N)
	// Consider the equivalence classes of a * 2^n where a is odd.
	// The key observations are that a) there are M equivalence classes, and b) we can only have at most one representative from each equivalence class
	// Thus, to get the M element set, we need a representative from each set
	sb := make([]bool,2*M+1); for _,a := range A { sb[a] = true }
	gr := make([][]int,2*M)
	for i:=1;i<=2*M;i+=2 {
		for j:=1;i*j<=2*M;j*=2 { if sb[i*j] { gr[i] = append(gr[i],i*j) } }
	}
	// Calculate the max value and min value that we can have for each eq class
	globgood := true
	maxv := make([]int,2*M)
	for i:=1;i<=2*M;i+=2 { l := len(gr[i]); maxv[i] = l-1 }
	for i:=1;i<2*M;i+=2 {
		if maxv[i] == -1 { globgood = false; break }
		v := gr[i][maxv[i]]
		for j:=3;i*j <= 2*M; j+= 2 {
			k := i*j
			for maxv[k] >= 0 && gr[k][maxv[k]] % v == 0 { maxv[k]-- }
		}
	}
	minv := make([]int,2*M)
	for i:=1;i<=2*M;i+=2 { minv[i] = 0 }
	for i:=2*M-1;i>=1;i-=2 {
		for j:=3;i*j <= 2*M; j+= 2 {
			k := i*j
			for minv[i] < len(gr[i]) && gr[k][minv[k]] % gr[i][minv[i]] == 0 { minv[i]++ }
		}
		if minv[i] == len(gr[i]) { globgood = false; break }
	}
	for _,a := range A {
		b := a; for b & 1 == 0 { b = b >> 1 }
		if globgood && gr[b][minv[b]] <= a && a <= gr[b][maxv[b]] { fmt.Fprintln(wrtr,"Yes") } else { fmt.Fprintln(wrtr,"No") }
	}
}

