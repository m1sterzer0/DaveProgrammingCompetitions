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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	gr := make([][]int,N)
	for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	distdown := iai(N,0)
	sz := iai(N,0)
	var dfs1 func(int,int)
	dfs1 = func(n,p int) {
		sz[n] = 1; distdown[n] = 0
		for _,c := range gr[n] {
			if c == p { continue }
			dfs1(c,n)
			sz[n] += sz[c]
			distdown[n] += sz[c] + distdown[c]
		}
	}
	dfs1(0,-1)
	distup := iai(N,0)
	var dfs2 func(int,int)
	dfs2 = func(n,p int) {
		if n > 0 { distup[n] = (N - sz[n]) + distdown[p] - distdown[n] - sz[n] + distup[p] }
		for _,c := range gr[n] { if c == p { continue }; dfs2(c,n) }
	}
	dfs2(0,-1)
	for i:=0;i<N;i++ { fmt.Fprintln(wrtr,distup[i]+distdown[i]) }
}



