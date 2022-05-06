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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K := gi(),gi(); A := gis(N)
	aidx := iai(N,-1)
	acnt := iai(N,0)
	rcnt := 0; aidx[0] = 0; idx := 0; loopFound := false
	for idx < K {
		rcnt += A[rcnt % N]; idx++
		if aidx[rcnt % N] >= 0 && !loopFound { // Shortcut around the loop
			loopFound = true
			loopLen := idx - aidx[rcnt % N]
			loopAmt := rcnt - acnt[rcnt % N]
			loopCnt := (K-idx) / loopLen
			rcnt += loopAmt*loopCnt
			idx += loopCnt*loopLen
		} else {
			aidx[rcnt % N] = idx; acnt[rcnt % N] = rcnt
		}
	}
	fmt.Println(rcnt)
}


