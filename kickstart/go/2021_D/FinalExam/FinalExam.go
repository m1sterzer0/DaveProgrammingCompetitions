package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type ival struct {l,r int}


type rbtreesetnode struct { left, right, up int32; red bool; key ival }
type rbtreeset struct {
	lessthan func(a, b ival) bool; tree []rbtreesetnode; root int32; recycler []int32; sz int; minidx int32
	maxidx int32
}
type rbtreesetIterator interface { Next() (ok bool); Prev() (ok bool); Key() ival }
type rbtreesetiter struct { cur int32; key ival; rbtree *rbtreeset }
func (i *rbtreesetiter) Key() ival { return i.key }
func (i *rbtreesetiter) Next() bool {
	rbtree := i.rbtree; v := rbtree.nextidx(i.cur); if v == 0 { return false }; i.cur, i.key = v, rbtree.tree[v].key
	return true
}
func (i *rbtreesetiter) Prev() bool {
	rbtree := i.rbtree; v := rbtree.previdx(i.cur); if v == 0 { return false }; i.cur, i.key = v, rbtree.tree[v].key
	return true
}
func Newrbtreeset(lessthan func(a, b ival) bool) *rbtreeset {
	q := &rbtreeset{lessthan, make([]rbtreesetnode, 2), int32(0), make([]int32, 0), 0, 0, 0}
	q.tree[0].left, q.tree[0].right, q.tree[0].up, q.tree[0].red = 0, 0, 0, false; q.recycler = append(q.recycler, 1)
	return q
}
func (q *rbtreeset) Add(k ival) {
	if q.sz == 0 {
		z := q.getNewNodenum(); tree := q.tree; q.minidx, q.maxidx, q.sz, q.root = z, z, q.sz+1, z
		tree[z].key, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 0, 0, 0, false; return
	}
	y, cmp := q.findInsertionPoint(k); if cmp == 0 { return }; z := q.getNewNodenum(); q.sz += 1; tree := q.tree
	tree[z].key, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, y, 0, 0, true
	if cmp < 0 { tree[y].left = z } else { tree[y].right = z }
	if q.sz == 0 || q.lessthan(k, tree[q.minidx].key) { q.minidx = z }
	if q.sz == 0 || q.lessthan(tree[q.maxidx].key, k) { q.maxidx = z }; var p, g, u int32
	for p = tree[z].up; tree[p].red; p = tree[z].up { 
		g = tree[p].up ; if g == 0 { break } 
		if p == tree[g].left {
			u = tree[g].right
			if tree[u].red { tree[p].red, tree[u].red, tree[g].red, z = false, false, true, g; continue }
			if z == tree[p].right { z = p; q.rotleft(z); p = tree[z].up }; q.rotright(g)
			tree[g].red, tree[p].red = true, false
		} else { 
			u = tree[g].left
			if tree[u].red { tree[p].red, tree[u].red, tree[g].red, z = false, false, true, g; continue }
			if z == tree[p].left { z = p; q.rotright(z); p = tree[z].up }; q.rotleft(g)
			tree[g].red, tree[p].red = true, false
		}
	}
	tree[q.root].red = false
}
func (q *rbtreeset) Delete(k ival) bool {
	if q.sz == 0 { return false }; z, cmp := q.findInsertionPoint(k); if cmp != 0 { return false }; q.sz--
	q.recycler = append(q.recycler, z)
	if q.sz > 0 && !q.lessthan(q.tree[q.minidx].key, k) { q.minidx = q.nextidx(q.minidx) }
	if q.sz > 0 && !q.lessthan(k, q.tree[q.maxidx].key) { q.maxidx = q.previdx(q.maxidx) }
	if q.sz == 0 { q.root = 0; return true }; tree := q.tree; var x int32; y, y_orig_red := z, tree[z].red
	if tree[z].left == 0 {
		x = tree[z].right; q.rbTransplant(z, x)
	} else if tree[z].right == 0 {
		x = tree[z].left; q.rbTransplant(z, x)
	} else {
		y = q.findminidx(tree[z].right); y_orig_red = tree[y].red; x = tree[y].right
		if tree[y].up == z {
			tree[x].up = y 
		} else {
			q.rbTransplant(y, x); tree[y].right = tree[z].right; tree[tree[y].right].up = y
		}
		q.rbTransplant(z, y); tree[y].left = tree[z].left; tree[tree[y].left].up = y; tree[y].red = tree[z].red
	}
	if !y_orig_red {
		for q.root != x && !tree[x].red {
			p := tree[x].up
			if tree[p].left == x {
				s := tree[p].right 
				if tree[s].red { tree[s].red = false; tree[p].red = true; q.rotleft(p); s = tree[p].right }
				c := tree[s].left; d := tree[s].right
				if !tree[c].red && !tree[d].red {
					tree[s].red = true; x = p
				} else {
					if !tree[d].red { tree[c].red = false; tree[s].red = true; q.rotright(s); s = tree[p].right }
					tree[s].red = tree[p].red; tree[p].red = false; tree[tree[s].right].red = false; q.rotleft(p)
					x = q.root
				}
			} else {
				s := tree[p].left 
				if tree[s].red { tree[s].red = false; tree[p].red = true; q.rotright(p); s = tree[p].left }
				c := tree[s].right; d := tree[s].left
				if !tree[c].red && !tree[d].red {
					tree[s].red = true; x = p
				} else {
					if !tree[d].red { tree[c].red = false; tree[s].red = true; q.rotleft(s); s = tree[p].left }
					tree[s].red = tree[p].red; tree[p].red = false; tree[tree[s].left].red = false; q.rotright(p)
					x = q.root
				}
			}
		}
		tree[x].red = false
	}
	return true
}
func (q *rbtreeset) Clear() {
	q.tree, q.root, q.recycler, q.sz = q.tree[:2], 0, q.recycler[:0], 0; q.recycler = append(q.recycler, int32(1))
}
func (q *rbtreeset) IsEmpty() bool { return q.sz == 0 }
func (q *rbtreeset) Contains(k ival) bool { _, cmp := q.findInsertionPoint(k); return cmp == 0 }
func (q *rbtreeset) Count(k ival) int { _, cmp := q.findInsertionPoint(k); if cmp == 0 { return 1 }; return 0 }
func (q *rbtreeset) Len() int { return q.sz }
func (q *rbtreeset) MinKey() (k ival) {
	if q.sz == 0 { panic("Called MinKey on an empty rbtreeset") }; return q.tree[q.minidx].key
}
func (q *rbtreeset) MaxKey() (k ival) {
	if q.sz == 0 { panic("Called MaxKey on an empty rbtreeset") }; return q.tree[q.maxidx].key
}
func (q *rbtreeset) findLtIdx(k ival) (int32, bool) {
	if q.sz == 0 || !q.lessthan(q.tree[q.minidx].key, k) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos != 1 { idx = q.previdx(idx) }; return idx, true
}
func (q *rbtreeset) findLeIdx(k ival) (int32, bool) {
	if q.sz == 0 || q.lessthan(k, q.tree[q.minidx].key) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos == -1 { idx = q.previdx(idx) }; return idx, true
}
func (q *rbtreeset) findGtIdx(k ival) (int32, bool) {
	if q.sz == 0 || !q.lessthan(k, q.tree[q.maxidx].key) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos != -1 { idx = q.nextidx(idx) }; return idx, true
}
func (q *rbtreeset) findGeIdx(k ival) (int32, bool) {
	if q.sz == 0 || q.lessthan(q.tree[q.maxidx].key, k) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos == 1 { idx = q.nextidx(idx) }; return idx, true
}
func (q *rbtreeset) FindLt(k ival) (ival, bool) {
	var ans ival; idx, ok := q.findLtIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreeset) FindLe(k ival) (ival, bool) {
	var ans ival; idx, ok := q.findLeIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreeset) FindGt(k ival) (ival, bool) {
	var ans ival; idx, ok := q.findGtIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreeset) FindGe(k ival) (ival, bool) {
	var ans ival; idx, ok := q.findGeIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreeset) FindLtIter(k ival) (rbtreesetIterator, bool) {
	var ans *rbtreesetiter; idx, ok := q.findLtIdx(k); if ok { ans = &rbtreesetiter{idx, q.tree[idx].key, q} }
	return ans, ok
}
func (q *rbtreeset) FindLeIter(k ival) (rbtreesetIterator, bool) {
	var ans *rbtreesetiter; idx, ok := q.findLeIdx(k); if ok { ans = &rbtreesetiter{idx, q.tree[idx].key, q} }
	return ans, ok
}
func (q *rbtreeset) FindGtIter(k ival) (rbtreesetIterator, bool) {
	var ans *rbtreesetiter; idx, ok := q.findGtIdx(k); if ok { ans = &rbtreesetiter{idx, q.tree[idx].key, q} }
	return ans, ok
}
func (q *rbtreeset) FindGeIter(k ival) (rbtreesetIterator, bool) {
	var ans *rbtreesetiter; idx, ok := q.findGeIdx(k); if ok { ans = &rbtreesetiter{idx, q.tree[idx].key, q} }
	return ans, ok
}
func (q *rbtreeset) FindIter(k ival) (rbtreesetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx, pos := q.findInsertionPoint(k); if pos != 0 { return nil, false }
	return &rbtreesetiter{idx, q.tree[idx].key, q}, true
}
func (q *rbtreeset) MinIter() (rbtreesetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx := q.findminidx(q.root)
	return &rbtreesetiter{idx, q.tree[idx].key, q}, true
}
func (q *rbtreeset) MaxIter() (rbtreesetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx := q.findmaxidx(q.root)
	return &rbtreesetiter{idx, q.tree[idx].key, q}, true
}
func (q *rbtreeset) rbTransplant(u, v int32) {
	tree := q.tree
	if tree[u].up == 0 {
		q.root = v
	} else {
		p := tree[u].up; if u == tree[p].left { tree[p].left = v } else { tree[p].right = v }
	}
	tree[v].up = tree[u].up
}
func (q *rbtreeset) findInsertionPoint(k ival) (int32, int8) {
	n, lt, tree := q.root, q.lessthan, q.tree
	for {
		nkey := tree[n].key
		if lt(nkey, k) {
			r := tree[n].right; if r == 0 { return n, 1 }; n = r
		} else if lt(k, nkey) {
			l := tree[n].left; if l == 0 { return n, -1 }; n = l
		} else {
			return n, 0
		}
	}
}
func (q *rbtreeset) findmaxidx(n1 int32) int32 {
	tree := q.tree; for { xx := tree[n1].right; if xx == 0 { break }; n1 = xx }; return n1
}
func (q *rbtreeset) findminidx(n1 int32) int32 {
	tree := q.tree; for { xx := tree[n1].left; if xx == 0 { break }; n1 = xx }; return n1
}
func (q *rbtreeset) nextidx(cur int32) int32 {
	last := int32(-2); tree := q.tree; rr := tree[cur].right; if rr > 0 { return q.findminidx(rr) }
	for { last, cur = cur, tree[cur].up; if cur == 0 || tree[cur].left == last { break } }; return cur
}
func (q *rbtreeset) previdx(cur int32) int32 {
	last := int32(0); tree := q.tree; ll := tree[cur].left; if ll > 0 { return q.findmaxidx(ll) }
	for { last, cur = cur, tree[cur].up; if cur == 0 || tree[cur].right == last { break } }; return cur
}
func (q *rbtreeset) rotleft(x int32) {
	tree := q.tree; y := tree[x].right; p := tree[x].up; tree[x].right = tree[y].left
	if tree[y].left != 0 { tree[tree[y].left].up = x }; tree[y].up = p
	if p == 0 {
		q.root = y
	} else if x == tree[p].left {
		tree[p].left = y
	} else {
		tree[p].right = y
	}
	tree[y].left = x; tree[x].up = y
}
func (q *rbtreeset) rotright(x int32) {
	tree := q.tree; y := tree[x].left; p := tree[x].up; tree[x].left = tree[y].right
	if tree[y].right != 0 { tree[tree[y].right].up = x }; tree[y].up = p
	if p == 0 {
		q.root = y
	} else if x == tree[p].right {
		tree[p].right = y
	} else {
		tree[p].left = y
	}
	tree[y].right = x; tree[x].up = y
}
func (q *rbtreeset) getNewNodenum() int32 {
	l := len(q.recycler); newnode := q.recycler[l-1]; q.recycler = q.recycler[:l-1]
	if l == 1 { q.tree = append(q.tree, rbtreesetnode{}); q.recycler = append(q.recycler, int32(len(q.tree)-1)) }
	return newnode
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M := gi2(); A,B := fill2(N); S := gis(M)
		rset := Newrbtreeset(func(a,b ival) bool { return a.l < b.l } )
		for i:=0;i<N;i++ { rset.Add(ival{A[i],B[i]}) }
		ans := ia(M)
		for i:=0;i<M;i++ {
			lb,ok  := rset.FindLe(ival{S[i],0})
			ub,ok2 := rset.FindGt(ival{S[i],0})
			if ok && S[i] <= lb.r {
				ans[i] = S[i]
				rset.Delete(lb)
				if lb.l < S[i] { rset.Add(ival{lb.l,S[i]-1}) }
				if lb.r > S[i] { rset.Add(ival{S[i]+1,lb.r}) }
			} else if ok && (!ok2 || S[i]-lb.r <= ub.l-S[i]) {
				ans[i] = lb.r
				rset.Delete(lb)
				if lb.l != lb.r { rset.Add(ival{lb.l,lb.r-1}) }
			} else {
				ans[i] = ub.l
				rset.Delete(ub)
				if ub.l != ub.r { rset.Add(ival{ub.l+1,ub.r}) }
			}
		}
		ansstr := vecintstring(ans)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

