package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
const MOD int = 1000000007
var fact [1000010]int
var factinv [1000010]int

// Still don't know why this works, but multiple people submitted something similar to this, so 
// I need to dig in to understand the inclusion/exclusion here.
func solve(M,K int) int {
	twoinv := powmod(2,MOD-2,MOD)
	temp := make([]int,M/2+5)
	temp[0] = 1
	sum := 0
	for i:=1;i<=M/2;i++ {
		sum += 2*M - 4*i + 1; sum %= MOD
		temp[i] = temp[i-1] * (M-2*i+2) % MOD * (M-2*i+1) % MOD * twoinv % MOD * powmod(sum,MOD-2,MOD) % MOD
	}
	ans := 0
	for i:=K;i<=M/2;i++ {
		sgn := -1; if i&1==K&1 { sgn = 1 }
		ans += sgn * temp[i] * fact[i] % MOD * factinv[K] % MOD * factinv[i-K] % MOD
		ans += MOD; ans %= MOD
	}
	return ans
}
func solveSmall(M,K int) int {
	dp := make([]int,M/2+5)
	ndp := make([]int,M/2+5)
	dp[1] = 1
	for ne:=2;ne<=M-1;ne++ {
		for i:=0;i<M/2+5;i++ { ndp[i] = 0 }
		for ni := 1; ni<M/2+5; ni++ {
			if dp[ni] == 0 { continue }
			nodesInIslands := 2*ni+(ne-1-ni)
			nodesLeftOver  := M - nodesInIslands
			if nodesLeftOver < 2 { ndp[ni] += dp[ni]; ndp[ni] %= MOD; continue }
			newIslandEdges := nodesLeftOver * (nodesLeftOver-1) / 2
			oldIslandEdges := nodesInIslands * (nodesInIslands-1) / 2
			joinIslandEdges := M * (M-1) / 2 - newIslandEdges - oldIslandEdges
			denominv := powmod(joinIslandEdges + newIslandEdges,MOD-2,MOD)
			ndp[ni+1] += dp[ni] * newIslandEdges  % MOD * denominv % MOD; ndp[ni+1] %= MOD
			ndp[ni]   += dp[ni] * joinIslandEdges % MOD * denominv % MOD; ndp[ni] %= MOD
		}
		dp,ndp = ndp,dp
	}
	return dp[K]
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	nmax := 1000000
	fact[0] = 1; for i:=1;i<=nmax;i++ { fact[i] = fact[i-1] * i % MOD }
	factinv[nmax] = powmod(fact[nmax],MOD-2,MOD); for i:=nmax-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % MOD }
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		M,K := gi(),gi()
		//ans := solveSmall(M,K)
		ans := solve(M,K)
		fmt.Printf("Case #%v: %v\n",tt,ans)
    }
}

