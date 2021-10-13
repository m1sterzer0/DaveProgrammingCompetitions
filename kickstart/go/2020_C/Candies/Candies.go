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
func gi2() (int,int) { return gi(),gi() }
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
type query struct { typ byte; f,s int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2(); A := ia(N+1); for i:=1;i<=N;i++ { A[i] = gi() } 
		queries := make([]query,Q)
		for i:=0;i<Q;i++ { queries[i] = query{gs()[0],gi(),gi()} }
		ft1 := NewFenwick(N)  // Stores A[1] - 2A[2] + 3A[3] - 4A[4] + 5A[5] ...
		ft2 := NewFenwick(N)  // Stores A[1] - A[2] + A[3] - A[4] + A[5] - A[6] ...
		for i:=1;i<=N;i++ { 
			sgn := 1; if i % 2 == 0 { sgn = -1 }
			ft1.Inc(i,i*sgn*A[i])
			ft2.Inc(i,sgn*A[i])
		}
		sumq := 0
		for _,q := range queries {
			if q.typ == 'U' {
				i := q.f
				sgn := 1; if i % 2 == 0 { sgn = -1 }
				ft1.Dec(i,i*sgn*A[i])
				ft2.Dec(i,sgn*A[i])
				A[i] = q.s
				ft1.Inc(i,i*sgn*A[i])
				ft2.Inc(i,sgn*A[i])
			} else {
				s1 := ft1.Rangesum(q.f,q.s)
				s2 := ft2.Rangesum(q.f,q.s)
				s3 := s1 - (q.f-1)*s2
				if q.f % 2 == 0 { s3 *= -1 }
				sumq += s3
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,sumq)
    }
}
