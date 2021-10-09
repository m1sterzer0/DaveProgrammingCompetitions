package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//import "runtime/pprof"

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
const BUFSIZE = 10000000
var rdr = newScanner(bufio.NewReaderSize(os.Stdin,BUFSIZE))
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }

func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

type msetnode struct { left, right, up int32; red bool; key int; count int }
type mset struct {
	lessthan func(a, b int) bool; tree []msetnode; root int32; recycler []int32; sz int; minidx int32
	maxidx int32
}
type msetIterator interface { Next() (ok bool); Prev() (ok bool); Key() int; Count() int }
type msetiter struct { cur int32; key int; count int; rbtree *mset }
func (i *msetiter) Key() int { return i.key }
func (i *msetiter) Count() int { return i.count }
func (i *msetiter) Next() bool {
	rbtree := i.rbtree; v := rbtree.nextidx(i.cur); if v == 0 { return false }
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count; return true
}
func (i *msetiter) Prev() bool {
	rbtree := i.rbtree; v := rbtree.previdx(i.cur); if v == 0 { return false }
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count; return true
}
func Newmset(lessthan func(a, b int) bool) *mset {
	q := &mset{lessthan, make([]msetnode, 2), int32(0), make([]int32, 0), 0, 0, 0}
	q.tree[0].left, q.tree[0].right, q.tree[0].up, q.tree[0].red = 0, 0, 0, false; q.recycler = append(q.recycler, 1)
	return q
}
func (q *mset) Add(k int) {
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
func (q *mset) Delete(k int) bool {
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
func (q *mset) Clear() {
	q.tree, q.root, q.recycler, q.sz = q.tree[:2], 0, q.recycler[:0], 0; q.recycler = append(q.recycler, int32(1))
}
func (q *mset) IsEmpty() bool { return q.sz == 0 }
func (q *mset) Contains(k int) bool { _, cmp := q.findInsertionPoint(k); return cmp == 0 }
func (q *mset) Count(k int) int {
	z, cmp := q.findInsertionPoint(k); if cmp != 0 { return 0 }; return q.tree[z].count
}
func (q *mset) Len() int { return q.sz }
func (q *mset) MinKey() (k int) {
	if q.sz == 0 { panic("Called MinKey on an empty mset") }; return q.tree[q.minidx].key
}
func (q *mset) MaxKey() (k int) {
	if q.sz == 0 { panic("Called MaxKey on an empty mset") }; return q.tree[q.maxidx].key
}
func (q *mset) LowerBound(k int) (int, bool) {
	var def int; if q.sz == 0 { return def, false }; idx, pos := q.findInsertionPoint(k)
	if pos == 1 { idx = q.nextidx(idx) }; if idx <= 0 { return def, false }; return q.tree[idx].key, true
}
func (q *mset) UpperBound(k int) (int, bool) {
	var def int; if q.sz == 0 { return def, false }; idx, pos := q.findInsertionPoint(k)
	if pos != -1 { idx = q.nextidx(idx) }; if idx <= 0 { return def, false }; return q.tree[idx].key, true
}
func (q *mset) LowerBoundIter(k int) (msetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx, pos := q.findInsertionPoint(k); if pos == 1 { idx = q.nextidx(idx) }
	if idx <= 0 { return nil, false }; return &msetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *mset) UpperBoundIter(k int) (msetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx, pos := q.findInsertionPoint(k); if pos != -1 { idx = q.nextidx(idx) }
	if idx <= 0 { return nil, false }; return &msetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *mset) FindIter(k int) (msetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx, pos := q.findInsertionPoint(k); if pos != 0 { return nil, false }
	return &msetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *mset) MinIter() (msetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx := q.findminidx(q.root)
	return &msetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *mset) MaxIter() (msetIterator, bool) {
	if q.sz == 0 { return nil, false }; idx := q.findmaxidx(q.root)
	return &msetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *mset) rbTransplant(u, v int32) {
	tree := q.tree
	if tree[u].up == 0 {
		q.root = v
	} else {
		p := tree[u].up; if u == tree[p].left { tree[p].left = v } else { tree[p].right = v }
	}
	tree[v].up = tree[u].up
}
func (q *mset) findInsertionPoint(k int) (int32, int8) {
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
func (q *mset) findmaxidx(n1 int32) int32 {
	tree := q.tree; for { xx := tree[n1].right; if xx == 0 { break }; n1 = xx }; return n1
}
func (q *mset) findminidx(n1 int32) int32 {
	tree := q.tree; for { xx := tree[n1].left; if xx == 0 { break }; n1 = xx }; return n1
}
func (q *mset) nextidx(cur int32) int32 {
	last := int32(-2); tree := q.tree; rr := tree[cur].right; if rr > 0 { return q.findminidx(rr) }
	for { last, cur = cur, tree[cur].up; if cur == 0 || tree[cur].left == last { break } }; return cur
}
func (q *mset) previdx(cur int32) int32 {
	last := int32(0); tree := q.tree; ll := tree[cur].left; if ll > 0 { return q.findmaxidx(ll) }
	for { last, cur = cur, tree[cur].up; if cur == 0 || tree[cur].right == last { break } }; return cur
}
func (q *mset) rotleft(x int32) {
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
func (q *mset) rotright(x int32) {
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
func (q *mset) getNewNodenum() int32 {
	l := len(q.recycler); newnode := q.recycler[l-1]; q.recycler = q.recycler[:l-1]
	if l == 1 { q.tree = append(q.tree, msetnode{}); q.recycler = append(q.recycler, int32(len(q.tree)-1)) }
	return newnode
}

func solve(N,Q int, A,B,C,D []int) []int {
	schools := make([]*mset,200_001)
	lt := func (a,b int) bool {return a < b }
	for i:=1;i<=200_000;i++ { schools[i] = Newmset(lt) }
	master := Newmset(lt)
	kid2school := make([]int,200_001)
	for i:=0;i<N;i++ { schools[B[i]].Add(A[i]); kid2school[i+1] = B[i] }
	for i:=1;i<=200_000;i++ { 
		if schools[i].Len() > 0 { 
			v := schools[i].MaxKey()
			master.Add(v)
		}
	}
	ansarr := make([]int,Q)
	for i:=0;i<Q;i++ {
		// Take the kid away
		kid := C[i]; rating := A[kid-1]; oldschoolid := kid2school[kid]; oldschool := schools[oldschoolid]; newschoolid := D[i]; newschool := schools[newschoolid]
		//fmt.Fprintf(wrtr,"DBG: i:%v kid:%v rating:%v oldschool:%v newschool:%v\n",i,kid,rating,oldschoolid,newschoolid)
		oldmax := oldschool.MaxKey() 
		oldschool.Delete(rating)
		newmax := -1; if oldschool.Len() > 0 { newmax = oldschool.MaxKey() }
		if oldmax != newmax { 
			master.Delete(oldmax)
			if newmax >= 0 { master.Add(newmax)}
		}
		//fmt.Fprintf(wrtr,"    oldschool oldmax:%v newmax:%v\n",oldmax,newmax)

		oldmax = -1; if newschool.Len() > 0 { oldmax = newschool.MaxKey() }
		newschool.Add(rating)
		newmax = newschool.MaxKey()
		if oldmax != newmax { 
			if oldmax >= 0 { master.Delete(oldmax) }
			master.Add(newmax)
		}

		//fmt.Fprintf(wrtr,"    newschool oldmax:%v newmax:%v\n",oldmax,newmax)
		ansarr[i] = master.MinKey()
		kid2school[kid] = D[i]
	}
	return ansarr
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile();
	//debug.SetGCPercent(-1)
	//rand.Seed(8675309)
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,Q := gi(),gi()
	A,B := fill2(N)
	C,D := fill2(Q)
	res := solve(N,Q,A,B,C,D)
	ansstr := make([]string,len(res))
	for i:=0;i<len(res);i++ { ansstr[i] = strconv.Itoa(res[i]) }
	final := strings.Join(ansstr,"\n")
	fmt.Fprintln(wrtr,final)
}


