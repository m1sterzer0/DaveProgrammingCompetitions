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
const MOD int = 998244353
func comblow(a,b int) int {
	ans := 1
	for i:=0;i<b;i++  { ans = (a-i) % MOD * ans % MOD }
	for i:=b;i>=2;i-- { ans = ans * powmod(i,MOD-2,MOD) % MOD }
	return ans
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	// One messy way is to try to do inclusion exclusion with the unique cases of (6,51,42,411,33,321,3111,222,2211,21111,11111).
	// I think the simpler way is to use Burside's lemma.
	// If we didn't have to deal with rotations, we could just use (stars and bars) * 6! and we would be done.
	// However, we need to cancel out copies that are the same.
	// Burside's Lemma to the rescue (https://en.wikipedia.org/wiki/Burside%27s_lemma)
	// Burside says to count the number of orbits (i.e. equivalence classes of solutions w.r.t. cube rotation) we need to sum
	// the number of pairs (g,X) where X is a solution "fixed" by rotation g, and then finally divide by the sum by the number of rotations.
	// There are 24 rotations of the cube (6 ways to pick top side, and 4 ways to pick front face)
	// - 1 identity "non-rotation" -- all solutions are "fixed" by this
	// - 6 90 degree face rotations (3 for picking pair that is up/down, and 2 for rotating left vs. right).  These are "fixed"
	//   by a set of 4 identical faces on the parts that are rotating.
	// - 3 180 degree face rotations -- these are "fixed" by having 2 pairs of identical elements on opposing side amongst the
	//   4 rotating faces
	// - 8 120 degree vertex rotations (taking two opposite corners and then spinning the cube on that axis).  These are fixed
	//   by 2 sets of 3
	// - 6 180 degree edge rotations (grip cube on center of one edge and center of opposite edge and spin cube on that axis).
	//   These are fixed by 3 pairs.
	//  Thus we need to count (1/24) * (ALL + 6 * "411 cases" + 3 * "2211 cases" + 8 * "33 cases" + 6 * "222 cases")
	S := gi(); S -= 6
	ans := comblow(S+5,5) // Stars and bars for identity case
	// 3-3 case
	if S % 3 == 0 { ans += 8 * comblow(S/3+1,1) }
	// 2-2-2 case
	if S % 2 == 0 { ans += 6 * comblow(S/2+2,2) }
	// 4-1-1 case
	for a:=0;a<4;a++ {
		for b:=0;b<4;b++ {
			n := S-a-b
			if n < 0 || n%4 != 0 { continue }
			ans += 6 * comblow(n/4+2,2)
		}
	}
	// 2-2-1-1 case
	for a:=0;a<2;a++ {
		for b:=0;b<2;b++ {
			n := S-a-b
			if n < 0 || n%2 != 0 { continue }
			ans += 3 * comblow(n/2+3,3)
		}
	}
	ans %= MOD; ans *= powmod(24,MOD-2,MOD); ans %= MOD
	fmt.Println(ans)
}
