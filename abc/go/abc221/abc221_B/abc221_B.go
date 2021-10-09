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
	S := gs(); T := gs()
	good := true; used := false
	for i:=0;i<len(S);i++ {
		if S[i] == T[i] { continue }
		if S[i] != T[i] && i+1 < len(S) && S[i] == T[i+1] && S[i+1] == T[i] && !used { used = true; i++; continue }
		good = false; break
	}
	if good { fmt.Println("Yes") } else { fmt.Println("No") }
}

