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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	// * Want to find largest a s.t. a <= A[0] and a xor (S - a) == nimber -->
	//      S-a == nimber xor a --> a xor nimber + a == S 
	// * MAGIC FORMULA RELATING SUM AND XOR :  a + b = a xor b + 2 * a & b
	// * therefore a xor a xor nimber + 2 * (a & (a xor nimber)) == S
	// * 2 * (a & (a xor nimber)) == (S-nimber)
	// * (a & a xor nimber) == (S-nimber) / 2  == target
	// Now we can do this bitwise
	// nimber == 1, target == 1 -- BAD
	// nimber == 1, target == 0 -- FREE CHOICE
	// nimber == 0, target == 1 -- a == 1
	// nimber == 0, target == 0 -- a == 0
	// This leads to the following greedy approach
	// -- Check to see if we are bad either because
	//    a) (S-nimber) isn't divisible by 2
	//    b) we have a bad condition with nimber and target
	//    c) our forced ones already put us over our limit
	// -- If we pass there, we greedily fill in ones
	N := gi(); A := gis(N); nimber := 0
	for i:=2;i<N;i++ { nimber ^= A[i] }
	S := A[0] + A[1]
	good := true
	if nimber > S || (S-nimber) & 1 == 1 { good = false }
	targ := (S - nimber) / 2
	must := 0
	for i:=62;i>=0;i-- {
		mask := 1 << i
		if nimber & mask != 0 && targ & mask != 0 { good = false }
		if nimber & mask == 0 && targ & mask != 0 { must += mask }
	}
	if must > A[0] { good = false }
	a0size := must
	for i:=62;i>=0;i-- {
		mask := 1 << i
		if nimber & mask != 0 && targ & mask == 0 && a0size + mask < A[0] { a0size += mask }
	}
	if a0size == 0 { good = false }
	ans := A[0] - a0size
	if !good { ans = -1 }
	fmt.Println(ans)	
}



