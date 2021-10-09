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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	A,B,K := gi3()
	comb := twodi(61,61,0)
	for i:=0;i<=60;i++ { 
		comb[i][0] = 1; comb[i][i] = 1
		for j:=1;j<i;j++ { comb[i][j] = comb[i-1][j-1] + comb[i-1][j] } 
	}
	bs := make([]byte,A+B); a,b,k := A,B,K
	for i:=0;i<A+B;i++ {
		if a > 0 && comb[a-1+b][b] >= k { bs[i] = 'a'; a -= 1 } else { bs[i] = 'b'; k -= comb[a-1+b][b]; b -= 1 }
	}
	ans := string(bs)
	fmt.Println(ans)
}



