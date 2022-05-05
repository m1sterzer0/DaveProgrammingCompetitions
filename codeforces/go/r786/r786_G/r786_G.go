package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func I(a int) int64 { return int64(a) }
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int64 { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func gi2() (int64,int64) { return gi(),gi() }
func gi3() (int64,int64,int64) { return gi(),gi(),gi() }
func gi4() (int64,int64,int64,int64) { return gi(),gi(),gi(),gi() }
func gis(n int64) []int64 { res := make([]int64,n); for i:=I(0);i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int64) []float64  { res := make([]float64,n); for i:=I(0);i<n;i++ { res[i] = gf() }; return res }
func gss(n int64) []string  { res := make([]string,n); for i:=I(0);i<n;i++ { res[i] = gs() }; return res }
func ia(m int64) []int64 { return make([]int64,m) }
func iai(m int64,v int64) []int64 { a := make([]int64,m); for i:=I(0);i<m;i++ { a[i] = v }; return a }
func twodi(n int64,m int64,v int64) [][]int64 {
	r := make([][]int64,n); for i:=I(0);i<n;i++ { x := make([]int64,m); for j:=I(0);j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int64) ([]int64,[]int64) { a,b := ia(m),ia(m); for i:=I(0);i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int64) ([]int64,[]int64,[]int64) { a,b,c := ia(m),ia(m),ia(m); for i:=I(0);i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int64) ([]int64,[]int64,[]int64,[]int64) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=I(0);i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int64) int64 { if a < 0 { return -a }; return a }
func rev(a []int64) { i,j := I(0),I(len(a)-1); for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int64) int64 { if a > b { return a }; return b }
func min(a,b int64) int64 { if a > b { return b }; return a }
func tern(cond bool, a int64, b int64) int64 { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int64) int64 { ans := I(0); for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int64) { for i:=I(0); i<I(len(a)); i++ { a[i] = 0 } }
func powmod(a,e,mod int64) int64 { res, m := I(1), a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int64) int64 { res, m := I(1), a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int64) int64 { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int64) (int64,int64,int64) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int64) (int64,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func makefact(n int64,mod int64) ([]int64,[]int64) {
	fact,factinv := make([]int64,n+1),make([]int64,n+1)
	fact[0] = 1; for i:=I(1);i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int64 = 1000000000000000000 
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); U,V := fill2(M); for i:=I(0);i<M;i++ { U[i]--; V[i]-- }
	gr:=make([][]int64,N); grrev := make([][]int64,N)
	for i:=I(0);i<M;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); grrev[v] = append(grrev[v],u) }
	numdeps := make([]int64,N); for i:=I(0);i<N;i++ { numdeps[i] = I(len(gr[i])) }
	q := make([]int64,0); for i:=I(0);i<N;i++ { if numdeps[i] == 0 { q = append(q,i) } }
	ans := I(1)
	maxsuffix := make([]int64,N)
	for len(q) != 0 {
		n := q[0]; q = q[1:]
		startok := len(gr[n]) >= 2
		endok   := len(grrev[n]) >= 2
		cand := I(1)
		if startok {
			for _,c := range gr[n] { cand = max(cand,1+maxsuffix[c]) }
			ans = max(cand,ans)
		} 
		if endok { maxsuffix[n] = 1 }
		if startok && endok { maxsuffix[n] = cand }
		for _,p := range grrev[n] { numdeps[p]--; if numdeps[p] == 0 { q = append(q,p) } }
	}
	fmt.Println(ans)
}
