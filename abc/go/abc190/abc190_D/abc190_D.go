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
	N := gi()
	// Sum of seq starting with a with n terms is a*n + n * (n-1) / 2
	// ** If n is odd, then n-1 is even, and thus sum is n * (a + (n-1)/2), so odd factors of N work
	// ** If n is even, then sum is (n/2) (2*a + n - 1).  Right term is always odd, and right term and sum determine system
	// Thus ans is 2 * number of odd factors
	i := 1; ans := 0
	for i*i < N {
		if N % i == 0 {
			if i & 1 == 1     { ans += 2 }
			if (N/i) & 1 == 1 { ans += 2 }
		}
		i++
	}
	if i*i == N && i & 1 == 1 { ans += 2 }
	fmt.Println(ans)

}



