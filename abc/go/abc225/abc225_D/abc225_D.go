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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
type query struct { t,x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi2()
	qarr := make([]query,Q)
	for i:=0;i<Q;i++ { 
		t := gi(); x := gi(); y := 0; if t != 3 { y = gi() }; qarr[i] = query{t,x,y}
	}
	front := ia(N+1); back := ia(N+1); ansarr := []int{}
	for _,q := range qarr {
		if q.t == 1 { front[q.y] = q.x; back[q.x] = q.y }
		if q.t == 2 { front[q.y] = 0; back[q.x] = 0 }
		if q.t == 3 {
			f := q.x; for front[f] != 0 { f = front[f] }
			ansarr = ansarr[:0]; ansarr = append(ansarr,f)
			for back[f] != 0 { f = back[f]; ansarr = append(ansarr,f) }
			l := len(ansarr); ansstr := vecintstring(ansarr)
			fmt.Fprintf(wrtr,"%v %v\n",l,ansstr)
		}
	}
}

