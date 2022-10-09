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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); V := make([]string,N); for i:=0;i<N;i++ { V[i] = gs() }
		Q := gi(); W := make([]string,Q); for i:=0;i<Q;i++ { W[i] = gs() }
		L := len(V[0])
		mod1,mod2,mod3 := 998244353,1000000007,1000000009
		type hashval struct { h1,h2,h3 int }
		charval := func(c byte) int { 
			if c == 'm' { return 1 } else if c == 'e' { return 2 } else if c == 't' { return 3 } else { return 4 }
		}
		pv1,pv2,pv3 := ia(L),ia(L),ia(L)
		pv1[0] = 1; pv2[0] = 1; pv3[0] = 1
		for i:=1;i<L;i++ { 
			pv1[i] = pv1[i-1]*5 % mod1
			pv2[i] = pv2[i-1]*5 % mod2
			pv3[i] = pv3[i-1]*5 % mod3
		}
		calcHashvalFromString := func(s string) hashval {
			h1,h2,h3 := 0,0,0
			for i,cc := range s { 
				c := byte(cc); cv := charval(c)
				h1 += pv1[i] * cv % mod1
				h2 += pv2[i] * cv % mod2
				h3 += pv3[i] * cv % mod3
			}
			return hashval{ h1%mod1,h2%mod2,h3%mod3 }
		}
		VH := make([]hashval,N); for i:=0;i<N;i++ { VH[i] = calcHashvalFromString(V[i]) }
		WH := make([]hashval,Q); for i:=0;i<Q;i++ { WH[i] = calcHashvalFromString(W[i]) }
		dd := make(map[hashval]int)
		// Now add all of the strings from WH that have exactly one edit
		for i:=0;i<L;i++ {
			pp1,pp2,pp3 := pv1[i],pv2[i],pv3[i]
			for _,cc := range "meta" {
				c := byte(cc); cv := charval(c)
				for j:=0;j<Q;j++ {
					if W[j][i] == c { continue }
					hh1 := (WH[j].h1 + (mod1+cv-charval(W[j][i]))*pp1) % mod1
					hh2 := (WH[j].h2 + (mod2+cv-charval(W[j][i]))*pp2) % mod2
					hh3 := (WH[j].h3 + (mod3+cv-charval(W[j][i]))*pp3) % mod3
					dd[hashval{hh1,hh2,hh3}]++
				}
			}
		}
		ans := 0
		for i:=0;i<L;i++ {
			pp1,pp2,pp3 := pv1[i],pv2[i],pv3[i]
			for _,cc := range "meta" {
				c := byte(cc); cv := charval(c)
				for j:=0;j<Q;j++ {
					if W[j][i] == c { continue }
					hh1 := (WH[j].h1 + (mod1+cv-charval(W[j][i]))*pp1) % mod1
					hh2 := (WH[j].h2 + (mod2+cv-charval(W[j][i]))*pp2) % mod2
					hh3 := (WH[j].h3 + (mod3+cv-charval(W[j][i]))*pp3) % mod3
					dd[hashval{hh1,hh2,hh3}]--
				}
				for j:=0;j<N;j++ {
					if V[j][i] == c { continue }
					hh1 := (VH[j].h1 + (mod1+cv-charval(V[j][i]))*pp1) % mod1
					hh2 := (VH[j].h2 + (mod2+cv-charval(V[j][i]))*pp2) % mod2
					hh3 := (VH[j].h3 + (mod3+cv-charval(V[j][i]))*pp3) % mod3
					ans += dd[hashval{hh1,hh2,hh3}]
				}
			}
		}
		fmt.Printf("Case #%v: %v\n",tt,ans)
	}
}
