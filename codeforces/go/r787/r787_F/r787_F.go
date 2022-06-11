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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N,K := gi(),gi(); X,Y := gi(),gi(); X--; Y--; A := gis(K); for i:=0;i<K;i++ { A[i]-- }; U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
		//Build the graph
		gr := make([][]int,N)
		for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
		//Mark the path down to the target
		needToVisit := make([]bool,N); for _,a := range A { needToVisit[a] = true }
		pathToY := make([]bool,N)
		var dfs1 func(n,p int) bool
		dfs1 = func(n,p int) bool {
			ans := false
			if n == Y { ans = true }
			for _,c := range gr[n] { if c != p && dfs1(c,n) { ans = true } }
			pathToY[n] = ans
			return ans
		}
		dfs1(X,-1)
		var dfs2 func(n,p int) bool
		dfs2 = func(n,p int) bool {
			ans := false
			if needToVisit[n] || pathToY[n] { ans = true }
			for _,c := range gr[n] { if c != p && dfs2(c,n) { ans = true } }
			needToVisit[n] = ans
			return ans
		}
		dfs2(X,-1)
		var dfs3 func(n,p int) int
		dfs3 = func(n,p int) int {
			ans := 0
			for _,c := range gr[n] { 
				if c == p || !needToVisit[c] { continue }
				cans := dfs3(c,n)
				if pathToY[c] { ans += 1 + cans } else { ans += 2 + cans }
			}
			return ans
		}
		ans := dfs3(X,-1)
		fmt.Fprintln(wrtr,ans)
	}
}

