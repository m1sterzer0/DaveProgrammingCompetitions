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
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}


type lstfunc struct { typ,val int }
type lstnode struct { sum,min,max,len int }
type lazysegtreebeats struct {
	n, size, log int
	op func(lstnode, lstnode) lstnode
	mapping func(lstfunc, lstnode) (lstnode,bool)
	composition func(lstfunc, lstfunc) lstfunc
	e lstnode
	id lstfunc
	d []lstnode
	lz []lstfunc
}
func Newlazysegtreebeats(n int, op func(lstnode, lstnode) lstnode, mapping func(lstfunc, lstnode) (lstnode,bool), composition func(lstfunc, lstfunc) lstfunc,
	e lstnode, id lstfunc) *lazysegtreebeats {
		v := make([]lstnode, n); for i := 0; i < n; i++ { v[i] = e }
		return NewlazysegtreebeatsVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreebeatsVec(v []lstnode, op func(lstnode, lstnode) lstnode, mapping func(lstfunc, lstnode) (lstnode,bool), composition func(lstfunc, lstfunc) lstfunc,
	e lstnode, id lstfunc) *lazysegtreebeats {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]lstnode, 2*sz)
	lz := make([]lstfunc, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtreebeats{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtreebeats) Set(p int, v lstnode) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtreebeats) Get(p int) lstnode {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtreebeats) Prod(l int, r int) lstnode {
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
func (q *lazysegtreebeats) Allprod() lstnode { return q.d[1] }
func (q *lazysegtreebeats) ApplyRange(l int, r int, f lstfunc) {
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := q.log; i >= 1; i-- {
		if ((l >> uint(i)) << uint(i)) != l { q.push(l >> uint(i)) }
		if ((r >> uint(i)) << uint(i)) != r { q.push((r - 1) >> uint(i)) }
	}
	l2, r2 := l, r
	for l < r { if l&1 != 0 { q.allApply(l, f); l += 1 }; if r&1 != 0 { r -= 1; q.allApply(r, f) }; l >>= 1; r >>= 1 }
	l, r = l2, r2
	for i := 1; i <= q.log; i++ {
		if ((l >> uint(i)) << uint(i)) != l { q.update(l >> uint(i)) }
		if ((r >> uint(i)) << uint(i)) != r { q.update((r - 1) >> uint(i)) }
	}
}
func (q *lazysegtreebeats) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *lazysegtreebeats) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }
func (q *lazysegtreebeats) Apply(p int, f lstfunc) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p],_ = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtreebeats) allApply(k int, f lstfunc) {
	var ok bool
	q.d[k],ok = q.mapping(f, q.d[k])
	if k < q.size { q.lz[k] = q.composition(f, q.lz[k]); if !ok { q.push(k); q.update(k) } }
}

func lstop (a,b lstnode) lstnode { return lstnode{a.sum+b.sum,min(a.min,b.min),max(a.max,b.max),a.len+b.len} }
func lstmap (f lstfunc, a lstnode) (lstnode,bool) { 
	if f.typ == 0 { return a,true }
	if f.typ == 2 { return lstnode{a.len*f.val,f.val,f.val,a.len},true }
	v := a.min/f.val
	if a.max/f.val == v { 
	    return lstnode{a.len*v,v,v,a.len},true
	} else {
		return lstnode{0,0,0,0},false
	}
}
func lstcomp (f lstfunc, g lstfunc) lstfunc {
	if f.typ == 0 { return g }
	if g.typ == 0 { return f }
	if f.typ == 2 { return f }
	if g.typ == 2 { return lstfunc{2,g.val/f.val } }
	if f.val * g.val > 100000 { return lstfunc{2,0} }
	return lstfunc{1,f.val*g.val}
}

const inf = 1000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); A := gis(N)
	v := make([]lstnode,N)
	for i:=0;i<N;i++ { v[i] = lstnode{A[i],A[i],A[i],1} }
	lst := NewlazysegtreebeatsVec( v,lstop,lstmap,lstcomp,lstnode{0,inf,0,0},lstfunc{0,0} )
	for i:=0;i<Q;i++ { 
		t := gi()
		if t == 1 || t == 2 {
			l,r,x := gi(),gi(),gi(); lst.ApplyRange(l-1,r-1,lstfunc{t,x})
		} else {
			l,r := gi(),gi(); ans := lst.Prod(l-1,r-1); fmt.Fprintln(wrtr,ans.sum)
		}
	}
}
