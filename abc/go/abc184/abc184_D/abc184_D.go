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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	A,B,C := gi3()
	dp := [101][101][101]float64{}
	for a:=100;a>=0;a-- {
		for b:=100;b>=0;b-- {
			for c:=100;c>=0;c-- {
				if a == 100 || b == 100 || c == 100 { dp[a][b][c] = 0; continue}
				if a == 0 && b == 0 && c == 0 { dp[a][b][c] = 0; continue }
				s := a+b+c
				dp[a][b][c] = 1.0 + (float64(a)*dp[a+1][b][c] + float64(b)*dp[a][b+1][c] + float64(c)*dp[a][b][c+1]) / float64(s)
			}
		}
	}
	fmt.Println(dp[A][B][C])
}



