package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type query struct { idx,u,d int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); P := gis(N-1); Q := gi(); U,D := fill2(Q); for i:=0;i<Q;i++ { U[i]-- }; for i:=0;i<N-1;i++ { P[i]-- }
	gr := make([][]int,N)
	for i:=1;i<N;i++ { p := P[i-1]; gr[p] = append(gr[p],i) }
	qbynode := make([][]query,N)
	for i:=0;i<Q;i++ { u,d := U[i],D[i]; qbynode[u] = append(qbynode[u],query{i,u,d}) }
	dcnt := iai(N,0); ans := iai(Q,0)
	var dfs func(int,int)
	dfs = func(n,dp int) {
		for _,q := range qbynode[n] { ans[q.idx] -= dcnt[q.d] }
		dcnt[dp]++
		for _,c := range gr[n] { dfs(c,dp+1) }
		for _,q := range qbynode[n] { ans[q.idx] += dcnt[q.d] }
	}
	dfs(0,0)
	for _,a := range ans { fmt.Fprintln(wrtr,a) }
}

