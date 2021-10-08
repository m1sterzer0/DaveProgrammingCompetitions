package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type query struct { t,pos,val,s,l,r int }
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
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	//test(100,500,1000,500,1000,1000)

    T := gi()
    // Ok, we need the relatively esoteric "Lifting the Exponent" Lemma.  Let v_p(x) be the degree of p in the prime factorization of x
	// Part A: Assume p does NOT divide a and p does NOT divide b
	//   1) If p is an ODD PRIME, then v_p(a^n-b^n) = v_p(a-b) + v_p(n)
	//   2) If 4 | a-b, then v_2(a^n-b^n) = v_2(a-b) + v_2(n)
	//   3) If n is even and 2 | a-b, then v_2(a^n-b^n) = v_2(a-b) + v_2(a+b) + v_2(n) - 1
	// Part B 
	//   1) For odd n, if p | a+b, then v_p(a^n+b^n) = v_p(a+b) + v_p(n)
	//   2) for n with gcd(p,n) == 1, if p | a-b, then v_p(a^n-b^n) = v_p(a-b)
	//   3) If n is odd, gcd(p,n) = 1, and p | a+b, then we have v_p(a^n+b^n) = v_p(a+b)
	//If n is odd and 2 | a-b, then v_2(a^n-b^n) = v_2(a-b) = v-2(a-b) + v_2(n) (by B-2 below)
    
	// Now for the problem, we keep 2 fenwick trees
	// a) One that sums v_p(a-b)
	// b) one that sums v_2(a+b)
	
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q,P := gi3(); A := gis(N)
		qq := make([]query,Q)
		for i:=0;i<Q;i++ {
			tt := gi(); if tt == 1 { qq[i] = query{1,gi(),gi(),0,0,0} } else { qq[i] = query{2,0,0,gi(),gi(),gi()} }
		}
		ft1 := NewFenwick(N+10)  // stores v(a-a%P) when a != a%P and a%P != 0 
		ft2 := NewFenwick(N+10)  // stores v(a+a%P) when a != a%P and a%P != 0
		ft3 := NewFenwick(N+10)  // stores v(a) when a%P == 0
		ft4 := NewFenwick(N+10)  // stores 1 when a != a%P and a%P != 0 
		ft1sb,ft2sb,ft3sb,ft4sb := ia(N+10),ia(N+10),ia(N+10),ia(N+10)
		ans := ia(0)
		doit := func(x int) int {
			if x == 0 { return 0 } 
			v := 0; for x % P == 0 { v++; x /= P }; return v
		}
		doquery := func (s,l,r int) int { 
			lans := s * ft3.Rangesum(l,r)
			if P == 2 && s % 2 == 0 { return lans + ft1.Rangesum(l,r) + ft2.Rangesum(l,r) + ft4.Rangesum(l,r) * (doit(s)-1) } 
			return lans + ft1.Rangesum(l,r) + ft4.Rangesum(l,r) * doit(s)
		}
		doupdate := func(idx,a int) {
			if ft1sb[idx] > 0 { ft1.Dec(idx,ft1sb[idx]); ft1sb[idx] = 0 }
			if ft2sb[idx] > 0 { ft2.Dec(idx,ft2sb[idx]); ft2sb[idx] = 0 }
			if ft3sb[idx] > 0 { ft3.Dec(idx,ft3sb[idx]); ft3sb[idx] = 0 }
			if ft4sb[idx] > 0 { ft4.Dec(idx,ft4sb[idx]); ft4sb[idx] = 0 }
			if a < P { return }
			if a % P == 0 {
				ft3sb[idx] = doit(a);   ft3.Inc(idx,ft3sb[idx])
			} else {
				ft1sb[idx] = doit(a-a%P); ft1.Inc(idx,ft1sb[idx])
				ft2sb[idx] = doit(a+a%P); ft2.Inc(idx,ft2sb[idx])
				ft4sb[idx] = 1;           ft4.Inc(idx,ft4sb[idx])
			}
		}
		for i,a := range A { doupdate(i+1,a) }
		for i:=0;i<Q;i++ {
			if qq[i].t == 2 { 
				v := doquery(qq[i].s,qq[i].l,qq[i].r)
				ans = append(ans,v)
			} else {
				doupdate(qq[i].pos,qq[i].val)
			}
		}
		ansstr := vecintstring(ans)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

