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
func gi2() (int,int) { return gi(),gi() }
func max(a,b int) int { if a > b { return a }; return b }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
    defer wrtr.Flush()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	N,T := gi2()
	A := gis(N)
	if N == 1 { ans:=0; if A[0] <= T { ans = A[0]}; fmt.Println(ans); return }
	m := N/2
	solveside := func(l,r int) []int {
		left := make([]int,0); left = append(left,0)
		nleft := make([]int,0)
		working := make([]int,0)
		for i:=l;i<=r;i++ {
			nleft = nleft[:0]; working = working[:0]
			a := A[i]
			for _,k := range left { v := a+k; if v > T { break }; working = append(working,v) }
			j,k,ll,lw := 0,0,len(left),len(working)
			for j < ll || k < lw {
				if j < ll && (k == lw || left[j] < working[k]) { 
					nleft = append(nleft,left[j]); j++
				} else {
					nleft = append(nleft,working[k]); k++
				}
			}
			left,nleft = nleft,left
		}
		return left
	}

	left := solveside(0,m-1)
	right := solveside(m,N-1)
	ans := 0
	ridx := len(right)-1
	for _,l := range left {
		for l+right[ridx] > T { ridx-- }
		ans = max(ans,l+right[ridx])
	}
	fmt.Println(ans)
}



