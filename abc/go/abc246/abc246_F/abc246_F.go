package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,L := gi(),gi(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
	mypow := make([]int,27); for i:=0;i<=26;i++ { mypow[i] = powmod(i,L,MOD) }
	bmarr := make([]int,1<<N)
	ans := 0
	for bm:=uint(1);bm<1<<N;bm++ {
		nOnes := bits.OnesCount(bm)
		sgn := 1; if nOnes & 1 == 0 { sgn = -1 }
		if nOnes == 1 {
			pos := bits.TrailingZeros(bm)
			nbm := 0; for _,c := range S[pos] { cc := int(c-'a'); nbm |= 1 << uint(cc) }
			bmarr[bm] = nbm
		} else {
			pos := bits.TrailingZeros(bm)
			nbm1,nbm2 := bmarr[1<<uint(pos)],bmarr[bm ^ (1<<uint(pos))]
			nbm := nbm1 & nbm2
			bmarr[bm] = nbm
		}
		xx := bits.OnesCount(uint(bmarr[bm]))
		ans += MOD + sgn * mypow[xx] % MOD
	}
	ans %= MOD
	fmt.Println(ans)
}

