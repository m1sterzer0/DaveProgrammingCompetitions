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

	doInsert := func(ref string, c byte) string {
		for i,c2 := range(ref) {
			if c == '0' && i == 0 { continue }
			if byte(c2) > c {
				if i == 0 { return string(c) + ref } else { return ref[0:i] + string(c) + ref[i:] }  
			}
		}
		return ref + string(c)
	}

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		ref := gs()
		sumdig := 0
		for _,c := range ref { sumdig += int(byte(c)-'0') }
		ans := "A"
		for i:=0;i<=9;i++ {
			if (sumdig+i) % 9 != 0 { continue }
			res := doInsert(ref,'0'+byte(i))
			if res < ans { ans = res }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

