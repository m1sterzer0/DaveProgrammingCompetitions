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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type dnode struct {s,n,d int }
type node struct {s,n int }
type minheap struct { buf []dnode; less func(dnode, dnode) bool }
func Newminheap(f func(dnode, dnode) bool) *minheap { buf := make([]dnode, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool { return len(q.buf) == 0 }
func (q *minheap) Clear() { q.buf = q.buf[:0] }
func (q *minheap) Len() int { return len(q.buf) }
func (q *minheap) Push(v dnode) { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() dnode { return q.buf[0] }
func (q *minheap) Pop() dnode {
	v1 := q.buf[0]; l := len(q.buf)
	if l == 1 { q.buf = q.buf[:0] } else { l--; q.buf[0] = q.buf[l]; q.buf = q.buf[:l]; q.siftup(0) }; return v1
}
func (q *minheap) Heapify(pri []dnode) {
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,S,R := gi4()
		U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
		stones := make([][]int,N)
		for i:=0;i<N;i++ {
			k := gi()
			for j:=0;j<k;j++ { stones[i] = append(stones[i],gi()) }
		}
		reci,recr := make([][]int,R),ia(R)
		for i:=0;i<R;i++ {
			k := gi()
			for j:=0;j<k;j++ { reci[i] = append(reci[i],gi()) }
			recr[i] = gi()
		}

		gr := make([][]int,N)
		for i:=0;i<M;i++ {
			u,v := U[i],V[i]
			gr[u] = append(gr[u],v)
			gr[v] = append(gr[v],u)
		}

		// To save time, we need a list of recipes for each stone
		s2r := make([][]int,S+1)
		for i:=1;i<=S;i++ {
			for j:=0;j<R;j++ {
				found := false
				for _,k := range reci[j] { if i == k { found = true } }
				if found { s2r[i] = append(s2r[i],j) }
			}
		}

		// The nodes here are represented by stone,node pairs
		dist := make(map[node]int)
		mh := Newminheap(func (a,b dnode) bool { return a.d < b.d })
		for i:=0;i<N;i++ {
			for _,s := range stones[i] { mh.Push(dnode{s,i,0}) }
		}
		
		ans := 0
		for !mh.IsEmpty() {
			dd := mh.Pop()
			s,n,d := dd.s,dd.n,dd.d
			_,ok := dist[node{s,n}]; if ok { continue }
			dist[node{s,n}] = d
			if s == 1 { ans = d; break }
			if d >= 1000000000000 { ans = -1; break }
			for _,n2 := range gr[n] { mh.Push(dnode{s,n2,d+1}) }
			for _,rid := range s2r[s] {
				targ := recr[rid]
				_,ok := dist[node{targ,n}]; if ok { continue }
				good,cost := true,0
				for _,ing := range reci[rid] {
					v,ok := dist[node{ing,n}]
					if !ok { good = false; break }
					cost += v
				}
				if good { mh.Push(dnode{targ,n,cost}) }
			}
		}
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

