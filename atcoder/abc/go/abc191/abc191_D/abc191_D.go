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
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	X := gf(); Y := gf(); R := gf()
	if X < 0 { X = -X }; if Y < 0 { Y = -Y }
	XX := int(10000.0*X+0.5)
	YY := int(10000.0*Y+0.5)
	RR := int(10000.0*R+0.5)
	x1 := (XX+9999) / 10000 * 10000; x2 := x1-10000

	ans := 0
	// First pass -- going to right from x1
	ya := (YY + RR + 9999) / 10000 * 10000 + 30000; yb := (YY - RR - 9999) / 10000 * 10000 - 30000
	for {
		for ya >= YY { if (ya-YY)*(ya-YY) + (x1-XX)*(x1-XX) <= RR * RR { break }; ya -= 10000 }
		for yb <= YY { if (yb-YY)*(yb-YY) + (x1-XX)*(x1-XX) <= RR * RR { break }; yb += 10000 }
		if yb > ya { break }
		ans += ((ya-yb)/10000+1)
		x1 += 10000
	}
	// Second pass -- going left from x2
	ya = (YY + RR + 9999) / 10000 * 10000 + 30000; yb = (YY - RR - 9999) / 10000 * 10000 - 30000
	for {
		for ya >= YY { if (ya-YY)*(ya-YY) + (x2-XX)*(x2-XX) <= RR * RR { break }; ya -= 10000 }
		for yb <= YY { if (yb-YY)*(yb-YY) + (x2-XX)*(x2-XX) <= RR * RR { break }; yb += 10000 }
		if yb > ya { break }
		ans += ((ya-yb)/10000+1)
		x2 -= 10000
	}

	fmt.Println(ans)
}



