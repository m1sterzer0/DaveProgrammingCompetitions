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
func min(a,b int) int { if a > b { return b }; return a }
type st struct {idx,amt int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,X := gi(),gi(); A := gis(N)
	var solveit func(idx,amt int) int
	cache := make(map[st]int)
	solveit = func(idx,amt int) int {
		s := st{idx,amt}
		v,ok := cache[s]
		if !ok {
			if idx == 0 {
				v = amt
			} else {
				c := amt / A[idx]
				if A[idx] * c == amt {
					v = c
				} else {
					v1 := c + solveit(idx-1,amt-A[idx]*c)
					v2 := c+1 + solveit(idx-1,A[idx]*(c+1)-amt)
					v = min(v1,v2)
				}
			}
			cache[s] = v
		}
		return v
	}
	ans := solveit(N-1,X)
	fmt.Println(ans)
}

