import argparse
import os.path
from pathlib import Path
import shutil
import sys

def getQueue(classname,datatype) :
    template = '''
type CLASSNAME struct {	buf []DATATYPE; head,tail,sz,bm,l int }
func NewCLASSNAME() *CLASSNAME { buf := make([]DATATYPE,8); return &CLASSNAME{buf,0,0,8,7,0} }
func (q *CLASSNAME) IsEmpty() bool { return q.l == 0 }
func (q *CLASSNAME) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *CLASSNAME) Len() int { return q.l }
func (q *CLASSNAME) Push(x DATATYPE) {
	if q.l == q.sz { q.sizeup() }
	if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *CLASSNAME) Pop() DATATYPE {
	if q.l == 0 { panic("Empty CLASSNAME Pop()") }
	v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }
	return v
}
func (q *CLASSNAME) Head() DATATYPE {if q.l == 0 { panic("Empty CLASSNAME Head()") }; return q.buf[q.head] }
func (q *CLASSNAME) Tail() DATATYPE {if q.l == 0 { panic("Empty CLASSNAME Tail()") }; return q.buf[q.tail] }
func (q *CLASSNAME) sizeup() {
	buf := make([]DATATYPE, 2*q.sz)
	for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }
	q.buf = buf; q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
'''
    template = template.replace("DATATYPE",datatype)
    template = template.replace("CLASSNAME",classname)
    return template

def getDeque(classname,datatype) :
    template = '''
type CLASSNAME struct { buf []DATATYPE; head, tail, sz, bm, l int }
func NewCLASSNAME() *CLASSNAME { buf := make([]DATATYPE, 8); return &CLASSNAME{buf, 0, 0, 8, 7, 0} }
func (q *CLASSNAME) IsEmpty() bool { return q.l == 0 }
func (q *CLASSNAME) Clear()      { q.head = 0; q.tail = 0; q.l = 0 }
func (q *CLASSNAME) PushFront(x DATATYPE) {
	if q.l == q.sz { q.sizeup() }
	if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *CLASSNAME) PushBack(x DATATYPE) {
	if q.l == q.sz { q.sizeup() }
	if q.l > 0 { q.tail = (q.tail + 1) & q.bm }; q.l++; q.buf[q.tail] = x
}
func (q *CLASSNAME) PopFront() DATATYPE {
	if q.l == 0 { panic("Empty CLASSNAME PopFront()") }
	v := q.buf[q.head]; q.l--
	if q.l > 0 { q.head = (q.head + 1) & q.bm } else { q.Clear() }
	return v
}
func (q *CLASSNAME) PopBack() DATATYPE {
	if q.l == 0 { panic("Empty CLASSNAME PopBack()") }
	v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }
	return v
}
func (q *CLASSNAME) Len() int { return q.l }
func (q *CLASSNAME) Head() DATATYPE { if q.l == 0 { panic("Empty CLASSNAME Head()") }; return q.buf[q.head] }
func (q *CLASSNAME) Tail() DATATYPE { if q.l == 0 { panic("Empty CLASSNAME Tail()") }; return q.buf[q.tail] }
func (q *CLASSNAME) sizeup() {
	buf := make([]DATATYPE, 2*q.sz)
	for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }
	q.buf = buf; q.head = 0; q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
'''
    template = template.replace("DATATYPE",datatype)
    template = template.replace("CLASSNAME",classname)
    return template

def getStack(classname,datatype) :
    template = '''
type CLASSNAME struct { buf []DATATYPE; l int }
func NewCLASSNAME() *CLASSNAME { buf := make([]DATATYPE,0); return &CLASSNAME{buf,0} }
func (q *CLASSNAME) IsEmpty() bool { return q.l == 0 }
func (q *CLASSNAME) Clear() { q.buf = q.buf[:0]; q.l = 0 }
func (q *CLASSNAME) Len() int { return q.l }
func (q *CLASSNAME) Push(x DATATYPE) { q.buf = append(q.buf,x); q.l++ }
func (q *CLASSNAME) Pop() DATATYPE { if q.l == 0 { panic("Empty CLASSNAME Pop()") }; v := q.buf[q.l-1]; q.l--; q.buf=q.buf[:q.l]; return v }
func (q *CLASSNAME) Head() DATATYPE {if q.l == 0 { panic("Empty CLASSNAME Head()") }; return q.buf[q.l-1] }
func (q *CLASSNAME) Top() DATATYPE { return q.Head() }
'''
    template = template.replace("DATATYPE",datatype)
    template = template.replace("CLASSNAME",classname)
    return template

def getMinheap(classname,datatype) :
    template = '''
type CLASSNAME struct { buf []DATATYPE; less func(DATATYPE,DATATYPE)bool }
func NewCLASSNAME(f func(DATATYPE,DATATYPE)bool) *CLASSNAME { buf := make([]DATATYPE, 0); return &CLASSNAME{buf,f} }
func (q *CLASSNAME) IsEmpty() bool { return len(q.buf) == 0 }
func (q *CLASSNAME) Clear() { q.buf = q.buf[:0] }
func (q *CLASSNAME) Len() int { return len(q.buf) }
func (q *CLASSNAME) Push(v DATATYPE) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *CLASSNAME) Head() DATATYPE { return q.buf[0] }
func (q *CLASSNAME) Pop() DATATYPE {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }
	return v1
}
func (q *CLASSNAME) Heapify(pri []DATATYPE) { q.buf=append(q.buf,pri...); n:=len(q.buf); for i:=n/2-1;i>=0;i-- { q.siftup(i) } }
func (q *CLASSNAME) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos { ppos:=(pos-1)>>1; p:=q.buf[ppos]; if !q.less(newitem,p) { break } ;q.buf[pos], pos = p, ppos }
	q.buf[pos] = newitem
}
func (q *CLASSNAME) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos+1; if rtpos < endpos && !q.less(q.buf[chpos],q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos],pos = q.buf[chpos],chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}
'''
    template = template.replace("DATATYPE",datatype)
    template = template.replace("CLASSNAME",classname)
    return template

def getSegtree(classname,datatype) :
    template = '''
type CLASSNAME struct { n,size,log int; op func(DATATYPE,DATATYPE) int; e DATATYPE; d []DATATYPE }
func NewCLASSNAME(n int, op func(DATATYPE,DATATYPE) DATATYPE, e DATATYPE) *CLASSNAME { v := make([]DATATYPE, n); for i:=0; i<n; i++ { v[i] = e }; return NewCLASSNAMEVec(v, op, e) }
func NewCLASSNAMEVec(v []DATATYPE, op func(DATATYPE,DATATYPE) DATATYPE, e DATATYPE) *CLASSNAME {
	n,sz,log := len(v),1,0; for n < sz { sz <<= 1; log += 1 }
	d := make([]DATATYPE, 2*sz); d[0] = e; for i := 0; i < n; i++ { d[sz+i] = v[i] }
	st := &CLASSNAME{n, sz, log, op, e, d}
	for i := sz - 1; i >= 1; i-- { st.update(i) }
	return st
}
func (q *CLASSNAME) Set(p int, v DATATYPE) { p += q.size; q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> i) } }
func (q *CLASSNAME) Get(p int) DATATYPE { return q.d[p+q.size] }
// Gives product from l to r inclusive
func (q *CLASSNAME) Prod(l int, r int) DATATYPE {
	// We add 1 to right vs. atcoder, as we want to get all the points from l->r inclusive
	if r < l { return q.e }; r += 1; sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }
		if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }
		l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *CLASSNAME) Allprod() DATATYPE { return q.d[1] }
// Given monotone f, finds maximum r such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *CLASSNAME) MaxRight(l int, f func(DATATYPE)bool) int{
	if l == q.n { return q.n-1 }; l += q.size; sm := q.e;
	for {
		for l % 2 == 0 { l >>= 1 }
		if !f(q.op(sm,q.d[l])) { for l < q.size { l *= 2; if f(q.op(sm,q.d[l])) { sm = q.op(sm,q.d[l]); l++ } }; return l - q.size - 1 }
		sm = q.op(sm,q.d[l]); l++; if l & -l == l { break }
	}
	return q.n-1 
}
// Given monotone f, finds minimum l such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *CLASSNAME) MinLeft(r int, f func(DATATYPE)bool) int{
	if r < 0 { return 0 }; r += q.size; sm := q.e; r++ //r++ for the fully closed vs. half open
	for {
		r--; for r > 1 && r % 2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r],sm)) { for r < q.size { r = 2*r+1; if f(q.op(q.d[r],sm)) { sm = q.op(q.d[r],sm); r-- } }; return r+1-q.size }
		sm = q.op(q.d[r],sm); if r & -r == r { break }
	}
	return 0
}
func (q *CLASSNAME) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
'''
    template = template.replace("DATATYPE",datatype)
    template = template.replace("CLASSNAME",classname)
    return template
 
def getLazySegtree(classname,datatype,functype) :
    template = '''
type CLASSNAME struct { n,size,log int; op func(DATATYPE,DATATYPE)DATATYPE; mapping func(FUNCTYPE,DATATYPE)DATATYPE; composition func(FUNCTYPE,FUNCTYPE)FUNCTYPE; e DATATYPE; id FUNCTYPE; d []DATATYPE; lz []FUNCTYPE }
func NewCLASSNAME(n int, op func(DATATYPE,DATATYPE) DATATYPE, mapping func(FUNCTYPE,DATATYPE) DATATYPE, composition func(FUNCTYPE,FUNCTYPE) FUNCTYPE, e DATATYPE, id FUNCTYPE) *CLASSNAME {
	v := make([]DATATYPE, n); for i := 0; i < n; i++ { v[i] = e }; return NewCLASSNAMEVec(v, op, mapping, composition, e, id)
}
func NewCLASSNAMEVec(v []DATATYPE, op func(DATATYPE,DATATYPE) DATATYPE, mapping func(FUNCTYPE,DATATYPE) DATATYPE, composition func(FUNCTYPE,FUNCTYPE) FUNCTYPE, e DATATYPE, id FUNCTYPE) *CLASSNAME {
	n,sz,log := len(v),1,0; for n < sz { sz <<= 1; log += 1 }; d := make([]DATATYPE, 2*sz); lz := make([]FUNCTYPE, sz); d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &CLASSNAME{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }
	return st
}
func (q *CLASSNAME) Set(p int, v DATATYPE) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }
	q.d[p] = v; for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *CLASSNAME) Get(p int) DATATYPE { p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; return q.d[p] }
// Gets the product from l to r incluseive
func (q *CLASSNAME) Prod(l int, r int) DATATYPE {
	if r < l { return q.e }; l += q.size; r += q.size; r += 1 // r+1 for close right end interval
	for i := q.log; i >= 1; i-- { 
		if ((l >> i) << i) != l { q.push(l >> i) }
		if ((r >> i) << i) != r { q.push((r - 1) >> i) }
	}
	sml, smr := q.e, q.e; l += q.size; r += q.size
	for l < r {
		if l&1 != 0 { sml = q.op(sml, q.d[l]); l++ }
		if r&1 != 0 { r--; smr = q.op(q.d[r], smr) }
		l >>= 1; r >>= 1
	}
	return q.op(sml, smr)
}
func (q *CLASSNAME) Allprod() DATATYPE { return q.d[1] }
func (q *CLASSNAME) Apply(p int, f FUNCTYPE) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }
	q.d[p] = q.mapping(f, q.d[p]); for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *CLASSNAME) ApplyRange(l int, r int, f FUNCTYPE) {
	// Add one, as our versioin applies from l to r inclusive
	if r < l { return }; r += 1; l += q.size; r += q.size
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l { q.push(l >> i) }
		if ((r >> i) << i) != r { q.push((r - 1) >> i) }
	}
	l2, r2 := l, r
	for l < r {
		if l&1 != 0 { q.allApply(l, f); l += 1 }
		if r&1 != 0 { r -= 1; q.allApply(r, f) }
		l >>= 1; r >>= 1
	}
	l, r = l2, r2
	for i := q.log; i >= 1; i-- {
		if ((l >> i) << i) != l { q.update(l >> i) }
		if ((r >> i) << i) != r { q.update((r - 1) >> i) }
	}
}
// Given monotone f, finds maximum r such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *CLASSNAME) MaxRight(l int, f func(DATATYPE)bool) int{
	if l == q.n { return q.n-1 }; l += q.size; for i:=q.log; i >= 1; i-- { q.push(l>>i) }; sm := q.e;
	for {
		for l % 2 == 0 { l >>= 1 }
		if !f(q.op(sm,q.d[l])) { for l < q.size { q.push(l); l *= 2; if f(q.op(sm,q.d[l])) { sm = q.op(sm,q.d[l]); l++ } }; return l - q.size - 1 }
		sm = q.op(sm,q.d[l]); l++; if l & -l == l { break }
	}
	return q.n-1 
}
// Given monotone f, finds minimum l such that f(op(a[l],a[l+1],...,a[r])) = true
func (q *CLASSNAME) MinLeft(r int, f func(DATATYPE)bool) int{
	if r < 0 { return 0 }; r += q.size; r++; for i:=q.log; i >= 1; i-- { q.push((r-1)>>i) }; sm := q.e  //r++ for the fully closed vs. half open
	for {
		r--; for r > 1 && r % 2 == 1 { r >>= 1 }
		if !f(q.op(q.d[r],sm)) { for r < q.size { q.push(r); r = 2*r+1; if f(q.op(q.d[r],sm)) { sm = q.op(q.d[r],sm); r-- } }; return r+1-q.size }
		sm = q.op(q.d[r],sm); if r & -r == r { break }
	}
	return 0
}
func (q *CLASSNAME) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *CLASSNAME) allApply(k int, f FUNCTYPE) { q.d[k] = q.mapping(f, q.d[k]);  if k < q.size { q.lz[k] = q.composition(f, q.lz[k]) } }
func (q *CLASSNAME) push(k int) { q.allApply(2*k, q.lz[k]);  q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id }
'''
    template = template.replace("DATATYPE",datatype)
    template = template.replace("CLASSNAME",classname)
    template = template.replace("FUNCTYPE",functype)
    return template

def getConvolver() :
    template = '''
type Convolver struct { mod, root, rootinv, rootpw int }
// (998244353,31,23) works
func NewConvolver(mod, root, rootdepth int) *Convolver {
	rootinv, e, m := 1, mod-2, root
	for e > 0 { if e&1 != 0 { rootinv = rootinv * m % mod }; m = m * m % mod; e >>= 1 }
	return &Convolver{mod, root, rootinv, 1 << rootdepth}
}
// Not as fancy as atcoder version which takes double-steps to save modulus operations. Instead modelled after cpalgorithms version
func (q *Convolver) NTT(a []int, invert bool) {
	mod := q.mod; n := len(a)
	for i, j := 1, 0; i < n; i++ {
		bit := n >> 1; for ; j&bit != 0; bit >>= 1 { j ^= bit }; j ^= bit
		if i < j { a[i], a[j] = a[j], a[i] }
	}
	for ll := 2; ll <= n; ll <<= 1 {
		wlen := q.root; if invert { wlen = q.rootinv }
		for i := ll; i < q.rootpw; i <<= 1 { wlen = wlen * wlen % mod }
		for i := 0; i < n; i += ll {
			w := 1; lover2 := ll >> 1
			for j := 0; j < lover2; j++ {
				idx1 := i + j; idx2 := idx1 + lover2; u := a[idx1]; v := a[idx2] * w % mod; v1 := u + v;  v2 := u - v
				if v1 >= mod { v1 -= mod }; if v2 < 0 { v2 += mod }
				a[idx1], a[idx2] = v1, v2; w = w * wlen % mod
			}
		}
	}
	if invert {
		ninv, e, m := 1, mod-2, n
		for e > 0 { if e&1 != 0 { ninv = ninv * m % mod }; m = m * m % mod; e >>= 1 }
		for i := 0; i < n; i++ { a[i] *= ninv; a[i] %= mod }
	}
}
func (q *Convolver) Convolve(a []int, b []int) []int {
	mod := q.mod; finalsz := len(a) + len(b) - 1; z := 1; for z < finalsz { z *= 2 }
	lena, lenb := len(a), len(b); la := make([]int, z); lb := make([]int, z)
	for i := 0; i < lena; i++ { la[i] = a[i] }
	for i := 0; i < lenb; i++ { lb[i] = b[i] }
	q.NTT(la, false); q.NTT(lb, false)
	for i := 0; i < z; i++ { la[i] *= lb[i]; la[i] %= mod }
	q.NTT(la, true)
	return la[:finalsz]
}
'''
    return template

def getFenwick() :
    template = '''
type Fenwick struct { n,tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx); }; q.tot += val }
func (q *Fenwick) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *Fenwick) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *Fenwick) Prefixsum(idx int) int { if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans }
func (q *Fenwick) Suffixsum(idx int) int { return q.tot - q.Prefixsum(idx-1) }
func (q *Fenwick) Rangesum(left int, right int) int { if right < left { return 0 }; return q.Prefixsum(right) - q.Prefixsum(left-1) }
'''
    return template

def getMaxflow() :
    template = '''
type mfpreedge struct { to,rev,cap int }
type mfedge struct { from,to,cap,flow int }
type mfpos struct { x,y int }
type Mfgraph struct { n int; pos []mfpos; g [][]mfpreedge }
func NewMfgraph(n int) *Mfgraph { g := make([][]mfpreedge, n); pos := make([]mfpos, 0); return &Mfgraph{n, pos, g} }
func (q *Mfgraph) Addedge(from, to, cap int) int {
	m := len(q.pos); fromid := len(q.g[from]); toid := len(q.g[to])
	q.pos = append(q.pos, mfpos{from, fromid})
	if from == to { toid++ }
	q.g[from] = append(q.g[from], mfpreedge{to, toid, cap})
	q.g[to] = append(q.g[to], mfpreedge{from, fromid, 0})
	return m
}
func (q *Mfgraph) Getedge(i int) mfedge { e := q.g[q.pos[i].x][q.pos[i].y]; re := q.g[e.to][e.rev]; return mfedge{q.pos[i].x, e.to, e.cap + re.cap, re.cap} }
func (q *Mfgraph) Edges() []mfedge { m := len(q.pos); res := make([]mfedge, 0); for i := 0; i < m; i++ { res = append(res, q.Getedge(i)) }; return res }
func (q *Mfgraph) Changeedge(i int, newcap int, newflow int) { e := &(q.g[q.pos[i].x][q.pos[i].y]); re := &(q.g[e.to][e.rev]); e.cap = newcap - newflow; re.cap = newflow }
func (q *Mfgraph) Flow(s, t int) int { return q.FlowCapped(s, t, 1_000_000_000_000_000_000) }
func (q *Mfgraph) FlowCapped(s int, t int, flowlimit int) int {
	level := make([]int, q.n); iter := make([]int, q.n)
	bfs := func() {
		for i := 0; i < q.n; i++ { level[i] = -1 }; level[s] = 0
		que := make([]int,0,q.n); que = append(que,s)
		for len(que) > 0 {
			v := que[0]; que = que[1:]
			for _, e := range q.g[v] {
				if e.cap == 0 || level[e.to] >= 0 { continue }
				level[e.to] = level[v] + 1; if e.to == t { return }; que = append(que,e.to)
			}
		}
	}
	var dfs func(int, int) int
	dfs = func(v int, up int) int {
		if v == s { return up }
		res := 0; level_v := level[v]
		for i := iter[v]; i < len(q.g[v]); i++ {
			e := q.g[v][i]; cap := q.g[e.to][e.rev].cap
			if level_v <= level[e.to] || cap == 0 { continue }
			newup := up - res; if cap < up-res { newup = cap }
			d := dfs(e.to, newup)
			if d <= 0 { continue }
			q.g[v][i].cap += d; q.g[e.to][e.rev].cap -= d; res += d
			if res == up { return res }
		}
		level[v] = q.n
		return res
	}
	flow := 0
	for flow < flowlimit {
		bfs(); if level[t] == -1 { break }
		for i := 0; i < q.n; i++ { iter[i] = 0 }
		f := dfs(t, flowlimit-flow); if f == 0 { break }; flow += f
	}
	return flow
}
func (q *Mfgraph) Mincut(s int) []bool {
	visited := make([]bool, q.n); que := make([]int,0,q.n); que = append(que,s)
	for len(que) > 0 {
		p := que[0]; que = que[1:]; visited[p] = true
		for _, e := range q.g[p] { if e.cap > 0 && !visited[e.to] { visited[e.to] = true; que = append(que,e.to) } }
	}
	return visited
}
'''
    return template

def getDsu() :
    template = '''
type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n);  for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int { if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans }
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }; ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ { 
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}
'''
    return template

def getDsuSparse() :
	template = '''
type DsuSparse struct {	n int; parentOrSize map[int]int }
func NewDsuSparse() *DsuSparse { mm := make(map[int]int); return &DsuSparse{0, mm} }
func (q *DsuSparse) Add(x int) { q.n++; q.parentOrSize[x] = -1 }
func (q *DsuSparse) Leader(a int) int { if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]);  q.parentOrSize[a] = ans; return ans }
func (q *DsuSparse) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x 
}
func (q *DsuSparse) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *DsuSparse) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *DsuSparse) Groups() [][]int {
	numgroups := 0; leader2idx := make(map[int]int); ans := make([][]int, 0);
	for i := 0; i <= q.n; i++ {
		l := q.Leader(i)
		v, ok := leader2idx[l]
		if !ok { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; v = numgroups; numgroups += 1 }
		ans[v] = append(ans[v], i)
	}
	return ans
}
'''
	return template

def getMinCostFlow() :
    template = '''
type MinCostFlowPI struct { c,v int }
type MinHeapMinCostFlow struct { buf []MinCostFlowPI; less func(MinCostFlowPI,MinCostFlowPI)bool }
func NewMinHeapMinCostFlow(f func(MinCostFlowPI,MinCostFlowPI)bool) *MinHeapMinCostFlow { buf := make([]MinCostFlowPI, 0); return &MinHeapMinCostFlow{buf,f} }
func (q *MinHeapMinCostFlow) IsEmpty() bool { return len(q.buf) == 0 }
func (q *MinHeapMinCostFlow) Push(v MinCostFlowPI) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *MinHeapMinCostFlow) Pop() MinCostFlowPI {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else {	l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }
	return v1
}
func (q *MinHeapMinCostFlow) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos { ppos:=(pos-1)>>1; p:=q.buf[ppos]; if !q.less(newitem,p) { break } ;q.buf[pos], pos = p, ppos }
	q.buf[pos] = newitem
}
func (q *MinHeapMinCostFlow) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos+1; if rtpos < endpos && !q.less(q.buf[chpos],q.buf[rtpos]) {	chpos = rtpos }
		q.buf[pos],pos = q.buf[chpos],chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}
type MinCostFlow struct { n,numedges int; g [][]int; to,cap,cost []int }
func NewMinCostFlow(n int) *MinCostFlow {
	g := make([][]int, n); to := make([]int, 0); cap := make([]int, 0); cost := make([]int, 0);  return &MinCostFlow{n, 0, g, to, cap, cost}
}
func (q *MinCostFlow) AddEdge(fr, to, cap, cost int) {
	q.to = append(q.to, to); q.to = append(q.to, fr); q.cap = append(q.cap, cap); q.cap = append(q.cap, 0)
	q.cost = append(q.cost, cost); q.cost = append(q.cost, -cost); q.g[fr] = append(q.g[fr], q.numedges); q.g[to] = append(q.g[to], q.numedges+1)
	q.numedges += 2
}
// Successive shortest paths
// Requirement -- no negative cycles
// In theory -- O(n*m+m*log(m)*B) where B bounds the total flow
// but with potentials and positive costs at first, it gets to
// O(m*log(m)*B)
func (q *MinCostFlow) Flowssp(s, t int) (int, int) {
	inf := 1_000_000_000_000_000_000; res := 0; h := make([]int, q.n); prv_v := make([]int, q.n); prv_e := make([]int, q.n); f := 0
	dist := make([]int, q.n); for i := 0; i < q.n; i++ { dist[i] = inf }
	for {
		for i := 0; i < q.n; i++ { dist[i] = inf }; dist[s] = 0
		que := NewMinHeapMinCostFlow(func(a,b MinCostFlowPI) bool { return a.c < b.c }); que.Push(MinCostFlowPI{0,s})
		for !que.IsEmpty() {
			xx := que.Pop(); c,v := xx.c,xx.v; if dist[v] < c {	continue }; r0 := dist[v] + h[v]
			for _, e := range q.g[v] {
				w, cap, cost := q.to[e], q.cap[e], q.cost[e]
				if cap > 0 && r0+cost-h[w] < dist[w] { r := r0 + cost - h[w]; dist[w] = r; prv_v[w] = v; prv_e[w] = e; que.Push(MinCostFlowPI{r, w}) }
			}
		}
		if dist[t] == inf {	return f, res }
		for i := 0; i < q.n; i++ { h[i] += dist[i] }
		d := inf
		v := t
		for v != s { dcand := q.cap[prv_e[v]]; if dcand < d { d = dcand }; v = prv_v[v] }
		f += d; res += d * h[t]; v = t
		for v != s { e := prv_e[v]; e2 := e ^ 1; q.cap[e] -= d; q.cap[e2] += d; v = prv_v[v] }
	}
}
'''
    return template

def getScc() :
	template = '''
func Kosaraju(n int, diredges []PI) (int, []int) {
	g, grev, visited, visitedInv, scc, s, counter := make([][]int, n), make([][]int, n), make([]bool, n), make([]bool, n), make([]int, n), make([]int,0,n), 0
	var dfs1, dfs2 func(int)
	for _, xx := range diredges { x, y := xx.x, xx.y; g[x] = append(g[x], y); grev[y] = append(grev[y], x) }
	dfs1 = func(u int) { if !visited[u]    { visited[u] = true;    for _, c := range g[u]    { dfs1(c) }; s = append(s,u) } }
	for i := 0; i < n; i++ { dfs1(i) }
	dfs2 = func(u int) { if !visitedInv[u] { visitedInv[u] = true; for _, c := range grev[u] { dfs2(c) }; scc[u] = counter } }
	for i := n-1; i >= 0; i-- { nn := s[i]; if !visitedInv[nn] { dfs2(nn); counter += 1 } }
	return counter, scc
}
'''
	return template

def getTwoSat() :
	s1 = getScc()
	template = '''
type Twosat struct { n int; answer []bool; edgelist []PI }
func NewTwosat(n int) *Twosat { answer := make([]bool, n); edgelist := make([]PI, 0); return &Twosat{n, answer, edgelist} }
func (q *Twosat) AddClause(i int, f bool, j int, g bool) {
	n1,n2,n3,n4 := 2*i, 2*j, 2*j, 2*i
	if f { n4 += 1 } else { n1 += 1	}; if g { n2 += 1 } else { n3 += 1 }
	q.edgelist = append(q.edgelist, PI{n1, n2}); q.edgelist = append(q.edgelist, PI{n3, n4})
}
func (q *Twosat) Satisfiable() (bool, []bool) {
	_, id := Kosaraju(2*q.n, q.edgelist)
	for i := 0; i < q.n; i++ { if id[2*i] == id[2*i+1] { return false, q.answer }; q.answer[i] = id[2*i] < id[2*i+1] }
	return true, q.answer
}
'''
	return s1+template

def appendStringToFile(s,fn) :
    if fn.lower() == "stdout" : print(s); return
    with open(fn,"at") as fp : print(s,file=fp)

helpstring = ("gosnippets.py generates competitive programming code with user-customized classes for certain elements." +
              "USAGE:\n" +
              "    python3 gosnippets.py queue       CLASSNAME DATATYPE FILENAME\n" +
              "    python3 gosnippets.py stack       CLASSNAME DATATYPE FILENAME\n" +
              "    python3 gosnippets.py deque       CLASSNAME DATATYPE FILENAME\n" +
              "    python3 gosnippets.py minheap     CLASSNAME DATATYPE FILENAME\n" +
              "    python3 gosnippets.py segtree     CLASSNAME DATATYPE FILENAME\n" +
              "    python3 gosnippets.py lazysegtree CLASSNAME DATATYPE FUNCTIONTYPE FILENAME\n" +
              "    python3 gosnippets.py convolver   FILENAME\n" +
              "    python3 gosnippets.py fenwick     FILENAME\n" +
              "    python3 gosnippets.py maxflow     FILENAME\n" +
              "    python3 gosnippets.py dsu         FILENAME\n" +
              "    python3 gosnippets.py dsusparse   FILENAME\n" +
              "    python3 gosnippets.py scc         FILENAME\n" +
              "    python3 gosnippets.py twosat      FILENAME\n" +
              "NOTE: gosnippets.py will APPEND the results to the given filenme.\n" +
			  "Use 'stdout' for filename if you just want the codesnippet on stdout.")

## Minheap check -- abc164 E


if __name__ == "__main__" :
    good = True
    if len(sys.argv) <= 1 : good = False
    if len(sys.argv) > 1 :
        if   sys.argv[1] == "queue" and len(sys.argv) == 5 :       s = getQueue(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        elif sys.argv[1] == "stack" and len(sys.argv) == 5 :       s = getStack(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        elif sys.argv[1] == "deque" and len(sys.argv) == 5 :       s = getDeque(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        elif sys.argv[1] == "minheap" and len(sys.argv) == 5 :     s = getMinheap(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        elif sys.argv[1] == "segtree" and len(sys.argv) == 5 :     s = getSegtree(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        elif sys.argv[1] == "lazysegtree" and len(sys.argv) == 6 : s = getLazySegtree(sys.argv[2],sys.argv[3],sys.argv[4]); appendStringToFile(s,sys.argv[5])
        elif sys.argv[1] == "convolver" and len(sys.argv) == 3 :   s = getConvolver(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "fenwick" and len(sys.argv) == 3 :     s = getFenwick(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "maxflow" and len(sys.argv) == 3 :     s = getMaxflow(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "dsu" and len(sys.argv) == 3 :         s = getDsu(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "dsusparse" and len(sys.argv) == 3 :   s = getDsuSparse(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "mincostflow" and len(sys.argv) == 3 : s = getMinCostFlow(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "scc" and len(sys.argv) == 3 :         s = getScc(); appendStringToFile(s,sys.argv[2])
        elif sys.argv[1] == "twosat" and len(sys.argv) == 3 :      s = getTwoSat(); appendStringToFile(s,sys.argv[2])
        else : good = False
    if not good : print(helpstring)




         
