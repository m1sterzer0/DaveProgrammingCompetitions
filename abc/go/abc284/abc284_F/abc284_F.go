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

type shash struct { pv,x,xinv,p,val int }
func Newshash(x,p int) *shash { xinv := powmod(x,p-2,p); return &shash{1,x,xinv,p,0} }
func (q *shash) pushleft(c int) { q.val += q.pv*c % q.p; q.val %= q.p; q.pv *= q.x; q.pv %= q.p }
func (q *shash) popleft(c int)  { q.pv *= q.xinv; q.pv %= q.p;  q.val += q.p - (q.pv*c % q.p); q.val %= q.p }
func (q *shash) pushright(c int) { q.pv *= q.x; q.pv %= q.p; q.val *= q.x; q.val %= q.p; q.val += c; q.val %= q.p }
func (q *shash) popright(c int)  { q.val += q.p-c; q.val %= q.p; q.val *= q.xinv; q.val %= q.p; q.pv *= q.xinv; q.pv %= q.p }


func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N := gi()
	T := gs()
	alph := "abcdefghijklmnopqrstuvwxyz"
	dval := make(map[byte]int)
	mm := 1000000007
	for i,c := range alph { dval[byte(c)] = i+1 }
	h1,h2,h3 := Newshash(31,mm),Newshash(31,mm),Newshash(31,mm)
	ans := -1
	ansstr := ""

	check := func(i int) bool {
		s1 := T[0:i] + T[i+N:2*N]
		bs2 := make([]byte,N)
		for j:=0;j<N;j++ { bs2[j] = byte(T[i+N-1-j]) }
		s2 := string(bs2)
		return s1 == s2
	}

	for i:=0;i<=N;i++ {
		if i == 0 {
			for k:=N-1;k>=0;k-- { h1.pushright(dval[byte(T[k])]) }
			for k:=N;k<2*N;k++ { h3.pushright(dval[byte(T[k])]) }
		} else {
			h1.popright(dval[byte(T[i-1])])
			h2.pushright(dval[byte(T[i-1])])
			h3.popleft(dval[byte(T[i+N-1])])
			h1.pushleft(dval[byte(T[i+N-1])])
		}
		cand := (h2.val * powmod(31,N-i,mm) % mm + h3.val) % mm
		if cand == h1.val && check(i) { ansstr = T[0:i] + T[i+N:2*N]; ans = i; break }
	}
	if ans == -1 {
		fmt.Println(ans)
	} else {
		fmt.Println(ansstr); fmt.Println(ans)
	}
}


