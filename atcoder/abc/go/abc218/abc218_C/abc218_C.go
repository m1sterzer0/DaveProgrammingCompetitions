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
func gbs() []byte { return []byte(gs()) }
func count(S [][]byte, N int) int {
	cnt := 0
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			if S[i][j] == '#' { cnt++ }
		}
	}
	return cnt
}
func rot90(T [][]byte, U [][]byte, N int) {
	for i:=0;i<N;i++ {
		for j:=0;j<N;j++ {
			U[j][N-1-i] = T[i][j]
		}
	}
}
func check(S,T [][]byte, N int, cnts int) bool {
	c := 0
	i1 := 0; for emptyRow(S,N,i1) { i1++ }
	i2 := 0; for emptyRow(T,N,i2) { i2++ }
	j1 := 0; for emptyCol(S,N,j1) { j1++ }
	j2 := 0; for emptyCol(T,N,j2) { j2++ }
	for di:=0;di<N;di++ {
		if i1+di >= N || i2+di >= N { break }
		for dj:=0;dj<N;dj++ {
			if j1+dj >= N || j2+dj >= N { break }
			if S[i1+di][j1+dj] == '#' && T[i2+di][j2+dj] == '#' {
				c++
			} else if S[i1+di][j1+dj] == '.' && T[i2+di][j2+dj] == '.' {
				continue
			} else {
				return false
			}
		}
	}
	return c == cnts
}
func emptyRow(S [][]byte, N int, i int) bool {
	for j:=0;j<N;j++ { if S[i][j] == '#' { return false } }
	return true
}
func emptyCol(S [][]byte, N int, j int) bool {
	for i:=0;i<N;i++ { if S[i][j] == '#' { return false } }
	return true
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); S := make([][]byte,N); T := make([][]byte,N); U := make([][]byte,N)
	for i:=0;i<N;i++ { S[i] = gbs() }
	for i:=0;i<N;i++ { T[i] = gbs() }
	for i:=0;i<N;i++ { U[i] = make([]byte,N) }
	cnts := count(S,N); cntt := count(T,N)
	if cnts != cntt { fmt.Println("No"); return }
	ans := "No"
	for i:=0;i<4;i++ {
		if check(S,T,N,cnts) { ans = "Yes"; break }
		rot90(T,U,N); U,T = T,U
	}
	fmt.Println(ans)
}

