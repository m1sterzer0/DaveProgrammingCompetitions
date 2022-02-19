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
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	dp := [19][9]int{}
	dp[0][0] = 1
	for i:=1;i<19;i++ {
		for j:=0;j<=8;j++ {
			for k:=0;k<=8;k++ {
				dp[i][(j+k)%9] += dp[i-1][j]
			}
		}
	}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		F,L := gi2()
		var solveit func(string,int) int
		solveit = func(numstr string, remsofar int) int {
			ans := 0; d := int(numstr[0]-'0'); numdig := len(numstr)
			if len(numstr) == 1 {
				x,_ := strconv.Atoi(numstr)
				ans := 0
				for i:=0;i<=8;i++ { if i <= x && (remsofar+i) % 9 != 0 { ans++ } }
				return ans
			}
			for i:=0;i<=8;i++ {
				for j:=0;j<=8;j++ {
					if i < d && (remsofar + i + j) % 9 != 0 {
						ans += dp[numdig-1][j]
					}
				}
			}
			if d < 9 { ans += solveit(numstr[1:],(remsofar+d)%9) }
			return ans
		}
		ans := solveit(strconv.Itoa(L),0) - solveit(strconv.Itoa(F-1),0)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

