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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
type edge struct {n2,c int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A,B,C := fill3(N-1); for i:=0;i<N-1;i++ { A[i]--; B[i]-- }; D := gis(N)
	gr := make([][]edge,N)
	for i:=0;i<N-1;i++ {
		a,b,c := A[i],B[i],C[i]
		gr[a] = append(gr[a],edge{b,c})
		gr[b] = append(gr[b],edge{a,c})
	}
	bestdn1 := ia(N)
	bestdn2 := ia(N)
	dn1idx := iai(N,-1)
	dn2idx := iai(N,-1)
	ans := ia(N)
	var dfs1 func(n,p int)
	dfs1 = func(n,p int) {
		for _,e := range gr[n] {
			if e.n2 == p { continue }
			dfs1(e.n2,n)
			cand := e.c + max(D[e.n2],bestdn1[e.n2])
			if cand > bestdn1[n] {
				bestdn2[n],dn2idx[n] = bestdn1[n],dn1idx[n]
				bestdn1[n],dn1idx[n] = cand,e.n2
			} else if cand > bestdn2[n] {
				bestdn2[n],dn2idx[n] = cand,e.n2
			}
		}
	}
	dfs1(0,-1)
	var dfs2 func(n,p,bestup int)
	dfs2 = func(n,p,bestup int) {
		ans[n] = max(bestup,bestdn1[n])
		bestup = max(bestup,D[n])
		for _,e := range gr[n] {
			if e.n2 == p { continue }
			lbest := bestup
			if dn1idx[n] != e.n2 { lbest = max(lbest,bestdn1[n]) }
			if dn2idx[n] != e.n2 { lbest = max(lbest,bestdn2[n]) }
			lbest += e.c
			dfs2(e.n2,n,lbest)
		}
	}
	dfs2(0,-1,0)
	for _,a := range ans { fmt.Fprintln(wrtr,a) }
}

