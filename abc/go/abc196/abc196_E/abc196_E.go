package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type ff struct {lb,adder,ub int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); A,T := fill2(N); Q := gi(); X := gis(Q)
	f := ff{-1_000_000_000_000_000_000,0,1_000_000_000_000_000_000}
	for i:=0;i<N;i++ {
		a,t := A[i],T[i]
		if t == 1 { f.adder += a; f.lb += a; f.ub += a } else if t == 2 { f.lb,f.ub = max(f.lb,a),max(f.ub,a) } else { f.lb,f.ub = min(f.lb,a),min(f.ub,a) }
	}
	for _,x := range X {
		y := x + f.adder; if y < f.lb { y = f.lb }; if y > f.ub { y = f.ub }
		fmt.Fprintln(wrtr,y) 
	}
}



