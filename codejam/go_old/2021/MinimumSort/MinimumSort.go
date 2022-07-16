package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T,N := gi(),gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		for i:=1;i<=N-1;i++ {
			fmt.Fprintf(os.Stderr,"M %v %v\n",i,N)
			fmt.Fprintf(wrtr,"M %v %v\n",i,N); wrtr.Flush()
			idx := gi()
			if idx != i { 
				fmt.Fprintf(os.Stderr,"S %v %v\n",i,idx)
				fmt.Fprintf(wrtr,"S %v %v\n",i,idx); wrtr.Flush()
				res := gi(); if res != 1 { fmt.Fprintf(os.Stderr,"SOMETHING BAD HAPPENED1\n"); os.Exit(1) }
			}
		}
		fmt.Fprintf(wrtr,"D\n"); wrtr.Flush(); res := gi()
		if res != 1 { fmt.Fprintf(os.Stderr,"SOMETHING BAD HAPPENED2\n"); os.Exit(1) }
	}
}

