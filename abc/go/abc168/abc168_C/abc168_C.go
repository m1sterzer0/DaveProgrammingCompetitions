package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func main() {
	defer wrtr.Flush()
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	A,B,H,M := gi(),gi(),gi(),gi()
	// Each minute, hour hand goes 1/720 of way around clock
	// Each minute, minute hand goes 1/60 of way around clock = 12/720 of way around clock
	hrpos,mnpos := H*60+M,12*M
	fa,fb,fdiff := float64(A),float64(B),float64(hrpos-mnpos)
	// Law of cosines for the answer
	ans := math.Sqrt(fa*fa+fb*fb-2.0*fa*fb*math.Cos(math.Pi*2.0*fdiff/720.0))
    fmt.Fprintln(wrtr, ans)
}



