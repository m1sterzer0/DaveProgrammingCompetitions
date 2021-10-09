package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
type query struct { i,k int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,Q := gi2(); A := gis(N); K := ia(Q); for i:=0;i<Q;i++ { K[i] = gi() }
	ma := make(map[int]bool); for _,a := range A { ma[a] = true }
	A2 := make([]int,0); for k := range ma { A2 = append(A2,k) }
	sort.Slice(A2,func(i,j int)bool{return A2[i] < A2[j]})
	ans := ia(Q)
	q := make([]query,Q); for i:=0;i<Q;i++ { q[i] = query{i,K[i]} }
	sort.Slice(q,func(i,j int)bool{return q[i].k < q[j].k})
	qidx := 0; numtoskip := 0
	for _,a := range(A2) {
		for qidx < Q && q[qidx].k+numtoskip < a {
			ans[q[qidx].i] = q[qidx].k+numtoskip
			qidx++
		}
		numtoskip++
	}
	for qidx < Q { ans[q[qidx].i] = q[qidx].k + len(A2); qidx++ }
	for _,aa := range ans {	fmt.Fprintln(wrtr,aa) }
}



