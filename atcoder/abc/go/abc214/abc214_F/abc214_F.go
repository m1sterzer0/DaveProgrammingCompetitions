package main

import (
	"bufio"
	"fmt"
	"os"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
const MOD int = 1_000_000_007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	S := gs(); 
	N := len(S); last := iai(26,0)
	dp1 := ia(N+1); dp2 := ia(N+1)
	dp2[0] = 1
	// dp1[i+1] is the number of valid substrings through character i such that we mark element i that are not already counted in dp2[i+1]
	// dp2[i+1] is the number of valid substrings through character i such that we do not mark element i
	for i:=0;i<N;i++ {
		cval := int(S[i]-'a')
		dp2[i+1] = (dp1[i] + dp2[i]) % MOD
		dp1[i+1] = (dp2[i] + MOD - last[cval]) % MOD
		last[cval] = dp2[i]
	}
	ans := (dp1[N] + dp2[N] + MOD - 1) % MOD
	fmt.Println(ans)
}

