package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
type query struct {idx,l,r,x int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi2(); A := gis(N); L,R,X := fill3(Q); for i:=0;i<Q;i++ { L[i]--; R[i]-- }
	qq := make([]query,Q)
	for i:=0;i<Q;i++ { qq[i] = query{i,L[i],R[i],X[i]} }
	sort.Slice(qq,func(i,j int) bool { return qq[i].r < qq[j].r} )
	prev := make([]int,0,65); previdx := make([]int,0,65); prevbit := make([]int,0,65)
	cur := make([]int,0,65); curidx := make([]int,0,65); curbit := make([]int,0,65)
	qptr := 0
	ansarr := make([]string,Q)
	for i:=0;i<N;i++ {
		cur = cur[:0]; curidx = curidx[:0]; curbit = curbit[:0]
		cur = append(cur,A[i]); curidx = append(curidx,i); curbit = append(curbit,bits.TrailingZeros64(uint64(A[i])))
		pcnt := len(prev)
		for j:=0;j<pcnt;j++ {
			bb := prev[j]; ccnt := len(cur)
			for k:=0;k<ccnt;k++ {
				if bb & (1 << curbit[k]) != 0 { bb ^= cur[k] }
			}
			if bb != 0 {
				idx := previdx[j]
				cur = append(cur,bb); curidx = append(curidx,idx); curbit = append(curbit, bits.TrailingZeros64(uint64(bb)))
			}
		}
		for qptr < Q && qq[qptr].r == i {
			idx,l,x := qq[qptr].idx,qq[qptr].l,qq[qptr].x
			for j,ii := range curidx {
				if ii < l { break }
				if x & (1 << curbit[j]) != 0 { x ^= cur[j] }
			}
			ansarr[idx] = "No"; if x == 0 { ansarr[idx] = "Yes" }
			qptr++
		}
		cur,curidx,curbit,prev,previdx,prevbit = prev,previdx,prevbit,cur,curidx,curbit
	}
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}
