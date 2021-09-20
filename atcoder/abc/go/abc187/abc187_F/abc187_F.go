package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M := gi2(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	gr := make([][]bool,N); for i:=0;i<N;i++ { gr[i] = make([]bool,N) }
	for i:=0;i<M;i++ { a,b := A[i],B[i]; gr[a][b] = true; gr[b][a] = true }
	inf := 1_000_000_000_000_000_000
	dp := iai(1<<N,inf); dp[0] = 0; bmmax := 1<<N
	isclique := func (bm int) bool {
		for i:=0;i<N;i++ {
			if bm & (1<<i) == 0 { continue }
			for j:=i+1;j<N;j++ {
				if bm & (1<<j) == 0 { continue }
				if !gr[i][j] { return false }
			}
		}
		return true
	}
	for bm:=1;bm<bmmax;bm++ {
		msb := 63-bits.LeadingZeros(uint(bm)); msbm := 1<<msb; residual := bm ^ msbm
		if residual == 0 { dp[bm] = 1;  continue }
		if isclique(bm)  { dp[bm] = 1 ; continue }
		v := inf
		x := (bm-1)&bm
		for x != 0 { 
			v2 := dp[x]+dp[bm^x]
			if v2 < v { v = v2 }
			x = (x-1)&bm 
		}
		dp[bm] = v
	}
	fmt.Println(dp[(1<<N)-1])
}



