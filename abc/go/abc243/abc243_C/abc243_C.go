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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
type pers struct { x int; d byte }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N); S := gs()
	rows := make(map[int][]pers)
	for i:=0;i<N;i++ { x,y := X[i],Y[i]; rows[y] = append(rows[y],pers{x,S[i]}) }
	ans := "No"
	for _,v := range rows {
		sort.Slice(v,func(i,j int) bool { return v[i].x < v[j].x} )
		st := 0
		for _,x := range v { if st == 0 && x.d == 'R' { st = 1 } else if st == 1 && x.d == 'L' { st = 2 } }
		if st == 2 { ans = "Yes" }
	}
	fmt.Println(ans)
}

