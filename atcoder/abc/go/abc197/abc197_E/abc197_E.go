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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); X,C := fill2(N)
	inf := 1_000_000_000_000_000_000
	minc,maxc := iai(N+1,inf),iai(N+1,-inf)
	for i:=0;i<N;i++ {
		idx := C[i]; minc[idx] = min(minc[idx],X[i]); maxc[idx] = max(maxc[idx],X[i]) 
	}
	lpos,ltime,rpos,rtime := 0,0,0,0
	for i:=0;i<=N;i++ {
		if minc[i] > maxc[i] { continue }
		newltime := min(ltime + abs(lpos-maxc[i]), rtime + abs(rpos-maxc[i])) + (maxc[i]-minc[i])
		newrtime := min(ltime + abs(lpos-minc[i]), rtime + abs(rpos-minc[i])) + (maxc[i]-minc[i])
		lpos,ltime,rpos,rtime = minc[i],newltime,maxc[i],newrtime
	}
	ans := min(ltime+abs(lpos),rtime+abs(rpos))
	fmt.Println(ans)
}



