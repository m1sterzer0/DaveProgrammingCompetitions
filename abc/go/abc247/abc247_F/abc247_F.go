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
	N := gi(); P := gis(N); Q := gis(N); for i:=0;i<N;i++ { P[i]--; Q[i]-- }
	// Cards form a series of cycles
	// We must choose cards from the cycle such that there is not a two card gap.  Single card cycle is a corner case
	// Use a DP to calculate the ways 
	cycways := make([]int,N+1); cycways[1] = 1
	dp00,dp01,dp10,dp11 := 1,0,0,1
	for i:=2;i<=N;i++ {
		dp00,dp01,dp10,dp11 = dp01,dp00+dp01,dp11,dp10+dp11; dp01 %= MOD; dp11 %= MOD
		cycways[i] = (dp01+dp10+dp11) % MOD
	}
	PP := make([]int,N); for i,p := range P { PP[p] = i }
	ans := 1
	visited := make([]bool,N)
	for i:=0;i<N;i++ {
		cnt,curs := 0,i
		for !visited[curs] { 
			cnt++
			visited[curs] = true
			curs = PP[Q[curs]]
		}
		if cnt > 0 { ans *= cycways[cnt]; ans %= MOD }
	}
	fmt.Println(ans)
}
