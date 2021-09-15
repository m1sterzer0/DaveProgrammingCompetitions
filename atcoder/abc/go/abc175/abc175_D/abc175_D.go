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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,K := gi(),gi(); P := gis(N); C := gis(N); for i:=0;i<N;i++ { P[i]-- }
	best := -2_000_000_000_000_000_000

	// Possible choices
	// One partial loop
	// N-1 Full loops plus a suffix
	for start:=0; start<N; start++ {
		m := 0
		n := start
		running := 0
		for m < K {
			n = P[n]
			running += C[n]
			m++
			if running > best { best = running }
			if n == start { break }
		}
		if m == K || running <= 0 { continue }
		loopsleft := (K-m) / m
		if loopsleft >= 2 {
			running += (loopsleft-1)*running
			if running > best { best = running }
			m += (loopsleft-1) * m
		}
		for m < K {
			n = P[n]
			running += C[n]
			m += 1
			if running > best { best = running }
		}
	}

	fmt.Println(best)
}



