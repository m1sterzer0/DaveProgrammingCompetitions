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


type line struct { a,b,c,r float64; idx int }
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
		RX,RY := make([]float64,N),make([]float64,N)
		lines := make([]line,N)
		st := make([]line,0)

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
			rotstart := angs[bestidx] - 0.5 * bestgap; rotend := 0.5 * math.Pi(); rotang := rotend - rotstart;
			for i:=0;i<N;i++ { 
				rx,ry := rot2d(float64(PX[i]-QX[idx]),float64(PY[i]-QY[idx]),rotang)
				a,b,c := invertCircleToLine(rx,ry)
				lines[i] = line{a,b,c,math.Sqrt(rx*rx+ry*ry),i}
			}
			// Step4 : Sort lines and do lower envelope stack (alternatively could use point line duality, but this works too)
			sort.Slice(lines,func(i,j int) bool { return lineSlopeGt(lines[i],lines[j]) } )
			st := st[:0]; ls := 0
			for _,l := range lines { 
				for ls >= 2 { 
					x1,y1 := intersection(st[ls-2],st[ls-1])
					x2,y2 := intersection(st[ls-2],l)
					if x2 >= x1 { break } else { ls--; st = st[:ls] }
				}
				st = append(st,l)
			}



					
					st = append(st,l); ls++; break }





				for {
					x1,y1 := intersection(st[ls-2],st[ls-1])
					x2,y2 := intersection(st[ls-1],l)
					x2,y2 := intersection(st[ls-1],l)

				}
			}



			// Step5 : Calculate the area
			









		








			qx,qy := QX[idx],QY[idx]
			// Invert circles to lines
			for i:=0;i<N;i++ {
				xc,yc := float64(PX[i]-qx), float64(PY[i]-qy)
				r2 := xc*xc+yc*yc; r := math.Sqrt(r2)
				vx,vy := -yc/r,xc/r
				p1x,p1y,p2x,p2y := xc+r*vx,yc+r*vy,xc-r*vx,yc-r*vy
				r1s := p1x*p1x+p1y*p1y; r2s := p2x*p2x+p2y*p2y
				q1x,q1y,q2x,q2y := p1x/r1s,p1y/r1s,p2x/r2s,p2y/r2s
				A[i] = q1y - q2y; B[i] = q2x - q1x; C[i] = A[i] * q1x + B[i] * q1y
			}



			


				
				





				r2 := X[i]*X[i]+Y[i]*Y[i]
				XX[i] = X[i]*(r2/r); YY[i] = Y[i]*(r2/r)

			}
			// Check angle range, and bail if all lines aren't within 180 degrees
			// Do the rotation
			// Sort the lines by slope
			// Do the lower envelope thing
			// Invert back
			// Polygon Area
			// Circle parts
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,vecfloatstring(ansarr))
    }
}

