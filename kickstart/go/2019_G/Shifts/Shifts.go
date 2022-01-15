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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
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

type happy struct {ha,hb int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,H := gi2(); A := gis(N); B := gis(N); ans := 0
		n1 := (N+1)/2; n2 := N-n1
		var doHappy func(a,b []int) []happy
		doHappy = func(a,b []int) []happy {
			if len(a) == 1 {
				return []happy{{a[0],0},{0,b[0]},{a[0],b[0]}}
			} else {
				h1 := doHappy(a[1:],b[1:])
				res := make([]happy,0,3*len(h1))
				for _,h := range h1 {
					res = append(res,happy{h.ha+a[0],h.hb})
					res = append(res,happy{h.ha,     h.hb+b[0]})
					res = append(res,happy{h.ha+a[0],h.hb+b[0]})
				}
				return res
			}
		}

		h1 := doHappy(A[0:n1],B[0:n1])
		if n2 == 0 {
			for _,h := range h1 { if h.ha >= H && h.hb >= H { ans++ } }
		} else {
			h2 := doHappy(A[n1:N],B[n1:N])
			sort.Slice(h1,func(i,j int) bool { return h1[i].ha < h1[j].ha} )
			sort.Slice(h2,func(i,j int) bool { return h2[i].ha > h2[j].ha} )

			// coordinate compression
			sb := make(map[int]bool)
			for _,h := range h2 { sb[h.hb] = true }
			bcoords := make([]int,0,len(sb))
			for k := range sb { bcoords = append(bcoords,k) }
			sort.Slice(bcoords,func(i,j int) bool { return bcoords[i] < bcoords[j]} )
			old2new := make(map[int]int)
			for i,v := range bcoords { old2new[v] = i }

			h2ptr := 0
			ft := NewFenwick(len(bcoords)+10)
			lastBcoord := bcoords[len(bcoords)-1]
			for _,h := range h1 {
				for h2ptr < len(h2) && h.ha + h2[h2ptr].ha >= H {
					newv := old2new[h2[h2ptr].hb]
					ft.Inc(newv+1,1)
					h2ptr++
				}
				if h.hb + lastBcoord < H { continue }
				l,u := -1,len(bcoords)-1
				for u-l > 1 { m := (u+l)>>1; if bcoords[m] + h.hb >= H { u = m } else { l = m } }
				ans += ft.Suffixsum(u+1)
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

