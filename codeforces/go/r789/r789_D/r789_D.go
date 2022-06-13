package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N,M := gi(),gi(); S := gs()
		NM := N*M
		ans := make([]int,NM)
		colflag := make([]bool,M); goodcol := 0
		rowsum := make([]int,M); numgood := 0; i:=0
		for j:=0;j<N;j++ {
			for k:=0;k<M;k++ {
				if j == 0 {
					if S[i] == '1' { colflag[k] = true; goodcol++ }
					if S[i] == '1' { numgood++ }
					if numgood > 0 { rowsum[k]++ }
					ans[i] = goodcol + rowsum[k]
				} else {
					if S[i] == '1' && !colflag[k] { colflag[k] = true; goodcol++}
					if S[i] == '1' { numgood++ }
					if S[i-M] == '1' { numgood--}
					if numgood > 0 { rowsum[k]++ }
					ans[i] = goodcol + rowsum[k]
				}
			i++
			}
		}
		fmt.Fprintln(wrtr,vecintstring(ans))
	}
}

