package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

type rect struct { x1,x2,y1,y2 int }
type pt struct { x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	Sx,Sy,Tx,Ty := gi(),gi(),gi(),gi()
	a,b,c,d := gi(),gi(),gi(),gi()
	ans := "No"; cnt := 0; rects := make([]rect,0,1000001); rects = append(rects,rect{Sx,Sx,Sy,Sy})

	project := func(x1,x2,a,b int) (int,int) {
		z1 := x1 + 2*(a-x1)
		z2 := x2 + 2*(a-x2)
		z3 := x1 + 2*(b-x1)
		z4 := x2 + 2*(b-x2)
		return min(min(z1,z2),min(z3,z4)),max(max(z1,z2),max(z3,z4))
	}

	solveReflectionPoint := func(end,a,b,rx,ry int) (int,int) {
		xa := end+2*(a-end)
		if xa >= rx && rx <= ry { return a,end+2*(a-end) }
		if xa > ry { os.Exit(1) }
		apt := a+(rx-xa)/2
		newpt := end + 2*(apt-end)
		return apt,newpt
	}

	if (Sx==Tx && Sy==Ty) {
		ans = "Yes"
	} else if (Sx-Tx)%2==0 && (Sy-Ty)%2==0 {
		for cnt < 1000000 {
			cnt++
			x1,x2 := project(rects[cnt-1].x1,rects[cnt-1].x2,a,b)
			y1,y2 := project(rects[cnt-1].y1,rects[cnt-1].y2,c,d)
			rects = append(rects,rect{x1,x2,y1,y2})
			if Tx >= x1 && Tx <= x2 && Ty >= y1 && Ty <= y2 { ans = "Yes"; break }
		}
	}	
	fmt.Fprintln(wrtr,ans)
	if ans == "Yes" {
		pts := make([]pt,cnt)
		x,y := Tx,Ty; rx,ry := 0,0
		for i:=cnt-1;i>=0;i-- {
			rx,x = solveReflectionPoint(x,a,b,rects[i].x1,rects[i].x2)
			ry,y = solveReflectionPoint(y,c,d,rects[i].y1,rects[i].y2)
			pts[i] = pt{rx,ry}
		}
		for _,p := range pts { fmt.Fprintf(wrtr,"%v %v\n",p.x,p.y) }
	}
}
