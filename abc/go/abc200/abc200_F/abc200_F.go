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
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	S := gs(); K := gi()
	if len(S) == 1 && K == 1 { fmt.Println(0); return }
	inversions := 0; q := 0
	for _,c := range S { if c == '?' { q++ } }
	for i,c := range S {
		nc := S[0]
		if i+1 < len(S) { nc = S[i+1] }
		if (c == '1' && nc == '1' || c == '0' && nc == '0') { continue }
		if (c == '1' && nc == '0' || c == '0' && nc == '1') { inversions += K*powmod(2,K*q,MOD) % MOD; continue }
		inversions += powmod(2,K*q-1,MOD) * K % MOD
	}
	inversions %= MOD
	ans := inversions * powmod(2,MOD-2,MOD) % MOD
	fmt.Println(ans)
}



