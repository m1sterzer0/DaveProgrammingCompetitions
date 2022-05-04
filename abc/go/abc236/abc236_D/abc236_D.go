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
	N := gi(); A := make([][]int,2*N); for i:=0;i<2*N;i++ { A[i] = make([]int,2*N) }
	for i:=0;i<2*N;i++ { for j:=i+1;j<2*N;j++ { A[i][j] = gi(); A[j][i] = A[i][j] } }
	ans := 0; twoN := 2*N
	var search func(lastused,usedmask,matchessofar,scoresofar int)
	search = func(lastused,usedmask,matchessofar,scoresofar int) {
		if matchessofar == N { if ans < scoresofar { ans = scoresofar }; return }
		first := lastused+1; for usedmask & (1<<uint(first)) != 0 { first++ }
		for second:=first+1;second<twoN;second++ {
			if usedmask & (1<<uint(second)) != 0 { continue }
			search(first, usedmask | (1<<uint(first)) | (1<<uint(second)),matchessofar+1,scoresofar ^ A[first][second])
		}
	}
	search(-1,0,0,0)
	fmt.Println(ans)
}
