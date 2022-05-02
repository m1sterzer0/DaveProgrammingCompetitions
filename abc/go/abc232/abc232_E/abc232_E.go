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
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W,K := gi(),gi(),gi(); x1,y1,x2,y2 := gi(),gi(),gi(),gi()
	dpnomat,dprmat,dpcmat,dpmat := 0,0,0,0
	if x1==x2 && y1==y2 { dpmat++ } else if x1==x2 { dprmat++ } else if y1==y2 { dpcmat++ } else { dpnomat++ }
	for i:=0;i<K;i++ {
		ndpmat  :=  (dpnomat * 0       + dprmat * 1     + dpcmat * 1     + dpmat * 0)     % MOD
		ndprmat :=  (dpnomat * 1       + dprmat * (W-2) + dpcmat * 0     + dpmat * (W-1)) % MOD
		ndpcmat :=  (dpnomat * 1       + dprmat * 0     + dpcmat * (H-2) + dpmat * (H-1)) % MOD
		ndpnomat := (dpnomat * (W+H-4) + dprmat * (H-1) + dpcmat * (W-1) + dpmat * 0)     % MOD
		dpmat,dprmat,dpcmat,dpnomat = ndpmat,ndprmat,ndpcmat,ndpnomat
	}
	fmt.Println(dpmat)
}

