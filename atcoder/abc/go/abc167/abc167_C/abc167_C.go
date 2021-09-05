package main

import (
	"bufio"
	"fmt"
	"os"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewReaderSize(f, BUFSIZE) }
	
    // NON-BOILERPLATE STARTS HERE
	N,M,X := gti(),gti(),gti()
	C := [15]int{}
	A := [15][15]int{}
	for i:=0;i<N;i++ { C[i] = gti(); for j:=0;j<M;j++ { A[i][j] = gti() } }
	myinf := 1_000_000_000_000_000_000
	best := myinf
	for bm:=0;bm<(1<<N);bm++ {
		good := true
		cost := 0
		for j:=0; j<M;j++ {
			mast := 0
			for i:=0;i<N;i++ {
				if bm & (1<<i) == 0 { continue }
				if j==0 { cost += C[i] } //Once per book
				mast += A[i][j]
			}
			if mast < X { good = false; break }
		}
		if good { best = min(best,cost) }
	}
	ans := tern(best==myinf,-1,best,)
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



