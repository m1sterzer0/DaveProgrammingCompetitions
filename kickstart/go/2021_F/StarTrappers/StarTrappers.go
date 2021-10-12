package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
const inf float64 = 1e99
type Pt2 struct{ x, y int }
func ptsub(a, b Pt2) Pt2 { return Pt2{a.x - b.x, a.y - b.y} }
func dot2(a, b Pt2) int { return a.x*b.x + a.y*b.y }
func cross2(a, b Pt2) int { return a.x*b.y - a.y*b.x }
func normb(orig,a Pt2) float64 { x := ptsub(a, orig); return math.Sqrt(float64(dot2(x, x))) }
func doTriangles(N int, pts []Pt2) float64 {
	best := inf
	for i:=0;i<N;i++ {
		p1 := pts[i]
		for j:=i+1;j<N;j++ {
			p2 := pts[j]
			c1 := cross2(p2,p1)
			for k:=j+1;k<N;k++ {
				p3 := pts[k]
				c2 := cross2(p3,p2)
				c3 := cross2(p1,p3)
				if c1 > 0 && c2 > 0 && c3 > 0 || c1 < 0 && c2 < 0 && c3 < 0 { 
					cand := normb(p1,p2) + normb(p2,p3) + normb(p3,p1)
					if cand < best { best = cand }
				}
			}
		}
	}
	return best
}
func doQuads(N int, pts []Pt2) float64 {
	best := inf
	for i:=0;i<N;i++ {
		p1 := pts[i]
		for j:=i+1;j<N;j++ {
			p2 := pts[j]
			if cross2(p1,p2) != 0 { continue }
			if p1.x > 0 && p2.x > 0 || p1.x < 0 && p2.x < 0 { continue }
			if p1.y > 0 && p2.y > 0 || p1.y < 0 && p2.y < 0 { continue }
			s1 := inf; s2 := inf; v1 := ptsub(p2,p1)
			for k:=0;k<N;k++ {
				p3 := pts[k]
				xx := cross2(v1,ptsub(p3,p1))
				if xx == 0 { continue }
				d := normb(p3,p1) + normb(p3,p2)
				if xx > 0 && d < s1 { s1 = d }
				if xx < 0 && d < s2 { s2 = d }
			}
			if s1+s2 < best { best = s1+s2 }
		}
	}
	return best
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y := fill2(N); Xs,Ys := gi2(); for i:=0;i<N;i++ { X[i] -= Xs; Y[i] -= Ys }
		pts := make([]Pt2,N); for i:=0;i<N;i++ { pts[i] = Pt2{X[i],Y[i]} }
		ans1 := doTriangles(N,pts)
		ans2 := doQuads(N,pts)
		if ans2 < ans1 { ans1 = ans2 }
		if ans1 == inf {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans1)
		}
    }
}

