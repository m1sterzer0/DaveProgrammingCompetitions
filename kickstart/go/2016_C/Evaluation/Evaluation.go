package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }

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
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		good := true; strlookup := make(map[string]int); nstr := 0; elist := make([]PI,0)
		getStrNum := func(s string) int { v,ok := strlookup[s]; if !ok { v = nstr; nstr++; strlookup[s] = v }; return v }

		parseExpr := func(s string) (int,[]int) {
			idx := 0
			for s[idx] != '=' { idx++ }; l := getStrNum(s[0:idx])
			for s[idx] != '(' { idx++ }; argstart := idx+1
			for s[idx] != ')' { idx++ }; argend := idx
			if argstart == argend { return l,[]int{} }
			ss := strings.Split(s[argstart:argend],",")
			rarr := make([]int,len(ss))
			for i,sss := range ss { rarr[i] = getStrNum(sss) }
			return l,rarr
		}

		for i:=0;i<N;i++ {
			l,rarr := parseExpr(gs())
			for _,r := range rarr { if l==r { good = false; break }; elist = append(elist,PI{l,r}) }
		}
		if nstr > N { good = false }  // This means we have more variables than equations, so something can't get set
		if good {
			cnt,_ := Kosaraju(nstr,elist)
			if cnt < nstr { good = false }
		}
		ans := "BAD"; if good { ans = "GOOD" }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

