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

type Fenwick2 struct { xmax,ymax int; bit2 []int32 }
func NewFenwick2(xmax,ymax int) *Fenwick2 { return &Fenwick2{xmax,ymax,make([]int32,xmax*ymax)} }
func (q *Fenwick2) Inc(x,y,val int) { 
	v := int32(val)
	for xx:=x; xx<q.xmax; xx += (xx & -xx) {
		for yy:=y; yy < q.ymax; yy += (yy & -yy) {	q.bit2[xx*q.ymax+yy] += v }
	}
}
func (q *Fenwick2) Prefixsum(x,y int) int {
	ans := int32(0)
	for xx:=x; xx>0; xx -= (xx & -xx) {
		for yy:=y; yy>0; yy -= (yy & -yy) {	ans += q.bit2[xx*q.ymax+yy] }
	}
	return int(ans)
}

type Fenwick struct { n, tot int; bit []int }
func NewFenwick(n int) *Fenwick { buf := make([]int, n+1); return &Fenwick{n, 0, buf} }
func (q *Fenwick) Clear() { for i := 0; i <= q.n; i++ { q.bit[i] = 0 }; q.tot = 0 }
func (q *Fenwick) Inc(idx int, val int) { for idx <= q.n { q.bit[idx] += val; idx += idx & (-idx) }; q.tot += val }
func (q *Fenwick) Prefixsum(idx int) int {
	if idx < 1 { return 0 }; ans := 0; for idx > 0 { ans += q.bit[idx]; idx -= idx & (-idx) }; return ans
}

type pileSum struct { a,b,c int32 }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()

	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		A := make([]int32,0); for i:=0;i<3*N;i++ { A = append(A,int32(gi())) }
		B := make([]int32,0); for i:=0;i<3*N;i++ { B = append(B,int32(gi())) }
				
		var dfs1 func(idx int)
		var dfs2 func(idx int)
		sum1 := int32(0); sum2 := int32(0); sum3 := int32(0)
		cnt1 := 0; cnt2 := 0; cnt3 := 0
		threeN := 3 * N

		a,b := make([]pileSum,0),make([]pileSum,0)
		dfs1 = func(idx int) {
			if idx == threeN { a = append(a,pileSum{sum1,sum2,sum3}); return }
			if cnt1 < N { sum1 += A[idx]; cnt1++; dfs1(idx+1); cnt1--; sum1 -= A[idx] }
			if cnt2 < N { sum2 += A[idx]; cnt2++; dfs1(idx+1); cnt2--; sum2 -= A[idx] }
			if cnt3 < N { sum3 += A[idx]; cnt3++; dfs1(idx+1); cnt3--; sum3 -= A[idx] }
		}
		dfs1(0)
		dfs2 = func(idx int) {
			if idx == threeN { b = append(b,pileSum{sum1,sum2,sum3}); return }
			if cnt1 < N { sum1 += B[idx]; cnt1++; dfs2(idx+1); cnt1--; sum1 -= B[idx] }
			if cnt2 < N { sum2 += B[idx]; cnt2++; dfs2(idx+1); cnt2--; sum2 -= B[idx] }
			if cnt3 < N { sum3 += B[idx]; cnt3++; dfs2(idx+1); cnt3--; sum3 -= B[idx] }
		}
		dfs2(0)
		for i,aa := range a {
			if aa.a <= aa.b { 
				if aa.c <= aa.a {
					a[i] = pileSum{aa.c,aa.a,aa.b}
				} else if aa.b <= aa.c {
					a[i] = pileSum{aa.a,aa.b,aa.c}
				} else {
					a[i] = pileSum{aa.a,aa.c,aa.b}
				}
			} else {
				if aa.c <= aa.b {
					a[i] = pileSum{aa.c,aa.b,aa.a}
				} else if aa.a <= aa.c {
					a[i] = pileSum{aa.b,aa.a,aa.c}
				} else {
					a[i] = pileSum{aa.b,aa.c,aa.a}
				}
			}
		}
		//sort.Slice(a,func(i,j int) bool { return a[i].a < a[j].a || a[i].a == a[j].a && a[i].b < a[j].b || a[i].a == a[j].a && a[i].b == a[j].b && a[i].c < a[j].c })
		//sort.Slice(b,func(i,j int) bool { return b[i].a < b[j].a || b[i].a == b[j].a && b[i].b < b[j].b || b[i].a == b[j].a && b[i].b == b[j].b && b[i].c < b[j].c })
		sort.Slice(a,func(i,j int) bool { return a[i].a < a[j].a })
		sort.Slice(b,func(i,j int) bool { return b[i].a < b[j].a })
		bvals := make([]int32,0); lastbval := int32(-1)
		for _,bb := range b {
			if bb.a == lastbval { continue }
			lastbval = bb.a
			bvals = append(bvals,lastbval)
		}
		bvals2idx := make([]int,bvals[len(bvals)-1]+1)
		for i,bb := range bvals { bvals2idx[bb] = i+1 }

		numbvals := len(bvals)
		cache := make(map[int32]int)
		findltidx := func(v int32) int {
			res,ok := cache[v]
			if !ok {
				if v <= bvals[0] { 
					res = 0 
				} else {
					l,r := 0,numbvals
					for r-l > 1 { 
						m := (r+l)>>1
						if v > bvals[m] { l = m } else { r = m }
					}
					res = l + 1
				}
				cache[v] = res
			}
			return res
		}

		findidx := func(v int32) int { return bvals2idx[v] }

		sb := make([]int,len(a))
		// sb = wins on (a,b) + wins on (a,c) - 2 * wins(a,b,c) + wins on (b,c)
		blen := len(b); bidx := 0
		ft2 := NewFenwick2(numbvals+2,numbvals+2)
		ftb := NewFenwick(numbvals+1)
		ftc := NewFenwick(numbvals+1)
		for i,aa := range a {
			if i > 0 && aa.a == a[i-1].a && aa.b == a[i-1].b && aa.c == a[i-1].c { continue } 
			for bidx < blen && b[bidx].a < aa.a { 
				bb,cc := findidx(b[bidx].b),findidx(b[bidx].c)
				ftb.Inc(bb,1); ftc.Inc(cc,1); ft2.Inc(bb,cc,1)
				bidx++
			}
			abb := findltidx(aa.b); acc := findltidx(aa.c)
			sb[i] = ftb.Prefixsum(abb) + ftc.Prefixsum(acc) - 2*ft2.Prefixsum(abb,acc)
		}
		// Finish out the 2d segtree
		for bidx < blen { 
			bb,cc := findidx(b[bidx].b),findidx(b[bidx].c)
			ftb.Inc(bb,1); ftc.Inc(cc,1); ft2.Inc(bb,cc,1)
			bidx++
		}
		// Don't know if resorting or using the 2d tree is faster
		// Try reusing the 2d tree first
		for i,aa := range a {
			if i > 0 && aa.a == a[i-1].a && aa.b == a[i-1].b && aa.c == a[i-1].c { continue } 
			abb := findltidx(aa.b); acc := findltidx(aa.c)
			sb[i] += ft2.Prefixsum(abb,acc)
		}
		best := 0
		for _,s := range sb { if s > best { best = s } }
		ans := float64(best) / float64(blen)
		fmt.Printf("Case #%v: %.12f\n",tt,ans)

	}
}

