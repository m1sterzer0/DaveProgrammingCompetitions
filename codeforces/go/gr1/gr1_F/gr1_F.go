package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
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
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type query struct { v,l,r,idx int }
type edge struct { n2,w int }
const inf int64 = 1000000000000000000

type lazysegtree struct {
	n, size, log int; op func(int64, int64) int64; mapping func(int64, int64) int64
	composition func(int64, int64) int64; e int64; id int64; d []int64; lz []int64
}
func Newlazysegtree(n int, op func(int64, int64) int64, mapping func(int64, int64) int64, composition func(int64, int64) int64, e int64, id int64) *lazysegtree {
	v := make([]int64, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []int64, op func(int64, int64) int64, mapping func(int64, int64) int64, composition func(int64, int64) int64, e int64, id int64) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int64, 2*sz)
	lz := make([]int64, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v int64) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) int64 {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtree) Prod(l int, r int) int64 {
	if r < l { return q.e }; l += q.size; r += q.size; r += 1 
	for i := q.log; i >= 1; i-- {
		if ((l >> uint(i)) << uint(i)) != l { q.push(l >> uint(i)) }
		if ((r >> uint(i)) << uint(i)) != r { q.push((r - 1) >> uint(i)) }
	}
	sml, smr := q.e, q.e
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *lazysegtree) Allprod() int64 { return q.d[1] }
func (q *lazysegtree) Apply(p int, f int64) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f int64) {
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := uint(q.log); i >= 1; i-- {
		if ((l >> i) << i) != l { xx := l >> i;       if q.lz[xx] != q.id { q.push(xx) } }
		if ((r >> i) << i) != r { xx := (r - 1) >> i; if q.lz[xx] != q.id { q.push(xx) } }
	}
	l2, r2 := l, r
	for l < r { if l&1 != 0 { q.allApply(l, f); l += 1 }; if r&1 != 0 { r -= 1; q.allApply(r, f) }; l >>= 1; r >>= 1 }
	l, r = l2, r2
	for i := uint(1); i <= uint(q.log); i++ {
		if ((l >> i) << i) != l { q.update(l >> i) }
		if ((r >> i) << i) != r { q.update((r - 1) >> i) }
	}
}
func (q *lazysegtree) MaxRight(l int, f func(int64) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; for i := q.log; i >= 1; i-- { q.push(l >> uint(i)) }; sm := q.e
	for {
		for l%2 == 0 { l >>= 1 }
		if !f(q.op(sm, q.d[l])) {
			for l < q.size { q.push(l); l *= 2; if f(q.op(sm, q.d[l])) { sm = q.op(sm, q.d[l]); l++ } }
			return l - q.size - 1
		}
		sm = q.op(sm, q.d[l]); l++; if l&-l == l { break }
	}
	return q.n - 1
}
func (q *lazysegtree) MinLeft(r int, f func(int64) bool) int {
	if r < 0 { return 0 }; r += q.size; r++; for i := q.log; i >= 1; i-- { q.push((r - 1) >> uint(i)) }; sm := q.e 
	for {
		r--; for r > 1 && r%2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r], sm)) {
			for r < q.size { q.push(r); r = 2*r + 1; if f(q.op(q.d[r], sm)) { sm = q.op(q.d[r], sm); r-- } }
			return r + 1 - q.size
		}
		sm = q.op(q.d[r], sm); if r&-r == r { break }
	}
	return 0
}
func (q *lazysegtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *lazysegtree) allApply(k int, f int64) {
	q.d[k] = q.mapping(f, q.d[k]); if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) }
}
func (q *lazysegtree) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); P,W := fill2(N-1); V,L,R := fill3(Q)
	// Fix this P/W to be usable to give the parent and weight of the edge leading to node n
	PP,WW := make([]int,N),make([]int64,N)
	for i:=1;i<N;i++ { PP[i] = P[i-1]-1; WW[i] = int64(W[i-1]) }
	for i:=0;i<Q;i++ { V[i]--; L[i]--; R[i]-- }
	qq := make([][]int,N); for i:=0;i<Q;i++ { v := V[i]; qq[v] = append(qq[v],i) }
	isLeaf := make([]bool,N); for i:=0;i<N;i++ { isLeaf[i] = true }; for i:=1;i<N;i++ { isLeaf[PP[i]] = false }
	sz := iai(N,1); for i:=N-1;i>=1;i-- { sz[PP[i]] += sz[i] }  // take advantage of preorder eulerian traversal
	dist := make([]int64,N); for i:=1;i<N;i++ { dist[i] = dist[PP[i]] + WW[i] }
	leafd := make([]int64,N); for i:=0;i<N;i++ { if isLeaf[i] { leafd[i] = dist[i] } else { leafd[i] = inf } }
	lstmin := func(a,b int64) int64 { if b < a { return b }; return a }
	lstsum := func(a,b int64) int64 { return a + b }
	lst := NewlazysegtreeVec(leafd,lstmin,lstsum,lstsum,inf,0)
	ansarr := make([]int64,Q); curs := 0; adder := int64(0)
	for i:=0;i<N;i++ {
		if i != 0 {
			for curs != PP[i] { w := WW[curs]; adder -= w; lst.ApplyRange(curs,curs+sz[curs]-1,2*w); curs = PP[curs] }
			adder += WW[i]; lst.ApplyRange(i,i+sz[i]-1,-2*WW[i]); curs = i 
		}
		for _,qidx := range qq[i] {
			ansarr[qidx] = adder + lst.Prod(L[qidx],R[qidx])
		}
	}
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}
