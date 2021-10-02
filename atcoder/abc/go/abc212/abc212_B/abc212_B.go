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
	S := gs(); ans := "Strong"
	SS := []int{0,0,0,0}; for i:=0;i<4;i++ { SS[i] = int(S[i]-'0') }
	if SS[0] == SS[1] && SS[0] == SS[2] && SS[0] == SS[3] { ans = "Weak" }
	if SS[1] == (SS[0]+1)%10 && SS[2] == (SS[1]+1)%10 && SS[3] == (SS[2]+1)%10 { ans = "Weak" }
	fmt.Println(ans)
}

