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
	A,B,C,D := gi(),gi(),gi(),gi()
	n := 210; s := make([]bool,n+1)
	for i:=2;i<=n;i++ { s[i] = true }
	for i:=4;i<=n;i+=2 { s[i] = false }
	for i:=3;i*i<=n;i+=2 { 
		if !s[i] { continue }
		for j:=i*i;j<=n;j+=2*i { s[j] = false }
	}
	checkAoki := func(t int) bool {
		for x:=C;x<=D;x++ { if s[t+x] { return true } }
		return false
	}
	ans := "Aoki"
	for t:=A;t<=B;t++ {
		if !checkAoki(t) { ans = "Takahashi"; break }
	}
	fmt.Println(ans)
}

