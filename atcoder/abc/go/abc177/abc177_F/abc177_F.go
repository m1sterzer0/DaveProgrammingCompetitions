package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }

type lst struct {
	n, size, log int; op func(ststate, ststate) ststate; mapping func(stmap, ststate) ststate
	composition func(stmap, stmap) stmap; e ststate; id stmap; d []ststate; lz []stmap
}
func Newlst(n int, op func(ststate, ststate) ststate, mapping func(stmap, ststate) ststate, composition func(stmap, stmap) stmap, e ststate, id stmap) *lst {
	v := make([]ststate, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlstVec(v, op, mapping, composition, e, id)
}
func NewlstVec(v []ststate, op func(ststate, ststate) ststate, mapping func(stmap, ststate) ststate, composition func(stmap, stmap) stmap, e ststate, id stmap) *lst {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]ststate, 2*sz)
	lz := make([]stmap, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lst{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lst) Set(p int, v ststate) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *lst) Get(p int) ststate { p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; return q.d[p] }
func (q *lst) Prod(l int, r int) ststate {
	if r < l { return q.e }; l += q.size; r += q.size; r += 1 
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l { q.push(l >> i) }; if ((r >> i) << i) != r { q.push((r - 1) >> i) }
	}
	sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }; if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }; l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *lst) Allprod() ststate { return q.d[1] }
func (q *lst) Apply(p int, f stmap) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *lst) ApplyRange(l int, r int, f stmap) {
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l { q.push(l >> i) }; if ((r >> i) << i) != r { q.push((r - 1) >> i) }
	}
	l2, r2 := l, r
	for l < r { if l&1 != 0 { q.allApply(l, f); l += 1 }; if r&1 != 0 { r -= 1; q.allApply(r, f) }; l >>= 1; r >>= 1 }
	l, r = l2, r2
	for i := 1; i <= q.log; i++ {
		if ((l >> i) << i) != l { q.update(l >> i) }; if ((r >> i) << i) != r { q.update((r - 1) >> i) }
	}
}
func (q *lst) MaxRight(l int, f func(ststate) bool) int {
	if l == q.n { return q.n - 1 }; l += q.size; for i := q.log; i >= 1; i-- { q.push(l >> i) }; sm := q.e
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
func (q *lst) MinLeft(r int, f func(ststate) bool) int {
	if r < 0 { return 0 }; r += q.size; r++; for i := q.log; i >= 1; i-- { q.push((r - 1) >> i) }; sm := q.e 
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
func (q *lst) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *lst) allApply(k int, f stmap) {
	q.d[k] = q.mapping(f, q.d[k]); if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) }
}
func (q *lst) push(k int) { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }

const inf = 1_000_000_000_000_000_000
type ststate struct { xl,v int }
type stmap struct { left int }
func lstop(a,b ststate) ststate { return ststate{min(a.xl,b.xl),min(a.v,b.v)} } 
func lstmap(f stmap, x ststate) ststate { 
	if f.left < 1 { x.v = inf } else if f.left < inf { x.v = x.xl-f.left }
	return x
}
func lstcomp(f stmap, g stmap) stmap { return stmap{min(f.left,g.left)} }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2(); A,B := fill2(H)
	initvec := make([]ststate,W+1)
	initvec[0] = ststate{0,inf}
	for i:=1;i<=W;i++ { initvec[i] = ststate{i,0} }
	st := NewlstVec(initvec,lstop,lstmap,lstcomp,ststate{inf,inf},stmap{inf})
	ansarr := make([]string,H)
	for i:=0;i<H;i++ {
		st.ApplyRange(A[i],B[i],stmap{A[i]-1-st.Get(A[i]-1).v})
		vv := st.Allprod().v + i + 1; if vv >= inf { vv = -1 }
		ansarr[i] = strconv.Itoa(vv)
	}
	ans := strings.Join(ansarr,"\n")
	fmt.Println(ans)
}

