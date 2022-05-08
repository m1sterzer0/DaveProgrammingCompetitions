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
	S := gs(); m := make(map[byte]bool); hasUpper := false; hasLower := false
	for _,c := range S { if c >= 'a' && c <= 'z' { hasLower = true } else { hasUpper = true }; m[byte(c)] = true }
	ans := "No"; if hasLower && hasUpper && len(m) == len(S) { ans = "Yes" }
	fmt.Println(ans)
}

