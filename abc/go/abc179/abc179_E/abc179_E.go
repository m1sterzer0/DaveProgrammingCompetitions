package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
func gs() string  { rdr.sc.Scan(); return rdr.sc.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,X,M := gi3()
	if X == 0 || M == 1 { fmt.Println(0); return }
	A := iai(M+3,0)
	S := iai(M+3,0)
	last := iai(M+3,-1)
	x := X; A[0] = x; last[x] = 0; S[0] = X; ans := 0; i:=0
	for i=1;i<N;i++ {
		x = x*x % M
		A[i] = x; S[i] = (S[i-1] + x)
		if last[A[i]] > -1 {
			loopstart := last[A[i]]
			looplen := i - loopstart
			numloops := (N-i)/looplen
			loopsum := S[i] - S[loopstart]
			remlen := N - (i + looplen * numloops)
			remsum := 0; if remlen > 0 { remsum = S[loopstart+remlen-1]-S[loopstart]+A[loopstart] }
			ans = S[i-1] + loopsum * numloops + remsum
			break
		}
		last[A[i]] = i
	}
	if i == N { ans = S[N-1] }
	fmt.Println(ans)
}


