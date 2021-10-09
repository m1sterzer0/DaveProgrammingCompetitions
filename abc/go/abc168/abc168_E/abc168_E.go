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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }

const MOD = 1_000_000_007
type frac struct {n,d int}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	
    // NON-BOILERPLATE STARTS HERE
	// Dot product is zero if the two vectors are perpendicular
	// -- Record points by the rational slope (in lowest terms) of the line through the origin that contains the point
	// -- Special case the origin
	// -- Record 0 as (0,1).  Record infinite slope as (1,0)
	// Then loop through all of the slopes and either take terms from one or the other 
	N := gi()
	A,B := fill2(N)
	scnt := make(map[frac]int)
	for i:=0;i<N;i++ { 
		a,b := A[i],B[i]
		if a==0 && b==0 { 
			scnt[frac{0,0}]++
		} else if a == 0 {
			scnt[frac{0,1}]++
		} else if b == 0 {
			scnt[frac{1,0}]++
		} else {
			sgn := 1
			if a < 0 { sgn *= -1 ; a *= -1}
			if b < 0 { sgn *= -1 ; b *= -1}
			d := gcd(a,b)
			scnt[frac{sgn*a/d,b/d}]++
		}
	}
	ways := 1
	// To avoid double counting and not modifying map while iterating through, when two 
	// terms conflict, we only count when the numerator is positive
	// only look at terms with a positive numerator.  This picks up one of the zero cases, and one of each pair.
	// Origin is handled separately at the end

	for k,v := range scnt {
		if k.n == 0 && k.d == 0 { 
			continue
		} else if k.n == 0 {
			if _,ok := scnt[frac{1,0}]; ok { continue }
			ways = ways * powmod(2,v,MOD) % MOD
		} else if k.n < 0 {
			if _,ok := scnt[frac{k.d,-k.n}]; ok { continue }
			ways = ways * powmod(2,v,MOD) % MOD
		} else {
			l := powmod(2,v,MOD)
			if v2,ok := scnt[frac{-k.d,k.n}]; ok { l = (l + powmod(2,v2,MOD) - 1) % MOD }
			ways = ways * l % MOD
		}
	}
	// Cleanup -- deal with origin and empty set
	ans := (MOD + ways - 1 + scnt[frac{0,0}]) % MOD
	fmt.Fprintln(wrtr,ans)
}
