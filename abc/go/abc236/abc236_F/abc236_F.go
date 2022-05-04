package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
type bvec struct {m,idx int}
type vec struct { m,c int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// I *think* greedy works.  We should only need N spices to form a basis
	N := gi(); V := make([]vec,0)
	for i:=1;i<1<<N;i++ { c := gi(); V = append(V,vec{i,c})}
	basis := make([]bvec,0)
	sort.Slice(V,func(i,j int) bool { return V[i].c < V[j].c } )
	ans := 0
	for i:=0;len(basis)<N;i++ {
		m := V[i].m
		for _,b := range basis { if m & (1<<b.idx) != 0 { m = m ^ b.m } }
		if m == 0 { continue }
		ans += V[i].c
		for i:=0;i<N;i++ { if m & (1 << i) != 0 { basis = append(basis,bvec{m,i}); break } }
	}
	fmt.Println(ans)
}

