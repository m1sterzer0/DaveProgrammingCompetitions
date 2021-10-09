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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type baggage struct {w,v int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	// No need to be fancy
	N,M,Q := gi3(); W,V := fill2(N); X := gis(M); L,R := fill2(Q)
	b := make([]baggage,N); for i:=0;i<N;i++ { b = append(b,baggage{w:W[i],v:V[i]}) }
	sort.Slice(b,func(i,j int)bool{return b[i].v > b[j].v || b[i].v == b[j].v && b[i].w < b[j].w})
	avail := [50]bool{}
	for i:=0;i<Q;i++ {
		l,r := L[i]-1,R[i]-1
		for i:=0;i<M;i++ { avail[i] = true }
		for i:=l;i<=r;i++ { avail[i] = false }
		ans := 0
		for _,bb := range b {
			best := -1; bestsize := 1_000_000_000_000_000_000
			for j:=0;j<M;j++ {
				if avail[j] && X[j] >= bb.w && X[j] < bestsize { best = j; bestsize = X[j] }
			}
			if best >= 0 { avail[best] = false; ans += bb.v }
		}
		fmt.Println(ans)
	}
}
