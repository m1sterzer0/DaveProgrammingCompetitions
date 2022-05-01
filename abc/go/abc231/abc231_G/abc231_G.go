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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); A := gis(N)
	// Two balls : E[prod(Xi+Ai)] = E[X0X1+ A0X1 + A1X0 + A0A1] = E[X0X1] + A0 * E[X1] + A1 * E[X0] + A0A1
	//                            = E[X0X1] + (A0+A1)E[X0] + A0A1
	// Three balls : E[prod(Xi+Ai)] = E[X0X1X2+ A0X1X2 + X0A1X2 + X0X1A2 + A0A1X2 + A0X1A2 + A0A1X2 + X0X1X2]
	//                              = E[X0X1X2] + (A0+A1+A2)E[X0X1] + (A0A1+A0A2+A12)E[X1] + (A0A1A2)
	// Thus, we need two things:
	// A) Calculate the coefficients { 1, (sum Ai), (sum AiAj), ..., A0A1...An-1 }
	poly := []int{1}
	for _,a := range A {
		newpoly := make([]int,len(poly)+1)
		for i,p := range poly { newpoly[i+1] = p }
		for i,p := range poly { newpoly[i] += a*p; newpoly[i] %= MOD }
		poly = newpoly
	}
	// B) Calculate the coefficients { E[X0X1..Xn-1], E[X1X2...Xn-1], ..., E[Xn-1], 1 }
	// Decompose E[Xi] = E[sum(Yij)]  where Yij is 1 if the jth ball went into the ith slot and zero otherwise
	// E[XiXj] = E[sum(Yik)sum(Yjl)] = E[sum(YikYjl)].  Assuming i!=j, YikYjl is zero if k==l, and otherwise it is (1/N)^2
	//           Thus E[XiXj] = (1/N^2) * (# pairs (k,l) such that k!=l) = (1/N^2)*(K choices for k) * (K-1 choices for l)
	poly2 := make([]int,N+1); poly2[0] = 1; ninv := powmod(N,MOD-2,MOD)
	for i:=1;i<=N;i++ { poly2[i] = poly2[i-1] * ninv % MOD * (K + 1 - i) % MOD }
	ans := 0
	for i:=0;i<=N;i++ { ans += poly[i]*poly2[i]; ans %= MOD }
	fmt.Println(ans)
}

