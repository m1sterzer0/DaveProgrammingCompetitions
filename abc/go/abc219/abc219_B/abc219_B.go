package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	S1,S2,S3,T := gs(),gs(),gs(),gs()
	ansarr := make([]string,len(T))
	for i,c := range T { 
		if c == '1' { ansarr[i] = S1 }
		if c == '2' { ansarr[i] = S2 }
		if c == '3' { ansarr[i] = S3 }
	}
	ans := strings.Join(ansarr,"")
	fmt.Println(ans)
}



