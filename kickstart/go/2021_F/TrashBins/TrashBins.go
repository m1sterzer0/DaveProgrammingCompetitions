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
func ia(m int) []int { return make([]int,m) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); S := gs()
		t := ia(0); for i:=0;i<N;i++ { if S[i] == '1' { t = append(t,i) } }
		ans := 0
		n1 := t[0]; ans += n1 * (n1+1) / 2
		n2 := N-1 - t[len(t)-1]; ans += n2 * (n2+1) / 2
		for i:=0;i<len(t)-1;i++ {
			n3 := t[i+1]-t[i]-1
			n3a := n3/2; n3b := n3 - n3a
			ans += n3a * (n3a+1) / 2
			ans += n3b * (n3b+1) / 2
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

