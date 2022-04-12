package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type pair struct {x,y int}

func cross(x1,y1,x2,y2 int) int { return x1*y2-y1*x2 }

func solve(N int, X,Y,P,Q []int) []pair {
	pp0,qq0,pp1,qq1 := P[0]-1,Q[0]-1,P[1]-1,Q[1]-1

	// Set it up so that pp0,qq0 is first cut, and pp1,qq1 is second cut
	if pp0 != pp1 && pp0 != qq1 && qq0 != pp1 && qq0 != qq1 { // No common endpoint
		c1 := cross(X[qq0]-X[pp0],Y[qq0]-Y[pp0],X[pp1]-X[pp0],Y[pp1]-Y[pp0])
		c2 := cross(X[qq0]-X[pp0],Y[qq0]-Y[pp0],X[qq1]-X[pp0],Y[qq1]-Y[pp0])
		if c1 > 0 && c2 < 0 || c1 < 0 && c2 > 0 { pp0,qq0,pp1,qq1 = pp1,qq1,pp0,qq0 }
	}
	pts := make([]int,N); for i:=0;i<N;i++ { pts[i] = i }

	isRightTurn := func(p1,p2,p3 int) bool {
		return cross(X[p2]-X[p1],Y[p2]-Y[p1],X[p3]-X[p1],Y[p3]-Y[p1]) < 0
	}

	tricombine := func(f1,f2 []pair, h1,h2 []int, p1,p2 int) ([]pair,[]int) {
		f := []pair{}
		for _,ff := range f1 { f = append(f,ff) }
		for _,ff := range f2 { 
			if ff.x == p1 && ff.y == p2 || ff.x == p2 && ff.y == p1 { continue }
			f = append(f,ff)
		}
		// Rotate the hulls so that they stop at the endpoints
		for h1[0] != p1 && h1[0] != p2 { h1 = append(h1,h1[0]); h1 = h1[1:] }
		for h1[1] == p1 || h1[1] == p2 { h1 = append(h1,h1[0]); h1 = h1[1:] }
		for h2[0] != p1 && h2[0] != p2 { h2 = append(h2,h2[0]); h2 = h2[1:] }
		for h2[1] == p1 || h2[1] == p2 { h2 = append(h2,h2[0]); h2 = h2[1:] }

		// Do the "bottom" first
		for {
			lh1 := len(h1)
			if h1[lh1-1] == h2[0] { h2 = h2[1:]; continue }
			if isRightTurn(h1[lh1-2],h1[lh1-1],h2[0]) { 
				f = append(f,pair{h1[lh1-2],h2[0]})
				h1 = h1[:lh1-1]
				continue
			}
			if isRightTurn(h1[lh1-1],h2[0],h2[1]) {
				f = append(f,pair{h1[lh1-1],h2[1]})
				h2 = h2[1:]
				continue
			}
			break
		}

		// Now make the full hull chain
		h := []int{}
		for _,hh := range h1 { h = append(h,hh) }
		for _,hh := range h2 { h = append(h,hh) }

		// Now fix the endpoints
		for {
			lh := len(h)
			if h[lh-1] == h[0] { h = h[1:]; continue }
			if isRightTurn(h[lh-2],h[lh-1],h[0]) { 
				f = append(f,pair{h[lh-2],h[0]})
				h = h[:lh-1]
				continue
			}
			if isRightTurn(h[lh-1],h[0],h[1]) {
				f = append(f,pair{h[lh-1],h[1]})
				h = h[1:]
				continue
			}
			break
		}
		return f,h
	}

	var triangulate func(pts []int, depth int) ([]pair,[]int)
	triangulate = func(pts []int, depth int) ([]pair,[]int) {
		npts := len(pts)
		if npts == 3 {
			res1 := []pair{}
			res1 = append(res1,pair{pts[0],pts[1]})
			res1 = append(res1,pair{pts[1],pts[2]})
			res1 = append(res1,pair{pts[2],pts[0]})
			// TODO: Make the hull counterclockwise
			c1 := cross(X[pts[2]]-X[pts[0]],Y[pts[2]]-Y[pts[0]],X[pts[1]]-X[pts[0]],Y[pts[1]]-Y[pts[0]])
			var res2 []int
			if c1 < 0 { res2 = []int{pts[0],pts[1],pts[2]} } else { res2 = []int{pts[0],pts[2],pts[1]} }
			return res1,res2
		}
		p1,p2 := -1,-1
		if depth == 1 {
			p1,p2 = pp0,qq0
		} else if depth == 2 {
			fpp1,fqq1 := false,false
			for _,p := range pts { if p == pp1 { fpp1 = true }; if p == qq1 { fqq1 = true } }
			if fpp1 && fqq1 { p1,p2 = pp1,qq1 }
		}
		if p1 < 0 {
			// Default case of picking two random points
			p1 = pts[rand.Intn(npts)]
			p2 = p1
			for p1==p2 { p2 = pts[rand.Intn(npts)] }
		}

		left,right := []int{},[]int{}
		for _,p := range pts {
			c1 := cross(X[p2]-X[p1],Y[p2]-Y[p1],X[p]-X[p1],Y[p]-Y[p1])
			if c1 > 0 { left = append(left,p) } else if c1 < 0 { right = append(right,p) }
		}
		if len(left) == 0   { right = append(right,p1); right = append(right,p2); return triangulate(right,depth+1) }
		if len(right) == 0  {  left = append( left,p1);  left = append( left,p2); return triangulate(left,depth+1)  }
		left = append(left,p1); left = append(left,p2)
		right = append(right,p1); right = append(right,p2)
		f1,h1 := triangulate(left,depth+1)
		f2,h2 := triangulate(right,depth+1)
		f,h := tricombine(f1,f2,h1,h2,p1,p2)
		return f,h
	}

	pref,_ := triangulate(pts,1)
	ans := []pair{}
	for _,pp := range pref {
		if pp.x == pp0 && pp.y == qq0 { continue }
		if pp.x == pp1 && pp.y == qq1 { continue }
		if pp.x == qq0 && pp.y == pp0 { continue }
		if pp.x == qq1 && pp.y == pp1 { continue }
		ans = append(ans,pair{pp.x+1,pp.y+1})
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	rand.Seed(8675309)
	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y := fill2(N); P,Q := fill2(2)
		ans := solve(N,X,Y,P,Q)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,len(ans))
		for _,pp := range ans {	fmt.Fprintf(wrtr,"%v %v\n",pp.x,pp.y)}
		wrtr.Flush()
	}
}
