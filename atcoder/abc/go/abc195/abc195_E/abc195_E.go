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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); S := gs(); X := gs()
	dp := [200001][7]byte{}
	dp[N][0] = 'T'; for i:=1;i<7;i++ { dp[N][i] = 'A' }
	for i:=N-1;i>=0;i-- {
		for j:=0;j<7;j++ {
			t1 := (10*j)%7; t2 := (10*j+int(S[i]-'0')) % 7
			if dp[i+1][t1] == X[i] { dp[i][j] = X[i]; continue }
			if dp[i+1][t2] == X[i] { dp[i][j] = X[i]; continue }
			if X[i] == 'T' { dp[i][j] = 'A' } else { dp[i][j] = 'T' }
		}
	}
	ans := "Takahashi"; if dp[0][0] == 'A' { ans = "Aoki" }
	fmt.Println(ans)
}



