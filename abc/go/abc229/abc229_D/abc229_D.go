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
	S := gs(); K := gi(); L := len(S)
	check := func(m int) bool {
		dots := 0
		for i:=0;i<m;i++ { if S[i] == '.' { dots++ } }
		if dots <= K { return true }
		for i:=m;i<L;i++ {
			if S[i]   == '.' { dots++ }
			if S[i-m] == '.' { dots-- }
			if dots <= K { return true }
		}
		return false
	}
	l,u := 0,L+1
	for u-l > 1 {
		m := (l+u)>>1
		if check(m) { l = m } else { u = m }
	}
	fmt.Println(l)
}

