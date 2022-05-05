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
	// x = a + b, y = a + c, where a^b = a^c = b^c = 0
	// s = x + y = 2a + b + c
	// Might as well take x = a and y = a+b+c = s-a
	T := gi()
	for tt:=1;tt<=T;tt++ {
		a,s := gi(),gi(); ans := "No"; if s >= a && a & (s-a) == a { ans = "Yes" }
		fmt.Fprintln(wrtr,ans)
	}
}

