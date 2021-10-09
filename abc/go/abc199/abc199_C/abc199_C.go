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
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); S := gs(); Q := gi(); T,A,B := fill3(Q); for i:=0;i<Q;i++ { A[i]--; B[i]-- }
	SS := make([]byte,2*N); for i:=0;i<2*N;i++ { SS[i] = S[i] }
	flipped := false
	for i:=0;i<Q;i++ {
		t,a,b := T[i],A[i],B[i]
		if t == 2 { flipped = !flipped }
		if t == 1 {
			if flipped {a,b = (a + N)%(2*N),(b+N)%(2*N) }
			SS[a],SS[b] = SS[b],SS[a]
		}
	}
	s1 := string(SS[:N]); s2 := string(SS[N:])
	if !flipped { fmt.Println(s1+s2) } else { fmt.Println(s2+s1) }
}



