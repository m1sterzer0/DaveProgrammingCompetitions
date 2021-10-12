package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
type rbtreemultisetnode struct { left, right, up int32; red bool; key int; count int }
type rbtreemultiset struct {
	lessthan func(a, b int) bool; tree []rbtreemultisetnode; root int32; recycler []int32; sz int; minidx int32
	maxidx int32
}
type rbtreemultisetIterator interface { Next() (ok bool); Prev() (ok bool); Key() int; Count() int }
type rbtreemultisetiter struct { cur int32; key int; count int; rbtree *rbtreemultiset }
func (i *rbtreemultisetiter) Key() int { return i.key }
func (i *rbtreemultisetiter) Count() int { return i.count }
func (i *rbtreemultisetiter) Next() bool {
	rbtree := i.rbtree; v := rbtree.nextidx(i.cur); if v == 0 { return false }
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count; return true
}
func (i *rbtreemultisetiter) Prev() bool {
	rbtree := i.rbtree; v := rbtree.previdx(i.cur); if v == 0 { return false }
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count; return true
}
func Newrbtreemultiset(lessthan func(a, b int) bool) *rbtreemultiset {
	q := &rbtreemultiset{lessthan, make([]rbtreemultisetnode, 2), int32(0), make([]int32, 0), 0, 0, 0}
	q.tree[0].left, q.tree[0].right, q.tree[0].up, q.tree[0].red = 0, 0, 0, false; q.recycler = append(q.recycler, 1)
	return q
}
func (q *rbtreemultiset) Add(k int) {
	if q.sz == 0 {
		z := q.getNewNodenum(); tree := q.tree; q.minidx, q.maxidx, q.sz, q.root = z, z, q.sz+1, z
		tree[z].key, tree[z].count, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 1, 0, 0, 0, false; return
	}
	y, cmp := q.findInsertionPoint(k); if cmp == 0 { q.tree[y].count++; q.sz += 1; return }; z := q.getNewNodenum()
	q.sz += 1; tree := q.tree
	tree[z].key, tree[z].count, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 1, y, 0, 0, true
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
func (q *rbtreemultiset) Delete(k int) bool {
	if q.sz == 0 { return false }; z, cmp := q.findInsertionPoint(k)
	if cmp != 0 { return false } else if q.tree[z].count > 1 { q.tree[z].count--; q.sz--; return true }; q.sz--
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
func (q *rbtreemultiset) Clear() {
	q.tree, q.root, q.recycler, q.sz = q.tree[:2], 0, q.recycler[:0], 0; q.recycler = append(q.recycler, int32(1))
}
func (q *rbtreemultiset) IsEmpty() bool { return q.sz == 0 }
func (q *rbtreemultiset) Contains(k int) bool { _, cmp := q.findInsertionPoint(k); return cmp == 0 }
func (q *rbtreemultiset) Count(k int) int {
	z, cmp := q.findInsertionPoint(k); if cmp != 0 { return 0 }; return q.tree[z].count
}
func (q *rbtreemultiset) Len() int { return q.sz }
func (q *rbtreemultiset) MinKey() (k int) {
	if q.sz == 0 { panic("Called MinKey on an empty rbtreemultiset") }; return q.tree[q.minidx].key
}
func (q *rbtreemultiset) MaxKey() (k int) {
	if q.sz == 0 { panic("Called MaxKey on an empty rbtreemultiset") }; return q.tree[q.maxidx].key
}
func (q *rbtreemultiset) findLtIdx(k int) (int32, bool) {
	if q.sz == 0 || !q.lessthan(q.tree[q.minidx].key, k) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos != 1 { idx = q.previdx(idx) }; return idx, true
}
func (q *rbtreemultiset) findLeIdx(k int) (int32, bool) {
	if q.sz == 0 || q.lessthan(k, q.tree[q.minidx].key) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos == -1 { idx = q.previdx(idx) }; return idx, true
}
func (q *rbtreemultiset) findGtIdx(k int) (int32, bool) {
	if q.sz == 0 || !q.lessthan(k, q.tree[q.maxidx].key) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos != -1 { idx = q.nextidx(idx) }; return idx, true
}
func (q *rbtreemultiset) findGeIdx(k int) (int32, bool) {
	if q.sz == 0 || q.lessthan(q.tree[q.maxidx].key, k) { return 0, false }; idx, pos := q.findInsertionPoint(k)
	if pos == 1 { idx = q.nextidx(idx) }; return idx, true
}
func (q *rbtreemultiset) FindLt(k int) (int, bool) {
	var ans int; idx, ok := q.findLtIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreemultiset) FindLe(k int) (int, bool) {
	var ans int; idx, ok := q.findLeIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreemultiset) FindGt(k int) (int, bool) {
	var ans int; idx, ok := q.findGtIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreemultiset) FindGe(k int) (int, bool) {
	var ans int; idx, ok := q.findGeIdx(k); if ok { ans = q.tree[idx].key }; return ans, ok
}
func (q *rbtreemultiset) FindLtIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter; idx, ok := q.findLtIdx(k)
	if ok { ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q} }; return ans, ok
}
func (q *rbtreemultiset) FindLeIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter; idx, ok := q.findLeIdx(k)
	if ok { ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q} }; return ans, ok
}
func (q *rbtreemultiset) FindGtIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter; idx, ok := q.findGtIdx(k)
	if ok { ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q} }; return ans, ok
}
func (q *rbtreemultiset) FindGeIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter; idx, ok := q.findGeIdx(k)
	if ok { ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q} }; return ans, ok
}
func (q *rbtreemultiset) FindIter(k int) (rbtreemultisetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx, pos := q.findInsertionPoint(k); if pos != 0 { return nil, false }
	return &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *rbtreemultiset) MinIter() (rbtreemultisetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx := q.findminidx(q.root)
	return &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *rbtreemultiset) MaxIter() (rbtreemultisetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx := q.findmaxidx(q.root)
	return &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *rbtreemultiset) rbTransplant(u, v int32) {
	tree := q.tree
	if tree[u].up == 0 {
		q.root = v
	} else {
		p := tree[u].up; if u == tree[p].left { tree[p].left = v } else { tree[p].right = v }
	}
	tree[v].up = tree[u].up
}
func (q *rbtreemultiset) findInsertionPoint(k int) (int32, int8) {
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
func (q *rbtreemultiset) findmaxidx(n1 int32) int32 {
	tree := q.tree; for { xx := tree[n1].right; if xx == 0 { break }; n1 = xx }; return n1
}
func (q *rbtreemultiset) findminidx(n1 int32) int32 {
	tree := q.tree; for { xx := tree[n1].left; if xx == 0 { break }; n1 = xx }; return n1
}
func (q *rbtreemultiset) nextidx(cur int32) int32 {
	last := int32(-2); tree := q.tree; rr := tree[cur].right; if rr > 0 { return q.findminidx(rr) }
	for { last, cur = cur, tree[cur].up; if cur == 0 || tree[cur].left == last { break } }; return cur
}
func (q *rbtreemultiset) previdx(cur int32) int32 {
	last := int32(0); tree := q.tree; ll := tree[cur].left; if ll > 0 { return q.findmaxidx(ll) }
	for { last, cur = cur, tree[cur].up; if cur == 0 || tree[cur].right == last { break } }; return cur
}
func (q *rbtreemultiset) rotleft(x int32) {
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
func (q *rbtreemultiset) rotright(x int32) {
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
func (q *rbtreemultiset) getNewNodenum() int32 {
	l := len(q.recycler); newnode := q.recycler[l-1]; q.recycler = q.recycler[:l-1]
	if l == 1 { q.tree = append(q.tree, rbtreemultisetnode{}); q.recycler = append(q.recycler, int32(len(q.tree)-1)) }
	return newnode
}

type bestn struct {k int; rt,lf *rbtreemultiset; rtsum int}
func Newbestn (k int) *bestn {
	lf := Newrbtreemultiset(func(a,b int) bool { return a < b })
	rt := Newrbtreemultiset(func(a,b int) bool { return a < b })
	return &bestn{k,rt,lf,0}
}
func (q *bestn) Add(n int) {
	if q.rt.Len() < q.k { q.rtsum += n; q.rt.Add(n); return }
	if n <= q.rt.MinKey() { q.lf.Add(n); return }
	q.rtsum += n; q.rt.Add(n); mk := q.rt.MinKey(); q.rtsum -= mk; q.rt.Delete(mk); q.lf.Add(mk)
}
func (q *bestn) Delete(n int) {
	if q.lf.Len() > 0 && n <= q.lf.MaxKey() { q.lf.Delete(n); return }
	q.rtsum -= n; q.rt.Delete(n)
	if q.lf.Len() > 0 { mk := q.lf.MaxKey(); q.rtsum += mk; q.lf.Delete(mk); q.rt.Add(mk) }
}
func (q *bestn) Query() int { return q.rtsum }
type event struct {t, d, n int}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		D,N,K := gi3()
		H,S,E := fill3(N)
		ee := []event{}
		for i:=0;i<N;i++ {
			ee = append(ee,event{1,S[i],H[i]})
			ee = append(ee,event{2,E[i]+1,H[i]})
		}
		sort.Slice(ee,func(i,j int) bool { return ee[i].d < ee[j].d})
		eeptr := 0
		bb := Newbestn(K)
		best := 0
		for i:=1;i<=D;i++ {
			for eeptr < 2*N && ee[eeptr].d == i {
				if ee[eeptr].t == 1 { bb.Add(ee[eeptr].n) } else { bb.Delete(ee[eeptr].n) }
				eeptr++
			}
			best = max(best,bb.Query())
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}




