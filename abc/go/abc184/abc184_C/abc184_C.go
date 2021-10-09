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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func abs(a int) int { if a >= 0 { return a }; return -a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	r1,c1,r2,c2 := gi4(); ans := 3
	if r1 == r2 && c1 == c2 { 
		ans = 0
	} else if abs(r1-r2)+abs(c1-c2)  <= 3 || r1+c1 == r2+c2 || r1-c1 == r2-c2 { 
		ans = 1 
	} else if abs(r1-r2)+abs(c1-c2)  <= 6 || abs(r1+c1-r2-c2) <= 3 || abs(r1-c1-r2+c2) <= 3 || (r1+c1)%2 == (r2+c2)%2 { 
		ans = 2
	}
	fmt.Println(ans)
}



