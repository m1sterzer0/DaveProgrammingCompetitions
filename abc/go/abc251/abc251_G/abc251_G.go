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

// Inspired by Gennady's Library
type Pt struct {x,y int}
func addpt(a,b Pt) Pt { return Pt{a.x+b.x,a.y+b.y} }
func subpt(a,b Pt) Pt { return Pt{a.x-b.x,a.y-b.y} }
func scalept(a int,b Pt) Pt { return Pt{a*b.x,a*b.y} }
func dprod(a,b Pt) int { return a.x*b.x + a.y*b.y }
func cprod(a,b Pt) int { return a.x*b.y - a.y*b.x }
func abs2(a Pt) int { return a.x*a.x+a.y*a.y }
func ltpt(a,b Pt) bool { return a.x < b.x || a.x == b.x && a.y < b.y }
func isUpper(a Pt) bool { return a.y > 0 || a.y == 0 && a.x > 0 }
func cmpPolar(a,b Pt) int { 
	aa,bb := isUpper(a),isUpper(b)
	if aa != bb { if aa { return -1 } else { return 1 } }
	v := cprod(a,b); if v < 0 { return -1 } else if v > 0 { return 1 } else { return 0 }
}
type hp struct { p1,p2 Pt }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N); M := gi(); U,V := fill2(M); Q := gi(); A,B := fill2(Q)
	darr := make([]Pt,M)
	for i:=0;i<M;i++ { darr[i] = Pt{U[i],V[i]} }
	hparr := make([]hp,N) //half plane array
	for i:=0;i<N;i++ {
		j := (i+1)%N
		p1,p2 := Pt{X[i],Y[i]},Pt{X[j],Y[j]}
		cp1,cp2 := addpt(p1,darr[0]),addpt(p2,darr[0])
		for _,d := range darr {
			if cprod(subpt(cp2,cp1),subpt(addpt(p1,d),cp1)) > 0 { cp1,cp2=addpt(p1,d),addpt(p2,d) }
		}
		hparr[i] = hp{cp1,cp2}
	}
	ansarr := make([]string,Q)
	for i:=0;i<Q;i++ {
		t := Pt{A[i],B[i]}
		good := true
		for _,h := range hparr {
			pp1,pp2 := h.p1,h.p2
			if cprod(subpt(pp2,pp1),subpt(t,pp1)) < 0 { good = false; break }
		}
		if good { ansarr[i] = "Yes" } else { ansarr[i] = "No" }
	}
	for _,s := range ansarr { fmt.Fprintln(wrtr,s) }
}

