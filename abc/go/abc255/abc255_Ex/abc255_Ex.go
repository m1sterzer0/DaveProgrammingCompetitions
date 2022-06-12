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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func sortUniqueIntarr(a []int) []int {
	sort.Slice(a,func (i,j int) bool { return a[i] < a[j] })
	i,j,la := 0,0,len(a)
	for ;i<la;i++ { if i == 0 || a[i] != a[i-1] { a[j] = a[i]; j++ } }
	return a[:j]
}
const MOD = 998244353
type node struct { perday,tot int }
type stfunc struct { typ,val int }

type lazysegtree struct {
	n, size, log int; op func(node, node) node; mapping func(stfunc, node) node
	composition func(stfunc, stfunc) stfunc; e node; id stfunc; d []node; lz []stfunc
}
func Newlazysegtree(n int, op func(node, node) node, mapping func(stfunc, node) node, composition func(stfunc, stfunc) stfunc, e node, id stfunc) *lazysegtree {
	v := make([]node, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []node, op func(node, node) node, mapping func(stfunc, node) node, composition func(stfunc, stfunc) stfunc, e node, id stfunc) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]node, 2*sz)
	lz := make([]stfunc, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v node) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) node {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtree) Prod(l int, r int) node {
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
func (q *lazysegtree) Allprod() node { return q.d[1] }
func (q *lazysegtree) Apply(p int, f stfunc) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f stfunc) {
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
func (q *lazysegtree) MaxRight(l int, f func(node) bool) int {
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
func (q *lazysegtree) MinLeft(r int, f func(node) bool) int {
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
func (q *lazysegtree) allApply(k int, f stfunc) {
	q.d[k] = q.mapping(f, q.d[k]); if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) }
}
func (q *lazysegtree) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }

func lstop(a,b node) node { 
	return node{(a.perday+b.perday)%MOD,(a.tot+b.tot)%MOD}
}
func lstmap(f stfunc, a node) node {
	if f.typ == 0 { return a } // Identity function
	if f.typ == 1 { return node{a.perday,0} }     // Clear
	if f.typ == 2 { inc := f.val % MOD * a.perday % MOD; return node{a.perday, (a.tot+inc) % MOD} } // Increment
	if f.typ == 3 { inc := f.val % MOD * a.perday % MOD; return node{a.perday, inc} } // Set
	return node{0,0} // shouldn't get here
}
func lstcomp(f stfunc, g stfunc) stfunc {
	if f.typ == 0 {
		return g
	} else if g.typ == 0 {
		return f
	} else if f.typ == 1 || f.typ == 3 {
		return f
	} else if f.typ == 2 {
		if g.typ == 1 { return stfunc{3,f.val} }
		if g.typ == 2 { return stfunc{2,(f.val+g.val)%MOD} }
		if g.typ == 3 { return stfunc{3,(f.val+g.val)%MOD} }
	}
	return stfunc{0,0} // shouldn't get here
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	gi(); //N
	Q := gi()
	D,L,R := fill3(Q)
	// Clearly need some coordinate compression
	c := make([]int,0,2*Q)
	for _,l := range L { c = append(c,l) }
	for _,r := range R { c = append(c,r+1) }
	c = sortUniqueIntarr(c); nc := len(c)
	revc := make(map[int]int)
	for i,cc := range c { revc[cc] = i }
	twoinv,_ := modinv(2,MOD)
	nodes := make([]node,0,nc)
	for i:=0;i<nc-1;i++ {
		l := (c[i]) % MOD; rp1 := c[i+1]; r := (rp1-1) % MOD
		perday := r * (r+1) % MOD * twoinv % MOD + MOD
		perday -= (l-1) * l % MOD * twoinv % MOD
		perday %= MOD
		nodes = append(nodes,node{perday,0} )
	}
	lst := NewlazysegtreeVec(nodes,lstop,lstmap,lstcomp,node{0,0},stfunc{0,0})
	lastd := 0
	for i:=0;i<Q;i++ {
		d,l,r := D[i],L[i],R[i]
		lst.ApplyRange(0,nc-2,stfunc{2,(d-lastd)%MOD})
		lastd = d
		l1 := revc[l]; r1p1 := revc[r+1]; r1 := r1p1-1
		ans := lst.Prod(l1,r1)
		fmt.Fprintln(wrtr,ans.tot)
		lst.ApplyRange(l1,r1,stfunc{1,0})
	}
}

