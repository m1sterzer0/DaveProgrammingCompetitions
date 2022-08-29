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

type lstnode struct { n01,n02,n12,n10,n20,n21,n0,n1,n2 int }
type lstfunc struct { s,t,u int }
type lazysegtree struct {
	n, size, log int; op func(lstnode, lstnode) lstnode; mapping func(lstfunc, lstnode) lstnode
	composition func(lstfunc, lstfunc) lstfunc; e lstnode; id lstfunc; d []lstnode; lz []lstfunc
}
func Newlazysegtree(n int, op func(lstnode, lstnode) lstnode, mapping func(lstfunc, lstnode) lstnode, composition func(lstfunc, lstfunc) lstfunc, e lstnode, id lstfunc) *lazysegtree {
	v := make([]lstnode, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []lstnode, op func(lstnode, lstnode) lstnode, mapping func(lstfunc, lstnode) lstnode, composition func(lstfunc, lstfunc) lstfunc, e lstnode, id lstfunc) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]lstnode, 2*sz)
	lz := make([]lstfunc, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v lstnode) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) lstnode {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtree) Prod(l int, r int) lstnode {
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
func (q *lazysegtree) Allprod() lstnode { return q.d[1] }
func (q *lazysegtree) Apply(p int, f lstfunc) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f lstfunc) {
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
func (q *lazysegtree) MaxRight(l int, f func(lstnode) bool) int {
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
func (q *lazysegtree) MinLeft(r int, f func(lstnode) bool) int {
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
func (q *lazysegtree) allApply(k int, f lstfunc) {
	q.d[k] = q.mapping(f, q.d[k]); if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) }
}
func (q *lazysegtree) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); A := gis(N)
	vec := make([]lstnode,0,N)
	for _,a := range(A) {
		n0,n1,n2 := 0,0,0 
		if a == 0 { n0++ } else if a == 1 { n1++ } else { n2++ }
		vec = append(vec,lstnode{0,0,0,0,0,0,n0,n1,n2})
	}

	lstop := func(a,b lstnode) lstnode {
		n01 := a.n01+b.n01+a.n0*b.n1
		n02 := a.n02+b.n02+a.n0*b.n2
		n12 := a.n12+b.n12+a.n1*b.n2
		n10 := a.n10+b.n10+a.n1*b.n0
		n20 := a.n20+b.n20+a.n2*b.n0
		n21 := a.n21+b.n21+a.n2*b.n1
		n0 := a.n0+b.n0
		n1 := a.n1+b.n1
		n2 := a.n2+b.n2
		return lstnode{n01,n02,n12,n10,n20,n21,n0,n1,n2}
	}
	
	lstmap := func(f lstfunc, a lstnode) lstnode {
		if f.s == 0 && f.t == 1 && f.u == 2 { return a }
		n01,n02,n12,n10,n20,n21,n0,n1,n2 := 0,0,0,0,0,0,0,0,0
	
		if f.s == 0 { n0 += a.n0 } else if f.s == 1 { n1 += a.n0 } else { n2 += a.n0 }
		if f.t == 0 { n0 += a.n1 } else if f.t == 1 { n1 += a.n1 } else { n2 += a.n1 }
		if f.u == 0 { n0 += a.n2 } else if f.u == 1 { n1 += a.n2 } else { n2 += a.n2 }
	
		if f.s == 0 && f.t == 1 { n01 += a.n01; n10 += a.n10 }
		if f.s == 0 && f.t == 2 { n02 += a.n01; n20 += a.n10 }
		if f.s == 1 && f.t == 2 { n12 += a.n01; n21 += a.n10 }
		if f.s == 1 && f.t == 0 { n10 += a.n01; n01 += a.n10 }
		if f.s == 2 && f.t == 0 { n20 += a.n01; n02 += a.n10 }
		if f.s == 2 && f.t == 1 { n21 += a.n01; n12 += a.n10 }
	
		if f.s == 0 && f.u == 1 { n01 += a.n02; n10 += a.n20 }
		if f.s == 0 && f.u == 2 { n02 += a.n02; n20 += a.n20 }
		if f.s == 1 && f.u == 2 { n12 += a.n02; n21 += a.n20 }
		if f.s == 1 && f.u == 0 { n10 += a.n02; n01 += a.n20 }
		if f.s == 2 && f.u == 0 { n20 += a.n02; n02 += a.n20 }
		if f.s == 2 && f.u == 1 { n21 += a.n02; n12 += a.n20 }
	
		if f.t == 0 && f.u == 1 { n01 += a.n12; n10 += a.n21 }
		if f.t == 0 && f.u == 2 { n02 += a.n12; n20 += a.n21 }
		if f.t == 1 && f.u == 2 { n12 += a.n12; n21 += a.n21 }
		if f.t == 1 && f.u == 0 { n10 += a.n12; n01 += a.n21 }
		if f.t == 2 && f.u == 0 { n20 += a.n12; n02 += a.n21 }
		if f.t == 2 && f.u == 1 { n21 += a.n12; n12 += a.n21 }
	
		return lstnode{n01,n02,n12,n10,n20,n21,n0,n1,n2}
	}
	
	lstcomp := func(f lstfunc, g lstfunc) lstfunc {
		s := 0; if g.s == 0 { s = f.s } else if g.s == 1 { s = f.t } else { s = f.u }
		t := 0; if g.t == 0 { t = f.s } else if g.t == 1 { t = f.t } else { t = f.u }
		u := 0; if g.u == 0 { u = f.s } else if g.u == 1 { u = f.t } else { u = f.u }
		return lstfunc{s,t,u}
	}
	
	lst := NewlazysegtreeVec(vec,lstop,lstmap,lstcomp,lstnode{0,0,0,0,0,0,0,0,0},lstfunc{0,1,2})
	for i:=0;i<Q;i++ {
		t,l,r := gi(),gi(),gi(); l--; r--
		if t == 1 {
			aa := lst.Prod(l,r)
			ans := aa.n10+aa.n20+aa.n21
			fmt.Fprintln(wrtr,ans)
		} else {
			s,t,u := gi(),gi(),gi()
			lst.ApplyRange(l,r,lstfunc{s,t,u})
		}
	}
}

