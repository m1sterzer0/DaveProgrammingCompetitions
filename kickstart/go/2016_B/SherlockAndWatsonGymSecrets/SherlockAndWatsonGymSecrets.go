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
func ia(m int) []int { return make([]int,m) }
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
		A,B,N,K := gi(),gi(),gi(),gi()
		sba := ia(K); sbb := ia(K); ans := 0; cycles := N / K
		for i:=1;i<=K;i++ {
			n := cycles; if n*K+i<=N { n++ }
			k := i%K
			ka := 1 % K; if A > 0 { ka = powmod(k,A,K) }
			kb := 1 % K; if B > 0 { kb = powmod(k,B,K) }
			if (ka+kb) % K == 0 { ans += MOD - n % MOD } //subtract out the same number matches
			sba[ka] += n; sba[ka] %= MOD
			sbb[kb] += n; sbb[kb] %= MOD
		}
		for i:=0;i<K;i++ { ans += sba[i] * sbb[(K-i)%K] % MOD }
		ans %= MOD
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

