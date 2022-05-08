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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); B := 1000000
	s := make([]bool,B+1); s[2] = true; for i:=3;i<=B;i+=2 { s[i] = true }
	for i:=3;i*i<=B;i+=2 { if s[i] { for j:=i*i;j<=B;j+=2*i { s[j] = false } } }
	cnt := ia(B+1)
	cnt[0] = 0; for i:=1;i<=B;i++ { cnt[i] = cnt[i-1]; if s[i] { cnt[i]++ } }
	ans := 0
	for i:=1;i<=B;i++ {
		if !s[i] { continue }
		qqq := i*i*i
		if qqq > N { break }
		pmax := min(i-1,N/qqq)
		ans += cnt[pmax]
	}
	fmt.Println(ans)
}

