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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	gr := make([][]int,N)
	for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }
	colors := make([]byte,N); for i:=0;i<N;i++ { colors[i] = '.' }
	visited := make([]bool,N); order := make([]int,0); check := make([][]int,N)
	var dfs1 func(int)
	dfs1 = func(n int) {
		if visited[n] { return }
		visited[n] = true
		order = append(order,n)
		for _,c := range gr[n] { if visited[c] { check[n] = append(check[n],c) } }
		for _,c := range gr[n] { if !visited[c] { dfs1(c) } }
	}
	tmpans := 0; lastidx :=0
	var dfs2 func(int,byte)
	dfs2 = func(oidx int, c byte) {
		n := order[oidx]
		for _,n2 := range check[n] { if colors[n2] == c { return } }
		if oidx == lastidx { tmpans++; return }
		colors[n] = c
		dfs2(oidx+1,'r')
		dfs2(oidx+1,'g')
		dfs2(oidx+1,'b')
		colors[n] = '.'
	}
	ans := 1
	for i:=0;i<N;i++ {
		if visited[i] { continue }
		order = order[:0]
		dfs1(i)
		tmpans = 0; lastidx = len(order)-1
		dfs2(0,'r')
		ans *= 3 * tmpans
	}
	fmt.Println(ans)
}

