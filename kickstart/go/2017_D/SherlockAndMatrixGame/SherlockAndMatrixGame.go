package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
const MOD int = 1000000007

func getSortedSubarraySums(a []int) []int {
	n := len(a)
	cumsum := ia(n); cumsum[0] = a[0]
	for i:=1;i<n;i++ { cumsum[i] = cumsum[i-1] + a[i] }
	res := make([]int,0,n*(n+1)/2)
	for i:=0;i<n;i++ {
		res = append(res,cumsum[i])
		for j:=0;j<i;j++ { res = append(res,cumsum[i]-cumsum[j]) }
	}
	sort.Slice(res,func(i,j int) bool { return res[i] < res[j] })
	return res
}

func countGe(b []int, v int) int {
	if b[0] >= v { return len(b) }
	if b[len(b)-1] < v { return 0 }
	l,r := 0,len(b)-1
	for r-l > 1 {
		m := (r+l)>>1
		if b[m] < v { l = m } else {r = m }
	}
	return len(b)-(l+1)
}

func countLe(b []int, v int) int {
	if b[0] > v { return 0 }
	if b[len(b)-1] <= v { return len(b) }
	l,r := 0,len(b)
	for r-l > 1 {
		m := (r+l)>>1
		if b[m] <= v { l = m } else {r = m }
	}
	return l+1
}

func smallCheck(m int, a,b []int, k int) bool {
	cnt := 0
	for _,x := range a {
		if m > 0 {
			if x > 0 {
				cutoff := (m+x-1)/x
				cnt += countGe(b,cutoff)
			} else if x < 0 {
				cutoff := (m + (-x)-1)/(-x)
				cnt += countLe(b,-cutoff)
			}
		} else if m == 0 {
			if x > 0 {
				cnt += countGe(b,0) 
			} else if x == 0 {
				cnt += len(b)
			} else {
				cnt += countLe(b,0)
			}
		} else {
			if x > 0 {
				cutoff := m / x
				cnt += countGe(b,cutoff)
			} else if x == 0 {
				cnt += len(b)
			} else {
				cnt += countLe(b,(-m)/(-x))
			}
		}
	}
	return cnt >= k
}

func solveSmall(N,K int, A,B []int) int {
	AS := getSortedSubarraySums(A)
	BS := getSortedSubarraySums(B)
	// Biggest we can get is 1000 * 1000 * 200 * 200 = 40000000000 (overkill, since we can only get to 999)
	l,r := -40000000000,40000000000
	for r-l > 1 {
		m := (r+l)>>1
		if smallCheck(m,AS,BS,K) { l = m } else { r = m }
	}
	return l
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K,A1,B1,C,D,E1,E2,F := gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi()
		X,Y,R,S,A,B := ia(N),ia(N),ia(N),ia(N),ia(N),ia(N)
		X[0] = A1; Y[0] = B1; R[0] = 0; S[0] = 0; A[0] = A1; B[0] = B1
		for i:=1;i<N;i++ { 
			X[i] = (C*X[i-1]+D*Y[i-1]+E1) % F
			Y[i] = (D*X[i-1]+C*Y[i-1]+E2) % F
			R[i] = (C*R[i-1]+D*S[i-1]+E1) % 2
			S[i] = (D*R[i-1]+C*S[i-1]+E2) % 2
			A[i] = X[i]; if R[i] == 1 { A[i] *= -1 }
			B[i] = Y[i]; if S[i] == 1 { B[i] *= -1 }
		}

		// Very dumb solution: construct the matrix , construct a list of all the subset sums, sort, return answer
		// Small: Matrix is 200 x 200 = 40,000 entries.  Number of subarray sums is around 0.5 * (40,000 choose 2) -- way too big even with good prefix sums.
		//   Instead, can build list of subarray sums.  A and B each have C(200,2) = 19900 subarray sums (very tractable).  Can sort those, and then use a binary search strategy to find the
		//   Kth largest 2-d subarray
		if N <= 200 {
			ans := solveSmall(N,K,A,B)
        	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		} else {
        	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,0)
		}
    }
}

