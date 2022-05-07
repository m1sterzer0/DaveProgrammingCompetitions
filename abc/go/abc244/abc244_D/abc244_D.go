package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Even and odd permutations
	s1,s2,s3 := gs(),gs(),gs(); t1,t2,t3 := gs(),gs(),gs()
	s := s1+s2+s3; t := t1+t2+t3
	spar := 1; if s == "RGB" || s == "GBR" || s == "BRG" { spar = 0 }
	tpar := 1; if t == "RGB" || t == "GBR" || t == "BRG" { tpar = 0 }
	if spar == tpar { fmt.Println("Yes") } else { fmt.Println("No") }
}

