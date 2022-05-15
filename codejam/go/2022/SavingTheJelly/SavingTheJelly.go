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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

type ss struct {i,j int}
type PI struct{ x, y int }
type hopcroftKarpQueue struct { buf []int; head, tail, sz, bm, l int }
func NewhopcroftKarpQueue() *hopcroftKarpQueue { buf := make([]int, 8); return &hopcroftKarpQueue{buf, 0, 0, 8, 7, 0} }
func (q *hopcroftKarpQueue) IsEmpty() bool { return q.l == 0 }
func (q *hopcroftKarpQueue) Clear() { q.head = 0; q.tail = 0; q.l = 0 }
func (q *hopcroftKarpQueue) Len() int { return q.l }
func (q *hopcroftKarpQueue) Push(x int) {
	if q.l == q.sz { q.sizeup() }; if q.l > 0 { q.head = (q.head - 1) & q.bm }; q.l++; q.buf[q.head] = x
}
func (q *hopcroftKarpQueue) Pop() int {
	if q.l == 0 { panic("Empty hopcroftKarpQueue Pop()") }; v := q.buf[q.tail]; q.l--
	if q.l > 0 { q.tail = (q.tail - 1) & q.bm } else { q.Clear() }; return v
}
func (q *hopcroftKarpQueue) Head() int { if q.l == 0 { panic("Empty hopcroftKarpQueue Head()") }; return q.buf[q.head] }
func (q *hopcroftKarpQueue) Tail() int { if q.l == 0 { panic("Empty hopcroftKarpQueue Tail()") }; return q.buf[q.tail] }
func (q *hopcroftKarpQueue) sizeup() {
	buf := make([]int, 2*q.sz); for i := 0; i < q.l; i++ { buf[i] = q.buf[(q.head+i)&q.bm] }; q.buf = buf; q.head = 0
	q.tail = q.sz - 1; q.sz = 2 * q.sz; q.bm = q.sz - 1
}
func HopcroftKarp(N1, N2 int, adj [][]int) []PI {
	mynil := N1 + N2; pairu := make([]int, N1); pairv := make([]int, N2); dist := make([]int, N1+N2+1)
	myinf := 1000000000000000000; q := NewhopcroftKarpQueue()
	bfs := func() bool {
		for u := 0; u < N1; u++ { if pairu[u] == mynil { dist[u] = 0; q.Push(u) } else { dist[u] = myinf } }
		dist[mynil] = myinf
		for !q.IsEmpty() {
			u := q.Pop()
			if u != mynil && dist[u] < dist[mynil] {
				for _, v := range adj[u] { u2 := pairv[v]; if dist[u2] == myinf { dist[u2] = dist[u] + 1; q.Push(u2) } }
			}
		}
		return dist[mynil] != myinf
	}
	var dfs func(int) bool
	dfs = func(u int) bool {
		if u == mynil { return true }
		for _, v := range adj[u] {
			u2 := pairv[v]; if dist[u2] == dist[u]+1 && dfs(u2) { pairv[v], pairu[u] = u, v; return true }
		}
		dist[u] = myinf; return false
	}
	for i := 0; i < N1; i++ { pairu[i] = mynil }; for i := 0; i < N2; i++ { pairv[i] = mynil }
	for bfs() { for u := 0; u < N1; u++ { if pairu[u] == mynil { dfs(u) } } }; res := make([]PI, 0)
	for u := 0; u < N1; u++ { if pairu[u] != mynil { res = append(res, PI{u, pairu[u]}) } }; return res
}

func dist(x1,y1,x2,y2 int) int { return (x2-x1)*(x2-x1)+(y2-y1)*(y2-y1) }
type edge struct {n,d int }
func solve(N int, X,Y,CX,CY []int) []ss {
	gr := make([][]int,N)
	gr2 := make([][]edge,N)
	for i:=0;i<N;i++ {
		dd := dist(X[i],Y[i],CX[0],CY[0])
		for j:=0;j<N;j++ {
			d := dist(X[i],Y[i],CX[1+j],CY[1+j])
			if d <= dd { 
				gr[i] = append(gr[i],j)
				gr2[i] = append(gr2[i],edge{j,d})
			}
		}
	}
	pairs := HopcroftKarp(N,N,gr)
	matches    := make([]int,N); for _,p := range pairs {    matches[p.x] = p.y }
	revmatches := make([]int,N); for _,p := range pairs { revmatches[p.y] = p.x }
	if len(pairs) != N { return []ss{} }
	for k:=0;k<N;k++ { a := gr2[k]; sort.Slice(a,func(i,j int) bool { return a[i].d < a[j].d }) }
	used := make([]bool,N)
	usedc := make([]bool,N)
	order := make([]ss,0,N)
	ptr := ia(N)

	rejigger := func(n,c int) {
		used := make([]bool,N)
		usedc := make([]bool,N)
		used[n] = true
		usedc[c] = true
		cyc := []int{n,c}
		for {
			n = revmatches[c];     cyc = append(cyc,n); if used[n]  { break }; used[n] = true
			c = gr2[n][ptr[n]].n;  cyc = append(cyc,c); if usedc[c] { break }; usedc[c] = true
		}
		idx := len(cyc)-1; targ := cyc[len(cyc)-1]
		if len(cyc) % 2 == 0 {
			for {
				n,c := cyc[idx-1],cyc[idx]
				matches[n] = c; revmatches[c] = n
				idx -= 2
				if cyc[idx] == targ { break }
			}
		} else {
			for {
				idx -= 2
				n,c = cyc[idx],cyc[idx+1]
				matches[n] = c; revmatches[c] = n
				if cyc[idx] == targ { break }
			}
		}
	}
	
	for len(order) < N {
		found := false
		for i:=0;i<N;i++ {
			if used[i] { continue }
			for usedc[gr2[i][ptr[i]].n] { ptr[i]++ }
			if matches[i] == gr2[i][ptr[i]].n {
				found = true
				used[i] = true
				usedc[matches[i]] = true
				order = append(order,ss{i+1,matches[i]+2})
			}
		}
		if found { continue }
		// OK, no one is matched to their desired answer, so we need to rejigger some matches
		for i:=0;i<N;i++ {
			if used[i] { continue }
			rejigger(i,gr2[i][ptr[i]].n)
			break
		}
	}
	return order
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
 		N := gi(); X,Y := fill2(N); CX,CY := fill2(N+1)
		//ansarr := solveSmall(N,X,Y,CX,CY)
		ansarr := solve(N,X,Y,CX,CY)
		if len(ansarr) == N {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"POSSIBLE")
			for _,s := range ansarr {
				fmt.Fprintf(wrtr,"%v %v\n",s.i,s.j)
			}
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		}
	}
}

