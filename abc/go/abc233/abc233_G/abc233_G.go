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
func max(a,b int) int { if a > b { return a }; return b }
func min8(a,b int8) int8 { if a > b { return b }; return a }

var dp [50][50][50][50]int8
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }
	for xinc:=0;xinc<N;xinc++ {
		for yinc:=0;yinc<N;yinc++ {
			for x1:=0;x1+xinc<N;x1++ {
				for y1:=0;y1+yinc<N;y1++ {
					x2,y2 := x1+xinc,y1+yinc
					v := int8(0)
					if xinc == 0 && yinc == 0 {
						if bd[x1][y1] == '#' { v = 1 }
					} else {
						v = int8(max(x2-x1+1,y2-y1+1))
						if x2-x1 >= y2-y1 { for xm:=x1;xm<x2;xm++ { v = min8(v,dp[x1][y1][xm][y2]+dp[xm+1][y1][x2][y2]) } }
						if y2-y1 >= x2-x1 { for ym:=y1;ym<y2;ym++ { v = min8(v,dp[x1][y1][x2][ym]+dp[x1][ym+1][x2][y2]) } }
					}
					dp[x1][y1][x2][y2] = v
				}
			}
		}
	}
	fmt.Println(dp[0][0][N-1][N-1])
}
