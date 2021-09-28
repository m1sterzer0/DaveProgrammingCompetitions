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
type Pt2 struct{ x, y int }
func ptsub(a, b Pt2) Pt2 { return Pt2{a.x - b.x, a.y - b.y} }
func dot2(a, b Pt2) int { return a.x*b.x + a.y*b.y }
func cross2(a, b Pt2) int { return a.x*b.y - a.y*b.x }
func dot2b(orig, a, b Pt2) int { return dot2(ptsub(a, orig), ptsub(b, orig)) }
func cross2b(orig, a, b Pt2) int { return cross2(ptsub(a, orig), ptsub(b, orig)) }
func normsq2b(orig, a Pt2) int { x := ptsub(a, orig); return dot2(x, x) }
type sig struct {normsq,dp,cp int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A,B := fill2(N); C,D := fill2(N)
	if N == 1 { fmt.Println("Yes"); return }
	AA := make([]Pt2,N); for i:=0;i<N;i++ { AA[i] = Pt2{A[i],B[i]} }
	BB := make([]Pt2,N); for i:=0;i<N;i++ { BB[i] = Pt2{C[i],D[i]} }
	sigs := make(map[sig]bool)
	o1,v1 := AA[0],AA[1]; refnormsq := normsq2b(o1,v1)

	for i:=0;i<N;i++ { xx1,xx2,xx3 := normsq2b(o1,AA[i]),dot2b(o1,v1,AA[i]),cross2b(o1,v1,AA[i]); sigs[sig{xx1,xx2,xx3}] = true }
	good := false
	for i:=0;i<N;i++ {
		o2 := BB[i]
		for j:=0;j<N;j++ {
			v2 := BB[j]
			if normsq2b(o2,v2) != refnormsq { continue }
			good = true
			for k:=0;k<N;k++ {
				xx1,xx2,xx3 := normsq2b(o2,BB[k]),dot2b(o2,v2,BB[k]),cross2b(o2,v2,BB[k])
				if !sigs[sig{xx1,xx2,xx3}] { good = false; break }
			}
			if good { break }
		}
		if good { break }
	}
	ans := "No"; if good { ans = "Yes" }; fmt.Println(ans)
}

