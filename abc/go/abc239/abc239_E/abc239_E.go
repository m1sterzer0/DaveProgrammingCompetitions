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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); X := gis(N); A,B := fill2(N-1); V,K := fill2(Q)
	for i:=0;i<N-1;i++ { A[i]--; B[i]-- }; for i:=0;i<Q;i++ { V[i]--; K[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }
	scratch := make([]int,20)
	kmax := make([][]int,N); for i:=0;i<N;i++ { kmax[i] = append(kmax[i],X[i]) }
	domerge := func(n,c int) {
		scratch = scratch[:0]
		for _,x := range kmax[n] { scratch = append(scratch,x) }
		kmax[n] = kmax[n][:0]
		a1,a2 := scratch,kmax[c]; a1len,a2len := len(a1),len(a2); i1,i2,cnt := 0,0,0
		for cnt < 20 && (i1 < a1len || i2 < a2len) {
			if i1 < a1len && (i2 == a2len || a1[i1] >= a2[i2]) {
				kmax[n] = append(kmax[n],a1[i1]); i1++; cnt++
			} else {
				kmax[n] = append(kmax[n],a2[i2]); i2++; cnt++
			}
		}
	}
	var dfs func(n,p int)
	dfs = func(n,p int) {
		for _,c := range gr[n] {
			if c == p { continue }
			dfs(c,n)
			domerge(n,c)
		}
	}
	dfs(0,-1)
	for i:=0;i<Q;i++ { v,k := V[i],K[i]; fmt.Fprintln(wrtr,kmax[v][k]) }
}

