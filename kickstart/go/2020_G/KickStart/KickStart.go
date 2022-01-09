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
	st := make([]int,0)
	en := make([]int,0)
	for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		S := gs()
		st = st[:0]; en = en[:0]
		for i:=0;i<len(S);i++ {
			if i+3 < len(S) && S[i:i+4] == "KICK"  { st = append(st,i) }
			if i+4 < len(S) && S[i:i+5] == "START" { en = append(en,i) }
		}
		numen := len(en); j := 0
		ans := 0
		for _,v := range st {
			for numen > 0 && v > en[j] { numen--; j++ }
			ans += numen
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

