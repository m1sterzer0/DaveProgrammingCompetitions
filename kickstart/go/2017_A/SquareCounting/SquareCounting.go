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
func min(a,b int) int { if a > b { return b }; return a }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// 1: (R-1) * (C-1) * 1 = 1*RC - 1R - 1C + 1
		// 2: (R-2) * (C-2) * 2 = 2*RC - 4R - 4C + 8
		// 3: (R-3) * (C-3) * 3 = 3*RC - 9R - 9C + 27
		// 1+2+3+...+N = N * (N+1) / 2
		// 1+4+9+...+N^2 = N * (N+1) * (2N+1) / 6
		// 1+8+27+...+N^3 = N*N*(N+1)*(N+1)/4
		R,C := gi2(); m := min(R-1,C-1)
		ans := R * C % MOD * m % MOD * (m+1) % MOD * powmod(2,MOD-2,MOD) % MOD
		ans += MOD - (R+C) * m % MOD * (m+1) % MOD * (2*m+1) % MOD * powmod(6,MOD-2,MOD) % MOD 
		ans += m * m % MOD * (m+1) % MOD * (m+1) % MOD * powmod(4,MOD-2,MOD) % MOD
		ans %= MOD
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

