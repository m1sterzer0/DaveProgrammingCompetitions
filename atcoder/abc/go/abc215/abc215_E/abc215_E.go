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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
const MOD int = 998244353
type update struct { a,b,v int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := gs()
	dp := twodi(1<<10,10,0)
	updates := make([]update,0)
	for i:=0;i<N;i++ {
		id := int(S[i]-'A')
		updates = append(updates,update{1<<id,id,1})
		for j:=0;j<1<<10;j++ {
			for k:=0;k<10;k++ {
				if dp[j][k] == 0 { continue }
				if (1<<id) & j != 0 && id != k { continue }
				updates = append(updates,update{j | (1<<id),id,dp[j][k]})
			}
		}
		for _,u := range updates { dp[u.a][u.b] += u.v; dp[u.a][u.b] %= MOD }
		updates = updates[:0]
	}
	ans := 0
	for i:=0;i<1<<10;i++ { for j:=0;j<10;j++ { ans += dp[i][j] } }; ans %= MOD; fmt.Println(ans) 
}

