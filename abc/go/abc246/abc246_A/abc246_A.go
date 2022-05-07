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
	x1,y1,x2,y2,x3,y3 := gi(),gi(),gi(),gi(),gi(),gi()
	x:=0; if x1==x2 { x = x3 } else if x1==x3 { x = x2 } else if x2==x3 { x = x1 }
	y:=0; if y1==y2 { y = y3 } else if y1==y3 { y = y2 } else if y2==y3 { y = y1 }
	fmt.Printf("%v %v\n",x,y)
}

