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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W,N,M := gi4(); A,B := fill2(N); C,D := fill2(M); for i:=0;i<N;i++ { A[i]--; B[i]-- }; for i:=0;i<M;i++ { C[i]--; D[i]-- }
	bd := make([][]byte,H); for i:=0;i<H;i++ { bd[i] = make([]byte,W) }; for i:=0;i<H;i++ { for j:=0;j<W;j++ { bd[i][j] = ' ' } }
	for i:=0;i<N;i++ { bd[A[i]][B[i]] = '*' }
	for i:=0;i<M;i++ { bd[C[i]][D[i]] = '#' }
	cnt,state := 0,false
	for i:=0;i<H;i++ { state = false; for j:=0;j<W;j++ { 
		if bd[i][j] == '*' { state = true }; if bd[i][j] == '#' { state = false }; if bd[i][j] == ' ' && state { bd[i][j] = '.'; cnt++ }
	} }
	for i:=0;i<H;i++ { state = false; for j:=W-1;j>=0;j-- { 
		if bd[i][j] == '*' { state = true }; if bd[i][j] == '#' { state = false }; if bd[i][j] == ' ' && state { bd[i][j] = '.'; cnt++ }
	} }
	for j:=0;j<W;j++ { state = false; for i:=0;i<H;i++ { 
		if bd[i][j] == '*' { state = true }; if bd[i][j] == '#' { state = false }; if bd[i][j] == ' ' && state { bd[i][j] = '.'; cnt++ }
	} }
	for j:=0;j<W;j++ { state = false; for i:=H-1;i>=0;i-- { 
		if bd[i][j] == '*' { state = true }; if bd[i][j] == '#' { state = false }; if bd[i][j] == ' ' && state { bd[i][j] = '.'; cnt++ }
	} }
	cnt += N
	fmt.Println(cnt)
}



