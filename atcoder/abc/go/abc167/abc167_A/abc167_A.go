package main

import (
	"bufio"
	"fmt"
	"os"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func gts() string { var a string; fmt.Fscan(rdr,&a); return a }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewReaderSize(f, BUFSIZE) }
	
    // NON-BOILERPLATE STARTS HERE
	S := gts(); T := gts()
	n := len(S)
	ans := terns(T[:n]==S,"Yes","No") 
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



