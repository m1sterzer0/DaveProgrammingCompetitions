package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A,B := fill2(N-1); gr := make([][]int,N+1)
	for i:=0;i<N-1;i++ { a,b := A[i],B[i]; gr[a] = append(gr[a],b); gr[b] = append(gr[b],a) }
	for i:=1;i<=N;i++ { g := gr[i]; sort.Slice(g,func(i,j int)bool{ return g[i] < g[j]}) }
	ans := ia(0)
	var dfs func(n,p int)
	dfs = func(n,p int) {
		ans = append(ans,n)
		for _,c := range gr[n] {
			if c == p { continue }
			dfs(c,n)
			ans = append(ans,n)
		}
	}
	dfs(1,-1)
	a2 := make([]string,len(ans))
	for i:=0;i<len(ans);i++ { a2[i] = strconv.Itoa(ans[i]) }
	ans2 := strings.Join(a2," ")
	fmt.Println(ans2)
}

