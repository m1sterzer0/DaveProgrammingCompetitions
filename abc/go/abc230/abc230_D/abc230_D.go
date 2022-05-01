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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type wall struct {idx,l,r int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,D := gi(),gi(); L,R := fill2(N)
	walls  := make([]wall,N); for i:=0;i<N;i++ { walls[i]  = wall{i,L[i],R[i]} }
	walls2 := make([]wall,N); for i:=0;i<N;i++ { walls2[i] = wall{i,L[i],R[i]} }
	checklist := make([]bool,N)
	sort.Slice(walls,func(i,j int) bool { return walls[i].l < walls[j].l } )
	sort.Slice(walls2,func(i,j int) bool { return walls2[i].r < walls2[j].r } )
	ans,lptr,rptr := 0,0,0
	for rptr < N {
		ans++; xl := walls2[rptr].r; xr := xl+D-1
		for lptr < N && walls[lptr].l <= xr { checklist[walls[lptr].idx] = true; lptr++ }
		for rptr < N && checklist[walls2[rptr].idx] { rptr++ }
	}
	fmt.Println(ans)
}

