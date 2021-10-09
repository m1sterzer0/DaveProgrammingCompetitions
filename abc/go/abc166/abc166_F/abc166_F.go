package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {l, p, e := rdr.ReadLine(); if e != nil { fmt.Println(e.Error()); panic(e) }; buf = append(buf, l...); if !p { break } }
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }
func gti() int { var a int; fmt.Fscan(rdr,&a); return a }
func gtf() float64 { var a float64; fmt.Fscan(rdr,&a); return a }
func gts() string { var a string; fmt.Fscan(rdr,&a); return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewReaderSize(f, BUFSIZE) }
	
	// Greedy solution for the win!!

	// If all of the A,B,C are greater than zero, then we can pick anything (see below).  The most we will create
	// is one zero, and from the discussion below, one zero is winning.

	// Now all moves are forced at having 2 zeros, and 3 zeros is losing unless we are at the end of the game (should only encounter 3 zeros at start)

	// This leaves what to choose when we have one zero and have to choose between the other two, potentially creating a 2nd zero.  However, this also has
	// a greedy strategy.  We simply look ahead to the next move, and pick the option for the current move that won't lose us the game in the next move.
	// We will end up with one zero after the pair of moves.

    // NON-BOILERPLATE STARTS HERE
	N := gti(); A := gti(); B := gti(); C := gti()
	S := make([]string,N)
	for i:=0;i<N;i++ { S[i] = gts() }
	moves := make([]byte,0); good := true
	play := func(a byte, b byte) {
		if a == 'A' { A++ } else if a == 'B' { B++ } else { C++ }
		if b == 'A' { A-- } else if b == 'B' { B-- } else { C-- }
		moves = append(moves,a)
	}
	for i,s := range S {
		c1 := tern(s[0]=='A',A,tern(s[0]=='B',B,C))
		c2 := tern(s[1]=='A',A,tern(s[1]=='B',B,C))
		if c1 == 0 && c2 == 0 { good = false; break	} 
		if c1 == 0  { play(s[0],s[1]); continue }
		if c2 == 0  { play(s[1],s[0]); continue }
		if i+1 == N { play(s[0],s[1]); continue }
		if s[0] == S[i+1][0] || s[0] == S[i+1][1] { play(s[0],s[1]) } else { play(s[1],s[0]) } 
	}
	if good { 
		fmt.Fprintln(wrtr,"Yes")
		for _,b := range moves { fmt.Fprintln(wrtr,string(b)) }
	} else {
		fmt.Fprintln(wrtr,"No")
	}
	wrtr.Flush()
}

			
