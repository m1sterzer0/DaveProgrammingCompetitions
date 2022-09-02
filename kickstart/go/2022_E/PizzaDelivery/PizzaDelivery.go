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
type state struct { ok bool; vmax int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
		N,P,M,Ar,Ac := gi(),gi(),gi(),gi(),gi(); Ar--; Ac--
		buf := ""
		buf = gs(); on := buf[0]; kn := gi()
		buf = gs(); oe := buf[0]; ke := gi()
		buf = gs(); ow := buf[0]; kw := gi()
		buf = gs(); os := buf[0]; ks := gi()
		X,Y,C := fill3(P); for i:=0;i<P;i++ { X[i]--; Y[i]-- }
		dp := make([][][]state,N)
		for i:=0;i<N;i++ { dp[i] = make([][]state,N) }
		for i:=0;i<N;i++ { for j:=0;j<N;j++ { dp[i][j] = make([]state,1<<uint(P)) } }
		ndp := make([][][]state,N)
		for i:=0;i<N;i++ { ndp[i] = make([][]state,N) }
		for i:=0;i<N;i++ { for j:=0;j<N;j++ { ndp[i][j] = make([]state,1<<uint(P)) } }
		dp[Ar][Ac][0] = state{true,0}
		calcCoins := func(a int, op byte, b int) int {
			if op == '+' { return a+b }
			if op == '-' { return a-b }
			if op == '*' { return a*b }
			if a >= 0 || a%b == 0 { return a/b }
			return a/b - 1
		}
		updateState := func(s state, c int) state { if !s.ok { return state{true,c} }; return state{true,max(c,s.vmax) } }
		for mmm:=0;mmm<M;mmm++ {
			for k:=0;k<1<<uint(P);k++ {
				for i:=0;i<N;i++ {
					for j:=0;j<N;j++ {
						ndp[i][j][k] = dp[i][j][k]
					}
				}
			}
			// First we deal with movement -- then we deal with dropping off pizzas
			for k:=0;k<1<<uint(P);k++ {
				for i:=0;i<N;i++ {
					for j:=0;j<N;j++ {
						if !dp[i][j][k].ok { continue }
						if i != 0   { c := calcCoins(dp[i][j][k].vmax,on,kn); ndp[i-1][j][k] = updateState(ndp[i-1][j][k],c) }
						if i != N-1 { c := calcCoins(dp[i][j][k].vmax,os,ks); ndp[i+1][j][k] = updateState(ndp[i+1][j][k],c) }
						if j != 0   { c := calcCoins(dp[i][j][k].vmax,ow,kw); ndp[i][j-1][k] = updateState(ndp[i][j-1][k],c) }
						if j != N-1 { c := calcCoins(dp[i][j][k].vmax,oe,ke); ndp[i][j+1][k] = updateState(ndp[i][j+1][k],c) }
					}
				}
			}
			// Now we deal with pizza dropoffs
			for pp:=0;pp<P;pp++ {
				i,j := X[pp],Y[pp]
				for k:=0;k<1<<uint(P);k++ {
					if k & (1<<uint(pp)) != 0 { continue }
					if !ndp[i][j][k].ok { continue }
					ndp[i][j][k | (1<<uint(pp))] = updateState(ndp[i][j][k | (1<<uint(pp))], ndp[i][j][k].vmax+C[pp])
				}
			}
			ndp,dp = dp,ndp
			//for k:=0;k<1<<uint(P);k++ {
			//	for i:=0;i<N;i++ {
			//		for j:=0;j<N;j++ {
			//			if dp[i][j][k].ok == false { continue }
			//			fmt.Printf("DBG: tt:%v mmm:%v dp[%v][%v][%v] = %v\n",tt,mmm,i,j,k,dp[i][j][k])
			//		}
			//	}
			//}

		}
		inf := 1<<61
		ans := -inf
		kk := (1<<uint(P))-1
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				if !dp[i][j][kk].ok { continue }
				ans = max(ans,dp[i][j][kk].vmax)
			}
		}
		if ans == -inf {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
		}
    }
}

