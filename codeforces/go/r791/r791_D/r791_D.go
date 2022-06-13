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
func gi64() int64 { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M,K := gi(),gi(),gi64(); A := gis(N); U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }
	gr := make([][]int,N)
	for i:=0;i<M;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v) }
	visited := make([]bool,N)
	depth := make([]int64,N)
	check := func(v int) bool {
		for i:=0;i<N;i++ { visited[i] = false; depth[i] = -1 }
		var dfs func(n int)
		dfs = func(n int) {
			if depth[n] != -1 { return }
			if A[n] > v { depth[n] = 0; return }
			if visited[n] { depth[n] = 1000000000000000001; return }
			visited[n] = true; m := int64(1)
			for _,c := range gr[n] {
				dfs(c)
				if 1 + depth[c] > m { m = 1 + depth[c] }
			}
			depth[n] = m
		}
		for i:=0;i<N;i++ { dfs(i); if depth[i] >= K { return true } }
		return false
	}
	l,u := 0,1000000001
	for u-l > 1 { m := (u+l)>>1; if check(m) { u = m } else { l = m } }
	if u > 1000000000 { fmt.Println(-1) } else { fmt.Println(u) }
}
