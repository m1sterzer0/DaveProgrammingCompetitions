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
func findCenter(N int, gr [][]int) (int,int,int) {
	// BFS1 -- find the furthest node away from node 0
	visited := make([]bool,N); last := 0; q := []int{0}; visited[0] = true
	for len(q) > 0 { 
		n := q[0]; q = q[1:]; last = n
		for _,c := range gr[n] { if !visited[c] { visited[c] = true; q = append(q,c) } }
	}
	par := iai(N,-1); par[last] = last; q = []int{last}; first := last
	// BFS2 -- find the furthest node away from last and call it first
	for len(q) > 0 {
		n := q[0]; q = q[1:]; first = n
		for _,c := range gr[n] { if par[c] == -1 { par[c] = n; q = append(q,c) } }
	}
	// Construct the path from last to first and pick out a center
	path := []int{first}; n := first
	for par[n] != n { n = par[n]; path = append(path,n) }
	if len(path) % 2 == 1 { 
		return len(path)-1, path[len(path)>>1], path[len(path)>>1] 
	} else {
		return len(path)-1, path[len(path)>>1], path[(len(path)>>1) - 1] 
	}
}

const MOD int = 998244353

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); U,V := fill2(N-1); for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
	if N == 2 { fmt.Println(1); return }  // Special case for N==2.  now we are guaranteed center will have at least 2 children
	gr := make([][]int,N); for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	d,c1,c2 := findCenter(N,gr)

	var dfs func (n,p,d int) int
	dfs = func(n,p,d int) int {
		if d == 0 { return 1 }
		ans := 0
		for _,c := range gr[n] {
			if c != p { ans += dfs(c,n,d-1) }
		}
		return ans
	}

	ans := 0
	if d % 2 == 1 {
		n1 := dfs(c1,c2,d/2)
		n2 := dfs(c2,c1,d/2)
		ans = n1 * n2 % MOD
	} else {
		narr := []int{}
		for _,c := range gr[c1] {
			v := dfs(c,c1,d/2-1)
			if v > 0 { narr = append(narr,v) }
		}
		prod := 1; sum := 0
		for _,v := range narr { prod *= (v+1); prod %= MOD; sum += v }
		ans = prod - sum - 1; for ans < 0 { ans += MOD }
	}
	fmt.Println(ans) 
}

