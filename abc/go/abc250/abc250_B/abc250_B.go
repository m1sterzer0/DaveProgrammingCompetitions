package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,A,B := gi(),gi(),gi()
	ww := make([]byte,N*B)
	bb := make([]byte,N*B)
	w,b := byte('.'),byte('#')
	for i:=0;i<N;i++ {
		for j:=0;j<B;j++ { ww[B*i+j] = w; bb[B*i+j] = b }
		w,b = b,w
	}
	www := string(ww)
	bbb := string(bb)
	for i:=0;i<N;i++ {
		for j:=0;j<A;j++ { fmt.Fprintln(wrtr,www) }
		www,bbb = bbb,www
	}
}

