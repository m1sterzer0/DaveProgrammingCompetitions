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
func gi3() (int,int,int) { return gi(),gi(),gi() }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,M,K := gi3(); A := gis(K)
	badsq := make([]bool,N+1)
	for _,a := range A { badsq[a] = true }
	good := true; run := 0
	for _,bad := range badsq { 
		if bad { run++ } else { run = 0 }
		if run >= M { good = false; break }
	}
	if !good { fmt.Println(-1); return }
	scalar := make([]float64,N+M+10); coeff := make([]float64,N+M+10)
	rfm := 1.00/float64(M); ss := 0.00; sc := 0.00
	for i:=N-1;i>=0;i-- {
		if badsq[i] { 
			scalar[i] = 0.00; coeff[i] = 1.000
		} else {
			scalar[i] = 1.000 + rfm*ss; coeff[i] = rfm * sc
		}
		ss += scalar[i] - scalar[i+M]
		sc += coeff[i] - coeff[i+M]
	}
	ans := (scalar[0]) / (1.00 - coeff[0])
	fmt.Println(ans)
}



