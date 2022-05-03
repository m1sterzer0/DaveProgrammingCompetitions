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
type bucket struct {i,j int}
type pair struct {i,j int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); X,Y := fill2(N)
	h := make(map[bucket][]int)
	for i:=0;i<N;i++ { 
		x,y := X[i],Y[i]
		b := bucket{x/K,y/K}
		_,ok := h[b]
		if !ok { h[b] = make([]int,0) }
		h[b] = append(h[b],i)
	}
	ansarr := make([]pair,0)
	limit := K*K
	processCell := func(v []int) {
		lv := len(v)
		for i:=0;i<lv;i++ {
			idx1 := v[i]; x1,y1 := X[idx1],Y[idx1]
			for j:=i+1;j<lv;j++ {
				idx2 := v[j]; x2,y2 := X[idx2],Y[idx2]
				dx,dy := (x2-x1),(y2-y1)
				if dx*dx+dy*dy <= limit { ansarr = append(ansarr,pair{idx1,idx2}) }
			}
		}
	}
	processCellPair := func(v1,v2 []int) {
		for _,idx1 := range v1 {
			x1,y1 := X[idx1],Y[idx1]
			for _,idx2 := range v2 {
				x2,y2 := X[idx2],Y[idx2]
				dx,dy := (x2-x1),(y2-y1)
				if dx*dx+dy*dy <= limit { ansarr = append(ansarr,pair{idx1,idx2}) }
			}
		}
	}
	for b,v := range h {
		processCell(v)
		b1,b2,b3,b4 := bucket{b.i+1,b.j},bucket{b.i-1,b.j+1},bucket{b.i,b.j+1},bucket{b.i+1,b.j+1}
		v1,ok1 := h[b1]; if ok1 { processCellPair(v,v1) }
		v2,ok2 := h[b2]; if ok2 { processCellPair(v,v2) }
		v3,ok3 := h[b3]; if ok3 { processCellPair(v,v3) }
		v4,ok4 := h[b4]; if ok4 { processCellPair(v,v4) }
	}
	for i,v := range ansarr { if v.i > v.j { ansarr[i] = pair{v.j,v.i} } }
	sort.Slice(ansarr,func(i,j int) bool { return ansarr[i].i < ansarr[j].i || ansarr[i].i == ansarr[j].i && ansarr[i].j < ansarr[j].j } )
	fmt.Fprintln(wrtr,len(ansarr))
	for _,p := range ansarr { fmt.Fprintf(wrtr,"%v %v\n",p.i+1,p.j+1) }
}
