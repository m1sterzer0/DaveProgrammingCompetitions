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
const MOD int64 = 998244353
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N,K := gi(),gi()
		V := gis(N)
		bad := false; ans := int64(1)
		// First, the last K entries should be the last entries of P in sorted order, so Vi should be 0 or -1
		for i:=1;i<=K;i++ { ans *= int64(i); ans %= MOD }
		for i,v := range V {
			if i < N-K {
				if v == -1 { ans *= int64(i+K+1); ans %= MOD }
				if v == 0  { ans *= int64(K+1); ans %= MOD }
			} else {
				if v != -1 && v != 0 { bad = true }
			}
		}
		if bad { ans = 0 }
		fmt.Fprintln(wrtr,ans)
	}
}
