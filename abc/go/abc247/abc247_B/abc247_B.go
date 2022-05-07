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
	N := gi(); S := make([]string,N); T := make([]string,N); for i:=0;i<N;i++ { S[i] = gs(); T[i] = gs() }
	good := true
	for i:=0;i<N;i++ {
		sbad,tbad := false,false
		for j:=0;j<N;j++ { if i == j { continue }; if S[i] == S[j] || S[i] == T[j] { sbad = true }; if T[i] == S[j] || T[i] == T[j] { tbad = true } }
		if sbad && tbad { good = false; break }
	}
	ans := "No"; if good { ans = "Yes" }; fmt.Println(ans)
}

