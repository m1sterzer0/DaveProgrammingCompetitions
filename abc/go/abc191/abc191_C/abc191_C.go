package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2(); bd := make([]string,H); for i:=0;i<H;i++ { bd[i] = gs() }
	numsides := 0
	for i:=0;i<H-1;i++ { 
		for j:=0;j<W-1;j++ {
			n := 0
			for di:=0;di<=1;di++ {
				for dj:=0;dj<=1;dj++ {
					if bd[i+di][j+dj] == '#' { n++ }
				}
			}
			if n == 1 || n == 3 { numsides++ }
		}
	}
	fmt.Println(numsides)
}



