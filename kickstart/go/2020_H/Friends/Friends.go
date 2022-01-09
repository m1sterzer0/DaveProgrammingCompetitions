package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
const inf int = 1000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2()
		S := make([]string,N); for i:=0;i<N;i++ { S[i] = gs() }
		X,Y := fill2(Q)
		// Partition the names by letter
		gr := twodi(26,26,inf)
		for _,s := range S {
			for _,c1 := range s {
				for _,c2 := range s {
					gr[int(c1-'A')][int(c2-'A')] = 1
				}
			}
		}
		for i:=0;i<26;i++ { gr[i][i] = 0 }
		// Bellman ford for the distances
		for k:=0;k<26;k++ {
			for i:=0;i<26;i++ {
				for j:=0;j<26;j++ {
					gr[i][j] = min(gr[i][j],gr[i][k]+gr[k][j])
				}
			}
		}
		// Now for the queries 
		qans := ia(0)
		for i:=0;i<Q;i++ {
			x,y := X[i],Y[i]; x--; y--
			s1,s2 := S[x],S[y]
			best := inf
			for _,c1 := range s1 {
				for _,c2 := range s2 {
					best = min(best,gr[int(c1-'A')][int(c2-'A')])
				}
			}
			best += 2
			if best >= inf { best = -1 }
			qans = append(qans,best)
		}
		ans := vecintstring(qans)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

