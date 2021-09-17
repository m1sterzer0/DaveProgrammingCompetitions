package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
func reverseiarr(a []int) { i,j := 0,len(a)-1; for j > i { a[i],a[j] = a[j],a[i]; i++; j-- } }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); A,B := gis(N),gis(N)
	reverseiarr(B)
	good := true; aidx := 0; bidx := 0
	for aidx < N {
		if A[aidx] != B[aidx] { aidx++; continue }
		for bidx < N { if A[bidx] != A[aidx] && B[bidx] != A[aidx] { break }; bidx++ }
		if bidx == N { good = false; break}
		B[aidx],B[bidx] = B[bidx],B[aidx]
	}
	if !good { 
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
		ansstrarr := make([]string,N)
		for i:=0;i<N;i++ { ansstrarr[i] = strconv.Itoa(B[i]) }
		ans := strings.Join(ansstrarr," ")
		fmt.Println(ans)
	}
}
