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
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W := gi(),gi(); bd := make([]string,H); for i:=0;i<H;i++ { bd[i] = gs() }
	i1,j1,i2,j2 := -1,-1,-1,-1;
	for i:=0;i<H;i++ {
		for j:=0;j<W;j++ {
			if bd[i][j] == '-' { continue }
			if i1 < 0 { i1,j1 = i,j } else { i2,j2 = i,j }
		}
	}
	ans := abs(i1-i2) + abs(j1-j2)
	fmt.Println(ans)
}

