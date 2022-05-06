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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	pow26 := make([]int,1000010)
	pow26[0] = 1; for i:=1;i<1000010;i++ { pow26[i] = pow26[i-1] * 26 % MOD }
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); S := gs()
		numlet := (N+1)/2
		ans := 0
		for i:=0;i<numlet;i++ {
			l := int(S[i]-'A')
			ans += l * pow26[numlet-i-1] % MOD
		}
		// Construct the implied palindrome from the first numlet letters of S and see if it is less than or equal to S
		tt := make([]byte,N)
		for i,j:=0,N-1;i<=j;i,j=i+1,j-1 { tt[i] = S[i]; tt[j] = S[i] }
		s2 := string(tt)
		if s2 <= S { ans++ }
		ans %= MOD
		fmt.Fprintln(wrtr,ans)
	}

}

