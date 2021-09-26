package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,K := gi2(); 
	dp1 := iai(N+1,0); for i:=1;i<=N;i++ { dp1[i] = 1 }
	cdp1 := iai(N+1,0); s := 0; for i:=1;i<=N;i++ { s += dp1[i]; cdp1[i] = s }
	dp2 := iai(2*N+1,0); for i:=2;i<=2*N;i++ { lb,ub := max(1,i-N),min(i-1,N); dp2[i] = cdp1[ub]-cdp1[lb-1] }
	cdp2 := iai(2*N+1,0); s = 0; for i:=2;i<=2*N;i++ { s += dp2[i]; cdp2[i] = s }
	dp3 := iai(3*N+1,0); for i:=3;i<=3*N;i++ { lb,ub := max(2,i-N),min(i-1,2*N); dp3[i] = cdp2[ub]-cdp2[lb-1] }
	cdp3 := iai(3*N+1,0); s = 0; for i:=3;i<=3*N;i++ { s += dp3[i]; cdp3[i] = s }
	last := 0; s = 0;  beauty := 1; taste := 1; popularity := 1

	for i:=3;i<=3*N;i++ { 
		inc := dp3[i]
		if last+inc >= K { s = i ; break }
		last += inc
	}
	for i:=1;i<=N;i++ { 
		if s-i <= 2*N { 
			inc := dp2[s-i]
			if last+inc >= K { beauty = i; break }
			last += inc
		}
	}
	for i:=1;i<=N;i++ {
		if s-beauty-i <= N {
			inc := 1
			if last+inc >= K { taste = i; break }
			last += inc
		}
	}
	popularity = s - beauty - taste
	fmt.Printf("%v %v %v\n",beauty,taste,popularity)
}
