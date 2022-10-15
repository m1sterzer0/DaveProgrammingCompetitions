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


func dolcm(a,b int) int {
	c := gcd(a,b); m := a/c
	if b > 10000000000000000 / m { return 10000000000000001 } else { return m*b }
}
func solveRightEnd(l,L,H int) int {
	minmult,maxmult := (L+l-1)/l,H/l
	if minmult <= maxmult { return minmult * l } else { return -1 }
}
func solveit(f1,f2,l,g,L,H int) int {
	if f2 <= L || f1 > H { return -1 }
	ll,uu := max(L,f1),min(f2-1,H)
	if ll > g || uu < l || g%l != 0 { return -1 }
	m := g/l // need to check factors of m
	best := 1<<60
	for i:=1;i*i<=m;i++ {
		if m%i != 0 { continue }
		cc1,cc2 := i*l,(m/i)*l
		if cc1 < best && cc1 >= ll && cc1 <= uu { best = cc1 }
		if cc2 < best && cc2 >= ll && cc2 <= uu { best = cc2 }
	}
	if best == 1<<60 { return -1} else { return best }
}

func solve(N,L,H int, A []int) int {
	// Add 1 to the input array to take care of one special case
	A2 := make([]int,N); copy(A2,A); A2 = append(A2,1); sort.Slice(A2,func(i,j int) bool { return A2[i]<A2[j] } )
	lcmarr := ia(N+1); runlcm := 1; for i:=0;i<N+1;i++   { runlcm = dolcm(runlcm,A2[i]); lcmarr[i] = runlcm }
	gcdarr := ia(N+1); rungcd := A2[N]; for i:=N;i>=0;i-- { rungcd = gcd(rungcd,A2[i]); gcdarr[i] = rungcd }
	for i:=0;i<N;i++ {
		l,g := lcmarr[i],gcdarr[i+1]
		cand := solveit(A2[i],A2[i+1],l,g,L,H)
		if cand >= A2[i] && cand < A2[i+1] { return cand }
	}
	cand := solveRightEnd(lcmarr[N],L,H)
	return cand // will be -1 if we fail
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N,L,H := gi(),gi(),gi(); A := gis(N)
		ans := solve(N,L,H,A)
		if ans == -1 {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"NO")
		} else {
        	fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
    }
}

