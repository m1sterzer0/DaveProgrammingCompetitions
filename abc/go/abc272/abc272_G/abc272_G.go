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

func millerRabin(n int) bool {
	if n == 2 || n == 7 || n == 61 { return true }
	if n == 1 || n%2 == 0 { return false }
	if n >= 2147483647 { fmt.Println("ERROR: Don't support 128 bit mult yet"); return false }
	d := n-1; r := 0; for d & 1 == 0 { d >>= 1; r++ }
	w := []int{2,7,61}
	for _,a := range w {
		x := powmod(a,d,n)
		if x == 1 || x == n-1 { continue }
		posPrime := false
		for i:=0;i<r-1;i++ {
			x = (x*x) % n
			if x == n-1 { posPrime = true; break }
		}
		if !posPrime { return false }
	}
	return true
}

func getFactors(n int) []int {
	ansarr := make([]int,0)
	ansarr = append(ansarr,1)
	ansarr = append(ansarr,n)
	for i:=2;i*i<=n;i++ { 
		if n%i != 0 { continue }
		j := n/i
		ansarr = append(ansarr,i)
		if j != i { ansarr = append(ansarr,j) }
	}
	return ansarr
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N)
	// Main idea -- Assume there is a solution.  If you randomly pick two numbers, you have greater than a 1 in 4 chance
	// of finding two numbers in the majority
	// 10 trials -- chance of missing is < 0.06
	// 50 trials  -- chance of missing is < 5.7e-7
	// 100 trials -- chance of missing is <3.3e-13

	tryit := func(f int) int {
		sb := make(map[int]int)
		for _,a := range A { sb[a%f]++ }
		ans := 0
		for _,v := range sb { ans = max(ans,v) }
		return ans
	}
	ans := -1
	rand.Seed(8675309)
	for t:=0;t<100 && ans==-1;t++ {
		i1 := rand.Intn(N)
		i2 := i1; for i2==i1 { i2 = rand.Intn(N) }
		diff := abs(A[i1]-A[i2])
		if diff < 3 { continue }
		farr := getFactors(diff)
		for _,f := range(farr) {
			if f < 3 { continue }
			if f == 4 || millerRabin(f) {
				n := tryit(f)
				if n > N-n { ans = f; break }
			}
		}
	}
	fmt.Println(ans)
}

