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
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func max(a,b int) int { if a > b { return a }; return b }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type stone struct {s,e,l int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// If I have two stones, which should I eat first?
		// If I eat 1 then 2, I get e1 + max(0,e2-l2*s1)
		// If I eat 2 then 1, I get e2 + max(0,e1-l1*s2)
		// Assume that the clipping at zero doesn't happen -- will deal with that later
		// then e1 + e2 - l2*s1 > e1 + e2 - l1*s2 iff l1*s2 > l2*s1 iff l1/s1 > l2/s2
		// This suggests that we can sort the rocks, and then choose some subset to eat.
		// Rest is a simple DP based on time (Since S_i * N is quite small)
		N := gi(); S,E,L := fill3(N)
		ss := make([]stone,0)
		for i:=0;i<N;i++ { ss = append(ss,stone{S[i],E[i],L[i]}) }
		sort.Slice(ss,func(i,j int) bool { return ss[i].l * ss[j].s > ss[j].l*ss[i].s })
		tots := sumarr(S)
		cur,old := ia(tots+1),ia(tots+1)
		for _,s := range ss {
			cur,old = old,cur
			for i:=0;i<=tots;i++ { cur[i] = old[i] } // This is me skipping the current stone
			for i:=0;i<=tots-s.s;i++ {
				v := s.e - s.l*i
				if v <= 0 { break }
				cur[i+s.s] = max(cur[i+s.s],old[i]+v)
			}
		}
		ans := 0
		for i:=0;i<=tots;i++ {ans = max(ans,cur[i]) }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

