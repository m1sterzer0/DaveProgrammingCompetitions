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

func next_permutation(a []int) bool {
	la := len(a); var i,j int
	for i=la-2;i>=0;i-- { if a[i] < a[i+1] { break } }
	if i<0 { i,j = 0,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- } ; return false }
	for j=la-1;j>=0;j-- { if a[i] < a[j] { break } }
	a[i],a[j] = a[j],a[i]
	i,j = i+1,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- }
	return true
}

const inf = 1<<61

func solveBrute(N int, V1,V2 []int) int {
	v2 := make([]int,N); copy(v2,V2)
	sort.Slice(v2, func(i,j int) bool { return v2[i] < v2[j] } )
	best := inf
	for {
		cand := 0
		for i:=0;i<N;i++ { cand += V1[i] * v2[i] }
		best = min(best,cand)
		if !next_permutation(v2) { break }
	}
	return best
}

func solve(N int, V1,V2 []int) int {
	v1 := make([]int,N); copy(v1,V1)
	v2 := make([]int,N); copy(v2,V2)
	sort.Slice(v1,func(i,j int) bool { return v1[i] < v1[j] } )
	sort.Slice(v2,func(i,j int) bool { return v2[i] > v2[j] } )
	ans := 0
	for i:=0;i<N;i++ { ans += v1[i] * v2[i] }
	return ans
}

func test(ntc,nmin,nmax,vmin,vmax int) {
	npassed := 0
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := nmin + rand.Intn(nmax-nmin+1)
		V1 := make([]int,0,N)
		V2 := make([]int,0,N)
		for i:=0;i<N;i++ { V1 = append(V1,vmin+rand.Intn(vmax-vmin+1) ) }
		for i:=0;i<N;i++ { V2 = append(V2,vmin+rand.Intn(vmax-vmin+1) ) }
		ans1 := solveBrute(N,V1,V2)
		ans2 := solve(N,V1,V2)
		if ans1 == ans2 { 
			npassed++
		} else {
			fmt.Printf("ERROR tt:%v N:%v V1:%v V2:%v ans1:%v ans2:%v\n",tt,N,V1,V2,ans1,ans2)
		}
	}
	fmt.Printf("%v/%v passed\n",npassed,ntc)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	T := gi()
    for tt:=1;tt<=T;tt++ {
		N := gi(); V1 := gis(N); V2 := gis(N)
		ans := solve(N,V1,V2)
    	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

