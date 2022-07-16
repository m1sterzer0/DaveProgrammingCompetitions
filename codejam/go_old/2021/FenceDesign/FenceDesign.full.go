package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int = 2000000000000000000
const MOD int = 1000000007
type pair struct {x,y int}

func cross(x1,y1,x2,y2 int) int { return x1*y2-y1*x2 }

func test(ntc,Nmin,Nmax,Cmax int) {
	numPassed := 0
	for tt:=1;tt<=ntc;tt++{
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		s := make(map[pair]bool)
		X := make([]int,0)
		Y := make([]int,0)
		
		for len(s) < N {
			x := -Cmax + rand.Intn(2*Cmax+1)
			y := -Cmax + rand.Intn(2*Cmax+1)
			if s[pair{x,y}] { continue }
			// colinearity check -- what a pita
			good := true
			if len(s) >= 2 {
				for i:=0;i<len(s);i++ {
					for j:=i+1;j<len(s);j++ {
						x1,y1,x2,y2 := X[i],Y[i],X[j],Y[j]
						if cross(x2-x1,y2-y1,x-x1,y-y1) == 0 { good = false }
					}
				}
				if !good { continue }
			}
			X = append(X,x); Y = append(Y,y); s[pair{x,y}] = true
		}

		P := []int{0,0}; Q := []int{0,0}
		for {
			P[0] = 1+rand.Intn(N)
			P[1] = 1+rand.Intn(N)
			Q[0] = 1+rand.Intn(N)
			Q[1] = 1+rand.Intn(N)
			if P[0] == Q[0] || P[1] == Q[1] { continue } // Can't have the same coordinate on both ends of a fencepost
			if P[0] == P[1] && Q[0] == Q[1] { continue } // Can't have a duplicate fencepost
			if P[0] == Q[1] && Q[0] == P[1] { continue } // Can't have a duplicate fencepost
			if P[0] == P[1] || P[0] == Q[1] || Q[0] == P[1] || Q[0] == Q[1] { break }
			x1,y1 := X[P[0]-1],Y[P[0]-1]
			x2,y2 := X[Q[0]-1],Y[Q[0]-1]
			x3,y3 := X[P[1]-1],Y[P[1]-1]
			x4,y4 := X[Q[1]-1],Y[Q[1]-1]
			c1 := cross(x2-x1,y2-y1,x3-x1,y3-y1) 
			c2 := cross(x2-x1,y2-y1,x4-x1,y4-y1)
			c3 := cross(x4-x3,y4-y3,x1-x3,y1-y3)
			c4 := cross(x4-x3,y4-y3,x2-x3,y2-y3)
			if c1 < 0 && c2 < 0 { break }
			if c1 > 0 && c2 > 0 { break }
			if c3 < 0 && c4 < 0 { break }
			if c3 > 0 && c4 > 0 { break }
		}
		edgelist1 := solveSmall(N,X,Y,P,Q)
		edgelist2 := solve(N,X,Y,P,Q)
		if len(edgelist1) == len(edgelist2) {
			numPassed++
		} else {
			if N <= 10 { 
				fmt.Printf("ERROR tt:%v N:%v X:%v Y:%v P:%v Q:%v len1:%v len2:%v\n    edgelist1:%v\n    edgelist2:%v\n",tt,N,X,Y,P,Q,len(edgelist1),len(edgelist2),edgelist1,edgelist2)
			} else {
				fmt.Printf("ERROR tt:%v N:%v len1:%v len2:%v\n",tt,N,len(edgelist1),len(edgelist2))
			}
		}
	}
	fmt.Printf("%v/%v passed\n",numPassed, ntc)
}

func solve(N int, X,Y,P,Q []int) []pair {
	pp0,qq0,pp1,qq1 := P[0]-1,Q[0]-1,P[1]-1,Q[1]-1
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

	var triangulate func(pts []int, v1,v2 bool) ([]pair,[]int)
	triangulate = func(pts []int, v1,v2 bool) ([]pair,[]int) {
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
		p1,p2,o1,o2 := -1,-1,-1,-1
		if v1 && v2 {
			// First time -- pick and edge where the other edge is fully on one side
			p1,p2,o1,o2 = pp0,qq0,pp1,qq1
			if p1 != o1 && p1 != o2 && p2 != o1 && p2 != o2 {
				c1 := cross(X[qq0]-X[pp0],Y[qq0]-Y[pp0],X[pp1]-X[pp0],Y[pp1]-Y[pp0])
				c2 := cross(X[qq0]-X[pp0],Y[qq0]-Y[pp0],X[qq1]-X[pp0],Y[qq1]-Y[pp0])
				if c1 > 0 && c2 < 0 || c1 < 0 && c2 > 0 { p1,p2,o1,o2 = pp1,qq1,pp0,qq0 }
			}
		} else if v1 {
			p1,p2 = pp0,qq0
		} else if v2 {
			p1,p2 = pp1,qq1
		} else {
			p1 = pts[rand.Intn(npts)]
			p2 = p1
			for p1==p2 { p2 = pts[rand.Intn(npts)] }
		}
		left,right := []int{},[]int{}
		for _,p := range pts {
			if p == p1 || p == p2 { continue }
			if cross(X[p2]-X[p1],Y[p2]-Y[p1],X[p]-X[p1],Y[p]-Y[p1]) > 0 { left = append(left,p) } else { right = append(right,p) }
		}
		for !v1 && !v2 && (3*len(left) < len(right) || 3*len(right) < len(left)) {
			// On bad splits, try again
			p2 = p1
			for p1==p2 { p2 = pts[rand.Intn(npts)] }
			left = left[:0]; right = right[:0]
			for _,p := range pts {
				if p == p1 || p == p2 { continue }
				c1 := cross(X[p2]-X[p1],Y[p2]-Y[p1],X[p]-X[p1],Y[p]-Y[p1])
				if c1 > 0 { left = append(left,p) } else if c1 < 0 { right = append(right,p) }
			}
		}
		left = append(left,p1); left = append(left,p2)
		right = append(right,p1); right = append(right,p2)
		var f1,f2 []pair
		var h1,h2 []int
		nv1,nv2 := false,false
		if o1 == pp0 && o2 == qq0 { nv1 = true } else if o1 == pp1 && o2 == qq1 { nv2 = true }
		lefto1,lefto2,leftContains := false,false,false
		if o1 >= 0 {
			for _,pp := range left { if pp == o1 { lefto1 = true } ; if pp == o2 { lefto2 = true } }
			leftContains = lefto1 && lefto2
		}

		if len(left) == 2 {
			f,h := triangulate(right,nv1,nv2)
			return f,h 
		} else  if len(right) == 2 {
			f,h := triangulate(left,nv1,nv2)
			return f,h 
		} else if o1 < 0 {
			f1,h1 = triangulate(left,false,false)
			f2,h2 = triangulate(right,false,false)
		} else if leftContains {
			f1,h1 = triangulate(left,nv1,nv2)
			f2,h2 = triangulate(right,false,false)
		} else  {
			f1,h1 = triangulate(left,false,false)
			f2,h2 = triangulate(right,nv1,nv2)
		}
		p1h1,p1h2,p2h1,p2h2 := false,false,false,false
		for _,hh := range h1 { if hh == p1 { p1h1 = true; break } }
		for _,hh := range h1 { if hh == p2 { p2h1 = true; break } }
		for _,hh := range h2 { if hh == p1 { p1h2 = true; break } }
		for _,hh := range h2 { if hh == p2 { p2h2 = true; break } }
		if !(p1h1 && p1h2 && p2h1 && p2h2) {
			fmt.Printf("ERROR IN HULL: h1:%v h2:%v p1:%v p2:%v\n",h1,h2,p1,p2)
			os.Exit(1)
		}

		f,h := tricombine(f1,f2,h1,h2,p1,p2)
		return f,h
	}

	pref,_ := triangulate(pts,true,true)
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

// This version doesn't handle colinearity, and returns
// false if they merely touch at one or two endpoints
func doesIntersect(x1,y1,x2,y2,x3,y3,x4,y4 int) bool {
	cross1 := cross(x2-x1,y2-y1,x3-x1,y3-y1)
	cross2 := cross(x2-x1,y2-y1,x4-x1,y4-y1)
	if (cross1 <= 0) && (cross2 <= 0) { return false }
	if (cross1 >= 0) && (cross2 >= 0) { return false }
	cross3 := cross(x4-x3,y4-y3,x1-x3,y1-y3)
	cross4 := cross(x4-x3,y4-y3,x2-x3,y2-y3)
	if (cross3 <= 0) && (cross4 <= 0) { return false }
	if (cross3 >= 0) && (cross4 >= 0) { return false }
	return true
}

func solveSmall(N int, X,Y,P,Q []int) []pair {
	pp0,qq0,pp1,qq1 := P[0]-1,Q[0]-1,P[1]-1,Q[1]-1
	ans := []pair{{pp0,qq0},{pp1,qq1}}
	for i:=0;i<N;i++ {
		for j:=i+1;j<N;j++ {
			if i == pp0 && j == qq0 { continue }
			if i == pp1 && j == qq1 { continue }
			if i == qq0 && j == pp0 { continue }
			if i == qq1 && j == pp1 { continue }
			good := true
			x1,y1,x2,y2 := X[i],Y[i],X[j],Y[j]
			for _,p := range ans {
				x3,y3,x4,y4 := X[p.x],Y[p.x],X[p.y],Y[p.y]
				if doesIntersect(x1,y1,x2,y2,x3,y3,x4,y4) { good = false; break }
			}
			if good { ans = append(ans,pair{i,j}) }
		}
	}
	lans := len(ans)
	for i:=0;i<lans;i++ { ans[i].x++; ans[i].y++ }
	ans = ans[2:] // Trim out the given segments
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	rand.Seed(8675309)
	//test(10,4,50,10000)
	//test(100,4,50,10000)
	//test(1000,4,50,10000)
	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y := fill2(N); P,Q := fill2(2)
		//ans := solveSmall(N,X,Y,P,Q)
		ans := solve(N,X,Y,P,Q)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,len(ans))
		for _,pp := range ans {	fmt.Fprintf(wrtr,"%v %v\n",pp.x,pp.y)}
		wrtr.Flush()
	}
}
