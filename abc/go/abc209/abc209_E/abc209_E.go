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
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
type edge struct {n1,n2 int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := gss(N);
	numnodes := 0
	suffixes := make(map[string]int)
	for _,s := range S { suf := s[len(s)-3:]; _,ok := suffixes[suf]; if !ok { suffixes[suf] = numnodes; numnodes += 1 } }
	gr := make([][]int,numnodes)
	grev := make([][]int,numnodes)
	edges := make(map[edge]bool)  // Since we are using dependency counting, need to uniquify edges
	for _,s := range(S) { 
		pre,suf := s[:3],s[len(s)-3:]
		v,ok := suffixes[pre]
		if !ok { continue }
		n1,n2 := v,suffixes[suf]
		edges[edge{n1,n2}] = true 
	}
	for k := range edges {
		n1,n2 := k.n1,k.n2
		gr[n1] = append(gr[n1],n2)
		grev[n2] = append(grev[n2],n1)
	}
	winning,losing := make([]bool,numnodes),make([]bool,numnodes)
	nodest := ia(0)
	deps := iai(numnodes,0)
	for i:=0;i<numnodes;i++ { 
		deps[i] = len(gr[i])
		if deps[i] == 0 { losing[i] = true; nodest = append(nodest,i) }
	}
	for len(nodest) > 0 {
		idx := len(nodest)-1; n := nodest[idx]; nodest = nodest[:idx]
		if losing[n] {
			for _,c := range grev[n] { 
				if !winning[c] { winning[c] = true; nodest = append(nodest,c) }
			}
		} else {
			for _,c := range grev[n] { 
				if !winning[c] { 
					deps[c]--; 
					if deps[c] == 0 { losing[c] = true; nodest = append(nodest,c) }
				}
			}
		}
	}
	for _,s := range S {
		suf := s[len(s)-3:]
		n := suffixes[suf]
		ans := "Draw"
		if winning[n] { ans = "Aoki" }
		if losing[n]  { ans = "Takahashi" }
		fmt.Fprintln(wrtr,ans)		
	}
}

