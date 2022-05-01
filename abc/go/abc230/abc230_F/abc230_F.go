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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
const MOD = 998244353

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N)
	//N := 4; A := []int{1,2,3,0}
	cum := make([]int,N); cum[0] = A[0]; for i:=1;i<N;i++ { cum[i] = cum[i-1] + A[i] }
	dp := make([]int,N); last := make(map[int]int)
	dp[0] = 1
	for i:=1;i<N;i++ {
		dp[i] = 2 * dp[i-1] % MOD
		idx,ok := last[cum[i-1]]
		if ok {	dp[i] += MOD - dp[idx]; dp[i] %= MOD }
		last[cum[i-1]] = i-1
	}
	ans := dp[N-1]
	fmt.Println(ans)
}

