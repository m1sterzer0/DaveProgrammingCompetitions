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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		K := gi(); A := gis(K); st := 1; dir := 1; ans := 0
		for i:=1;i<K;i++ {
			if A[i] == A[i-1] { continue }
			if A[i] < A[i-1] {
				if dir == 1 { st += 1; if st > 4 { ans += 1; st = 1 } } else { dir = 1; st = 2 }
			} else {
				if dir == -1 { st += 1; if st > 4 { ans += 1; st = 1 } } else { dir = -1; st = 2 }
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

