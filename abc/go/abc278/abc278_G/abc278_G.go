package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)

func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
func sortUniq(a []int) []int {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}

// Very poor man bitset
type FixedLenBitset struct { m int; n int; c []uint64 }
func NewFixedLenBitset(max int) *FixedLenBitset { n := (max+63)/64; return &FixedLenBitset{max, n, make([]uint64, n) } }
func (q *FixedLenBitset) Copy() *FixedLenBitset { a := make([]uint64,q.m); copy(a,q.c); return &FixedLenBitset{q.m, q.n, a } }
func (q *FixedLenBitset) Ins(n int) { q.c[n/64] |= uint64(1) << uint(n % 64) }
func (q *FixedLenBitset) AndEq(a *FixedLenBitset) { for i:=0;i<q.n;i++ { q.c[i] &= a.c[i] } }
func (q *FixedLenBitset) OrEq(a *FixedLenBitset) { for i:=0;i<q.n;i++ { q.c[i] |= a.c[i] } }
func (q *FixedLenBitset) XorEq(a *FixedLenBitset) { for i:=0;i<q.n;i++ { q.c[i] ^= a.c[i] } }
func (q *FixedLenBitset) InvEq() { 
	for i:=0;i<q.n;i++ { q.c[i] = ^q.c[i] };
	if 64*q.n-1 > q.m { numlastbits := q.m - 64*q.n + 65; bm := (uint64(1)<<uint(numlastbits))-1; q.c[q.n-1] &= bm; }
}
func (q *FixedLenBitset) Test(a int) bool { return (q.c[a/64] >> q.c[a%64]) & 1 == 1 }
func (q *FixedLenBitset) CountOnes() int { ans := 0; for i:=0;i<q.n;i++ { ans += bits.OnesCount64(q.c[i]) }; return ans }
func (q *FixedLenBitset) CountZeros() int { return q.m - q.CountOnes() }
func FlbAndCount(a,b *FixedLenBitset) int { n := len(a.c); ans := 0; for i:=0;i<n;i++ { ans += bits.OnesCount64(a.c[i] & b.c[i]) }; return ans }
func FlbOrCount(a,b *FixedLenBitset) int { n := len(a.c); ans := 0; for i:=0;i<n;i++ { ans += bits.OnesCount64(a.c[i] | b.c[i]) }; return ans }
func FlbXorCount(a,b *FixedLenBitset) int { n := len(a.c); ans := 0; for i:=0;i<n;i++ { ans += bits.OnesCount64(a.c[i] ^ b.c[i]) }; return ans }
func FlbAnd(a,b *FixedLenBitset) *FixedLenBitset { q := NewFixedLenBitset(a.m); q.OrEq(a); q.AndEq(b); return q }
func FlbOr(a,b *FixedLenBitset) *FixedLenBitset { q := NewFixedLenBitset(a.m); q.OrEq(a); q.OrEq(b); return q }
func FlbXor(a,b *FixedLenBitset) *FixedLenBitset { q := NewFixedLenBitset(a.m); q.OrEq(a); q.XorEq(b); return q }
func FlbInv(a,b *FixedLenBitset) *FixedLenBitset { q := NewFixedLenBitset(a.m); q.OrEq(a); q.InvEq(); return q }
func (q *FixedLenBitset) FindFirstOne() int {
	cand := -1
	for i,qq := range(q.c) { z := bits.TrailingZeros64(qq); if z == 64 { continue }; cand = 64*i+z; break }
	if cand >= q.m { cand = -1 }
	return cand
}
func (q *FixedLenBitset) FindFirstZero() int {
	cand := -1
	for i,qq := range(q.c) { q2 := ^qq; z := bits.TrailingZeros64(q2); if z == 64 { continue }; cand = 64*i+z; break }
	if cand >= q.m { cand = -1 }
	return cand
}
func (q *FixedLenBitset) FindLastOne() int {
	cand := -1
	for i:=q.n-1;i>=0;i-- { qq := q.c[i]; z := bits.LeadingZeros64(qq); if z == 64 { continue }; cand = 64*i+(63-z); break }
	return cand
}
func (q *FixedLenBitset) FindLastZero() int {
	cand := -1
	for i:=q.n-1;i>=0;i-- { 
		qq := ^(q.c[i]); if 64*q.n-1 > q.m { numlastbits := q.m - 64*q.n + 65; bm := (uint64(1)<<uint(numlastbits))-1; qq |= ^bm; }
		z := bits.LeadingZeros64(qq); if z == 64 { continue }; cand = 64*i+(63-z); break }
	return cand
}




func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N,L,R := gi(),gi(),gi()

	updateBoard := func(board []int, x, y int) { for i:=0;i<y;i++ { board[x-1+i] = 0 } }
	findMove := func(board []int, nimber []int, L int) (int,int) {
		segs,ss := make([]int,0),make([]int,0)
		i:=0
		for i < len(board) {
			for i<len(board) && board[i] == 0 { i++ }
			if i == len(board) { break }
			j := i
			for j+1 < len(board) && board[j+1] == 1 { j++ }
			segs = append(segs,j-i+1); ss = append(ss,i)
			i = j+1
		}
		n := 0; for _,s := range segs { n ^= nimber[s] }
		for ii,sz := range segs {
			st := ss[ii]
			targ := nimber[sz] ^ n
			if targ < nimber[sz] {
				for j:=0;j+L<=sz;j++ { cand := nimber[j] ^ nimber[sz-j-L]; if cand != targ { continue }; return 1+st+j,L }
			}
		}
		return 0,0
	}

	// Split into two separate piles, and mimic
	if R > L || (N % 2 == R % 2) {
		fmt.Fprintln(wrtr,"First"); wrtr.Flush()
		x,y := 0,0
		if N % 2 == R % 2 { x = (N-R)/2+1; y = R } else { x = (N-R+1)/2+1; y = R-1; }
		fmt.Fprintf(wrtr,"%v %v\n",x,y); wrtr.Flush()
		for {
			a,b := gi(),gi()
			if a == 0 && b == 0 || a == -1 && b == -1 { break }
			if a < x+y  { fmt.Fprintf(wrtr,"%v %v\n",a+x+y-1,b); wrtr.Flush() }
			if a >= x+y { fmt.Fprintf(wrtr,"%v %v\n",a-x-y+1,b); wrtr.Flush() }
		}
	} else {
		// Now we have a fixed width L with different parity than N
		// We need to calculate the nimbers for all of the states
		nimber := ia(N+1)
		for i:=0;i<L;i++ { nimber[i] = 0 }
		for i:=L;i<=N;i++ {
			bs := NewFixedLenBitset(2047)
			for j:=0;j+L<=i;j++ { bs.Ins(nimber[j] ^ nimber[i-j-L]) }
			nimber[i] = bs.FindFirstZero()
		}

		board := iai(N,1)
		if nimber[N] == 0 {
			fmt.Fprintln(wrtr,"Second"); wrtr.Flush()
		} else {
			fmt.Fprintln(wrtr,"First"); wrtr.Flush()
			x,y := findMove(board,nimber,L)
			fmt.Fprintf(wrtr,"%v %v\n",x,y); wrtr.Flush()
			updateBoard(board,x,y)
		}
		for {
			a,b := gi(),gi()
			if a == 0 && b == 0 || a == -1 && b == -1 { break }
			updateBoard(board,a,b)
			x,y := findMove(board,nimber,L)
			fmt.Fprintf(wrtr,"%v %v\n",x,y); wrtr.Flush()
			updateBoard(board,x,y)
		}
	}
}

