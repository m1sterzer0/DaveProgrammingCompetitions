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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func vecset(a []int, v int) { la := len(a); for i:=0;i<la;i++ { a[i] = v } }
const MOD = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); mask := ia(10); for i:=0;i<M;i++ { u,v := gi(),gi(); mask[u] |= 1 << uint(v); mask[v] |= 1 << uint(u) }
	f := ia(520); g := ia(520); f[0] = 1
	for i:=1;i<=N;i++ {
		vecset(g,0)
		for j:=0;j<(1<<9);j++ { 
			if f[j] != 0 {
				for k:=0;k<10;k++ {
					if (j >> uint(k)) & 1 == 0 {
						g[(j | ((1<<uint(k))-1)) & mask[k]] += f[j]
						g[(j | ((1<<uint(k))-1)) & mask[k]] %= MOD
					}
				}

			}
		}
		copy(f,g)
	}
	ans := 0
	for i:=0;i<(1<<9);i++ { ans += f[i]; ans %= MOD }
	fmt.Println(ans)
}

