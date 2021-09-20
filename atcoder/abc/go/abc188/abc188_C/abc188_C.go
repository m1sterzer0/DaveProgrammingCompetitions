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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a < b { return a }; return b }
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); A := gis(1<<N)
	g := iai(1<<N,0); for i:=0;i<1<<N;i++ { g[i] = A[i] }
	rdsize := 1<<N
	for xx:=0;xx<N-1;xx++ {
		for i,ri:=0,0;i<rdsize;i,ri=i+2,ri+1 {
			g[ri] = max(g[i],g[i+1])
		}
		rdsize = rdsize>>1
	}
	targ := min(g[0],g[1])
	ans := -1; for i:=0;i<1<<N;i++ { if A[i] == targ { ans = i+1 } }
	fmt.Println(ans)
}



