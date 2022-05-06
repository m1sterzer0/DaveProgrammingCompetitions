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
	N,M,B,W := gi(),gi(),gi(),gi()
	fact,factinv := makefact(3000,MOD)
	binom := func(n,r int) int { if n < 0 || r < 0 || r > n { return 0 }; return fact[n] * factinv[r] % MOD * factinv[n-r] % MOD }
	makedp := func(n,m,w int) [][]int {
		ans := make([][]int,n+1); for i:=0;i<=n;i++ { ans[i] = make([]int,m+1) }
		for i:=0;i<=n;i++ {
			for j:=0;j<=m;j++ {
				if i*j < w || i > w || j > w { continue }
				r := binom(i*j,w)
				for k:=0;k<=i;k++ {
					for l:=0;l<=j;l++ {
						if i == k && j == l { continue }
						r += MOD - ans[k][l] * binom(i,k) % MOD * binom(j,l) % MOD
					}
				}
				ans[i][j] = r % MOD
			}
		}
		return ans
	}
	wdp := makedp(N,M,W)
	bdp := makedp(N,M,B)
	ans := 0
	for i:=1;i<=N;i++ {
		for j:=1;j<=M;j++ {
			f1 := binom(N,i) * binom(M,j) % MOD * bdp[i][j] % MOD
			for k:=1;k<=N-i;k++ {
				for l:=1;l<=M-j;l++ {
					if k > W || l > W || k*l < W { continue }
					f2 := binom(N-i,k) * binom(M-j,l) % MOD * wdp[k][l] % MOD
					ans += f1 * f2 % MOD
				}
			}
		}
	}
	ans %= MOD
	fmt.Println(ans)
}

