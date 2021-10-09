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
	for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b) }
	visited := make([]bool,N); cnt := 0; ans := 0
	var dfs func(n int)
	dfs = func(n int) {
		visited[n] = true; cnt += 1
		for _,c := range gr[n] { if !visited[c] { dfs(c) } }
	}
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ { visited[j] = false }
		cnt = 0;
		dfs(i)
		ans += cnt
	}
	fmt.Println(ans)
}



