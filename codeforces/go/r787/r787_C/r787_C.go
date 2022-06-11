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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		s := gs(); n := len(s)
		l1 := -1; f0 := -1
		// Can't have 101
		for i,c := range(s) { if c == '1' { l1 = i } else if c == '0' && f0 == -1 { f0 = i } }
		ans := 0
		if l1 == -1 && f0 == -1 { 
			ans = n
		} else if l1 == -1 {
			ans = f0+1
		} else if f0 == -1 {
			ans = n - l1
		} else {
			ans = f0-l1+1
		}
		fmt.Fprintln(wrtr,ans)
	}
}
