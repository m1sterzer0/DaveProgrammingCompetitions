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
		S := gs(); I,J := gi(),gi()
		ls := len(S); I--; J--
		mi,ri,mj,rj := I/ls,I%ls,J/ls,J%ls
		msum := 0; for _,c := range S { if c == 'B' { msum++ } }
		ans := 0
		if mi == mj {
			for i:=ri;i<=rj;i++ { if S[i] == 'B' { ans++ } }
		} else {
			for i:=ri;i<ls;i++ { if S[i] == 'B' { ans++ } }
			for j:=0;j<=rj;j++ { if S[j] == 'B' { ans++ } }
			ans += msum * (mj-mi-1)
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

