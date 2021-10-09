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
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); x0,y0 := gi2(); xother,yother := gi2()
	fx0 := float64(x0); fy0 := float64(y0)
	xc,yc := 0.5*float64(x0+xother),0.5*float64(y0+yother)
	ang := 2.0*math.Pi/float64(N)
	xd,yd := fx0-xc,fy0-yc
	xans := xc + xd * math.Cos(ang) - yd * math.Sin(ang)
	yans := yc + xd * math.Sin(ang) + yd * math.Cos(ang)
	fmt.Printf("%v %v\n",xans,yans)
}



