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
type Twosat struct { n int; answer []bool; edgelist []PI }
func NewTwosat(n int) *Twosat {
	answer := make([]bool, n); edgelist := make([]PI, 0); return &Twosat{n, answer, edgelist}
}
func (q *Twosat) AddClause(i int, f bool, j int, g bool) {
	n1, n2, n3, n4 := 2*i, 2*j, 2*j, 2*i; if f { n4 += 1 } else { n1 += 1 }; if g { n2 += 1 } else { n3 += 1 }
	q.edgelist = append(q.edgelist, PI{n1, n2}); q.edgelist = append(q.edgelist, PI{n3, n4})
}
func (q *Twosat) Satisfiable() (bool, []bool) {
	_, id := Kosaraju(2*q.n, q.edgelist)
	for i := 0; i < q.n; i++ { if id[2*i] == id[2*i+1] { return false, q.answer }; q.answer[i] = id[2*i] < id[2*i+1] }
	return true, q.answer
}

func sieve(n int) []int {
	s := make([]bool,n+1)
	for i:=0;i<=n;i++ { s[i] = true }
	s[0] = false; s[1] = false; for i:=4;i<=n;i+=2 { s[i] = false }
	for i:=3;i*i<=n;i+=2 {
		if !s[i] { continue }
		for j:=i*i;j<=n;j+=2*i { s[j] = false }
	}
	ans := []int{}
	for i:=2;i<=n;i++ { if s[i] { ans = append(ans,i)} }
	return ans
}

func doappend(xxx map[int][]int, p int, n int) {
	v,ok := xxx[p]; if !ok { v = []int{} }; xxx[p] = append(v,n)
}

func loadfactors(n int, a int, pr []int, xxx map[int][]int) {
	for _,p := range pr { 
		if a % p != 0 { continue }
		a /= p; for a % p == 0 { a /= p }
		doappend(xxx,p,n)
	}
	if a != 1 { doappend(xxx,a,n) }
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	//f2, _ := os.Create("mem.prof"); defer f2.Close(); runtime.GC(); pprof.WriteHeapProfile(f2)
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A,B := fill2(N)
	primes := sieve(1414)
	xxx := make(map[int][]int)
	for i:=0;i<N;i++ {
		loadfactors(2*i,A[i],primes,xxx)
		loadfactors(2*i+1,B[i],primes,xxx)
	}
	// We need to use the chain variable trick
	// Assume we have X1,X2,...,Xn -- at most one can be true
	// Introduce n new variable
	// Zn   = !Xn
	// Zn-1 = !Xn && !Xn-1
	// Zn-2 = !Xn && !Xn-1 && !Xn-2
	// Zn-3 = !Xn && !Xn-1 && !Xn-2 && !Xn-3
	// ...
	// Z1 = !X1 && !X2 && !X3 && ... && !Xn
	// Then we have the following 3 OR terms
	// ** Zi --> !Xi  ======== !Zi || !Xi
	// ** Xi --> Zi+1 ======== !Xi || Zi+1
	// ** Zi --> Zi+1 ======== !Zi || Zi+1

	// Pass 1 to count the number of variables
	numvars := N
	for _,v := range xxx { if len(v) > 1 { numvars += len(v) } }
	ts := NewTwosat(numvars)
	// Pass 2 to add the clauses
	numvars = N
	for _,v := range xxx {
		if len(v) == 1 { continue }
		for i,x := range v {
			ts.AddClause(numvars+i,false,x>>1,x&1==1)
			if i == len(v)-1 { continue }
			ts.AddClause(x>>1,x&1==1,numvars+i+1,true) 
			ts.AddClause(numvars+i,false,numvars+i+1,true) 
		}
		numvars += len(v)
	}
	b,_ := ts.Satisfiable()
	ans := "No"; if b { ans = "Yes" }; fmt.Println(ans)
}
