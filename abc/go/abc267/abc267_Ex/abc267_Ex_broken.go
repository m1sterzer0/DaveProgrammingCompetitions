package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"math/rand"
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

var docstr = "(998244353,3) works"
func CONVOLVERpowmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
type CONVOLVER struct{ mod, primroot, rank2 int; root,iroot,rate2,irate2,rate3,irate3 []int }
func NewCONVOLVER(mod, primroot int) *CONVOLVER {
	rank2 := bits.TrailingZeros(uint(mod-1))
	if rank2 < 3 { panic("Hard wired to work for a significantly large power of 2 in the modulus") }
	root := make([]int,rank2+1); iroot := make([]int,rank2+1); rate2 := make([]int,rank2-2+1)
	irate2 := make([]int,rank2-2+1); rate3 := make([]int,rank2-3+1); irate3 := make([]int,rank2-3+1)
	root[rank2] = CONVOLVERpowmod(primroot,(mod-1)>>rank2,mod); iroot[rank2] = CONVOLVERpowmod(root[rank2],mod-2,mod)
	for i:=rank2-1;i>=0;i-- { root[i] = root[i+1]*root[i+1] % mod; iroot[i] = iroot[i+1]*iroot[i+1] % mod }
	prod,iprod := 1,1
	for i:=0;i<=rank2-2;i++ {
		rate2[i] = root[i+2] * prod % mod; irate2[i] = iroot[i+2] * iprod % mod; prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod,iprod = 1,1
	for i:=0;i<=rank2-3;i++ {
		rate3[i] = root[i+3] * prod % mod; irate3[i] = iroot[i+3] * iprod % mod; prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &CONVOLVER{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}
func (q *CONVOLVER) butterfly(a []int) {
	mod := q.mod; n := len(a); h := 0; for (1<<h) < n { h++ }; ll := 0
	for ll < h {
		if (h - ll == 1) {
			p := 1 << (h-ll-1); rot := 1
			for s:=0; s < (1 << ll); s++ {
				offset := s << (h - ll)
				for i:=0;i<p;i++ {
					l := a[i+offset]; r := a[i+offset+p] * rot % mod; u := l + r; if u >= mod { u -= mod }
					v := l - r; if v < 0 { v += mod }; a[i+offset] = u; a[i+offset+p] = v
				}
				if s + 1 != (1 << ll) { rot = rot * q.rate2[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll++
		} else {
			p := 1 << (h-ll-2); rot := 1; imag := q.root[2]
			for s:=0; s < (1 << ll); s++ {
				rot2 := rot * rot % mod; rot3 := rot2 * rot % mod; offset := s << (h - ll)
				for i:=0;i<p;i++ {
					mod2 := mod * mod; a0 := a[i+offset]; a1 := a[i+offset+p] * rot; a2 := a[i+offset+2*p] * rot2
					a3 := a[i+offset+3*p] * rot3; a1na3imag := (a1+mod2-a3) % mod * imag; na2 := mod2 - a2
					a[i+offset] = (a0 + a2 + a1 + a3) % mod; a[i+offset+p] = (a0 + a2 + (2 * mod2 - a1 - a3)) % mod
					a[i+offset+2*p] = (a0 + na2 + a1na3imag) % mod
					a[i+offset+3*p] = (a0 + na2 + (mod2-a1na3imag)) % mod
				}
				if s + 1 != (1 << ll) { rot = rot * q.rate3[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll += 2
		}
	}
}
func (q *CONVOLVER) butterflyinv(a []int) {
	mod := q.mod; n := len(a); h := 0; for (1<<h) < n { h++ }; ll := h
	for ll > 0 {
		if (ll == 1) {
			p := 1 << (h-ll); irot := 1
			for s:=0; s < (1 << (ll-1)); s++ {
				offset := s << (h - ll + 1)
				for i:=0;i<p;i++ {
					l := a[i+offset]; r := a[i+offset+p]; u := l + r; if u >= mod { u -= mod }
					v := (mod+l-r) * irot % mod; a[i+offset] = u; a[i+offset+p] = v
				}
				if s + 1 != (1 << (ll-1)) { irot = irot * q.irate2[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll--
		} else {
			p := 1 << (h-ll); irot := 1; iimag := q.iroot[2]
			for s:=0; s < (1 << (ll-2)); s++ {
				irot2 := irot * irot % mod; irot3 := irot2 * irot % mod; offset := s << (h - ll + 2)
				for i:=0;i<p;i++ {
					a0 := a[i+offset]; a1 := a[i+offset+p]; a2 := a[i+offset+2*p]; a3 := a[i+offset+3*p]
					a2na3iimag := (mod + a2 - a3) * iimag % mod; a[i+offset] = (a0 + a1 + a2 + a3) % mod
					a[i+offset+p] = (a0 + (mod-a1) + a2na3iimag) * irot % mod
					a[i+offset+2*p] = (a0 + a1 + (mod-a2) + (mod-a3)) * irot2 % mod
					a[i+offset+3*p] = (a0 + (mod-a1) + (mod - a2na3iimag)) * irot3 % mod
				}
				if s + 1 != (1 << (ll-2)) { irot = irot * q.irate3[bits.TrailingZeros(^uint(s))] % mod }
			}
			ll -= 2
		}
	}
	iz := CONVOLVERpowmod(n,mod-2,mod); for i:=0;i<n;i++ { a[i] = a[i] * iz % mod }
}
func (q *CONVOLVER) convolvefft(a []int, b []int) []int {
	mod := q.mod; finalsz := len(a) + len(b) - 1; z := 1; for z < finalsz { z *= 2 }; lena, lenb := len(a), len(b)
	la := make([]int, z); lb := make([]int, z); for i := 0; i < lena; i++ { la[i] = a[i] }
	for i := 0; i < lenb; i++ { lb[i] = b[i] }; q.butterfly(la); q.butterfly(lb)
	for i := 0; i < z; i++ { la[i] *= lb[i]; la[i] %= mod }; q.butterflyinv(la); return la[:finalsz]
}
func (q *CONVOLVER) convolvenaive(a []int, b []int) []int {
	mod := q.mod; finalsz := len(a) + len(b) - 1; ans := make([]int, finalsz)
	for i,a := range a { for j,b := range b { ans[i+j] += a * b; ans[i+j] %= mod } }; return ans
}
func (q *CONVOLVER) Convolve(a []int, b []int) []int {
	lmin := len(a); if len(b) < lmin { lmin = len(b) }
	if lmin <= 60 { return q.convolvenaive(a,b) } else { return q.convolvefft(a,b) }
}

const MOD = 998244353

func test(ntc int) {
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := 90000 + rand.Intn(100000-90000+1)
		M := rand.Intn(1000000-900000+1)
		A := make([]int,N)
		for i:=0;i<N;i++ { A[i] = 1 + rand.Intn(10-1+1) }
		st := time.Now()
		ans := solve(N,M,A)
		dur := time.Since(st)
		fmt.Printf("tt:%v ans:%v time:%v\n",tt,ans,dur)
	}
}

func solve(N,M int, A []int) int {
	cnts := make([]int,11); for _,a := range A { cnts[a]++ }
	e := []int{1}; o := []int{0}
	fact,factinv := makefact(N+1,MOD)
	comb := func(n,r int) int { 
		if n < 0 || r < 0 || r > n { return 0; }
		return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD
	}
	cc := NewCONVOLVER(998244353,3)
	polymult := func(p1,p2 []int,maxlen int) []int {
		preres := cc.Convolve(p1,p2)
		if len(preres) <= maxlen { return preres }
		return preres[:maxlen]
	}
	polyadd := func (p1,p2 []int) []int {
		l := max(len(p1),len(p2))
		res := make([]int,l)
		for i,a := range p1 { res[i] = a }
		for i,a := range p2 { res[i] = (res[i]+a) % MOD }
		return res
	}
	for i,c := range cnts {
		if c == 0 { continue }
		e2 := make([]int,1+i*c)
		o2 := make([]int,1+i*c)
		for j:=0;j<=c;j++ {
			if j & 1 == 1 { o2[i*j] = comb(c,j) } else { e2[i*j] += comb(c,j) }
		}
		ne := polyadd(polymult(e,e2,M+1),polymult(o,o2,M+1))
		no := polyadd(polymult(e,o2,M+1),polymult(o,e2,M+1))
		e,o = ne,no
	}
	ans := 0; if len(o) >= M+1 { ans = o[M] }
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	//test(10)	
	// PROGRAM STARTS HERE
	N,M := gi(),gi()
	A := gis(N)
	ans := solve(N,M,A)
	fmt.Println(ans)
}

