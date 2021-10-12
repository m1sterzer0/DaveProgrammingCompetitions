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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
	adj := [16][16]bool{}
	dp := [1<<16]int{}
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,M,K := gi3(); L,R,A := fill3(N); X,Y := fill2(M)
		for i:=0;i<N;i++ {for j:=0; j<N; j++ { adj[i][j] = false } }
		for i:=0;i<M;i++ { adj[X[i]][Y[i]] = true; adj[Y[i]][X[i]] = true }
		ans := 0
		numbm := 1<<uint(N)
		for i:=0;i<numbm;i++ { dp[i] = 0 }
		for i:=0;i<N;i++ { dp[1<<uint(i)] = 1 }
		for bm:=0;bm<numbm;bm++ {
			if dp[bm] == 0 { continue }
			pts := 0
			for i:=0;i<N;i++ { if bm & (1<<uint(i)) != 0 { pts += A[i]} }
			if pts > K { continue }
			if pts == K { ans += dp[bm]; continue }
			for k:=0;k<N;k++ {
				bmk := 1<<uint(k)
				if bm & bmk != 0 { continue }
				if pts + A[k] > K || pts < L[k] || pts > R[k] { continue }
				good := false
				for i:=0;i<N;i++ { if bm & (1<<uint(i)) != 0 && adj[i][k] { good = true; break } }
				if good { dp[bm | bmk ] += dp[bm] }
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

