package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	N,K := gi(),gi(); A := gis(N);
	bd := make([][]int,K)
	for i,a := range(A) { bd[i%K] = append(bd[i%K],a) }
	for k:=0;k<K;k++ { sort.Slice(bd[k],func (i,j int) bool { return bd[k][i] < bd[k][j] }) }
	ans := "Yes"; last := -100
	for i:=0;i<N;i++ {
		next := bd[i%K][i/K]
		if next < last { ans = "No"; break }
		last = next
	}
	fmt.Println(ans)
}

