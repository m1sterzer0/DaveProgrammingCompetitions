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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

type Dsu struct { n int; parentOrSize []int }
func NewDsu(n int) *Dsu { buf := make([]int, n); for i := 0; i < n; i++ { buf[i] = -1 }; return &Dsu{n, buf} }
func (q *Dsu) Leader(a int) int {
	if q.parentOrSize[a] < 0 { return a }; ans := q.Leader(q.parentOrSize[a]); q.parentOrSize[a] = ans; return ans
}
func (q *Dsu) Merge(a int, b int) int {
	x := q.Leader(a); y := q.Leader(b); if x == y { return x }; if q.parentOrSize[y] < q.parentOrSize[x] { x, y = y, x }
	q.parentOrSize[x] += q.parentOrSize[y]; q.parentOrSize[y] = x; return x
}
func (q *Dsu) Same(a int, b int) bool { return q.Leader(a) == q.Leader(b) }
func (q *Dsu) Size(a int) int { l := q.Leader(a); return -q.parentOrSize[l] }
func (q *Dsu) Groups() [][]int {
	numgroups := 0; leader2idx := make([]int, q.n); for i := 0; i <= q.n; i++ { leader2idx[i] = -1 }
	ans := make([][]int, 0)
	for i := int(0); i <= int(q.n); i++ {
		l := q.Leader(i)
		if leader2idx[l] == -1 { ans = append(ans, make([]int, 0)); leader2idx[l] = numgroups; numgroups += 1 }
		ans[leader2idx[l]] = append(ans[leader2idx[l]], i)
	}
	return ans
}

type dsnode struct { n,f,l,d int}
type minheap struct { buf []dsnode; less func(dsnode, dsnode) bool }
func Newminheap(f func(dsnode, dsnode) bool) *minheap { buf := make([]dsnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v dsnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() dsnode { return q.buf[0] }
func (q *minheap) Pop() dsnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []dsnode) {
	q.buf = append(q.buf, pri...); n := len(q.buf); for i := n/2 - 1; i >= 0; i-- { q.siftup(i) }
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1; p := q.buf[ppos]; if !q.less(newitem, p) { break }; q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1; if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) { chpos = rtpos }
		q.buf[pos], pos = q.buf[chpos], chpos; chpos = 2*pos + 1
	}
	q.buf[pos] = newitem; q.siftdown(startpos, pos)
}

const inf = 2000000000000000000

type edge struct  { n2,d int  }
type edge2 struct { a,b,c int }
type query struct { idx,x,y,t int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Main idea
	// a) Reduce graph of N nodes down to a graph with just the houses
	// b) Sort the queries based on ti
	// c) Sort the edges based on length
	// For each query -- add the edges up to ti, and then see if source and dest are in same connected component
	//
	// The hard part is step a.  The realization is that
	// -- A spanning tree is as good as full network to determine iff I can get from pt A to pt B using only edges of length <= 
	// -- To find the spanning tree, Prim and Kruskal are too long as they need to find the complete K graph, but Boruvka only
	//    requires us to find the minimum distance between two connected components.
	// -- We can start a search at all of the houses in each connected component, and we only propagate through each node twice per round like abc245_G.

	N,M,K := gi(),gi(),gi(); A,B,C := fill3(M); Q := gi(); X,Y,T := fill3(Q)
	for i:=0;i<M;i++ { A[i]--; B[i]-- }; for i:=0;i<Q;i++ { X[i]--; Y[i]-- }
	elist := make([]edge2,0); uf := NewDsu(K)
	gr := make([][]edge,N)
	for i:=0;i<M;i++ { a,b,c := A[i],B[i],C[i]; gr[a] = append(gr[a],edge{b,c}); gr[b] = append(gr[b],edge{a,c}) }
	mh := Newminheap(func(a,b dsnode) bool { return a.d < b.d })
	for len(elist) < K-1 {
		d1arr := iai(N,inf); d2arr := iai(N,inf)
		f1arr := iai(N,-1);  f2arr := iai(N,-1)
		l1arr := iai(N,-1);	 l2arr := iai(N,-1)
		ldist := iai(N,inf); candlist := make([]edge2,0)

		for k:=0;k<K;k++ { mh.Push(dsnode{k,k,uf.Leader(k),0}) }
		for !mh.IsEmpty() {
			xx := mh.Pop(); n,f,l,d := xx.n,xx.f,xx.l,xx.d
			if l == l1arr[n] || l == l2arr[n] || l2arr[n] >= 0 { continue }
			if l1arr[n] == -1 { 
				l1arr[n] = l; d1arr[n] = d; f1arr[n] = f
			} else { 
				l2arr[n] = l; d2arr[n] = d; f2arr[n] = f
				if n < K {
					n1,n2,l := n,f,l1arr[n];
					if ldist[l] == inf { candlist = append(candlist,edge2{n1,n2,d}); ldist[l] = d }
				}
			}
			for _,ee := range gr[n] { mh.Push(dsnode{ee.n2,f,l,d+ee.d}) }
		}
		sort.Slice(candlist,func(i,j int) bool { return candlist[i].c < candlist[j].c} )
		for _,ee := range candlist {
			if uf.Leader(ee.a) != uf.Leader(ee.b) { elist = append(elist,ee); uf.Merge(ee.a,ee.b) }
		}
	}
	qq := make([]query,Q)
	ansarr := make([]string,Q)
	for i:=0;i<Q;i++ { qq[i] = query{i,X[i],Y[i],T[i]} }
	sort.Slice(qq,func (i,j int) bool { return qq[i].t < qq[j].t})
	sort.Slice(elist,func (i,j int) bool { return elist[i].c < elist[j].c})
	uf2 := NewDsu(K); eidx := 0
	for _,q := range qq {
		for eidx < K-1 && elist[eidx].c <= q.t { uf2.Merge(elist[eidx].a,elist[eidx].b); eidx++ }
		ans := "No"; if uf2.Same(q.x,q.y) { ans = "Yes" };  ansarr[q.idx] = ans
	}
	for _,s := range ansarr { fmt.Fprintln(wrtr,s) }
}

