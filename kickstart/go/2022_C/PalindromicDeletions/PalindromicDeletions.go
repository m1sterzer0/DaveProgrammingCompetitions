package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"math/rand"
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
const inf int = 2000000000000000000
const MOD int = 1000000007

var fact [500]int
var factinv [500]int
func makefact(n int,mod int) {
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
}

func solveSmall(N int,S string) int {
	combinv := func(n,r int) int { return fact[r] * fact[n-r] % MOD * factinv[n] % MOD }
	ans := 0
	pals := make([]int,N); pals[0] = 1
	bs := make([]byte,N)
	isPalindrome := func(a []byte) bool {
		i,j := 0,len(a)-1
		for i<j { if a[i] != a[j] { return false }; i++; j-- }
		return true
	}
	for bm := 1; bm < (1<<uint(N)) - 1; bm++ {
		oc := bits.OnesCount(uint(bm))
		bs = bs[:0]
		for i:=0;i<N;i++ { if bm & (1<<uint(i)) != 0 { bs = append(bs,S[i]) } }
		if isPalindrome(bs) { pals[oc]++ }
	}
	for i:=0;i<N;i++ { ans += pals[i] * combinv(N,i) % MOD; ans %= MOD }
	return ans
}

func solveLarge(N int,S string) int {
	combinv := func(n,r int) int { return fact[r] * fact[n-r] % MOD * factinv[n] % MOD }
	ans := 1
	if N > 1 { ans += 1}
	if N > 2 {
		numeven := 0
		for i:=0;i<N;i++ { for j:=i+1;j<N;j++ { if S[i] == S[j] { numeven++ } } }
		ans += numeven * combinv(N,2) % MOD
		ans %= MOD
	}
	if N > 3 {
		odd := twodi(N,N,0)
		even := twodi(N,N,0)
		cum := twodi(N,N,0)
		for i:=0;i<N;i++ { odd[i][i] = 1 }
		for i:=0;i<N;i++ { for j:=i+1;j<N;j++ { if S[i] == S[j] { even[i][j] = 1 } } }
		for n:=3;n<N;n++ {
			// do the cum array
			xx := odd; if n % 2 == 0 { xx = even }
			for i:=0;i<N;i++ {
				for j:=0;j<N;j++ {
					if i == 0 && j == 0 { cum[i][j] = xx[i][j]; continue }
					if i == 0           { cum[i][j] = (cum[i][j-1] + xx[i][j]) % MOD; continue }
					if j == 0           { cum[i][j] = (cum[i-1][j] + xx[i][j]) % MOD; continue }
					cum[i][j] = (MOD + cum[i-1][j] + cum[i][j-1] + xx[i][j] - cum[i-1][j-1]) % MOD
				}
			}
			mys := 0
			for i:=0;i<N;i++ {
				for j:=0;j<N;j++ {
					if j < i+n-1 || S[i] != S[j] { xx[i][j] = 0; continue  }
					xx[i][j] = (cum[j-1][j-1] + cum[i][i] + MOD - cum[i][j-1] + MOD - cum[j-1][i]) % MOD
					mys += xx[i][j]
				}
			}
			mys %= MOD
			ans += mys * combinv(N,n) % MOD
			ans %= MOD
		}
	}
	return ans
}

func test(ntc,Nmin,Nmax,Cmax int) {
	npassed := 0
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		bs := make([]byte,N)
		for i:=0;i<N;i++ { bs[i] = 'a' + byte(rand.Intn(Cmax+1)) }
		S := string(bs)
		ans1 := solveSmall(N,S)
		ans2 := solveLarge(N,S)
		if ans1 == ans2 {
			npassed++
		} else {
			fmt.Printf("ERROR S:%v ans1:%v ans2:%v\n",S,ans1,ans2)
		}
	}
	fmt.Printf("%v/%v passed\n",npassed,ntc)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	makefact(499,MOD)
    T := gi()
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); S := gs()
		ans := solveLarge(N,S)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}
