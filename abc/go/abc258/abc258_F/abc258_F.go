package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	type leg struct { d,x,y int }
	src := []leg{}
	dest := []leg{}
	fillCoords := func (res []leg, x,y,b,k int) []leg {
		res = res[:0]
		if x % b == 0 || y % b == 0 {
			res = append(res,leg{0,x,y}) 
		} else {
			mx := x % b; my := y % b
			res = append(res,leg{k*mx,x-mx,y}) 
			res = append(res,leg{k*(b-mx),x-mx+b,y}) 
			res = append(res,leg{k*my,x,y-my}) 
			res = append(res,leg{k*(b-my),x,y-my+b})
		}
		return res
	}
	calcMainRoadDist := func(x1,y1,x2,y2,B int) int {
		mh := abs(x2-x1)+abs(y2-y1)
		if x1 % B == 0 && y1 % B == 0 || x2 % B == 0 && y2 % B == 0 { return mh } // On a corner
		if x1 % B == 0 && y2 % B == 0 || x2 % B == 0 && y1 % B == 0 { return mh } // Horiz + Vert
		if x1 % B == 0 && x2 % B == 0 && x1 == x2     { return mh } // Same horiz road
		if y1 % B == 0 && y2 % B == 0 && y1 == y2     { return mh } // Same vert road
		if x1 % B == 0 && x2 % B == 0 { // two vertical roads
			if y1 / B != y2 / B { return mh } else { xd := abs(x2-x1); return min(xd+y1%B+y2%B,xd+B-y1%B+B-y2%B) }
		} else { // two horizontal roads
			if x1 / B != x2 / B { return mh } else { yd := abs(y2-y1); return min(yd+x1%B+x2%B,yd+B-x1%B+B-x2%B) }
		}
	}
	for tt:=1;tt<=T;tt++ {
		B,K,Sx,Sy,Gx,Gy := gi(),gi(),gi(),gi(),gi(),gi()
		best := K * (abs(Gx-Sx)+abs(Gy-Sy)) // Use no main roads
		src = fillCoords(src,Sx,Sy,B,K)
		dest = fillCoords(dest,Gx,Gy,B,K)
		for _,s := range src {
			for _,d := range dest {
				cand := s.d + d.d + calcMainRoadDist(s.x,s.y,d.x,d.y,B)
				best = min(best,cand)
			}
		}
		fmt.Fprintln(wrtr,best)
	}
}
