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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
const MOD int = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,K := gi3(); A := gis(N); X,Y := fill2(M); for i:=0;i<M;i++ { X[i]--; Y[i]-- }
	mat1,mat2,mat3 := twodi(N,N,0),twodi(N,N,0),twodi(N,N,0)
	for i:=0;i<N;i++ { mat1[i][i] = 2*M; mat2[i][i] = 1 }
	for i:=0;i<M;i++ { x,y := X[i],Y[i]; mat1[x][x]--; mat1[y][y]--; mat1[x][y]++; mat1[y][x]++ }
	k := K
	for k > 0 {
		if k & 1 == 1 { matmul(N,mat1,mat2,mat3); mat2,mat3 = mat3,mat2 }
		matmul(N,mat1,mat1,mat3); mat1,mat3 = mat3,mat1
		k >>= 1
	}
	preans := matvecmul(N,mat2,A)
	scalefactor := powmod(powmod(2*M,MOD-2,MOD),K,MOD)
	for _,p := range(preans) { fmt.Println(p * scalefactor % MOD ) }
}
func matmul(N int, A,B,C [][]int) {
	for i:=0;i<N;i++ {
		for j := 0; j < N; j++ {
			v := 0
			for k:=0;k<N;k++ {
				v += A[i][k] * B[k][j] % MOD
			}
			C[i][j] = v % MOD
		}
	}
}
func matvecmul(N int, A [][]int, X []int) []int {
	Y := make([]int,N)
	for i:=0;i<N;i++ {
		v := 0
		for j:=0;j<N;j++ {
			v += A[i][j] * X[j] % MOD
		}
		Y[i] = v % MOD
	}
	return Y
}

