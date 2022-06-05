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
func gbs() []byte { return []byte(gs()) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		gi(); p := gbs()
		upper,lower,number,special := false,false,false,false
		for _,c := range p { 
			if c >= 'A' && c <= 'Z' { upper = true }
			if c >= 'a' && c <= 'z' { lower = true }
			if c >= '0' && c <= '9' { number = true }
			if c == '#' || c == '@' || c == '*' || c == '&' { special = true }
		}
		if !upper { p = append(p,'A') }
		if !lower { p = append(p,'a') }
		if !number { p = append(p,'0') }
		if !special { p = append(p,'#') }
		for len(p) < 7 { p = append(p,'o') }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,string(p))
    }
}

