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
var dp [11][11][11]int
var ndp [11][11][11]int
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi()
	// Do the initial state
	dp[M][M][M] = 1
	for i:=0;i<N;i++ {
		for i1:=0;i1<=M;i1++ { for i2:=0;i2<=M;i2++ { for i3:=0;i3<=M;i3++ { ndp[i1][i2][i3] = 0 } } }
		for i1:=0;i1<=M;i1++ {
			for i2:=i1;i2<=M;i2++ {
				for i3:=i2;i3<=M;i3++ {
					for v:=0;v<M;v++ {
						if v <= i1 { 
							ndp[v][i2][i3] += dp[i1][i2][i3]
						} else if v <= i2 {
							ndp[i1][v][i3] += dp[i1][i2][i3]
						} else if v <= i3 {
							ndp[i1][i2][v] += dp[i1][i2][i3]
						}
					}

				}
			}
		}
		for i1:=0;i1<=M;i1++ { for i2:=0;i2<=M;i2++ { for i3:=0;i3<=M;i3++ { dp[i1][i2][i3] = ndp[i1][i2][i3] % MOD } } }
	}
	ans := 0
	for i1:=0;i1<M;i1++ {
		for i2:=i1+1;i2<M;i2++ {
			for i3:=i2+1;i3<M;i3++ {
				ans += dp[i1][i2][i3]
			}
		}
	}
	ans %= MOD
	fmt.Println(ans)
}
