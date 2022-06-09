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
func makeCopies(b,c int) int {
	sb := strconv.Itoa(b)
	s := ""; for i:=0;i<c;i++ { s = s + sb }
	res,_ := strconv.Atoi(s)
	return res
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		sn := gs(); N,_ := strconv.Atoi(sn)
		if N == 999999999999999999 { fmt.Fprintln(wrtr,999999999999999999); continue  }
		ans := 9; for 10*ans+9 <= N { ans = 10*ans+9 }
		for digcnt:=1;digcnt<=9;digcnt++ {
			copies := len(sn)/digcnt
			if copies < 2 || digcnt * copies != len(sn) { continue }
			b1,_ := strconv.Atoi(sn[0:digcnt])
			b2 := b1-1
			x1 := makeCopies(b1,copies)
			x2 := makeCopies(b2,copies)
			if x1 <= N && x1 > ans { ans = x1 }
			if x2 <= N && x2 > ans { ans = x2 }
		}
		fmt.Fprintln(wrtr,ans)
	}
}

