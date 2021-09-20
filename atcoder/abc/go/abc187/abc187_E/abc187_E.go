package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); A,B := fill2(N-1); Q := gi(); T,E,X := fill3(Q)
	for i:=0;i<N-1;i++ { A[i]--; B[i]-- }; for i:=0;i<Q;i++ { E[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }

	// Generate parent array
	par := iai(N,0)
	var dfs1 func(int,int)
	dfs1 = func (n,p int) { par[n] = p; for _,c := range gr[n] { if c != p { dfs1(c,n) } } }
	dfs1(0,-1)

	// Do the operations
	subtreesum := iai(N,0)
	for i:=0;i<Q;i++ {
		t,e,x := T[i],E[i],X[i]; a,b := A[e],B[e]; if t == 2 { a,b = b,a }
		if par[a] == b { subtreesum[a] += x } else { subtreesum[0] += x; subtreesum[b] += -x }
	}

	ansarr := iai(N,0)
	var dfs2 func(int,int,int) 
	dfs2 = func(n,p,s int) {
		s += subtreesum[n]
		ansarr[n] = s
		for _,c := range gr[n] { if c != p { dfs2(c,n,s) } }
	}
	dfs2(0,-1,0)
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}


