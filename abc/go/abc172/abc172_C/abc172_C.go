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

func max(a,b int) int { if a > b { return a }; return b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,K := gi(),gi(),gi()
	A := gis(N)
	B := gis(M)
	// First, assume we read as many as we can from A
	i := 0; j := 0; nummin := 0
	for i<N && A[i]+nummin <= K { nummin += A[i]; i++ }
	for j<M && B[j]+nummin <= K { nummin += B[j]; j++ }
	best := i+j
	for i > 0 { 
		i--; nummin -= A[i]
		for j<M && B[j]+nummin <= K { nummin += B[j]; j++ }
		best = max(best,i+j)
	}
	fmt.Println(best)
}



