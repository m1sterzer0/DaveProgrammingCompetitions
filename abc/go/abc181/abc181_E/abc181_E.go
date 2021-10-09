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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gi2() (int,int) { return gi(),gi() }
func min(a,b int) int { if a > b { return b }; return a }
func abs(a int) int { if a >= 0 { return a }; return -a }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); H := gis(N); W := gis(M)
	left := make([]int,0); right := make([]int,0)
	left = append(left,0); right = append(right,0); nl,nr := 0,0
	sort.Slice(H,func(i,j int)bool{return H[i]<H[j]})
	sort.Slice(W,func(i,j int)bool{return W[i]<W[j]})
	for i:=0;i+1<N;i+=2 { left = append(left,left[nl]+abs(H[i+1]-H[i])); nl++ }
	for i:=N-1;i-1>=0;i-=2 { right = append(right,right[nr]+abs(H[i]-H[i-1])); nr++	}
	widx := 0; best := 1_000_000_000_000_000_000
	for i:=0;i<N;i+=2 {
		for widx < M-2 { if W[widx+1] >= H[i] { break }; widx++ }
		cand := abs(H[i]-W[widx])
		if widx+1 < M  { cand = min(cand,abs(H[i]-W[widx+1])) }
		cand += left[i/2] + right[(N-1-i)/2]
		best = min(cand,best)
	}
	fmt.Println(best)
}




