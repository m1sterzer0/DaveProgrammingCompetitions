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
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi()
		K := gi()
		S := gs()
		lookup := make([]int,26); for i:=0;i<26;i++ { lookup[i] = i }; base := 0
		for _,c := range S {
			if lookup[c-'a'] == 0 { continue }
			if int(c-'a') < K  { for b:=base+1;b<=int(c-'a');b++ { lookup[b] = 0 }; base = int(c-'a'); continue }
			if int(c-'a') == K { for b:=base+1;b<=int(c-'a');b++ { lookup[b] = 0 }; break }
			movesLeft := K - base; targLetter := int(c-'a')-movesLeft; for b:=targLetter;b<=int(c-'a');b++ { lookup[b] = targLetter }; break
		}
		ansarr := make([]byte,N); for i,c := range S { ansarr[i] = 'a' + byte(lookup[int(c-'a')]) }
		fmt.Fprintln(wrtr,string(ansarr))
	}
}

