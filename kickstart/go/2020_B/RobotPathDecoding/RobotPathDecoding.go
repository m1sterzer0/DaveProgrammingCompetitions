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
const MOD int = 1000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	mult := [2010]int{}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S := gs(); x := 0; y := 0; mult[0] = 1; midx := 0
		for _,c := range S {
			if c == 'N' { 
				y += MOD - mult[midx]; y %= MOD
			} else if c == 'S' {
				y += mult[midx]; y %= MOD
			} else if c == 'W' {
				x += MOD - mult[midx]; x %= MOD
			} else if c == 'E' {
				x += mult[midx]; x %= MOD
			} else if c == ')' {
				midx--
			} else if c >= '0' && c <= '9' {
				d := int(c-'0')
				midx++; mult[midx] = mult[midx-1] * d % MOD
			}
		} 
        fmt.Fprintf(wrtr,"Case #%v: %v %v\n",tt,x+1,y+1)
    }
}

