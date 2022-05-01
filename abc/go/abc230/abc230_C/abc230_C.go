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
	_,A,B,P,Q,R,S := gi(),gi(),gi(),gi(),gi(),gi(),gi()
	ansarr := make([][]byte,Q-P+1); for i:=0;i<Q-P+1;i++ { ansarr[i] = make([]byte,S-R+1) }
	for i:=0;i<Q-P+1;i++ {
		for j:=0;j<S-R+1;j++ {
			x,y := P+i,R+j
			if x+y == A+B || x-y == A-B { ansarr[i][j] = '#' } else { ansarr[i][j] = '.' }
		}
		fmt.Fprintln(wrtr,string(ansarr[i]))
	}
}

