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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	reverse := func(a []int) { i,j := 0,len(a)-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j--} }
	for tt:=1;tt<=T;tt++ {
		N := gi(); P := gis(N); for i:=0;i<N;i++ { P[i]-- }
		gr := make([][]int,N); root := -1
		for i,p := range P { if i == p { root = p } else { gr[p] = append(gr[p],i) } }
		paths := make([][]int,0)
		var dfs func(int) []int
		dfs = func(n int) []int {
			if len(gr[n]) == 0 { return []int{n} }
			for i:=1;i<len(gr[n]);i++ { x := dfs(gr[n][i]); paths = append(paths,x) }
			x := dfs(gr[n][0]); x = append(x,n); return x
		}
		paths = append(paths,dfs(root))
		fmt.Fprintln(wrtr,len(paths))
		for _,p := range paths { 
			reverse(p)
			for i:=0;i<len(p);i++ { p[i]++ }
			fmt.Fprintln(wrtr,len(p))
			fmt.Fprintln(wrtr,vecintstring(p))
		}
		fmt.Fprintln(wrtr,"")
	}

}

