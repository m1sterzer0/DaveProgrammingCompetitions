package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	isPalindrome := func(i int) bool {
		s := strconv.Itoa(i)
		i,j := 0,len(s)-1
		for i<j { if s[i] != s[j] { return false }; i++; j-- }
		return true
	}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		A := gi(); ans := 0
		for f:=1;f*f<=A;f++ {
			if A%f != 0 { continue }
			if isPalindrome(f) { ans++ }
			f2 := A/f; if f2 > f && isPalindrome(f2) { ans++ }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

