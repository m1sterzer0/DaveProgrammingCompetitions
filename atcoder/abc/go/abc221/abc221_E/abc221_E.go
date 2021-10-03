package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func copyarr(a []int) []int { a2 := make([]int,len(a)); for i:=0;i<len(a);i++ { a2[i] = a[i] }; return a2 }
func uniqSorted(a []int) []int { 
	a2 := []int{}
	for i:=0;i<len(a);i++ { if i == 0 || a[i] != a[i-1] { a2 = append(a2,a[i]) } }
	return a2
}
const MOD int = 998244353
type Fenwick struct { n, tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx) }; q.tot += val }
func (q *Fenwick) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *Fenwick) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans
}
func (q *Fenwick) Suffixsum(idx int) int { return q.tot - q.Prefixsum(idx-1) }
func (q *Fenwick) Rangesum(left int, right int) int {
	if right < left { return 0 }; return q.Prefixsum(right) - q.Prefixsum(left-1)
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N);
	// Need to query the number of elements to the right of A[idx] that are  >= A[idx].  Call this n.
	// Then the number of sequences that start with A[idx] is 2^n - 1
	// Initially, this looks too big for a fenwick tree (or segment tree), but we can do coordinate
	// compression first before we start to get this down to a reasonable size.

	// Coordinate compression code
	A2 := copyarr(A)
	sort.Slice(A2,func(i,j int)bool{return A2[i] < A2[j]})
	A3 := uniqSorted(A2)
	cc := make(map[int]int)
	for i,a := range A3 { cc[a] = i+1 } //plus 1 for fenwick tree
	A4 := ia(N); for i:=0;i<N;i++ { A4[i] = cc[A[i]] }

	// Fenwick tree to count sequences.  Store 2^(k-1) at A[k], and then divide by 2^i to get number of sequences 
	ans := 0; ft := NewFenwick(len(A3)+10)
	for i:=N-1;i>=0;i-- {
		adder := ft.Suffixsum(A4[i]) % MOD * powmod(powmod(2,i,MOD),MOD-2,MOD) % MOD
		ans += adder
		if i > 0 { ft.Inc(A4[i],powmod(2,i-1,MOD)) }
	}
	ans %= MOD; fmt.Println(ans)
}

