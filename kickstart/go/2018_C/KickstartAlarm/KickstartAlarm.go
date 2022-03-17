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
		N,K,x1,y1,C,D,E1,E2,F := gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi()
		X,Y,A := ia(N),ia(N),ia(N); X[0],Y[0],A[0] = x1,y1,(x1+y1)%F
		for i:=1;i<N;i++ {
			X[i] = (C * X[i-1] + D * Y[i-1] + E1) % F
			Y[i] = (D * X[i-1] + C * Y[i-1] + E2) % F
			A[i] = (X[i] + Y[i]) % F
		}
		coeff := 0; ans := 0
		for i:=0;i<N;i++ {
			if i == 0 { 
				coeff += K 
			} else { 
				// (i+1) + (i+1)^2 + ... + (i+1)^K = ( (i+1)^(K+1) - 1 ) / ( (i+1)-1 ) - 1
				adder := (powmod(i+1,K+1,MOD) - 1) * powmod(i+1-1,MOD-2,MOD) % MOD + (MOD-1)
				coeff = (coeff+adder) % MOD
			}
			numendpoints := N - i
			adder := A[i] * coeff % MOD * numendpoints % MOD
			ans += adder; ans %= MOD
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

