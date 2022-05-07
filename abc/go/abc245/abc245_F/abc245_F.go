package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

type PI struct{ x, y int }
func Kosaraju(n int, diredges []PI) (int, []int) {
	g, grev, visited, visitedInv, scc, s, counter := make([][]int, n), make([][]int, n), make([]bool, n), make([]bool, n), make([]int, n), make([]int, 0, n), 0
	var dfs1, dfs2 func(int)
	for _, xx := range diredges { x, y := xx.x, xx.y; g[x] = append(g[x], y); grev[y] = append(grev[y], x) }
	dfs1 = func(u int) { if !visited[u] { visited[u] = true; for _, c := range g[u] { dfs1(c) }; s = append(s, u) } }
	for i := 0; i < n; i++ { dfs1(i) }
	dfs2 = func(u int) {
		if !visitedInv[u] { visitedInv[u] = true; for _, c := range grev[u] { dfs2(c) }; scc[u] = counter }
	}
	for i := n - 1; i >= 0; i-- { nn := s[i]; if !visitedInv[nn] { dfs2(nn); counter += 1 } }; return counter, scc
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi()
	U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
	edgelist := make([]PI,M); for i:=0;i<M;i++ { edgelist[i] = PI{U[i],V[i]} }
	nscc,scc := Kosaraju(N,edgelist)
	sccsz := make([]int,nscc); for _,s := range scc { sccsz[s]++ }
	// Now we need to make the graph of the sccs
	// This should be a DAG, so topological sort is possible
	// We want to process from bottom up
	numdeps := make([]int,nscc)
	gr := make([][]int,nscc)
	grrev := make([][]int,nscc)
	edgemap := make(map[PI]bool)
	for i:=0;i<M;i++ { 
		u,v := U[i],V[i]; su,sv := scc[u],scc[v]
		if su==sv || edgemap[PI{su,sv}] { continue }
		edgemap[PI{su,sv}] = true
		numdeps[su]++; gr[su] = append(gr[su],sv); grrev[sv] = append(grrev[sv],su) 
	}
	goodscc := make([]bool,nscc)
	q := make([]int,0,nscc)
	for i:=0;i<nscc;i++ { if numdeps[i] == 0 { q = append(q,i) } }
	for len(q) > 0 {
		s := q[0]; q = q[1:]
		if sccsz[s] > 1 { goodscc[s] = true }
		for _,n := range gr[s] { if goodscc[n] { goodscc[s] = true } }
		for _,n := range grrev[s] { numdeps[n]--; if numdeps[n] == 0 { q = append(q,n) } }
	}
	ans := 0; for i:=0;i<N;i++ { s := scc[i]; if goodscc[s] { ans++ } }
	fmt.Println(ans)
}

