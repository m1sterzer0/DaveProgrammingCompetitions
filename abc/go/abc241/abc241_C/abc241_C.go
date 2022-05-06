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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }
	ans := "No"
	a := make([]int,0)
	check := func(si,sj,di,dj int) bool {
		a = a[:0]; a = append(a,0); cnt := 0; lidx := 0
		for i,j:=si,sj;i >= 0 && i < N && j >= 0 && j < N; i,j = i+di,j+dj {
			if bd[i][j] == '#' { cnt++ }
			a = append(a,cnt); lidx++
			if lidx >= 6 && a[lidx]-a[lidx-6] >= 4 { return true }
		}
		return false
	}
	for i:=0;i<N;i++ { if check(i,0,0,1)     { ans = "Yes" } }
	for j:=0;j<N;j++ { if check(0,j,1,0)     { ans = "Yes" } }
	for i:=0;i<N;i++ { if check(i,0,1,1)     { ans = "Yes" } }
	for j:=0;j<N;j++ { if check(0,j,1,1)     { ans = "Yes" } }
	for i:=0;i<N;i++ { if check(i,0,-1,1)    { ans = "Yes" } }
	for j:=0;j<N;j++ { if check(N-1,j,-1,1)  { ans = "Yes" } }
	fmt.Println(ans)
}

