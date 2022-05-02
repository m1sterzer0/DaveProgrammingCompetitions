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
	S,T := gs(),gs(); shiftamt := int(T[0])-int(S[0]); if shiftamt < 0 { shiftamt += 26 }
	targ := make([]byte,len(S))
	for i,c := range S { a := (int(c-'a') + shiftamt) % 26; targ[i] = 'a' + byte(a) }
	cand := string(targ); 
	ans := "No"; if cand == T { ans = "Yes" }; fmt.Println(ans)
}

