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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewReaderSize(f, BUFSIZE) }
	
    // NON-BOILERPLATE STARTS HERE
	N,K := gti(),gti()
	A := make([]int,N)
	for i:=0;i<N;i++ { A[i] = gti()-1 }

	// Iterate until we find the loop
	ans := 0
	i2c := make([]int,N)
	c2i := make([]int,N); for i:=0;i<N;i++ { c2i[i] = -1 }
	c2i[0] = 0; i2c[0] = 0; city := 0
	for i:=1;;i++ {
		city = A[city]
		if i == K { ans = city+1; break }
		if c2i[city] == -1 { c2i[city] = i; i2c[i] = city; continue }
		loopsize := i - c2i[city]
		endidx := c2i[city] + (K-i) % loopsize
		ans = i2c[endidx]+1
		break
	}
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



