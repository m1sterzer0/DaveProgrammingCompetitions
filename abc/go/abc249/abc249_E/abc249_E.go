package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Let dp[i][j] = the number of ways to generate a string of length i with a RLE encoded string T of length j
	// Then we have dp[i][j] = 25 * sum over k dp[i-k][j-len(k)]
	// To facilitate the sum, we can deal with cumulative sums to make the calculation of dp[1][j] as O(log_10(k)).
	// At the end, we need to multiply by 26/25 to deal with the choice of the first element.
	N,P   := gi(),gi()
	// We really only need to keep 5 columns, but it is just more convenient to keep them all
	dp    := twodi(N+1,N+1,0); dp[0][0] = 1
	sumdp := twodi(N+1,N+1,0); for i:=0;i<=N;i++ { sumdp[i][0] = 1 }
	pow10 := []int{1,10,100,1000,10000,100000}
	ans := 0
	for t:=2;t<N;t++ {
		for i:=0;i<=N;i++ {
			for mt:=2;mt<=5 && t-mt>=0;mt++ {
				low := max(0,i - (pow10[mt-1]-1))
				high := i - (pow10[mt-2])
				if high < 0 { break }
				dp[i][t] += sumdp[high][t-mt]
				if low > 0 { dp[i][t] += (P - sumdp[low-1][t-mt]) }
			}
			dp[i][t] = (dp[i][t] % P * 25 % P )
			sumdp[i][t] = dp[i][t]; if i > 0 { sumdp[i][t] += sumdp[i-1][t]; if sumdp[i][t] >= P { sumdp[i][t] -= P } }
		}
		ans += dp[N][t]
	}
	ans %= P; ans *= 26 * powmod(25,P-2,P) % P; ans %= P // Multiply by 26/25 for the first choice.
	fmt.Println(ans)
}

