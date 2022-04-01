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
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,P := gi2(); prefixes := make([]string,P); for i:=0;i<P;i++ { prefixes[i] = gs() }
		redundant := make([]bool,P)
		for i:=0;i<P;i++ {
			li := len(prefixes[i])
			for j:=i+1;j<P;j++ {
				lj := len(prefixes[j])
				if li < lj && prefixes[j][0:li] == prefixes[i] { redundant[j] = true }
				if lj < li && prefixes[i][0:lj] == prefixes[j] { redundant[i] = true }
			}
		}
		ans := 1 << uint(N)
		for i:=0;i<P;i++ {
			if redundant[i] { continue }
			ans -= 1 << uint(N - len(prefixes[i]))
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

