package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi64() int64 { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi64()
		m1,m2 := int64(0),int64(0)
		if N % 2 == 1 || N < 4 { 
			m1 = -1
		} else { 
			n1 := N; for n1 % 6 != 0 { n1 -= 4; m1++ }; m1 += n1/6
			n2 := N; for n2 % 4 != 0 { n2 -= 6; m2++ }; m2 += n2/4
		}
		if m1 == -1 { fmt.Fprintln(wrtr,-1) } else { fmt.Fprintf(wrtr,"%v %v\n",m1,m2) }
	}
}

