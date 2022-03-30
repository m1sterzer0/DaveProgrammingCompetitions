package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
type edge struct {a,b,l int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		V,E := gi2(); A,B,L := fill3(E); for i:=0;i<E;i++ { A[i]--; B[i]-- }
		edges := make([]edge,E)
		for i:=0;i<E;i++ { edges[i] = edge{A[i],B[i],L[i]} }
		sort.Slice(edges,func (i,j int) bool { return edges[i].l < edges[j].l })
		special := make([]bool,V)
		visited := make([]bool,V)
		ans := 1
		for _,e := range edges {
			if e.l == 0 { 
				special[e.a] = true; special[e.b] = true; ans *= 2; visited[e.a] = true; visited[e.b] = true
			} else if !visited[e.a] && !visited[e.b] {
				ans *= 2; visited[e.a] = true; visited[e.b] = true
			} else if !visited[e.a] {
				if special[e.b] { ans *= 2 }; visited[e.a] = true
			} else if !visited[e.b] {
				if special[e.a] { ans *= 2}; visited[e.b] = true
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

