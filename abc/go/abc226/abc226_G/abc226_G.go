package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }

func check(A1,A2,A3,A4,A5,B1,B2,B3,B4,B5 int) string {
	if A5 > B5 { return "No" }
	B5 -= A5

	// For A4s, should we use a B4 or a B5? 
	// If we have both a B4 and a B5 available we would rathe have {B5} than {B4+B1}
	// STRATEGY: Use B4, then B5
	if A4 > B4+B5 { return "No" }
	if B4 >= A4 { B4 -= A4; A4 = 0 } else { A4 -= B4; B4 = 0 }
	B5 -= A4; B1 += A4

	// For A3s, should we use B3 or B4 or B5?
	// If we have {B3,B5} either is fine (B5 vs. B3+B2 both pick up 2 pairs plus a single)
	// If we have {B3,B4} we would rather have the B4 left than B3+B1
	// SURPRISE: If we have {B4,B5} we would rather have the B4+B2 left than B5+B1 (first can do 3 A2s, while latter can only do 2)
	// If we have {B3,B4,B5} we would rather have either B4+B5 or B3+B4+B2 (both can do 4 pairs + a single)
	// STRATEGY: Use B3, then B5, then B4  (alternatively B5, then B3, then B4 also works)
	if A3 > B3+B4+B5 { return "No" }
	if B3 >= A3 { B3 -= A3; A3 = 0 } else { A3 -= B3; B3 = 0 }
	if B5 >= A3 { B5 -= A3; B2 += A3; A3 = 0 } else { A3 -= B5; B2 += B5; B5 = 0}
	B4 -= A3; B1 += A3

	// Now we convert all remaining resources to B2s and B1s;
	B2 += 2*B5 + 2*B4 + B3
	B1 += B5 + B3
	if A2 > B2 { return "No" }
	B2 -= A2

	// Convert all remaining B2s to B1s
	B1 += 2 * B2
	if A1 > B1 { return "No" }
	return "Yes"
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		A1,A2,A3 := gi3(); A4,A5 := gi2()
		B1,B2,B3 := gi3(); B4,B5 := gi2()
		ans := check(A1,A2,A3,A4,A5,B1,B2,B3,B4,B5)
		fmt.Fprintln(wrtr,ans)
	}
}
