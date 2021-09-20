package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func min(a,b int) int { if a > b { return b }; return a }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,C := gi2(); A,B,CC := fill3(N)
	deltas := make(map[int]int)
	for i:=0;i<N;i++ { a,b,c := A[i],B[i],CC[i]; deltas[a] += c; deltas[b+1] -= c }
	events := make([]int,0); for k := range deltas { events = append(events,k) }
	sort.Slice(events,func(i,j int)bool { return events[i] < events[j] })
	last := -1; runningcosts := 0; ans := 0
	for _,e := range events {
		ans += (e-last) * min(C,runningcosts)
		runningcosts += deltas[e]
		last = e
	}
	fmt.Println(ans) 
}



