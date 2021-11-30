package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
    func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Solution transcription
	N := gi(); A,B := fill2(N-1); for i:=0;i<N-1;i++ { A[i]--; B[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }
	fact,factinv := makefact(N+1,MOD)
	sub := ia(N+1)
	dp := twodi(N,N+1,0)
	var dfs func(c,p int)
	dfs = func(n,p int) {
		sub[n] = 0
		cur,nxt := []int{1},[]int{}
		for _,a := range gr[n] {
			if a == p { continue }
			dfs(a,n)
			rui := ia(sub[a]+1)
			for i:=0;i<sub[a];i++ { rui[i+1] = dp[a][sub[a]-i] }
			for i:=0;i<sub[a];i++ { rui[i+1] += rui[i]; rui[i+1] %= MOD }
			nxt = nxt[:0]
			for i:=0;i<sub[n]+sub[a]+1;i++ { nxt = append(nxt,0) }
			for i:=0;i<=sub[n];i++ {
				if cur[i] == 0 { continue }
				for j:=0;j<=sub[a];j++ {
					if rui[j] == 0 { continue }
					nxt[i+j] += cur[i] * rui[j] % MOD * factinv[j] % MOD * factinv[sub[a]-j] % MOD
					nxt[i+j] %= MOD
				}
			}
			sub[n] += sub[a]
			cur,nxt = nxt,cur
		}
		sub[n]++
		for i:=0;i<sub[n];i++ {
			dp[n][i+1] = cur[i] * fact[i] % MOD * fact[sub[n]-i-1] % MOD
		}
	}
	dfs(0,-1)
	ans := 0
	for i:=0;i<N;i++ { ans += dp[0][i+1] }
	ans %= MOD; ans *= 2; ans %= MOD
	fmt.Println(ans)
}

