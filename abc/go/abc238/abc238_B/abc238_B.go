package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N)
	cuts := make([]int,0)
	cuts = append(cuts,0)
	curs := 0
	for _,a := range A { curs += a; curs %= 360; cuts = append(cuts,curs) }
	sort.Slice(cuts,func(i,j int) bool { return cuts[i] < cuts[j] } )
	best := 0
	for i,c := range cuts { 
		nc := 0; if i+1 < len(cuts) { nc = cuts[i+1] } else { nc = 360+cuts[0] }; best = max(best,nc-c)
	}
	fmt.Println(best)
}

