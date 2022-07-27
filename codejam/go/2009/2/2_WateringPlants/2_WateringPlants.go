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
func absf(a float64) float64 { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func maxf(a,b float64) float64 { if a > b { return a }; return b }
func minf(a,b float64) float64 { if a > b { return b }; return a }
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


func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N := gi(); X,Y,R := fill3(N)
		XF := make([]float64,N); for i,x := range X { XF[i] = float64(x); }
		YF := make([]float64,N); for i,y := range Y { YF[i] = float64(y); }
		RF := make([]float64,N); for i,r := range R { RF[i] = float64(r); }

		// Finding the center of a circle of Radius R that is tangent to 2 circles: (x1,y1,r1) and (x2,y2,r2):
		// Consider points relative to center x1,y1  -- unit vector v1 points from x1,r1 to x2,y2, and unit vector v2 is ortho to v1.
		// express the center of the circumcircle as (x1,y1) + u * v1 + v * v2, and express (x2,y2) as (x1,y1) + d * v1
		// then we have that
		//    u^2 + v^2 = (R-r1)^2 and (d-u)^2 + v^2 = (R-r2)^2 (Note that u can be negative).
		// Solving, we get that u = (d^2 - (R-r2)^2 + (R-r1)^2) / 2d
		type center struct {i,j int; x,y float64}
		centers := make([]center,0,N*N+N)
		// Binary search on R -- consider all possible circle centers which are either the center of an exisiting circle or
		// centers which tangentially cover a pair of existing circles.
		testr := func(R float64) bool {
			centers = centers[:0]
			for i:=0;i<N;i++ {
				if RF[i] <= R { centers = append(centers,center{i,i,XF[i],YF[i]}) }
				x1,y1,r1 := XF[i],YF[i],RF[i]
				for j:=i+1;j<N;j++ {
					x2,y2,r2 := XF[j],YF[j],RF[j]
					d := math.Sqrt((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1))
					if R < 0.5*(r1+d+r2) { continue }
					vx1 := (x2-x1)/d; vy1 := (y2-y1)/d; vx2 := -vy1; vy2 := vx1 // Basis vectors
					u := (d*d-(R-r2)*(R-r2)+(R-r1)*(R-r1))*0.5/d
					v := math.Sqrt((R-r1)*(R-r1)-u*u)
					xa,ya := x1+u*vx1+v*vx2,y1+u*vy1+v*vy2
					xb,yb := x1+u*vx1-v*vx2,y1+u*vy1-v*vy2
					centers = append(centers,center{i,j,xa,ya})
					centers = append(centers,center{i,j,xb,yb})
				}
			}
			nc := len(centers)
			for i:=0;i<nc;i++ {
				i1,j1,xa,ya := centers[i].i,centers[i].j,centers[i].x,centers[i].y
				for j:=i;j<nc;j++ { // Start at i to cover case of one plant w/o special case
					i2,j2,xb,yb := centers[j].i,centers[j].j,centers[j].x,centers[j].y
					good := true
					for k:=0;k<N;k++ {
						if k == i1 || k == i2 || k == j1 || k == j2 { continue }
						xk,yk,rk := XF[k],YF[k],RF[k]
						if (xk-xa)*(xk-xa)+(yk-ya)*(yk-ya) <= (R-rk)*(R-rk) { continue }
						if (xk-xb)*(xk-xb)+(yk-yb)*(yk-yb) <= (R-rk)*(R-rk) { continue }
						good = false; break
					}
					if good { return true }
				}
			}
			return false
		}
		l,r := 0.0,1200.0
		for i:=0;i<60;i++ { //1200 * 1/2^60 should be plenty good
			m := 0.5*(l+r)
			if testr(m) { r = m } else { l = m }
		}
		ans := 0.5*(l+r)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

