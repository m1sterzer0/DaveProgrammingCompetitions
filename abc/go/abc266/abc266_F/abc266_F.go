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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N := gi(); U,V := fill2(N); for i:=0;i<N;i++ { U[i]--; V[i]-- }
	Q := gi(); X,Y := fill2(Q); for i:=0;i<Q;i++ { X[i]--; Y[i]-- }
	gr := make([][]int, N)
	for i:=0;i<N;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }

	// Do first DFS to find a repeated node
	var dfs1 func(n,p int) int
	sb := make([]bool,N)
	dfs1 = func(n,p int) int {
		if sb[n] { return n }
		sb[n] = true
		for _,c := range gr[n] {
			if c == p { continue }
			r := dfs1(c,n)
			if r != -1 { return r }
		}
		return -1
	}
	repeatedNode := dfs1(0,-1)

	// Do second DFS from repeated node to construct cycle
	cycle := make([]int,0,N)
	parr := iai(N,-1)
	var dfs2 func(n,p int) bool
	dfs2 = func(n,p int) bool {
		if n == repeatedNode && p != -1 {
			x := p; for x >= 0 { cycle = append(cycle,x); x = parr[x] }; return true
		} else {
			parr[n] = p
			for _,c := range gr[n] { if c != p { if dfs2(c,n) { return true } } }
		}
		return false
	}
	dfs2(repeatedNode,-1)

	// Build the forest without the cycle nodes
	cyclesb := make([]bool,N); for _,c := range cycle { cyclesb[c] = true }
	gr2 := make([][]int, N)
	for i:=0;i<N;i++ { u,v := U[i],V[i]; if cyclesb[u] && cyclesb[v] { continue }; gr2[u] = append(gr2[u],v); gr2[v] = append(gr2[v],u) }

	// Color the nodes by traversing the forest.  Use cycle nodes as heads
	colors := iai(N,-1)
	var dfs3 func(n,p,cc int)
	dfs3 = func(n,p,cc int) {
		colors[n] = cc
		for _,c := range gr2[n] { if c != p { dfs3(c,n,cc) } }
	}
	for _,c := range cycle { dfs3(c,-1,c) }

	// Answer the queries
	for i:=0;i<Q;i++ { ans := "No"; if colors[X[i]] == colors[Y[i]] { ans = "Yes" }; fmt.Fprintln(wrtr,ans) }
}
