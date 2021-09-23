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
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func min(a,b int) int { if a > b { return b }; return a }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func crtsafemod(x, m int) int { x %= m; if x < 0 { x += m }; return x }
func crtinvgcd(a, b int) (int, int) {
	a = crtsafemod(a, b); if a == 0 { return b, 0 }; s, t, m0, m1 := b, a, 0, 1
	for t != 0 { u := s / t; s -= t * u; m0 -= m1 * u; s, t, m0, m1 = t, s, m1, m0 }; if m0 < 0 { m0 += b / s }
	return s, m0
}
func crt(r, m []int) (int, int) {
	if len(r) != len(m) { panic("Mismatched length in crt") }
	for _, mm := range m { if mm <= 0 { panic("CRT error -- non-positive modulus") } }; n, r0, m0 := len(r), 0, 1
	for i := 0; i < n; i++ {
		r1, m1 := crtsafemod(r[i], m[i]), m[i]; if m0 < m1 { r0, r1, m0, m1 = r1, r0, m1, m0 }
		if m0%m1 == 0 { if r0%m1 != r1 { return 0, 0 }; continue }; g, im := crtinvgcd(m0, m1); u1 := m1 / g
		if (r1-r0)%g != 0 { return 0, 0 }; x := (r1 - r0) / g % u1 * im % u1; r0 += x * m0; m0 *= u1
		if r0 < 0 { r0 += m0 }
	}
	return r0, m0
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	T := gi()
	for i:=0;i<T;i++ {
		X,Y,P,Q := gi4()
		myinf := 3*powint(10,18); ans := myinf; m := []int{2*X+2*Y,P+Q}; r := []int{0,0}
		for rem1:=0;rem1<Y;rem1++ { r[0] = X+rem1; r[1] = P;      rr,mm:=crt(r,m); if mm != 0 { ans = min(ans,rr) } }
		for rem2:=0;rem2<Q;rem2++ { r[0] = X;      r[1] = P+rem2; rr,mm:=crt(r,m); if mm != 0 { ans = min(ans,rr) } }
		if ans == myinf { fmt.Fprintln(wrtr,"infinity") } else { fmt.Fprintln(wrtr,ans) }
	}
}
