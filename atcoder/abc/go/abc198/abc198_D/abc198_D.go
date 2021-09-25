package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	S1 := gs(); S2 := gs(); S3 := gs(); cmap := make(map[byte]int); nxt := 0
	A1 := makeStrArray(S1,cmap,&nxt)
	A2 := makeStrArray(S2,cmap,&nxt)
	A3 := makeStrArray(S3,cmap,&nxt)
	if nxt > 10  { fmt.Println("UNSOLVABLE"); return }
	digperm := []int{0,1,2,3,4,5,6,7,8,9}
	for {
		if tryit(A1,A2,A3,digperm) { printit(A1,A2,A3,digperm); return }
		if !next_permutation(digperm) { break }
	}
	fmt.Println("UNSOLVABLE")
}

func makeStrArray(S string, cmap map[byte]int, nxt *int) []int {
	ans := make([]int,0)
	for i:=len(S)-1;i>=0;i-- { c := S[i]; v,ok := cmap[c]; if !ok { v = *nxt; cmap[c] = *nxt; *nxt++ }; ans = append(ans,v) }
	return ans 
}

func tryit(A1,A2,A3 []int, digperm []int) bool {
	if digperm[A1[len(A1)-1]] == 0 { return false }
	if digperm[A2[len(A2)-1]] == 0 { return false }
	if digperm[A3[len(A3)-1]] == 0 { return false }
	N1 := getval(A1,digperm)
	N2 := getval(A2,digperm)
	N3 := getval(A3,digperm)
	return N1+N2 == N3
}

func getval(A,digperm []int) int {
	pv := 1; res := 0
	for _,a := range A { res += pv * digperm[a]; pv *= 10}
	return res
}

func printit(A1,A2,A3,digperm []int) {
	N1 := getval(A1,digperm)
	N2 := getval(A2,digperm)
	N3 := getval(A3,digperm)
	fmt.Println(N1); fmt.Println(N2); fmt.Println(N3); 
}

func next_permutation(a []int) bool {
	la := len(a); var i,j int
	for i=la-2;i>=0;i-- { if a[i] < a[i+1] { break } }
	if i<0 { i,j = 0,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- } ; return false }
	for j=la-1;j>=0;j-- { if a[i] < a[j] { break } }
	a[i],a[j] = a[j],a[i]
	i,j = i+1,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- }
	return true
}