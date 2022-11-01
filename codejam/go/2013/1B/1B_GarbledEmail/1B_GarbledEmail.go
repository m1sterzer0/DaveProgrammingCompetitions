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

func makeDD(D []string) map[string]bool {
	ans := make(map[string]bool)
	for _,s := range D {
		ans[s] = true
		b := []byte(s)
		ls := len(s)
		for i:=0;i<ls;i++ {
			b[i] = '*'; ans[string(b)] = true
			for j:=i+5;j<ls;j++ { b[j] = '*'; ans[string(b)] = true; b[j] = s[j] }
			b[i] = s[i]
		}
	}
	return ans
}

func solve(S string, DD map[string]bool) int {
	inf := 1<<60; ls := len(S); dp := twodi(ls+1,5,inf); dp[0][4] = 0
	processPure := func(s string,idx int) {
		//fmt.Printf("DBG: pure processing s:%v\n",s)
		l := len(s)
		if l >= 4 { dp[idx+l][4] = min(dp[idx+l][4],dp[idx][0]); return }
		for i:=0;i+l<=4;i++ { dp[idx+l][i+l] = min(dp[idx+l][i+l],dp[idx][i]) }
	}
	process := func(s string, first, last, idx int) {
		//fmt.Printf("DBG: processing s:%v\n",s)
		inc := 1; if first != last { inc = 2 }
		l := len(s); dp[idx+l][min(4,l-last-1)] = min(dp[idx+l][min(4,l-last-1)],dp[idx][max(0,4-first)]+inc)
	}
	for idx:=0;idx<ls;idx++ {
		for j:=3;j>=0;j-- { dp[idx][j] = min(dp[idx][j],dp[idx][j+1]) }
		j := min(ls,idx+10); s := S[idx:j]; b := []byte(s)
		for l:=1;l<=len(s);l++ {
			ss := string(b[0:l]); if DD[ss] { processPure(ss,idx) }
			for i:=0;i<l;i++ {
				b[i] = '*'; ss = string(b[0:l]); if DD[ss] { process(ss,i,i,idx) }
				for j:=i+5;j<l;j++ { b[j] = '*'; ss = string(b[0:l]); if DD[ss] { process(ss,i,j,idx) }; b[j] = s[j] }
				b[i] = s[i]
			}
		}
	}
	for j:=3;j>=0;j-- { dp[ls][j] = min(dp[ls][j],dp[ls][j+1]) }
	return dp[ls][0]
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	W := gi(); D := make([]string,0,W); for i:=0;i<W;i++ { D = append(D,gs()) }; DD := makeDD(D)
    T := gi()
    for tt:=1;tt<=T;tt++ {
		S := gs()
		ans := solve(S,DD)
        fmt.Printf("Case #%v: %v\n",tt,ans)
    }
}

