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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S := gs()
		ansarr := make([]byte,0)
		last := S[0]; l,u := 1,2
		for i:=1;i<len(S);i++ {
			if S[i] == last {
				l++; u += 2
			} else if S[i] < last {
				for j:=0;j<l;j++ { ansarr = append(ansarr,byte(last)) }
				last,l,u = S[i],1,2
			} else {
				for j:=0;j<u;j++ { ansarr = append(ansarr,byte(last)) }
				last,l,u = S[i],1,2
			}
		}
		for j:=0;j<l;j++ { ansarr = append(ansarr,byte(last)) }
		ans := string(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

