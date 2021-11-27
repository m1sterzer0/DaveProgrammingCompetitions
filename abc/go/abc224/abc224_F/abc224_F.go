package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Iterate from the back
	// 2nd digit is in                   1s  10s 100s  1000s  10000s  100000s 1000000s
	//                         2nd digit  1   1
	//                         3rd digit  2   1   1
	//                         4th digit  4   2   1     1
	//                         5th digit  8   4   2     1       1
	S := gs()
	ans := int(S[len(S)-1]-'0'); pv := 11; cnt := 1
	for i:=len(S)-2;i>=0;i-- {
		v := int(S[i]-'0')
		ans = (2 * ans + pv * v) % MOD
		cnt = (2 * cnt) % MOD
		pv = (10 * pv + cnt) % MOD
	}
	fmt.Println(ans)
}

