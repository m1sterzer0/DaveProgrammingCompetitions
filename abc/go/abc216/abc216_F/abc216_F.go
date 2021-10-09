package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
const MOD int = 998244353
type pair struct {a,b int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); B := gis(N)
	q := make([]pair,N); for i:=0;i<N;i++ { q[i] = pair{A[i],B[i]} }
	sort.Slice(q,func(i,j int) bool { return q[i].a < q[j].a || q[i].a == q[j].a && q[i].b < q[j].b })
	dp := ia(5001); dp[0] = 1; ans := 0
	for _,pp := range q {
		a,b := pp.a,pp.b
		if a >= b {	s := 0; for i:=0;i<=a-b;i++ { s += dp[i] }; ans += s; ans %= MOD }
		for i:=5000;i-b>=0;i-- { dp[i] += dp[i-b]; dp[i] %= MOD }
	}
	fmt.Println(ans)
}

