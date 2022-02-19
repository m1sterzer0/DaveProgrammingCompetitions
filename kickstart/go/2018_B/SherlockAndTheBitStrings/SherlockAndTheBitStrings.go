package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func min(a,b int) int { if a > b { return b }; return a }
type constraint struct { a,b,c int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// dp[i][j] = number of valid strings starting at position i with 16 bit prefix j
		// Can calculate dp[i][*] from dp[i+1][*] by trying both possibilities and checking any constraints
		// that start at position i
		// After dp is constructed, feed forward to find the kth answer
		N,K,P := gi3(); A,B,C := fill3(K); for i:=0;i<K;i++ { A[i]--; B[i]-- }
		carr := make([]constraint,K)
		for i:=0;i<K;i++ { carr[i] = constraint{A[i],B[i],C[i]} }
		sort.Slice(carr,func(i,j int) bool { return carr[i].a > carr[j].a} )
		dp := twodi(N+1,1<<16,0); dp[N][0] = 1
		sb := make([]bool,1<<16); cidx := 0

		checkConst := func(mask int,cc constraint) bool {
			newmask := mask >> uint(15 - (cc.b-cc.a))
			return bits.OnesCount(uint(newmask)) == cc.c
		}

		for i:=N-1;i>=0;i-- {
			for cidx < K && carr[cidx].a > i { cidx++ }
			for j:=0;j<1<<16;j++ { sb[j] = true }
			for c := cidx; c < K && carr[c].a == i; c++ {
				for j:=0;j<1<<16;j++ {
					if !sb[j] { continue }
					res := checkConst(j,carr[c])
					sb[j] = sb[j] && res
				}
			}
			for j:=0;j<1<<16;j++ {
				if !sb[j] { continue }
				idx1 := (j << 1) & (0xffff)
				idx2 := idx1 | 1
				dp[i][j] =  min(P,dp[i+1][idx1] + dp[i+1][idx2]) // to prevent overflow
			}
		}
		// Reconstruct the final answer
		// First step is to find the first 16 bit prefix
		ansarr := make([]byte,0); left := P; lastj := 0
		for j:=0;j<1<<16;j++ {
			if dp[0][j] < left { left -= dp[0][j]; continue }
			for i:=0;i<16 && len(ansarr) < N;i++ {
				idx := uint(15-i)
				ansarr = append(ansarr,'0' + byte((j >> idx) & 1))
			}
			lastj = j
			break
		}

		idx := 1 
		for len(ansarr) < N {
			newj0 := (lastj << 1) & 0xffff; newj1 := newj0 | 1
			if dp[idx][newj0] >= left {
				ansarr = append(ansarr,'0'); lastj = newj0; idx++
			} else {
				left -= dp[idx][newj0]; ansarr = append(ansarr,'1'); lastj = newj1; idx++
			}
		}
		ans := string(ansarr)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

