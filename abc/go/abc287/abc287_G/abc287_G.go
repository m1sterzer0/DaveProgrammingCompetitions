package main

import (
	"bufio"
	"fmt"
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

type segtree struct { n, size, log int; op func(int, int) int; e int; d []int }
func Newsegtree(n int, op func(int, int) int, e int) *segtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }; return NewsegtreeVec(v, op, e)
}
func NewsegtreeVec(v []int, op func(int, int) int, e int) *segtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz); d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i] }; st := &segtree{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *segtree) Set(p int, v int) {
	p += q.size; q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *segtree) Get(p int) int { return q.d[p+q.size] }
func (q *segtree) Prod(l int, r int) int {
	if r < l { return q.e }; r += 1; sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *segtree) Allprod() int { return q.d[1] }
func (q *segtree) MaxRight(l int, f func(int) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; sm := q.e
	for {
		for l%2 == 0 { l >>= 1 }
		if !f(q.op(sm, q.d[l])) {
			for l < q.size { l *= 2; if f(q.op(sm, q.d[l])) { sm = q.op(sm, q.d[l]); l++ } }; return l - q.size - 1
		}
		sm = q.op(sm, q.d[l]); l++; if l&-l == l { break }
	}
	return q.n - 1
}
func (q *segtree) MinLeft(r int, f func(int) bool) int {
	if r < 0 { return 0 }; r += q.size; sm := q.e; r++ 
	for {
		r--; for r > 1 && r%2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r], sm)) {
			for r < q.size { r = 2*r + 1; if f(q.op(q.d[r], sm)) { sm = q.op(q.d[r], sm); r-- } }; return r + 1 - q.size
		}
		sm = q.op(q.d[r], sm); if r&-r == r { break }
	}
	return 0
}
func (q *segtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }



func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N := gi(); A,B := fill2(N); Q := gi(); T,X,Y := ia(Q),ia(Q),ia(Q)
	for i:=0;i<Q;i++ { T[i],X[i] = gi(),gi(); if T[i] != 3 { X[i]--; Y[i] = gi() } }

	// Now we need to compress the scores
	scores := ia(0)
	for _,a := range A { scores = append(scores,a) }
	for _,y := range Y { scores = append(scores,y) }
	scores = sortUniq(scores)
	score2idx := make(map[int]int); for i,s := range scores { score2idx[s] = i }

	// Segment tree for count and scoresum (could combine into one, but i'm lazy)
	stsum := func(a,b int) int { return a + b }
	cntst := Newsegtree(len(scores)+7,stsum,0)
	scoresumst := Newsegtree(len(scores)+7,stsum,0)

	// Initialize the counts and the scoresums
	for i:=0;i<N;i++ {
		idx := score2idx[A[i]]
		oldcnt,oldscoresum := cntst.Get(idx),scoresumst.Get(idx)
		cntst.Set(idx,oldcnt+B[i])
		scoresumst.Set(idx,oldscoresum+A[i]*B[i])
	}

	// Now loop through the queries
	for i,t := range T {
		x := X[i]
		if t == 1 {
			y := Y[i]
			idx1 := score2idx[A[x]]
			oldcnt1 := cntst.Get(idx1)
			oldscoresum1 := scoresumst.Get(idx1)
			cntst.Set(idx1,oldcnt1-B[x])
			scoresumst.Set(idx1,oldscoresum1-A[x]*B[x])
			A[x] = y
			idx2 := score2idx[A[x]]
			oldcnt2 := cntst.Get(idx2)
			oldscoresum2 := scoresumst.Get(idx2)
			cntst.Set(idx2,oldcnt2+B[x])
			scoresumst.Set(idx2,oldscoresum2+A[x]*B[x])
		} else if t == 2 {
			y := Y[i]
			idx := score2idx[A[x]]
			oldcnt := cntst.Get(idx)
			oldscoresum := scoresumst.Get(idx)
			cntst.Set(idx,oldcnt-B[x]+y)
			scoresumst.Set(idx,oldscoresum-A[x]*B[x]+A[x]*y)
			B[x] = y
		} else {
			ans := -1
			if cntst.Allprod() >= x {
				l,r := -1,len(scores)
				for (r-l > 1) {	m := (r+l)>>1; xx := cntst.Prod(m,len(scores)); if xx <= x { r = m } else { l = m } }
				ans = scoresumst.Prod(r,len(scores))
				if r > 0 { ans += scores[r-1]*(x-cntst.Prod(r,len(scores))) }
			}
			fmt.Fprintln(wrtr,ans)
		}
	}
}

