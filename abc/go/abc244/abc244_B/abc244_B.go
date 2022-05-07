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
	gi(); T := gs()
	x,y,dx,dy := 0,0,1,0
	dxarr := []int{1,0,-1,0}; dyarr := []int{0,-1,0,1}; idx := 0
	for _,c := range T {
		if c == 'S' { x+=dx;y+=dy } else { idx++; idx%=4; dx=dxarr[idx]; dy=dyarr[idx] }
	}
	fmt.Printf("%v %v\n",x,y)
}

