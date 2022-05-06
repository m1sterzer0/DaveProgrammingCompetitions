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
	gi(); X := gi(); S := gs()
	lim := 1000000000000000000
	x,extra := X,0
	for _,c := range S {
		if extra > 0 {
			if c == 'L' { extra++ } else if c == 'R' { extra++ } else { extra-- }
		} else if x > lim {
			if c == 'L' { extra++ } else if c == 'R' { extra++ } else { x /= 2 }
		} else {
			if c == 'L' { x *= 2 } else if c == 'R' { x *= 2; x++ } else { x /= 2 }
		}
	}
	fmt.Println(x)
}

