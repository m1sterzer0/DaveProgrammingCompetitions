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
func gi2() (int,int) { return gi(),gi() }
type player struct {id,wins int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); A := make([]string,2*N); for i:=0;i<2*N;i++ { A[i] = gs() }
	players := make([]player,2*N)
	for i:=0;i<2*N;i++ { players[i] = player{i,0} }
	for j:=0;j<M;j++ {
		for i:=0;i<2*N;i+=2 {
			s1 := A[players[i].id][j]
			s2 := A[players[i+1].id][j]
			if s1 == 'G' && s2 == 'C' || s1 == 'C' && s2 == 'P' || s1 == 'P' && s2 == 'G' { players[i].wins++ }
			if s2 == 'G' && s1 == 'C' || s2 == 'C' && s1 == 'P' || s2 == 'P' && s1 == 'G' { players[i+1].wins++ }
		}
		sort.Slice(players,func(i,j int) bool { return players[i].wins > players[j].wins || players[i].wins == players[j].wins && players[i].id < players[j].id} )
	}
	for i:=0;i<2*N;i++ { fmt.Fprintln(wrtr,players[i].id+1) }
}

