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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	S := gs(); Q := gi(); T,K := fill2(Q)
	solveit := func(c,t,k int) int {
		if t >= 64 { threesteps := (t-61)/3; t -= 3*threesteps }
		for t > 61 { c = (c+1)%3; t-- }
		for ;t > 0;t-- {
			sz := 1<<(t-1)
			if k >= sz { c = (c+2)%3; k -= sz } else { c = (c+1)%3 }
		}
		return c
	}
	for i:=0;i<Q;i++ {
		t,k := T[i],K[i]
		ans := 0
		if t >= 61 { 
			ans = solveit(int(S[0]-'A'),t,k-1)
		} else {
			endsize := 1 << t
			idx := (k-1) / endsize
			ans = solveit(int(S[idx]-'A'),t,k-1-idx*endsize)
		}
		ansbyte := 'A' + byte(ans)
		ansstring := string([]byte{ansbyte})
		fmt.Fprintln(wrtr,ansstring)
	}

}

