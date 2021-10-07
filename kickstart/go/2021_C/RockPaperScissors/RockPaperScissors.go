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
	dp := [61][61][61]float64{}
	choice := [61][61][61]byte{}
	T := gi()
	gi() // Throwaway X value
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		W,E := gi2(); fw := float64(W); fe := float64(E)
		var ch byte
		for sum:=60;sum>=0;sum-- {
			fsum := float64(sum)
			for r:=sum;r>=0;r-- {
				for p:=sum-r;p>=0;p-- {
					s:= sum-r-p
					if sum == 60 { 
						dp[r][p][s] = 0.00
					} else if sum == 0 {
						dp[r][p][s] = 1.0/3.0 * (fw + fe) + dp[1][0][0]
						choice[r][p][s] = 'R'
					} else {
						// Cand 1: play rock
						best := float64(W*p+E*s)/fsum + dp[r+1][p][s]; ch = 'R'
						// Cand 2: play paper
						cand := float64(W*s+E*r)/fsum + dp[r][p+1][s]
						if cand > best { best = cand; ch = 'P'}
						// Cand 3: play scissors
						cand = float64(W*r+E*p)/fsum + dp[r][p][s+1]
						if cand > best { best = cand; ch = 'S'}
						dp[r][p][s] = best; choice[r][p][s] = ch
					}
				}
			}
		}
		ans := make([]byte,60)
		r,p,s := 0,0,0;
		for i:=0;i<60;i++ {
			ch := choice[r][p][s]; ans[i] = ch
			if ch == 'R' { r++ } else if ch == 'P' { p ++ } else { s++ }
		}
		ansstr := string(ans)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ansstr)
    }
}

