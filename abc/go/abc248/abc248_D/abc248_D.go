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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
type qq struct { pos, sgn, idx, x int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); Q := gi(); L,R,X := fill3(Q); for i:=0;i<Q;i++ { L[i]--; R[i]-- }
	events := make([]qq,0,2*Q)
	for i:=0;i<Q;i++ { 
		l,r,x := L[i],R[i],X[i]
		if l != 0 { events = append(events,qq{l-1,-1,i,x}) }
		events = append(events,qq{r,1,i,x})
	}
	sort.Slice(events,func(i,j int) bool { return events[i].pos < events[j].pos })
	sb := ia(N+1); ansarr := ia(Q); ptr := 0
	for _,e := range events {
		for ptr <= e.pos { sb[A[ptr]]++; ptr++ }
		ansarr[e.idx] += e.sgn * sb[e.x]
	}
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}

