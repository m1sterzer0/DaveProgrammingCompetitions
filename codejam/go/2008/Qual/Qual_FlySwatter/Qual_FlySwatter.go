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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		f,R,t,r,g := gf(),gf(),gf(),gf(),gf()
		ans := 0.00
		if 2*f >= g {
			ans = 1.0
		} else {
			num := 0.0
			num2 := 0.0
			denom := 0.25 * math.Pi * R * R
			inner := R-t-f
			inner2 := inner*inner
			sliver := func(x1,y1,x2,y2 float64) float64{
				ang := math.Atan2(y2,x2)-math.Atan2(y1,x1)
				return 0.5*inner2*(ang - math.Sin(ang))
			}
			for x:=r+f; x < R-t; x+=2*r+g {
				for y:=r+f; x*x + y*y < inner2; y+=2*r+g {
					x2 := x-f+g-f; y2 := y-f+g-f
					if x2*x2+y2*y2 <= inner2 { // Entire square is good
						num += (x2-x)*(y2-x)
					} else {
						c1,c2 := x*x+y2*y2 < inner2, x2*x2+y*y < inner2
						if c1 && c2 {
							x3 := math.Sqrt(inner2-y2*y2)
							y3 := math.Sqrt(inner2-x2*x2)
							num += (x2-x)*(y3-y) + 0.5*(x2-x+x3-x)*(y2-y3)
							num2 += sliver(x2,y3,x3,y2)
						} else if !c1 && !c2 {
							x3 := math.Sqrt(inner2-y*y)
							y3 := math.Sqrt(inner2-x*x)
							num += 0.5*(x3-x)*(y3-y)
							num2 += sliver(x3,y,x,y3)
						} else if !c1 {
							y3 := math.Sqrt(inner2-x*x)
							y4 := math.Sqrt(inner2-x2*x2)
							num += 0.5*(y4-y+y3-y)*(x2-x)
							num2 += sliver(x2,y4,x,y3)
						} else {
							x3 := math.Sqrt(inner2-y*y)
							x4 := math.Sqrt(inner2-y2*y2)
							num += 0.5*(x4-x+x3-x)*(y2-y)
							num2 += sliver(x3,y,x4,y2)
						}
					}
				}
			}
			num += num2
			ans = (denom - num) / denom
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

