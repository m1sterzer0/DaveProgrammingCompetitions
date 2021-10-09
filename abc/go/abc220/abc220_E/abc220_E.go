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
func gi2() (int,int) { return gi(),gi() }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const MOD int = 998244353
func solve(N,D int) int {
	ans := 0
	pow2 := iai(N+1,0); pow2[0] = 1; for i:=1;i<=N;i++ { pow2[i] = 2 * pow2[i-1]; if pow2[i] >= MOD { pow2[i] -= MOD } }
	cumpow2 := iai(N+1,0); cumpow2[0] = 1; for i:=1;i<=N;i++ { cumpow2[i] = cumpow2[i-1] + pow2[i]; if cumpow2[i] > MOD { cumpow2[i] -= MOD } }
	for stdepth:=1;stdepth<=N;stdepth++ {
		if (stdepth-1) + (N-1) < D { continue }
		minup := max(0,(D-N+stdepth+1)/2)
		maxup := min(stdepth-1,D)
		numonlevel := pow2[stdepth-1]
		inc := 0
		if minup == 0 { inc += pow2[D]; minup++}
		if maxup == D { inc += 1 ; maxup-- }
		if maxup >= minup {
			maxpow2 := D - minup - 1
			minpow2 := D - maxup - 1
			inc += cumpow2[maxpow2]
			if minpow2 > 0 { inc += MOD-cumpow2[minpow2-1] }
		}
		inc %= MOD; inc *= numonlevel; inc %= MOD; ans += inc; ans %= MOD
	}
	return ans
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,D := gi2(); 
	ans := solve(N,D)
	fmt.Println(ans)
}
