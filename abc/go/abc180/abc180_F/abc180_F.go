package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a > b { return b }; return a }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
type PI struct { x,y int }
const MOD = 1_000_000_007

func solve(N,M,L int,fact,factinv,modinv []int, cache map[PI]int) int {
	v,ok := cache[PI{N,M}]; if ok { return v }
	ans := 0
	if N == 0 && M == 0 { 
		ans = 1
	} else if M > N {
		ans = 0
	} else {
		// Singleton
		ans = (ans + solve(N-1,M,L,fact,factinv,modinv,cache)) % MOD 
		for sz:=2;sz<=min(L,N);sz++ {
			ways := fact[N-1] * factinv[sz-1] % MOD * factinv[N-sz] % MOD
			// Chain
			if M >= sz-1 {
				c1 := solve(N-sz,M-(sz-1),L,fact,factinv,modinv,cache)
				if c1 != 0 {
					waysc := ways * fact[sz] % MOD * modinv[2] % MOD
					ans = (ans + waysc * c1 % MOD ) % MOD
				}
			}
			// Two-ring
			if sz == 2 && M >= sz {
				c2a := solve(N-sz,M-sz,L,fact,factinv,modinv,cache)
				if c2a != 0 { ans = (ans + ways * c2a % MOD ) % MOD }
			}

			// Ring
			if sz >= 3 && M >= sz {
				c2 := solve(N-sz,M-sz,L,fact,factinv,modinv,cache)
				if c2 != 0 {
					waysr := ways * fact[sz-1] % MOD * modinv[2] % MOD
					ans = (ans + waysr * c2 % MOD ) % MOD
				}
			}
		}
	}
	cache[PI{N,M}] = ans
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,L := gi3()
	fact,factinv := makefact(N+10,MOD)
	modinv := iai(N+10,1)
	for i:=2; i<N+10; i++ { modinv[i] = powmod(i,MOD-2,MOD) }
	cache := make(map[PI]int)
	ans1 := solve(N,M,L,fact,factinv,modinv,cache)
	cache = make(map[PI]int)
	ans2 := 0
	if L > 0 { ans2 = solve(N,M,L-1,fact,factinv,modinv,cache) }
	ans := (ans1 + MOD - ans2) % MOD
	fmt.Println(ans)
}
	
