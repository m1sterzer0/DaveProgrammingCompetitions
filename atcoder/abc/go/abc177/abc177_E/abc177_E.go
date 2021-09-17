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

func fsieve(n int) []int {
	fs := make([]int,n+1); fs[0] = 0; fs[1] = 1
	for i:=2;i<=n;i++ { fs[i] = -1 };
	for i:=2;i<=n;i+=2 { fs[i] = 2 }
	for i:=3;i<=n;i+=2 { if fs[i] == -1 { fs[i] = i; inc := 2*i; for k:=i*i;k<=n;k+=inc { if fs[k] == -1 { fs[k] = i } } } }
	return fs
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N := gi(); A := gis(N)
	// Build a factor sieve
	fs := fsieve(1_000_000)
	//fcnts := make(map[int]int)
	fcnts := make([]int,1_000_001)
	for _,a := range(A) {
		x := a; for x > 1 { f := fs[x]; fcnts[f]++; for x % f == 0 { x /= f } }
	}
	ans := "pairwise coprime"
	for _,v := range fcnts { if v > 1 { ans = "setwise coprime" }; if v == N { ans = "not coprime"; break } } 
	fmt.Println(ans)
}



