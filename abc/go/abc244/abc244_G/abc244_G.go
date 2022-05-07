package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); U,V := fill2(M); for i:=0;i<M;i++ { U[i]--; V[i]-- }; S := gs()
	targ := make([]int,N); for i:=0;i<N;i++ { if S[i] == '1' { targ[i] = 1 } }
	gr := make([][]int,N)
	for i:=0;i<M;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],v); gr[v] = append(gr[v],u) }
	cnt := make([]int,N); ans := make([]int,0)
	var dfs func(n,p int)
	dfs = func(n,p int) {
		cnt[n]++; ans = append(ans,n+1)
		for _,c := range gr[n] {
			if cnt[c] > 0 { continue }
			dfs(c,n)
			cnt[n]++; ans = append(ans,n+1)
		}
		if p >= 0 && cnt[n] & 1 != targ[n] {
			cnt[p]++; ans = append(ans,p+1)
			cnt[n]++; ans = append(ans,n+1)
		}
	}
	dfs(0,-1)
	if cnt[0] & 1 != targ[0] { ans = ans[1:] }
	K := len(ans); fmt.Fprintln(wrtr,K)
	ansstr := vecintstring(ans)
	fmt.Fprintln(wrtr,ansstr)
}

