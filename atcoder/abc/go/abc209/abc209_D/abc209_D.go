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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi2(); A,B := fill2(N-1); for i:=0;i<N-1;i++ { A[i]--; B[i]-- }; C,D := fill2(Q); for i:=0;i<Q;i++ { C[i]--; D[i]-- }
	gr := make([][]int,N); for i:=0;i<N-1;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }
	d := iai(N,0)
	var dfs func(n,p,dp int)
	dfs = func(n,p,dp int) {
		d[n] = dp
		for _,c := range gr[n] { if c != p { dfs(c,n,dp+1) } }
	}
	dfs(0,-1,0)
	for i:=0;i<Q;i++ { cc,dd := C[i],D[i]; ans := "Town"; if (d[cc]-d[dd]) & 1 == 1 { ans = "Road" }; fmt.Fprintln(wrtr,ans) }
}

