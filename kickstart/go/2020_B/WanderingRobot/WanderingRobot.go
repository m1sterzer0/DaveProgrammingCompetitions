package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func doit(n,l int) float64 {
	ans := 0.00
	logdenom := float64(n) * math.Log(2.0)
	lognum := 0.00
	for i:=0;i<=l;i++ {
		if i > 0 { lognum += math.Log(float64(n-i+1)) - math.Log(float64(i)) }
		ans += math.Exp(lognum-logdenom)
	} 
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		W,H,L := gi3(); U,R,D := gi3(); ans := 0.00
		if R < W && U > 1 { ans += doit(U-1+R-1,U-2) }
		if D < H && L > 1 { ans += doit(L-1+D-1,L-2) }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

