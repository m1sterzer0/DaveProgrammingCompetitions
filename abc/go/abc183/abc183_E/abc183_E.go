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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
const MOD = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2(); bd := make([]string,H); for i:=0;i<H;i++ { bd[i] = gs() }
	rowsum := iai(H,0); colsum := iai(W,0); diagsum := iai(H+W-1,0); sb := twodi(H,W,0)
	for i:=H-1;i>=0;i-- {
		for j:=W-1;j>=0;j-- {
			if bd[i][j] == '#' {
				rowsum[i] = 0; colsum[j] = 0; diagsum[W-1+i-j] = 0
			} else if i == H-1 && j == W-1 { 
				sb[i][j] = 1
			} else {
				sb[i][j] = (rowsum[i] + colsum[j] + diagsum[W-1+i-j]) % MOD
			} 
			rowsum[i] = (rowsum[i] + sb[i][j]) % MOD
			colsum[j] = (colsum[j] + sb[i][j]) % MOD
			diagsum[W-1+i-j] = (diagsum[W-1+i-j] + sb[i][j]) % MOD
		}
	}
	fmt.Println(sb[0][0])
}



