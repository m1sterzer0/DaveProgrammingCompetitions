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
func min(a,b int) int { if a > b { return b }; return a }
const inf = 1000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); S := gs()
		dp0,dp1,ans := 0,0,0
		for i:=0;i<N;i+=2 {
			s := S[i:i+2]
			if i == 0 {
				if s == "00" { 
					dp0,dp1 = 1,inf
				} else if s == "11" { 
					dp0,dp1 = inf,1
				} else { 
					ans++; dp0,dp1 = 1,1
				}
			} else {
				if s == "00" { 
					dp0,dp1 = min(dp0,1+dp1),inf
				} else if s == "11" { 
					dp0,dp1 = inf,min(dp0+1,dp1)
				} else { 
					ans++;
					dp0,dp1 = min(dp0,dp1+1),min(dp1,dp0+1)
				}
			}
		}
		fmt.Fprintf(wrtr,"%v %v\n",ans,min(dp0,dp1))
	}
}
