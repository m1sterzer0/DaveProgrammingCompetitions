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
		S := gs(); ansarr := make([]byte,0)
		curlevel := 0; n := len(S)
		for i:=0;i<n;i++ {
			d := int(S[i]-'0')
			for curlevel < d { ansarr = append(ansarr,'('); curlevel++ }
			for curlevel > d { ansarr = append(ansarr,')'); curlevel-- }
			ansarr = append(ansarr,S[i])
		}
		for curlevel > 0 { ansarr = append(ansarr,')'); curlevel-- }
		ans := string(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

