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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
	N,M,K := gi(),gi(),gi(); A := gis(N)
	s1 := Newrbtreemultiset(func(a,b int) bool { return a < b })
	s2 := Newrbtreemultiset(func(a,b int) bool { return a < b })
	sums1 := 0

	doAdd := func(x int) {
		if s1.Len() < K { 
			sums1 += x; s1.Add(x)
		} else {
			k := s1.MaxKey()
			if x < k { sums1 -= k; s1.Delete(k); s2.Add(k); s1.Add(x); sums1 += x } else { s2.Add(x) }
		}
	}

	doRemove := func(x int) {
		if s2.Contains(x) {
			s2.Delete(x)
		} else {
			sums1 -= x; s1.Delete(x)
			if s2.Len() > 0 { y := s2.MinKey(); s2.Delete(y); s1.Add(y); sums1 += y }
		}
	}

	ansarr := ia(0)
	for i:=0;i<N-M+1;i++ {
		if i==0 {
			for j:=0;j<M-1;j++ { doAdd(A[j]) }
		} else {
			doRemove(A[i-1])
		}
		doAdd(A[i+M-1])
		ansarr = append(ansarr,sums1)
	}
	ans := vecintstring(ansarr)
	fmt.Println(ans)
}
