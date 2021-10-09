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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func abs(a int) int { if a < 0 { return -a }; return a }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type Bitset struct { m int; c []uint64 }
func NewBitset(cap int) *Bitset { return &Bitset{0, make([]uint64, 0, cap)} }
func (q *Bitset) Copy() *Bitset {
	c2 := make([]uint64, len(q.c)); for i, x := range q.c { c2[i] = x }; return &Bitset{q.m, c2}
}
func (q *Bitset) Ins(n int) { for q.m <= n { q.c = append(q.c, 0); q.m += 64 }; q.c[n/64] |= 1 << (n % 64) }
func (q *Bitset) Del(n int) { if q.m > n { q.c[n/64] &= 0xffffffffffffffff ^ (1 << (n % 64)) } }
func (q *Bitset) Contains(n int) bool { if q.m <= n { return false }; return q.c[n/64] & (1 << (n % 64)) != 0 }
func (q *Bitset) Flip(n int) { for q.m <= n { q.c = append(q.c, 0); q.m += 64 }; q.c[n/64] ^= 1 << (n % 64) }
func (q *Bitset) Size() int { return q.m }
func (q *Bitset) Any() bool { for _, cc := range q.c { if cc != 0 { return true } }; return false }
func (q *Bitset) None() bool { for _, cc := range q.c { if cc != 0 { return false } }; return true }
func (q *Bitset) Count() int { ans := 0; for _, cc := range q.c { ans += bits.OnesCount64(cc) }; return ans }
func (q *Bitset) PadTo(a *Bitset) { for q.m < a.m { q.c = append(q.c, 0); q.m += 64 } }
func (q *Bitset) And(a *Bitset) { q.shrinkTo(q.m); la := len(a.c); for i := 0; i < la; i++ { q.c[i] &= a.c[i] } }
func (q *Bitset) Or(a *Bitset) { q.PadTo(a); la := len(a.c); for i := 0; i < la; i++ { q.c[i] |= a.c[i] } }
func (q *Bitset) Xor(a *Bitset) { q.PadTo(a); la := len(a.c); for i := 0; i < la; i++ { q.c[i] ^= a.c[i] } }
func (q *Bitset) Cap(n int) { q.shrinkTo(n) }
func (q *Bitset) Not() { lc := len(q.c); for i := 0; i < lc; i++ { q.c[i] = ^q.c[i] } }
func (q *Bitset) Shl(a int) {
	q.shrink(); if q.m == 0 { return }; mm := q.max() + 1; newmm := mm + a
	for q.m < newmm { q.c = append(q.c, 0); q.m += 64 }; g, b := a/64, a%64
	for i := len(q.c) - 1; i >= 0; i-- {
		if i-g < 0 {
			q.c[i] = 0
		} else {
			q.c[i] = q.c[i-g] << b; if i-g-1 >= 0 && b != 0 { q.c[i] |= q.c[i-g-1] >> (64 - b) }
		}
	}
}
func (q *Bitset) Shr(a int) {
	g, b, lc := a/64, a%64, len(q.c)
	for i := 0; i < lc; i++ {
		if i+g >= lc {
			q.c[i] = 0
		} else {
			q.c[i] = q.c[i+g] >> b; if i+g+1 < lc && b != 0 { q.c[i] |= q.c[i+g+1] << (64 - b) }
		}
	}
	q.shrink()
}
func (q *Bitset) GetBits() []int {
	base := 0; ans := []int{}
	for _, c := range q.c {
		for c != 0 { offset := bits.TrailingZeros64(c); ans = append(ans, base+offset); c ^= 1 << offset }; base += 64
	}
	return ans
}
func (q *Bitset) shrink() { for i := len(q.c) - 1; i >= 0 && q.c[i] == 0; i-- { q.c = q.c[:i]; q.m -= 64 } }
func (q *Bitset) shrinkTo(a int) { i := len(q.c) - 1; for q.m-64 > a { q.c = q.c[:i]; q.m -= 64 } }
func (q *Bitset) max() int {
	lc := len(q.c); if q.c[lc-1] == 0 { q.shrink(); lc = len(q.c) }; if lc == 0 { return -1 }
	return 64*lc - 1 - bits.LeadingZeros64(q.c[lc-1])
}
func BitsetAnd(a, b *Bitset) *Bitset { c := a.Copy(); c.And(b); return c }
func BitsetOr(a, b *Bitset) *Bitset { c := a.Copy(); c.Or(b); return c }
func BitsetXor(a, b *Bitset) *Bitset { c := a.Copy(); c.Xor(b); return c }
func BitsetShl(a *Bitset, n int) *Bitset { c := a.Copy(); c.Shl(n); return c }
func BitsetShr(a *Bitset, n int) *Bitset { c := a.Copy(); c.Shr(n); return c }

const MOD int = 998244353

func solve(N,M,K int, A,U,V []int) int {
	cnts := ia(N)
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	var dfs func(n,p int) *Bitset
	dfs = func(n,p int) *Bitset {
		res := NewBitset(N)
		res.Ins(n)
		for _,c := range gr[n] {
			if c == p { continue }
			cres := dfs(c,n)
			res.Or(cres)
		}
		trav := 0
		for i:=0;i<M-1;i++ {
			if res.Contains(A[i]) != res.Contains(A[i+1]) { trav++ }
		}
		cnts[n] = trav
		return res
	}
	dfs(0,-1)
	// Now we have counts, so it is time for subset sum DP
	sumcnt := sumarr(cnts); K = abs(K)
	if sumcnt < K || (sumcnt+K) % 2 == 1 { return 0 }
	targ := (sumcnt + K) / 2
	dp := ia(targ+1)
	dp[0] = 1
	for _,c := range cnts[1:] {
		for j:= targ-c; j >= 0; j-- {
			dp[j+c] += dp[j]; dp[j+c] %= MOD
		}
	}
	return dp[targ]
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,K := gi3(); A := gis(M); for i:=0;i<M;i++ { A[i]-- }; U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	ans := solve(N,M,K,A,U,V)
	fmt.Println(ans)
}

