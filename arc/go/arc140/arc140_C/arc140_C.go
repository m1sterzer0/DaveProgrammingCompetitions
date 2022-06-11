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
	N,X := gi(),gi(); X--
	P := make([]int,0,N)
	// Full N-1 solutions
	if N%2 == 1 && X == N/2 || N%2 == 0 && X == N/2-1 {
		l,r := X,X+1; for { if l < 0 { break }; P = append(P,l); l--; if r >= N { break }; P = append(P,r); r++ }
	} else if N%2 == 0 && X == N/2 {
		l,r := X-1,X; for { if r >= N { break }; P = append(P,r); r++; if l < 0 { break }; P = append(P,l); l-- }		
	} else {
		P2 := make([]int,0,N); for i:=0;i<N;i++ { if i != X { P2 = append(P2,i) } }
		lidx := 0; if len(P2) % 2 == 1 { lidx = len(P2)/2 } else { lidx = len(P2)/2-1 }; ridx := lidx+1
		P = append(P,X)
		for { if lidx < 0 { break }; P = append(P,P2[lidx]); lidx--; if ridx >= N-1 { break }; P = append(P,P2[ridx]); ridx++ }
	}
	for i:=0;i<N;i++ { P[i]++ }
	fmt.Println(vecintstring(P))
}

