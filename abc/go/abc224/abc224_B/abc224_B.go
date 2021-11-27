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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W := gi2(); A := make([][]int,H); for i:=0;i<H;i++ { A[i] = gis(W) }
	ans := "Yes"
	for i1:=0;i1<H;i1++ {
		for i2:=i1+1;i2<H;i2++ {
			for j1:=0;j1<W;j1++ {
				for j2:=j1+1;j2<W;j2++ {
					if A[i1][j1] + A[i2][j2] > A[i2][j1] + A[i1][j2] { ans = "No" }
				}
			}
		}
	}
	fmt.Println(ans)
}

