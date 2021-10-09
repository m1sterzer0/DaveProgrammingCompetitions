package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const BUFSIZE = 10000000
var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil { fmt.Println(e.Error()); panic(e) }
		buf = append(buf, l...)
		if !p { break }
	}
	return string(buf)
}

func gs() string    { return readLine() }
func gss() []string { return strings.Fields(gs()) }
func gi() int {	res, e := strconv.Atoi(gs()); if e != nil { panic(e) }; return res }
func gf() float64 {	res, e := strconv.ParseFloat(gs(), 64); if e != nil { panic(e) }; return float64(res) }
func gis() []int { res := make([]int, 0); 	for _, s := range gss() { v, e := strconv.Atoi(s); if e != nil { panic(e) }; res = append(res, int(v)) }; return res }
func gfs() []float64 { res := make([]float64, 0); 	for _, s := range gss() { v, _ := strconv.ParseFloat(s, 64); res = append(res, float64(v)) }; return res }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {
		f, e := os.Open(infn)
		if e != nil { panic(e) }
		rdr = bufio.NewReaderSize(f, BUFSIZE)
	}
	
    // NON-BOILERPLATE STARTS HERE
	N := gi()
	A := gis()
	// * For the children that move left, they will add Ax * (x - i),
	//   where i is the position they move to
	// * For the children that move right, they will add Ax * (i - x), where
	//   i is the position they move to.
	// * Suppose we fix the subset of children that move left.  Then the Ax*x terms for the
	//   left moving terms, and the -Ax*x terms for the right moving children are all fixed.  Then
	//   we just need to optimize the -Ax*i terms for the left and Ax*i terms for the right.
	//   Clearly this is optimized by assigning 1 to the largest leftward term, and N to the largest
	//   right moving term, and so on.  This means we can sort Ax and fill from outside in.
	myinf := 1_000_000_000_000_000_000
	AA := make([]PI,N)
	for i:=0;i<N;i++ { AA[i] = PI{A[i],i}}
	sort.Slice(AA,func(i,j int) bool { return AA[i].x > AA[j].x})
	dp := make([]int,N+1)
	ndp := make([]int,N+1)
	for i:=0;i<N;i++ { dp[i] = -myinf}
	dp[0] = 0
	for numfilled,aa := range(AA) {
		ax := aa.x; idx := aa.y
		for i:=0;i<N;i++ { ndp[i] = -myinf }
		for i:=0;i<N-1;i++ {
			if dp[i] == -myinf { continue }
			lidx := i
			ridx := N-1-numfilled+i
			if idx >= lidx { ndp[i+1] = max(ndp[i+1],dp[i]+ax*(idx-lidx)) }
			if idx <= ridx { ndp[i]   = max(ndp[i],dp[i]+ax*(ridx-idx)) }
		}
		dp,ndp = ndp,dp
	}
	ans := maxarr(dp)
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
}



