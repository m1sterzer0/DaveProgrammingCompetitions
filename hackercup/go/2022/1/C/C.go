package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)

func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
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
func sortUniq(a []int) []int {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}

type Pt2 struct{ x, y int }

func ptadd(a, b Pt2) Pt2         { return Pt2{a.x + b.x, a.y + b.y} }
func ptsub(a, b Pt2) Pt2         { return Pt2{a.x - b.x, a.y - b.y} }
func ptscale(n int, a Pt2) Pt2   { return Pt2{n * a.x, n * a.y} }
func dot2(a, b Pt2) int          { return a.x*b.x + a.y*b.y }
func cross2(a, b Pt2) int        { return a.x*b.y - a.y*b.x }
func normsq2(a Pt2) int          { return dot2(a, a) }
func dot2b(orig, a, b Pt2) int   { return dot2(ptsub(a, orig), ptsub(b, orig)) }
func cross2b(orig, a, b Pt2) int { return cross2(ptsub(a, orig), ptsub(b, orig)) }
func normsq2b(orig, a Pt2) int   { x := ptsub(a, orig); return dot2(x, x) }
func sortPt2xy(a []Pt2) {
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x || a[i].x == a[j].x && a[i].y < a[j].y })
}
func sortPt2yx(a []Pt2) {
	sort.Slice(a, func(i, j int) bool { return a[i].y < a[j].y || a[i].y == a[j].y && a[i].x < a[j].x })
}
func hullGraham(a []Pt2) ([]Pt2) {
	n := len(a); if n < 3 { return a }; m := 0
	for i := 1; i < n; i++ {
		if a[i].y < a[m].y || a[i].y == a[m].y && a[i].x < a[m].x { m = i }
	}
	cand := make([]int, 0, n-1)
	for i := 0; i < n; i++ { if i != m { cand = append(cand, i) } }
	sort.Slice(cand, func(i, j int) bool {
		x := cross2b(a[m], a[cand[i]], a[cand[j]])
		return x > 0 || x == 0 && normsq2b(a[m], a[cand[i]]) < normsq2b(a[m], a[cand[j]])
	})
	C := []int{m}; l := 1
	for _, c := range cand {
		for l > 1 && cross2b(a[C[l-2]], a[C[l-1]], a[c]) <= 0 {	C = C[:l-1]; l-- }
		C = append(C, c); l++
	}
	ans := make([]Pt2, l); for i := 0; i < l; i++ { ans[i] = a[C[i]] }; return ans
}

func solve(N,K,D int,X,Y []int) int {
	inpts := make([]Pt2,N); for i:=0;i<N;i++ { inpts[i] = Pt2{X[i],Y[i]} }
	pts := hullGraham(inpts)
	n := len(pts) // Should be less than 36000
	// Find the start and end points
	st,en := -1,-1
	for i,p := range pts { if p.x == X[0] && p.y == Y[0] { st = i}; if p.x == X[N-1] && p.y == Y[N-1] { en = i } }
	inf := 1<<61
	marked := make([]bool,n)
	darr := iai(n,inf)
	darr[st] = 0
	D2 := D*D
	for ii:=0;ii<n;ii++ {
		bidx,bdist := -1,inf
		for i,d := range darr { if !marked[i] && d < bdist { bidx = i; bdist = d } }
		if bidx == -1 || bidx == en { break }
		marked[bidx] = true
		d,x,y := darr[bidx],pts[bidx].x,pts[bidx].y
		for j,p := range pts {
			if marked[j] { continue }
			dd := (p.x-x)*(p.x-x)+(p.y-y)*(p.y-y)
			if dd > D2 { continue }
			if dd < K { dd = K }
			darr[j] = min(darr[j],d+dd)
		}
	}
	ans := -1; if darr[en] != inf { ans = darr[en] }
	return ans
}

func testBig() {
	ss := make([]Pt2,0,9000)
	xsum := 0; ysum := 0
	for ls := 2; ls <= 2000000; ls++ {
		for i:=1;i<ls;i++ {
			j:=ls-i
			if gcd(i,j) != 1 { continue }
			if xsum+i < 500000 && ysum+j < 500000 { ss = append(ss,Pt2{i,j}); xsum += i; ysum += j}
		}
		if xsum+ysum+ls >= 1000000 { break }
	}
	sort.Slice(ss,func(i,j int) bool { x := cross2(ss[i],ss[j]); return x < 0 } )
	nn := len(ss)
	pts := make([]Pt2,0,4*nn)
	x,y := 1,500000
	for i:=0;i<nn;i++ { x += ss[i].x; y += ss[i].y;    pts = append(pts,Pt2{x,y}) }
	for i:=nn-1;i>=0;i-- { x += ss[i].x; y -= ss[i].y; pts = append(pts,Pt2{x,y}) }
	for i:=0;i<nn;i++ { x -= ss[i].x; y -= ss[i].y;    pts = append(pts,Pt2{x,y}) }
	for i:=nn-1;i>=0;i-- { x -= ss[i].x; y += ss[i].y; pts = append(pts,Pt2{x,y}) }
	//for i,pp := range pts { fmt.Printf("DBG i:%v pp:%v\n",i,pp) }
	sort.Slice(pts,func(i,j int) bool { return pts[i].x < pts[j].x } )
	N := len(pts)
	X,Y := ia(N),ia(N)
	for i:=0;i<N;i++ { X[i] = pts[i].x; Y[i] = pts[i].y }
	K := 0; D := 1000000000000
	fmt.Printf("Starting Solve...\n")
	start := time.Now()
	solve(N,K,D,X,Y)
	duration := time.Since(start)
	fmt.Printf("Duration: %v\n",duration)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	// In one quadrant, we can have at most around 9000 points in the convex hull (looking at possible delx and dely value and relying on 1,000,000 coord limit)
	// So the convex hull contains less than 36,000 points
	// We will try dumb dense dijkstra first and test by constructing a really large convex hull and see how fast it runs
	// There are probably tricks to speed it along, but I think optimized n^2 will likely be fast enough
	//testBig()
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N,K,D := gi(),gi(),gi(); X,Y := fill2(N)
		ans := solve(N,K,D,X,Y)
		fmt.Printf("Case #%v: %v\n",tt,ans)
	}
}

