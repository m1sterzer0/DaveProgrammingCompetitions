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
	S := gs(); N := len(S)
	la,ra := 0,0
	for i:=0;i<N;i++    { if S[i] == 'a' { la++ } else { break } }
	for i:=N-1;i>=0;i-- { if S[i] == 'a' { ra++ } else { break } }
	ans := "Yes"; l,r := la,N-1-ra
	for l < r { if S[l] != S[r] { ans = "No" }; l++; r-- }
	if la > ra { ans = "No" }
	fmt.Println(ans)
}

