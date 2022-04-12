package main

import (
	"bufio"
	"fmt"
	"math"
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
const finf float64 = 1e70
const feps float64 = 1e-70
const MOD int = 1000000007

// Change of coordinates
// Room goes from -Q to Q in x.
// Minimize distance from (x,a(Q^2-x^2)) to (c,d)
// Distance Squared is (x-c)^2 + (a(Q^2-x^2) - d)^2 = a^2x^4 + (1 - 2*Q^2*a^2 + 2ad)x^2 - 2cx + (some constant terms)
// Setting derivative equal to zero, we get 4a^2x^3 + (2 - 4*Q^2*a^2 + 4ad)x - 2c
// Either have 1 or 3 monotonic regions of derivative.
// Monotonic region breakpoints at 12a^2x^2 + (2 - 4Q^2a^2 + 4ad) == 0
// If (2 - 4Q^2a^2 + 4ad) < 0 , then we have 3 regions, otherwise we have 1
// If we have one region, bin search for root
// If we have three regions, check to see if each brackets zero, and do a search
func findDistSquared(a,c,d,Q float64) float64 {
	best := finf
	c1 := 12*a*a; c2 := 2 * 4*Q*Q*a*a + 4*a*d
	if c2 < 0 {
		// Potentially 3 regions to check
		r2 := math.Sqrt(-c2/(c1+feps)); r1 := -r2
		if r1 < Q {
			for _,x := range []float64{-Q,r1,r2,Q} { cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand }}
			v1,v2,v3,v4 := evalDeriv(a,c,d,Q,-Q),evalDeriv(a,c,d,Q,r1),evalDeriv(a,c,d,Q,r2),evalDeriv(a,c,d,Q,Q)
			if v1*v2 < 0 { x := findZeroDeriv(a,c,d,Q,-Q,r1); cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand } }
			if v2*v3 < 0 { x := findZeroDeriv(a,c,d,Q,r1,r2); cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand } }
			if v3*v4 < 0 { x := findZeroDeriv(a,c,d,Q,r2,Q);  cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand } }
		} else {
			for _,x := range []float64{-Q,Q} { cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand }}
			v1,v2 := evalDeriv(a,c,d,Q,-Q),evalDeriv(a,c,d,Q,Q)
			if v1*v2 < 0 { x := findZeroDeriv(a,c,d,Q,-Q,Q); cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand } }
		}
	} else {
		for _,x := range []float64{-Q,Q} { cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand }}
		v1,v2 := evalDeriv(a,c,d,Q,-Q),evalDeriv(a,c,d,Q,Q)
		if v1*v2 < 0 { x := findZeroDeriv(a,c,d,Q,-Q,Q); cand := evalDistSquared(a,c,d,Q,x); if cand < best { best = cand } }
	}
	return best
}

func evalDeriv(a,c,d,Q,x float64) float64 {	return 4*a*a*x*x*x + (2 - 4*Q*Q*a*a + 4*a*d)*x - 2*c }
func evalDistSquared(a,c,d,Q,x float64) float64 { return (x-c)*(x-c) + (a*Q*Q-a*x*x-d)*(a*Q*Q-a*x*x-d) }
func findZeroDeriv(a,c,d,Q,xl,xh float64) float64 {
	vl,vh := evalDeriv(a,c,d,Q,xl),evalDeriv(a,c,d,Q,xh)
	for xh-xl > 1e-10 { //Revisit if needed
		xm = 0.5*(xh+xl)
		y := evalDeriv(a,c,d,Q,xm)
		if y == 0 { return xm }
		if y  * vl > 0 { xl = xm } else { xh = xm }
	}
	return 0.5*(xl+xh)
}



// 







// Calculate distance from (x,ax^2+bx+c) to (m,n)
// Minimize (x-m)^2 + (ax^2+bx+c-n)^2
// Set derivative == 0 --> (4ax+2b)*(a*x*x+b*x+c-n)-2m+2x = 0
// ** Set 2nd derivative == to zero to find breakpoints in monotonic regions of derivative







func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,0)
    }
}

