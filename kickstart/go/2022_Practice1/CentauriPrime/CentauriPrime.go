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

	vowels := []byte{'a','e','i','o','u','A','E','I','O','U'}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		s := gs(); lastlet := byte(s[len(s)-1])
		ans := "Bob"
		if lastlet == 'y' || lastlet == 'Y' { ans = "nobody" }
		for _,c := range(vowels) {
			if lastlet == c { ans = "Alice"}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v is ruled by %v.\n",tt,s,ans)
    }
}

