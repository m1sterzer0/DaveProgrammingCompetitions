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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gi2() (int,int) { return gi(),gi() }
func min(a,b int) int { if a > b { return b }; return a }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func solve(N,X int, A[]int) int {
	dp := [101][101][101]int{}
	for i,a := range A {
		for numelem:=i;numelem>=0;numelem-- {
			for modulus:=1;modulus<=N;modulus++ {
				for rem:=0;rem<modulus;rem++ {
					oldv := dp[numelem][modulus][rem]
					if oldv == 0 && numelem > 0 { continue }
					newv := oldv + a
					newrem := newv % modulus
					if newv > dp[numelem+1][modulus][newrem] {dp[numelem+1][modulus][newrem] = newv }
				}
			}
		}
	}
	ans := 2*powint(10,18)
	for k:=1;k<=N;k++ {
		best := dp[k][k][X%k]
		if best > 0 {
			ans = min(ans,(X-best) / k)
		}
	}
	return ans
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,X := gi2()
	A := gis(N)
	ans := solve(N,X,A)
	fmt.Println(ans)
	//test(1000,1,10,1,1000,10000,1_000_000)
}



