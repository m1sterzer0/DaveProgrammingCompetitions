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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()

		ans := N-1
		for i:=min(N-1,1000000);i>=2;i-- {
			m,n := 1,1
			for N/m >= i && n < N { m*=i; n+=m }
			if n == N { ans = min(ans,i) }
		}
		if ans > 1000000 && N > 1 + 1000000 + 1000000000000 {
			l,u := 1000000,1000000000
			for u-l > 1 {
				m := (l+u)>>1
				if 1+m+m*m <= N { l = m } else { u = m }
			}
			if 1+l+l*l == N { ans = min(ans,l) }
		}
	
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

