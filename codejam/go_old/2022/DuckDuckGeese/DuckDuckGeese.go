package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

// Ensure f(x+y) = f(x) + f(y) (where '+' is the op)
type lazysegtree struct {
	n, size, log int; op func(stnode, stnode) stnode; mapping func(int, stnode) stnode
	composition func(int, int) int; e stnode; id int; d []stnode; lz []int
}
func Newlazysegtree(n int, op func(stnode, stnode) stnode, mapping func(int, stnode) stnode, composition func(int, int) int, e stnode, id int) *lazysegtree {
	v := make([]stnode, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []stnode, op func(stnode, stnode) stnode, mapping func(int, stnode) stnode, composition func(int, int) int, e stnode, id int) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]stnode, 2*sz)
	lz := make([]int, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v stnode) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) Get(p int) stnode {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; return q.d[p]
}
func (q *lazysegtree) Prod(l int, r int) stnode {
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
func (q *lazysegtree) Allprod() stnode { return q.d[1] }
func (q *lazysegtree) Apply(p int, f int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> uint(i)) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> uint(i)) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f int) {
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
func (q *lazysegtree) MaxRight(l int, f func(stnode) bool) int {
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
func (q *lazysegtree) MinLeft(r int, f func(stnode) bool) int {
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
func (q *lazysegtree) allApply(k int, f int) {
	q.d[k] = q.mapping(f, q.d[k]); if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) }
}
func (q *lazysegtree) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }

func solveSmall(N,C int, A,B,P []int) int {
	PP := make([]int,0,2*N)
	for i:=0;i<2;i++ { for _,p := range P { PP = append(PP,p-1) } }
	ans := 0; sb := make([]int,C)
	for i:=0;i<N;i++ {
		for j:=0;j<C;j++ { sb[j] = 0 }
		sb[PP[i]]++;
		badcnt := 0
		if 1 < A[PP[i]] || B[PP[i]] < 1 { badcnt++ }
		for sz:=2;sz<N;sz++ {
			j := i + sz - 1
			c := PP[j]
			oldbad := (sb[c] != 0) && (sb[c] < A[c] || sb[c] > B[c])
			newbad := sb[c]+1 < A[c] || sb[c]+1 > B[c]
			if oldbad && !newbad { badcnt-- } else if newbad && !oldbad { badcnt++ }
			if badcnt == 0 { ans++ }
			sb[c]++
		}
	}
	return ans
}

type stnode struct { maxval,cnt int }
func lstop(a,b stnode) stnode {
	mv := max(a.maxval,b.maxval);
	c := 0; if a.maxval == mv { c += a.cnt }; if b.maxval == mv {c += b.cnt }
	return stnode{mv,c}
}
func lstmap(f int, a stnode) stnode { return stnode{a.maxval+f,a.cnt} }
func lstcmp(f int, g int) int { return f + g }
 
func solveLarge(N,C int, A,B,P []int) int {
	PP := make([]int,0,2*N)
	for i:=0;i<2;i++ { for _,p := range P { PP = append(PP,p-1) } }
	v := make([]stnode,2*N)
	for i:=0;i<2*N;i++ { v[i] = stnode{0,1} }
	lst := NewlazysegtreeVec(v,lstop,lstmap,lstcmp,stnode{0,0},0)
	sb := make([][]int,C); for i,p := range PP { sb[p] = append(sb[p],i) }
	iarr := make([]int,C) // indices for the current color
	getWindow := func(idx,c int) (int,int) {
		//fmt.Printf("DBG: idx:%v c:%v\n",idx,c)
		if B[c] > 0 && len(sb[c]) > B[c]+1+idx { return sb[c][A[c]+idx],sb[c][B[c]+1+idx]-1 }
		if B[c] > 0 && len(sb[c]) > A[c]+idx   { return sb[c][A[c]+idx],2*N-1 }
		return -1,-1
	}
	// Fix the constraints
	for i:=0;i<C;i++ { if B[i] > 0 && A[i] == 0 { A[i] = 1} }
	// Initialize the segment tree
	for i:=0;i<C;i++ {
		if len(sb[i]) == 0 { lst.ApplyRange(0,2*N-1,1); continue }
		if sb[i][0] > 0 { lst.ApplyRange(0,sb[i][0]-1,1) }
		l,r := getWindow(-1,i)
		if l >= 0 { lst.ApplyRange(l,r,1) }
	}
	// Now do the work
	ans := 0
	for i:=0;i<N;i++ {
		v := lst.Prod(i+1,i+N-2)
		if v.maxval == C { ans += v.cnt }
		// Now we need to do the dirty business of removing element i
		c := PP[i]; idx := iarr[c]
		l0,r0 := -1,-1
		if idx+1 == len(sb[c]) { l0,r0 = i+1,2*N-1 }
		if idx+1 < len(sb[c]) && sb[c][idx+1]-sb[c][idx] > 1 { l0,r0 = sb[c][idx]+1,sb[c][idx+1]-1 }
		l1,r1 := getWindow(idx-1,c)
		l2,r2 := getWindow(idx,c)
		if l0 >= 0 { lst.ApplyRange(l0,r0,1) }
		if l1 >= 0 { lst.ApplyRange(l1,r1,-1) }
		if l2 >= 0 { lst.ApplyRange(l2,r2,1) }
		iarr[c]++
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,C := gi(),gi(); A,B := fill2(C); P := gis(N)
		//ans := solveSmall(N,C,A,B,P)
		ans := solveLarge(N,C,A,B,P)
        fmt.Printf("Case #%v: %v\n",tt,ans)
    }
}

