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
func sortUniqueIntarr(a []int) []int {
	sort.Slice(a,func (i,j int) bool { return a[i] < a[j] })
	i,j,la := 0,0,len(a)
	for ;i<la;i++ { if i == 0 || a[i] != a[i-1] { a[j] = a[i]; j++ } }
	return a[:j]
}
type rect struct { w,l int }
type Fenwick struct { n, tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx) }; q.tot += val }
func (q *Fenwick) Dec(idx int, val int) { q.Inc(idx, -val) }
func (q *Fenwick) IncDec(left int, right int, val int) { q.Inc(left, val); q.Dec(right, val) }
func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans
}
func (q *Fenwick) Suffixsum(idx int) int { return q.tot - q.Prefixsum(idx-1) }
func (q *Fenwick) Rangesum(left int, right int) int {
	if right < left { return 0 }; return q.Prefixsum(right) - q.Prefixsum(left-1)
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Too big for a 2D segment tree
	// Need to use time for one of the dimensions
	// Also need coordinate compression to fit
	N,M := gi(),gi()
	A := gis(N); B := gis(N); C := gis(M); D := gis(M)
	xx := make([]int,0); xx = append(xx,0)
	for i:=0;i<N;i++ { xx = append(xx,A[i]); xx = append(xx,B[i]) }
	for i:=0;i<M;i++ { xx = append(xx,C[i]); xx = append(xx,D[i]) }
	xx = sortUniqueIntarr(xx)
	lookup := make(map[int]int); for i,x := range xx { lookup[x] = i }
	boxes := make([]rect,M); choc := make([]rect,N)
	for i:=0;i<M;i++ { c,d := C[i],D[i]; boxes[i] = rect{lookup[c],lookup[d]} }
	for i:=0;i<N;i++ { a,b := A[i],B[i]; choc[i]  = rect{lookup[a],lookup[b]} }
	sort.Slice(boxes,func(i,j int) bool { return boxes[i].w > boxes[j].w })
	sort.Slice(choc, func(i,j int) bool { return choc[i].w > choc[j].w })
	good := true; bptr := 0; ft := NewFenwick(len(xx)+10)
	for _,c := range choc {
		for bptr < M && boxes[bptr].w >= c.w { ft.Inc(boxes[bptr].l,1); bptr++ }
		s := ft.Suffixsum(c.l); if s == 0 { good = false; break }
		l,r := c.l,len(xx)+2; for r-l > 1 { m := (l+r)>>1; if ft.Suffixsum(m) == s { l = m } else { r = m } }
		ft.Inc(r-1,-1)
	}
	ans := "No"; if good { ans = "Yes" }; fmt.Println(ans)
}

