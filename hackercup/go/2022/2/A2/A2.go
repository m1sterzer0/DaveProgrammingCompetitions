package main

import (
	"bufio"
	"fmt"
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

const MOD = 1<<60
type FenwickMod struct { n, tot, mod int; bit []int }
func NewFenwickMod(n,mod int) *FenwickMod { buf := make([]int, n+1); return &FenwickMod{n, mod, 0, buf} }
func (q *FenwickMod) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *FenwickMod) Inc(idx int, val int) { 
	for idx <= q.n { q.bit[idx] += val % q.mod + q.mod; q.bit[idx] %= q.mod; idx += idx & (-idx) }
	q.tot += val % q.mod + q.mod; q.tot &= q.mod
}
func (q *FenwickMod) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *FenwickMod) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *FenwickMod) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { 
		ans += q.bit[idx]; ans %= q.mod; idx -= idx & (-idx)
	}
	return ans
}
func (q *FenwickMod) Suffixsum(idx int) int { 
	return (q.tot - q.Prefixsum(idx-1) + q.mod) % q.mod
}
func (q *FenwickMod) Rangesum(left int, right int) int {
	if right < left { return 0 }; return (q.Prefixsum(right) - q.Prefixsum(left-1) + q.mod) % q.mod
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	T := gi()
	rand.Seed(8675309)
	r1 := ia(1000001)
	r2 := ia(1000001)
	for i:=0;i<=1000000;i++ { 
		r1[i] = rand.Intn(MOD); 
		r2[i] = rand.Intn(MOD);
	}
	for tt:=1;tt<=T;tt++ {
		N := gi(); A := gis(N); Q := gi(); T,X,Y := fill3(Q)
		ft1 := NewFenwickMod(N+5,MOD)
		ft2 := NewFenwickMod(N+5,MOD)
		ft3 := NewFenwickMod(N+5,MOD)
		for i,a := range A { ft1.Inc(i+1,a); ft2.Inc(i+1,r1[a]); ft3.Inc(i+1,r2[a]) }
		ans := 0
		check := func(a1,b1,a2,b2 int) bool {
			s1 := ft1.Rangesum(a1,b1); s2 := ft1.Rangesum(a2,b2); d := s1-s2
			if d <= 0 || d > 1000000 { return false }
			if (ft2.Rangesum(a1,b1)-r1[d]+MOD) % MOD != ft2.Rangesum(a2,b2) { return false }
			if (ft3.Rangesum(a1,b1)-r2[d]+MOD) % MOD != ft3.Rangesum(a2,b2) { return false }
			return true
		}
		for i:=0;i<Q;i++ {
			t,a,b := T[i],X[i],Y[i]
			if t == 1 {
				//fmt.Printf("DBG: tt:%v i:%v N:%v Q:%v t:%v a:%v b:%v len(A):%v\n",tt,i,N,Q,t,a,b,len(A))
				ft1.Inc(a,b-A[a-1])
				ft2.Inc(a,r1[b]-r1[A[a-1]])
				ft3.Inc(a,r2[b]-r2[A[a-1]])
				A[a-1] = b
			} else {
				if a%2 != b%2 { continue }
				if a == b { ans++; continue }
				c := (a+b)>>1
				// First check a-c and c+1-r
				if check(a,c,c+1,b) { ans++; continue }
				if check(c,b,a,c-1) { ans++; continue }
			}
		}
		fmt.Printf("Case #%v: %v\n",tt,ans)
	}
}

