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
func sortUniqueIntarr(a []int) []int {
	sort.Slice(a,func (i,j int) bool { return a[i] < a[j] })
	i,j,la := 0,0,len(a)
	for ;i<la;i++ { if i == 0 || a[i] != a[i-1] { a[j] = a[i]; j++ } }
	return a[:j]
}
type pair struct { i,j int }
const MOD int = 998244353
var comb [1000][1000]int
var dp [410][410]int

func test(ntc,Nmin,Nmax,Cmin,Cmax int) {
	rand.Seed(8675309)
	for tt:=1;tt<=ntc;tt++ {
		N := Nmin + rand.Intn(Nmax-Nmin+1)
		A := twodi(N,N,0)
		for i:=0;i<N;i++ {
			for j:=0;j<N;j++ {
				A[i][j] = Cmin + rand.Intn(Cmax-Cmin+1)
			}
		}
		solve(N,A)
	}
}

func solve(N int, A [][]int) int {
	for i:=0;i<=2*N;i++ { 
		comb[i][0] = 1; comb[i][i] = 1
		for j:=1;j<i;j++ { comb[i][j] = (comb[i-1][j-1]+comb[i-1][j]) % MOD }
	}
	// First do some coordinate compression
	c := make([]int,0,N*N); for i:=0;i<N;i++ { for j:=0;j<N;j++ { c = append(c,A[i][j]) } }
	c = sortUniqueIntarr(c)
	lu := make(map[int]int); for i,n := range c { lu[n] = i }
	A2 := twodi(N,N,0); for i:=0;i<N;i++ { for j:=0;j<N;j++ { A2[i][j] = lu[A[i][j]] } }
	gr := make([][]pair,len(c))
	for i:=0;i<N;i++ { for j:=0;j<N;j++ { a := A2[i][j]; gr[a] = append(gr[a],pair{i,j}) } }
	ans := N*N // precount the paths that start and end on the same square
	for v:=0;v<len(c);v++ {
		n := len(gr[v])
		if n * (n-1) / 2 < N*N {
			for i:=0;i<n;i++ {
				i1,j1 := gr[v][i].i,gr[v][i].j
				for j:=i+1;j<n;j++ {
					i2,j2 := gr[v][j].i,gr[v][j].j
					if i1 <= i2 && j1 <= j2 { ans += comb[i2-i1+j2-j1][i2-i1] }
					if i2 <= i1 && j2 <= j1 { ans += comb[i1-i2+j1-j2][i1-i2] }
				}
			}
			ans %= MOD
		} else {
			for i:=0;i<N;i++ {
				for j:=0;j<N;j++ {
					dp[i][j] = 0
					if i > 0 { dp[i][j] += dp[i-1][j]; if A2[i-1][j] == v { dp[i][j]++ } }
					if j > 0 { dp[i][j] += dp[i][j-1]; if A2[i][j-1] == v { dp[i][j]++ } }
					for dp[i][j] >= MOD { dp[i][j] -= MOD }
					if A2[i][j] == v { ans += dp[i][j] }					
				}
			}
			ans %= MOD
		}
	}
	return ans
}


func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	//test(100,1,400,1,30)
	N := gi(); A := make([][]int,N); for i:=0;i<N;i++ { A[i] = gis(N) }
	ans := solve(N,A)
	fmt.Println(ans)
}
