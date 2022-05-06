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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type query struct {l,r,idx int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// This is a sqrt decomp problem using Mo's algorithm.
	// I'm not familiar with Mo's so this is a first for me.
	N := gi()
	A := make([]int32,N); for i:=0;i<N;i++ { A[i] = int32(gi()) }
	Q := gi()
	L := make([]int32,Q); R := make([]int32,Q)
	for i:=0;i<Q;i++ { L[i] = int32(gi()); R[i] = int32(gi()) }
	for i:=0;i<Q;i++ { L[i]--; R[i]-- }
	qidx := make([]int32,Q); for i:=0;i<Q;i++ { qidx[i] = int32(i) }
	const bsize = 128 // Pick a power of 2 near sqrt(N)
	qcomp := func(i,j int) bool {  //Mo's algorithm
		lidx,ridx := qidx[i],qidx[j]
		ablock,bblock,ar,br := L[lidx]/bsize,L[ridx]/bsize,R[lidx],R[ridx]
		if ablock != bblock { return ablock < bblock }
		if ablock & 1 == 1 { return ar < br }
		return ar > br // Sort order optimization : https://cp-algorithms.com/data_structures/sqrt_decomposition.html
	}
	sort.Slice(qidx,qcomp)
	cnt := make([]int32,100010)
	res,l,r := int32(0),int32(0),int32(0)
	moadd := func(id int32) { a := A[id]; cnt[a]++; if cnt[a] & 1 == 0 { res++ } }
	morem := func(id int32) { a := A[id]; if cnt[a] & 1 == 0 { res-- }; cnt[a]-- }
	moquer := func(ll,rr int32) int32 {
		for r < rr { r++; moadd(r) }
		for ll < l { l--; moadd(l) }
		for l < ll { morem(l); l++ }
		for rr < r { morem(r); r-- }
		return res
	}
	moadd(0)
	ansarr := make([]int32,Q)
	for _,qq := range qidx { ansarr[qq] = moquer(L[qq],R[qq]) }
	for _,a := range ansarr { fmt.Fprintln(wrtr,a) }
}
