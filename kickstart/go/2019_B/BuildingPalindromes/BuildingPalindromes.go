package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	prefixSum := [100000][26]int{}
	T := gi()
    for tt:=1;tt<=T;tt++ {
		_,Q := gi2(); S := gs(); L,R := fill2(Q); for i:=0;i<Q;i++ { L[i]--; R[i]-- }
		sb := ia(26)
		for i,c := range S {
			sb[int(c-'A')]++
			for j:=0;j<26;j++ { prefixSum[i][j] = sb[j] }
		}
		ans := 0
		for i:=0;i<Q;i++ {
			l,r := L[i],R[i]
			numodd := 0
			for j:=0;j<26;j++ {
				s := prefixSum[r][j]
				if l > 0 { s -= prefixSum[l-1][j] }
				if s % 2 == 1 { numodd++ }
			}
			if numodd < 2 { ans++ }
		}
	    // PROGRAM STARTS HERE
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

