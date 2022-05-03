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
type P struct  { x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// I got the concept right, but fouled the bookkeeping up, so transcribing sugarr solution
	N := gi(); A := gis(N)
	b,c := make([]P,0),make([]P,0)
	sb,sc := A[0],A[0]
	b = append(b,P{A[0],1}); c = append(c,P{A[0],1})
	dp := make([]int,N+1)
	dp[0] = 1
	for i:=1;i<=N;i++ {
		dp[i] = (sb-sc+MOD) % MOD
		if i == N { break }

		sm := dp[i]
		for ll:=len(b);ll>0;ll-- {
			top := b[ll-1]
			if top.x > A[i] { break }
			sm += top.y; sm %= MOD
			sb += (A[i]-top.x) * top.y % MOD; sb += MOD; sb %= MOD
			b = b[:ll-1]
		}
		b = append(b,P{A[i],sm})

		sm = dp[i]
		for ll:=len(c);ll>0;ll-- {
			top := c[ll-1]
			if top.x < A[i] { break }
			sm += top.y; sm %= MOD
			sc += (A[i]-top.x) * top.y % MOD; sc += MOD; sc %= MOD
			c = c[:ll-1]
		}
		c = append(c,P{A[i],sm})
	}
	ans := dp[N]
	fmt.Println(ans)
}

