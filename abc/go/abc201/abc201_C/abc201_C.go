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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	S := gs(); ans := 0
	for i:=0;i<10000;i++ {
		ss := strconv.Itoa(i); for len(ss) < 4 { ss = "0" + ss }
		good := true
		for i,c := range S {
			tc := byte('0' + i)
			if c == '?' { continue }
			if c == 'x' { if ss[0] == tc || ss[1] == tc || ss[2] == tc || ss[3] == tc { good = false; break } }
			if c == 'o' { if ss[0] != tc && ss[1] != tc && ss[2] != tc && ss[3] != tc { good = false; break } }
		}
		if good { ans++ }
	}
	fmt.Println(ans)
}



