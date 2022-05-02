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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }

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

type pt struct {x,y int}
type query struct {x,y,k int}
type event struct { idx, x, yl, yr, sgn int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N); Q := gi(); A,B,K := fill3(Q)
	pts     := make([]pt,N); for i:=0;i<N;i++ { pts[i] = pt{X[i]-Y[i],X[i]+Y[i]} }
	queries := make([]query,Q); for i:=0;i<Q;i++ { queries[i] = query{A[i]-B[i],A[i]+B[i],K[i]} }
	sort.Slice(pts,func(i,j int) bool { return pts[i].x < pts[j].x })
	xl := make([]int,Q); for i:=0;i<Q;i++ { xl[i] = -1 }
	xr := make([]int,Q); for i:=0;i<Q;i++ { xr[i] = 200000 }
	xmid := make([]int,Q)
	cnts := make([]int,Q)
	events := make([]event,0)
	for {
		done := true
		for i:=0;i<Q;i++ {
			if xr[i]-xl[i] > 1 { done = false }
			xmid[i] = (xr[i]+xl[i])>>1
		}
		if done { break }
		events = events[:0]
		for i:=0;i<Q;i++ {
			x1 := queries[i].x-xmid[i]-1
			x2 := queries[i].x+xmid[i]
			y1 := queries[i].y-xmid[i]
			y2 := queries[i].y+xmid[i]
			events = append(events,event{i,x1,y1,y2,-1})
			events = append(events,event{i,x2,y1,y2,1})
			cnts[i] = 0
		}
		sort.Slice(events,func(i,j int) bool { return events[i].x < events[j].x })
		ft := NewFenwick(200010)
		ptr := 0
		for _,e := range events {
			for ptr < N && pts[ptr].x <= e.x { ft.Inc(pts[ptr].y+1,1); ptr++ }
			s := ft.Prefixsum(min(200001,e.yr+1))
			if e.yl >= 1 { s -= ft.Prefixsum(min(200001,e.yl)) }
			cnts[e.idx] += e.sgn * s
		}
		for i:=0;i<Q;i++ {
			if xr[i]-xl[i] <= 1 { continue }
			if cnts[i] >= K[i] { xr[i] = xmid[i] } else { xl[i] = xmid[i] }
		}
	}
	for _,v := range xr { fmt.Fprintln(wrtr,v) }
}
