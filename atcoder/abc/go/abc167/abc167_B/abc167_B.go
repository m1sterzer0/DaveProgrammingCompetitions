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
func tern(cond bool, a int, b int) int { if cond { return a }; return b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewReaderSize(f, BUFSIZE) }
	
    // NON-BOILERPLATE STARTS HERE
	A,B,C,K := gti(),gti(),gti(),gti(); C |= C // to eliminate unused error
	ans := tern(K<=A,K,tern(K<=A+B,A,A-(K-A-B)))
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



