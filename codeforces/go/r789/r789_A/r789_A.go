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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); A := gis(N)
		sb := make([]bool,101)
		cnt := 0; cntzero := 0
		for _,a := range A {
			if a == 0 { cntzero++ } 
			if !sb[a] { cnt++; sb[a] = true }
		}
		ans := 0
		if cntzero > 0 { 
			ans = N - cntzero // Can make at most one zero in each operation
		} else if cnt < N {
			ans = N // already have a matching pair, so can use that to make first zero, and then one zero per op afterwards
		} else {
			ans = N+1 // need to use min to create matching pair, then same as above
		}
		fmt.Fprintln(wrtr,ans)
	}
}

