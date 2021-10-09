package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	SN := gs(); num1,num2 := 0,0
	for _,c := range(SN) { 
		if c == '1' || c == '4' || c == '7' { num1++ } else if c == '2' || c == '5' || c == '8' { num2++ }
	}
	curans := (num1 + 2*num2) % 3
	ans := -1
	if curans == 0 { 
		ans = 0
	} else if curans == 1 && num1 > 0 && len(SN) > 1 {
		ans = 1
	} else if curans == 2 && num2 > 0 && len(SN) > 1 {
		ans = 1
	} else if curans == 1 && num2 > 1 && len(SN) > 2 {
		ans = 2
	} else if curans == 2 && num1 > 1 && len(SN) > 2 {
		ans = 2
	}
	fmt.Println(ans)
}



