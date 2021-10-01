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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type lazysegtree struct {
	n, size, log int; op func(int, int) int; mapping func(int, int) int
	composition func(int, int) int; e int; id int; d []int; lz []int
}
func Newlazysegtree(n int, op func(int, int) int, mapping func(int, int) int, composition func(int, int) int, e int, id int) *lazysegtree {
	v := make([]int, n); for i := 0; i < n; i++ { v[i] = e }
	return NewlazysegtreeVec(v, op, mapping, composition, e, id)
}
func NewlazysegtreeVec(v []int, op func(int, int) int, mapping func(int, int) int, composition func(int, int) int, e int, id int) *lazysegtree {
	n, sz, log := len(v), 1, 0; for sz < n { sz <<= 1; log += 1 }; d := make([]int, 2*sz)
	lz := make([]int, sz); for i := 0; i < 2*sz; i++ { d[i] = e }; for i := 0; i < sz; i++ { lz[i] = id }; d[0] = e
	for i := 0; i < n; i++ { d[sz+i] = v[i]; lz[i] = id }
	st := &lazysegtree{n, sz, log, op, mapping, composition, e, id, d, lz}
	for i := sz - 1; i >= 1; i-- { st.update(i) }; return st
}
func (q *lazysegtree) Set(p int, v int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; q.d[p] = v
	for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *lazysegtree) Get(p int) int { p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; return q.d[p] }
func (q *lazysegtree) Prod(l int, r int) int {
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
func (q *lazysegtree) Allprod() int { return q.d[1] }
func (q *lazysegtree) Apply(p int, f int) {
	p += q.size; for i := q.log; i >= 1; i-- { q.push(p >> i) }; q.d[p] = q.mapping(f, q.d[p])
	for i := 1; i <= q.log; i++ { q.update(p >> i) }
}
func (q *lazysegtree) ApplyRange(l int, r int, f int) {
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
func (q *lazysegtree) update(k int) { q.d[k] = q.op(q.d[2*k], q.d[2*k+1]) }
func (q *lazysegtree) allApply(k int, f int) {
	q.d[k] = q.mapping(f, q.d[k])
	if k < q.size { 
		q.lz[k] = q.composition(f, q.lz[k])
	}
}
func (q *lazysegtree) push(k int) { if q.lz[k] != q.id { q.allApply(2*k, q.lz[k]); q.allApply(2*k+1, q.lz[k]); q.lz[k] = q.id } }

func stadd (a,b int) int { return a + b }
type query struct {i,x,y int}
type PI struct {x,y int}
type vert struct {x,y1,y2 int}
type event struct {x,y1,y2,delta int}

func solve(N int,Marr [][]int,Q int,X []int,Y[]int) []int {
	poly := make([][]PI,N)
	for i:=0;i<N;i++ {
		m := Marr[i]; lm := len(m)/2
		poly[i] = make([]PI,lm)
		for j:=0;j<lm;j++ { poly[i][j] = PI{m[2*j],m[2*j+1]} }
	}
	quer := make([]query,Q)
	for i:=0;i<Q;i++ { x,y := X[i],Y[i]; quer[i] = query{i,x,y} }
	sort.Slice(quer,func(i,j int)bool{return quer[i].x < quer[j].x })
	ans := iai(Q,0)

	// First pass, classify each vertical edge to see if it adds or subtracts 1 to the running polygon count
	events := make([]event,0)
	lst := Newlazysegtree(100_010,stadd,stadd,stadd,0,0) // Can reuse this between polygons, since will end with zeros
	for i:=0;i<N;i++ {
		pp := poly[i]; v := make([]vert,0)
		for j:=0;j<len(pp);j+=2 { v = append(v, vert{poly[i][j].x,poly[i][j].y,poly[i][j+1].y}) }
		sort.Slice(v,func(i,j int)bool{return v[i].x < v[j].x})
		for _,vv := range v {
			y1,y2 := min(vv.y1,vv.y2),max(vv.y1,vv.y2)-1
			n := lst.Get(y1)
			if n == 0 {
				lst.ApplyRange(y1,y2,1); events = append(events,event{vv.x,y1,y2,1})
			} else {
				lst.ApplyRange(y1,y2,-1); events = append(events,event{vv.x,y1,y2,-1})
			}
		}
	}
	// Second pass now to loop through all of the vertical edges
	sort.Slice(events,func(i,j int)bool{return events[i].x < events[j].x})
	qptr := 0
	for _,ee := range events {
		for qptr < Q && quer[qptr].x < ee.x {
			q := quer[qptr]
			ans[q.i] = lst.Get(q.y)
			qptr++
		}
		lst.ApplyRange(ee.y1,ee.y2,ee.delta)
	}
	return ans
}

func main() {
	//test(2,1,10_000,1,10_000,0,100_000)
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); Marr := make([][]int,N) 
	for i:=0;i<N;i++ { M := gi(); Marr[i] = gis(2*M) }
	Q := gi(); X,Y := fill2(Q)
	ans := solve(N,Marr,Q,X,Y)
	for _,a := range ans { fmt.Fprintln(wrtr,a)	}
}
