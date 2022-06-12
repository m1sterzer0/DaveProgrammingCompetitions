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
type edge struct {n2,eidx int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		P := gi(); N := 1<<P; U,V := fill2(N-1)
		for i:=0;i<N-1;i++ { U[i]--; V[i]-- }
		gr := make([][]edge,N)
		for i:=0;i<N-1;i++ { u,v := U[i],V[i]; gr[u] = append(gr[u],edge{v,i}); gr[v] = append(gr[v],edge{u,i}) }
		nans := make([]int,N)
		eans := make([]int,N-1)
		nans[0] = 1 << P
		adder := 1; big := 1 << P
		var dfs func(n,p,v int)
		dfs = func(n,p,v int) {
			for _,e := range gr[n] {
				if e.n2 != p {
					if v ^ adder < big { 
						eans[e.eidx] = adder; nans[e.n2] = adder | big
					} else {
						eans[e.eidx] = adder | big; nans[e.n2] = adder
					}
					adder++
					dfs(e.n2,n,v ^ big)
				}
			}
		}
		dfs(0,-1,big)
		fmt.Fprintln(wrtr,1)
		fmt.Fprintln(wrtr,vecintstring(nans))
		fmt.Fprintln(wrtr,vecintstring(eans))
	}
}

