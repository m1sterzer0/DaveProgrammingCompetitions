package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
const inf int = 2000000000000000000
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanLines);	rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
		best,bestcnt := "",0; ca := make([]bool,26)
		for _,s := range S {
			for i:=0;i<26;i++ { ca[i] = false }; lcnt := 0
			for _,c := range s {
				if c >= 'A' && c <= 'Z' { cc:=int(c-'A'); if !ca[cc] { lcnt++; ca[cc] = true } }
			}
			if lcnt > bestcnt || lcnt == bestcnt && s < best { best = s; bestcnt = lcnt }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
    }
}

