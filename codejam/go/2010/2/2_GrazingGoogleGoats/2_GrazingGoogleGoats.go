package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
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

func vecfloatstring(a []float64) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = fmt.Sprintf("%.17g",a) }; return strings.Join(astr," ") }
func rot2d(x,y,ang float64) (float64,float64) { 
	return x*math.Cos(ang)-y*math.Sin(ang),x*math.Sin(ang)+y*math.Cos(ang)
}
func invertPoint(x,y float64) (float64,float64) {
	return x/(x*x+y*y),y/(x*x+y*y)
}
func invertCircleToLine(rx,ry float64) (float64,float64,float64) {
	l := math.Sqrt(rx*rx+ry*ry)
	A,B := rx/l,ry/l
	x,y := 2.0*rx,2.0*ry;
	xi,yi := invertPoint(x,y)
	if B < 0 { A,B = -A,-B }; C := A*xi+B*yi
	return A,B,C
}
func lineSlopeGt(a,b line) bool {
	if a.b < 1e-12 { return false } // Shouldn't happen
	if b.b < 1e-12 { return true  } // Shouldn't happen
	m1 := -a.a/a.b; m2 := -b.a/b.b; return m1 > m2;
}
func intersection(l1,l2 line) (float64,float64) {
	denom := l1.a*l2.b-l1.b*l2.a
	numx  := l1.c*l2.b-l1.b*l2.c
	numy  := l1.a*l2.c-l1.c*l2.a
	return numx/denom,numy/denom
}
func circleSegment(x1,y1,x2,y2,r float64) float64 {
	ang1 := math.Atan2(y1,x1)
	ang2 := math.Atan2(y2,x2)
	ang := ang2-ang1; for ang < 0 { ang += 2.0 * math.Pi }
	ans := 0.5*r*r*ang
	if ang > math.Pi { ans += 0.5*r*r*math.Sin(2.0*math.Pi-ang) } else { ans -= 0.5*r*r*math.Sin(ang) }
	return ans
}

func polyArea(a []pt) float64 {
	n := len(a)
	area := 0.0
	for i:=0;i<n-1;i++ { area += a[i].x*a[i+1].y-a[i].y*a[i+1].x }
	area += a[n-1].x*a[0].y-a[n-1].y*a[0].x
	if area < 0 { area = -area }
	return 0.5*area
}

type line struct { a,b,c,cx,cy,r float64 }
type pt struct { x,y float64 }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		eps := 1e-12;
		N,M := gi(),gi(); PX,PY := fill2(N); QX,QY := fill2(M)
		ansarr := make([]float64,M)
		angs := make([]float64,N)
		lines := make([]line,N)
		st := make([]line,0)
		poly := make([]pt,0)

		for idx:=0;idx<M;idx++ {
			// Step 1: make sure that the centers are in a 180 degree arc
			//         Why? : locally around Q, all of the circles will look like half planes.  There
			//         will not be an intersection unless all of these half planes overlap, meaning that the
			//         centers are contained in some 90 degree arc.
			for i:=0;i<N;i++ { angs[i] = math.Atan2(float64(PY[i]-QY[idx]),float64(PX[i]-QX[idx])) }
			sort.Slice(angs,func(i,j int) bool { return angs[i] < angs[j] } )
			bestgap := math.Pi * 2.0 + angs[0] - angs[N-1]; bestidx := 0
			for i:=1;i<N;i++ { cand := angs[i]-angs[i-1]; if cand > bestgap { bestgap = cand; bestidx = i } }
			if bestgap <= math.Pi+eps { ansarr[idx] = 0.0; continue }
			// Step 2: Rotate so the center of the gap is pointing up -- all points are below Q
			// Step3 : Invert all of the circles into lines of form Ax+By=C
			rotstart := angs[bestidx] - 0.5 * bestgap; rotend := 0.5 * math.Pi; rotang := rotend - rotstart;
			for i:=0;i<N;i++ { 
				rx,ry := rot2d(float64(PX[i]-QX[idx]),float64(PY[i]-QY[idx]),rotang)
				a,b,c := invertCircleToLine(rx,ry)
				lines[i] = line{a,b,c,rx,ry,math.Sqrt(rx*rx+ry*ry)}
			}
			// Step4 : Sort lines and do lower envelope stack (alternatively could use point line duality, but this works too)
			sort.Slice(lines,func(i,j int) bool { return lineSlopeGt(lines[i],lines[j]) } )
			st := st[:0]; ls := 0
			for _,l := range lines { 
				for ls >= 2 { 
					x1,_ := intersection(st[ls-2],st[ls-1])
					x2,_ := intersection(st[ls-2],l)
					if x2 >= x1 { break } else { ls--; st = st[:ls] }
				}
				st = append(st,l); ls++
			}
			// Step 5: Calculate Area
			lx,ly := 0.0,0.0; poly = poly[:0]; area := 0.00; stn := len(st)
			poly = append(poly,pt{0.0,0.0})
			for i:=0;i+1<stn;i++ {
				xx,yy := intersection(st[i],st[i+1])
				x,y := invertPoint(xx,yy)
				poly = append(poly,pt{x,y})
				adder := circleSegment(lx-st[i].cx,ly-st[i].cy,x-st[i].cx,y-st[i].cy,st[i].r)
				area += adder
				lx,ly = x,y
			}
			adder := circleSegment(lx-st[stn-1].cx,ly-st[stn-1].cy,0-st[stn-1].cx,0-st[stn-1].cy,st[stn-1].r) 
			area += adder
			area += polyArea(poly)
			ansarr[idx] = area
		}
        fmt.Printf("Case #%v: %v\n",tt,vecfloatstring(ansarr))
	}
}
