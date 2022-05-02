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
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type pt struct {i,j int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W := gi(),gi(); gr := make([]string,H); for i:=0;i<H;i++ { gr[i] = gs() }
	visited := make([][]bool,H); for i:=0;i<H;i++ { visited[i] = make([]bool,W) }
	visited[0][0] = true; q := []pt{{0,0}}
	for len(q) > 0 {
		p := q[0]; q = q[1:]
		if p.i+1 < H && gr[p.i+1][p.j] == '.' && !visited[p.i+1][p.j] { visited[p.i+1][p.j] = true; q = append(q,pt{p.i+1,p.j}) }
		if p.j+1 < W && gr[p.i][p.j+1] == '.' && !visited[p.i][p.j+1] { visited[p.i][p.j+1] = true; q = append(q,pt{p.i,p.j+1}) }
	}
	best := 1
	for i:=0;i<H;i++ { for j:=0;j<W;j++ { if visited[i][j] { best = max(best,1+i+j) } } }
	fmt.Println(best)
}

