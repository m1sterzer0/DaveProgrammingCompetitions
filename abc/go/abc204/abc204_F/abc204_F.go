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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
const MOD int = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	H,W := gi2()
	dp := twodi(1<<H,1<<H,0)
	full := (1<<H)-1
	sb := make([]bool,1<<H)
	for st:=0;st<=full;st++ {
		if st == full { dp[full][0] = 1; continue }
		for i:=0;i<1<<H;i++ { sb[i] = false }
		sb[st] = true
		// First, find all the ways we can deal with the double verticals
		for i:=0;i<H-1;i++ {
			for pos:=1<<H-1;pos>=0;pos-- {
				if !sb[pos] { continue }
				if pos >> i & 1 == 1 { continue }
				if pos >> (i+1) & 1 == 1 { continue }
				sb[pos | 1<<i | 1<<(i+1)] = true
			}
		}
		// Now loop through these starting positions and find out where the right columns can end up
		for pos:=0;pos<1<<H;pos++ {
			if !sb[pos] { continue }
			holes := full ^ pos
			for i:=0;i<1<<H;i++ {
				if i & holes == i { dp[st][i]++ }
			}
		}
	}
	dp2 := matrixExponentiation(1<<H,dp,W,MOD)
	ans := dp2[0][0]
	fmt.Println(ans)
}
func matrixExponentiation(N int,dp [][]int,K int,mod int) [][]int {
	// Mat1 is the multiplier, mat2 is the answer
	mat1 := make([][]int,N); for i:=0;i<N;i++ { mat1[i] = make([]int,N) }
	mat2 := make([][]int,N); for i:=0;i<N;i++ { mat2[i] = make([]int,N) }
	mat3 := make([][]int,N); for i:=0;i<N;i++ { mat3[i] = make([]int,N) }
	for i:=0;i<N;i++ { for j:=0;j<N;j++ { mat1[i][j] = dp[i][j] } }
	for i:=0;i<N;i++ { mat2[i][i] = 1 }
	for K > 0 {
		if K & 1 == 1 { matmul(N,mod,mat1,mat2,mat3); mat2,mat3 = mat3,mat2 }
		matmul(N,mod,mat1,mat1,mat3); mat1,mat3 = mat3,mat1
		K >>= 1
	}
	return mat2
}
func matmul(N,mod int, A,B,C [][]int) {
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			v := 0
			for k:=0;k<N;k++ { v += A[i][k]*B[k][j] % mod }
			C[i][j] = v % mod
		}
	}
}

