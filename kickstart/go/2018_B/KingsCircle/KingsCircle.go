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

type pt struct { x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE

		// Key observation -- a triple is valid iff all 3 point are on their axis-aligned bounding rectangle
		// This is somewhat hard to see, but the logic is that
		// -- if they aren't even on their rectangle, they can't be on a square
		// -- if they are on a rectangle, we can always displace one of the two "long sides" to make a square without affecting the inclusion.
		// Thus, we employ complementary counting.  The "bad triples" have a point strictly on the interior of bounding rectangle of the other two points.
		// Thus, we iterate through each point and count (num NW)(num SE) + (num NE)(num SW)
		// We can use a simple Fenwick tree for the queries.

		N,V1,H1,A,B,C,D,E,F,M := gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi(),gi()
		V := ia(N); H := ia(N); V[0] = V1; H[0] = H1
		for i:=1;i<N;i++ {
			V[i] = (A * V[i-1] + B * H[i-1] + C) % M
			H[i] = (D * V[i-1] + E * H[i-1] + F) % M
		}
		pts := make([]pt,N)
		for i:=0;i<N;i++ { pts[i] = pt{V[i],H[i]} }
		sort.Slice(pts, func(i,j int) bool { return pts[i].y < pts[j].y } )
		northbit := NewFenwick(M+1)
		southbit := NewFenwick(M+1)
		for _,p := range pts { southbit.Inc(p.x+1,1) }
		idx := 0; bad := 0
		for idx < N {
			a,b := idx,idx
			for b+1 < N && pts[b+1].y == pts[a].y { b++ }
			for i:=a; i<=b; i++ { southbit.Inc(pts[i].x+1,-1) }
			for i:=a; i<=b; i++ { 
				x := pts[i].x
				nw := northbit.Prefixsum(x+1-1)
				ne := northbit.Suffixsum(x+1+1)
				sw := southbit.Prefixsum(x+1-1)
				se := southbit.Suffixsum(x+1+1)
				bad += nw*se + ne*sw
			}
			for i:=a; i<=b; i++ { northbit.Inc(pts[i].x+1,1) }
			idx = b+1
		}
		ans := N * (N-1) * (N-2) / 6 - bad
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

