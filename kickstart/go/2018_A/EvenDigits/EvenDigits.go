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
		var solvecase func(int,int) int
		solvecase = func (n int, dir int) int {
			nn,pv := n,1
			for 10*pv <= nn { pv *= 10 }
			for pv > 0 {
				digit := nn / pv
				if digit % 2 != 0 {
					if dir == 0 {
						adder := (digit+1) * pv - nn
						return adder + solvecase(n+adder,dir)
					} else {
						subtrahend := nn - digit * pv + 1
						return subtrahend + solvecase(n-subtrahend,dir)
					}
				}
				nn -= digit * pv; pv /= 10
			}
			return 0
		}
		ans1 := solvecase(N,0)
		ans2 := solvecase(N,1)
		ans := min(ans1,ans2)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

