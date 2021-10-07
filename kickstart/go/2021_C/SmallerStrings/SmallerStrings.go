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
func ia(m int) []int { return make([]int,m) }
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
		N,K := gi2(); S := gs(); ans := 0; maxtochoose := 0
		powk := ia(N+1); powk[0] = 1; for i:=1;i<=N;i++ { powk[i] = K * powk[i-1] % MOD }
		if N%2 == 0 { maxtochoose = N / 2 } else { maxtochoose = (N+1)/2 }
		for i:=0;i<maxtochoose;i++ { ans += int(S[i]-'a') * powk[maxtochoose-i-1] % MOD }

		// One check to see if the "natural palidrome" induced by the string prefix is
		// less than the original string -- this needs to be added back in.
		SS2 := []byte(S); i,j := 0,N-1
		for i < j { SS2[j] = SS2[i]; i++; j-- }
		S2 := string(SS2)
		if S2 < S { ans++ }
		ans %= MOD
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

