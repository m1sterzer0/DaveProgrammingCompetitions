package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
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
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func fwht(a []int) {
	h := 1
	for h < len(a) {
		for i:=0;i<len(a);i+=2*h {
			for j:=i;j<i+h;j++ {
				x,y := a[j],a[j+h]
				a[j],a[j+h] = x+y,x-y
			}
		}
		h *= 2
	}
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }

	// Transcribed from python submission https://atcoder.jp/contests/abc220/submissions/26170898
	// I will go back and figure this out later.  There is no way I would have gotten this in a
	// timed contest.
	N,M := gi2(); A,B := fill2(M); for i:=0;i<M;i++ { A[i]--; B[i]-- }
	adj := make([][]bool,N); for i:=0;i<N;i++ { adj[i] = make([]bool,N) }
	for i:=0;i<M;i++ { a,b := A[i],B[i]; adj[a][b] = true; adj[b][a] = true }
	leftsz := N/2; rightsz := N-leftsz
	left_bitwise_adj := iai(leftsz,0)
	for i:=0;i<leftsz;i++ {
		for j:=i+1;j<leftsz;j++ {
			if adj[i][j] { left_bitwise_adj[i] |= 1<<j }
		}
	}
	right_bitwise_adj := iai(rightsz,0)
	for i:=0;i<rightsz;i++ {
		for j:=i+1;j<rightsz;j++ {
			if adj[leftsz+i][leftsz+j] { right_bitwise_adj[i] |= 1<<j}
		}
	}
	masks := iai(rightsz,0)
	for i:=0;i<rightsz;i++ {
		for j:=0;j<leftsz;j++ {
			if adj[leftsz+i][j] { masks[i] |= 1<<j }
		}
	}
	dpleft := iai(1<<leftsz,0)
	for bm:=0;bm<1<<leftsz;bm++ {
		p := 1
		for i:=0;i<leftsz;i++ {
			if (bm >> i) & 1 == 0 { continue }
			if bits.OnesCount(uint(left_bitwise_adj[i] & bm)) % 2 == 1 { p *= -1 }
		}
		dpleft[bm] = p
	}
	fwht(dpleft)
	ans := 0
	for bm:=0;bm<1<<rightsz;bm++ {
		p,mask := 1,0
		for i:=0;i<rightsz;i++ {
			if (bm >>i) & 1 == 0 { continue }
			if bits.OnesCount(uint(right_bitwise_adj[i] & bm)) % 2 == 1 { p *= -1 }
			mask ^= masks[i]
		}
		psum := dpleft[mask] * p
		ans += (psum + (1 << leftsz)) / 2
	}
	if M % 2 == 1 { ans = powint(2,N) - ans }
	fmt.Println(ans)
}



