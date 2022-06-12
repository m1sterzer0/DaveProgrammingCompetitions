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
func gi64() int64 { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func powmod64(a,e,mod int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func makefact64(n int64,mod int64) ([]int64,[]int64) {
	fact,factinv := make([]int64,n+1),make([]int64,n+1)
	fact[0] = 1; for i:=int64(1);i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod64(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const MOD int64 = 1000000007
type st struct {idx,v int64}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,L,R,Z := gi64(),gi64(),gi64(),gi64()
	fact,factinv := makefact64(N+10,MOD)
	combarr := make([]int64,N+1)
	for i:=int64(0);i<=N;i++ { combarr[i] = fact[N] * factinv[i] % MOD * factinv[N-i] % MOD  }
	var dp [61][2002]int64
	var bitdp func(idx int64, rem int64, val int64) int64
	bitdp = func(idx int64, rem int64, val int64) int64 {
		if rem > 2*N { rem = 2*N }
		if rem < 0 { return 0 }
		if idx < 0 { return 1 }
		if dp[idx][rem] != -1 { return dp[idx][rem] }
		ans := int64(0); st := int64(0); if Z & (int64(1) << uint64(idx)) != 0 { st++ }
		for i:=st;i<=N;i+=2 {
			cbs := int64(0); if val & (int64(1) << uint64(idx)) != 0 { cbs++ }
			adder := combarr[i] * bitdp(idx-1, 2*(rem+cbs-i), val) % MOD
			ans += adder; ans %= MOD
		}
		//fmt.Printf("DBG: bitdp(%v,%v,%v)=%v\n",idx,rem,val,ans)
		dp[idx][rem] = ans
		return ans
	}
	calc := func(ubound int64) int64 {
		for i:=0;i<=60;i++ { for j:=int64(0);j<=2*N;j++ { dp[i][j] = -1 } }
		return bitdp(60,0,ubound)
	}
	ans := (calc(R)+MOD-calc(L-1)) % MOD
	fmt.Println(ans)
}
