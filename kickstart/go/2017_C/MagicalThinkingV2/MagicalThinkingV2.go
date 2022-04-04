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
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2(); ans := 0
		if N == 1 {
			p1 := gs(); mine := gs(); score1 := gi()
			tt := 0; tf := 0
			for i:=0;i<Q;i++ { if p1[i] == mine[i] { tt++ } else { tf++ } }
			for a:=0;a<=tt;a++ {
				for b:=0;b<=tf;b++ {
					if a + (tf - b) != score1 { continue }
					ans = max(a+b,ans)
				}
			}
		} else {
			p1 := gs(); p2 := gs(); mine := gs(); score1 := gi(); score2 := gi()
			ttt := 0; tff := 0; ttf := 0; tft := 0
			for i:=0;i<Q;i++ { 
				if p1[i] == mine[i] && p2[i] == mine[i] { ttt++ }
				if p1[i] == mine[i] && p2[i] != mine[i] { ttf++ }
				if p1[i] != mine[i] && p2[i] == mine[i] { tft++ }
				if p1[i] != mine[i] && p2[i] != mine[i] { tff++ }
			}
			for a:=0;a<=ttt;a++ {
				for b:=0;b<=tff;b++ {
					for c:=0;c<=ttf;c++ {
						for d:=0;d<=tft;d++ {
							if a + (tff - b) + c + (tft - d) != score1 { continue }
							if a + (tff - b) + (ttf - c) + d != score2 { continue }
							ans = max(ans,a+b+c+d)
						}
					}
				}
			}
	
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

