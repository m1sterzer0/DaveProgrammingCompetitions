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
	A,B,C,D,E,F,X := gi(),gi(),gi(),gi(),gi(),gi(),gi()
	t := B*A * (X/(A+C)); a := E*D * (X/(D+F))
	if X % (A+C) >= A { t+=B*A } else { t += (X%(A+C))*B }
	if X % (D+F) >= D { a+=E*D } else { a += (X%(D+F))*E }
	ans := "Draw"; if t > a { ans = "Takahashi" }; if a > t { ans = "Aoki" }
	fmt.Println(ans)
}

