package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
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
type ival struct { v float64; idx int; st bool }

type Fenwick struct { n, tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx) }; q.tot += val }
func (q *Fenwick) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *Fenwick) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans
}
func (q *Fenwick) Suffixsum(idx int) int { return q.tot - q.Prefixsum(idx-1) }
func (q *Fenwick) Rangesum(left int, right int) int {
	if right < left { return 0 }; return q.Prefixsum(right) - q.Prefixsum(left-1)
}

type frac struct { n,d int }
func test(ntc,Cmax,Nmin,Nmax,Kmax int) {
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		ff := make(map[frac]bool)
		A := make([]int,0,N); B := make([]int,0,N)
		for len(A) < N {
			a := -Cmax + rand.Intn(2*Cmax+1)
			b := -Cmax + rand.Intn(2*Cmax+1)
			if a == 0 && b == 0 { continue }
			e,f := a,b
			if f < 0 { e *= -1; f *= -1 }
			if e == 0 { 
				e,f = 0,1
			} else if f == 0 {
				e,f = 1,0
			} else {
				g := gcd(abs(e),abs(f))
				e /= g; f /= g
			}
			_,ok := ff[frac{e,f}]
			if ok { continue }
			A = append(A,a); B = append(B,b)
		}
		C := make([]int,0,N)
		for i:=0;i<N;i++ { C = append(C,-Cmax+rand.Intn(2*Cmax+1)) }
		K := 1 + rand.Intn(Kmax)
		K = min(N*(N-1)/2,K)
		solve(N,K,A,B,C)
	}
}

func solve(N,K int, A,B,C []int) float64 {
	X,Y,D := make([]float64,N),make([]float64,N),make([]float64,N)
	XV,YV := make([]float64,N),make([]float64,N)
	for i:=0;i<N;i++ {
		a,b,c := float64(A[i]),float64(B[i]),float64(C[i])
		denom := math.Sqrt(a*a+b*b); denom2 := a*a+b*b
		D[i] = c / denom; if D[i] < 0 { D[i] = -D[i] }
		X[i] = -a*c/denom2
		Y[i] = -b*c/denom2
		XV[i] = -b/denom;
		YV[i] = a/denom
	}
	ee := make([]ival,2*N)
	lkup := make([]int,N)
	ft := NewFenwick(2*N+10)
	check := func(r float64) bool {
		ee := ee[:0]
		ft.Clear()
		for i:=0;i<N;i++ {
			d := D[i]
			if d > r { continue }
			rd := math.Sqrt(r*r-d*d)
			x1,y1 := X[i]+XV[i]*rd,Y[i]+YV[i]*rd
			x2,y2 := X[i]-XV[i]*rd,Y[i]-YV[i]*rd
			ang1 := math.Atan2(y1,x1)
			ang2 := math.Atan2(y2,x2)
			if ang1 > ang2 { ang1,ang2 = ang2,ang1 }
			ee = append(ee,ival{ang1,i,true})
			ee = append(ee,ival{ang2,i,false})
		}
		sort.Slice(ee,func(i,j int) bool { return ee[i].v < ee[j].v } )
		ans := 0; open := 0
		for i,e := range ee {
			if e.st {
				lkup[e.idx] = i; open++; ft.Inc(i+1,-1)
			} else {
				ii := lkup[e.idx]
				ans += open + ft.Prefixsum(ii+1)
				open--; ft.Inc(ii+1,1)
			}
		}
		return ans >= K
	}
	// Now we do our binary search
	l,u := 0.0,3000000.0
	for iter:=1;iter<=80;iter++ {
		m := 0.5*(l+u)
		if check(m) { u = m } else { l = m }
	}
	ans := 0.5*(l+u)
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	//test(1000,100,2,100,50)
	N,K := gi(),gi()
	A,B,C := fill3(N)
	ans := solve(N,K,A,B,C)
	fmt.Println(ans)
}
