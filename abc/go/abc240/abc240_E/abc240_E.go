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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	start := make([]int,N); stop := make([]int,N)
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	timer := 0
	var dfs func(n,p int)
	dfs = func(n,p int) {
		start[n] = timer+1
		if len(gr[n]) == 0 || len(gr[n]) == 1 && n != 0 { timer++ }
		for _,c := range gr[n] { if c != p { dfs(c,n) } }
		stop[n] = timer
	}
	dfs(0,-1)
	for i:=0;i<N;i++ { fmt.Fprintf(wrtr,"%v %v\n",start[i],stop[i]) }
}

