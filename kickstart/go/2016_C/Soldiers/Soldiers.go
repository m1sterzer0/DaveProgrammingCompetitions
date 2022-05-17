package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func min(a,b int) int { if a > b { return b }; return a }
func max(a,b int) int { if a < b { return b }; return a }
func sortUniqueIntarr(a []int) []int {
	sort.Slice(a,func (i,j int) bool { return a[i] < a[j] })
	i,j,la := 0,0,len(a)
	for ;i<la;i++ { if i == 0 || a[i] != a[i-1] { a[j] = a[i]; j++ } }
	return a[:j]
}
const inf int = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); A,D := fill2(N)
		AA := make([]int,N); copy(AA,A); AA = sortUniqueIntarr(AA)
		DD := make([]int,N); copy(DD,D); DD = sortUniqueIntarr(DD)
		atocomp := iai(10001,-1); for i,a := range AA { atocomp[a] = i } 
		dtocomp := iai(10001,-1); for i,d := range DD { dtocomp[d] = i }
		A2 := make([]int,N); for i,a := range A { A2[i] = atocomp[a] }
		D2 := make([]int,N); for i,d := range D { D2[i] = dtocomp[d] }
		na,nd := len(AA),len(DD)
		card := make([][]bool,na);    for i:=0;i<na;i++ { card[i] = make([]bool,nd) }
		winning := make([][]bool,na); for i:=0;i<na;i++ { winning[i] = make([]bool,nd) }
		for i:=0;i<N;i++ { card[A2[i]][D2[i]] = true }
		firstrow := iai(nd,inf); firstcol := iai(na,inf)
		for i:=0;i<N;i++ { ii,jj := A2[i],D2[i]; firstrow[jj] = min(firstrow[jj],ii); firstcol[ii] = min(firstcol[ii],jj) }
		alice := false
		for i:=na-1;i>=0;i-- {
			for j:=nd-1;j>=0;j-- {
				if winning[i][j] { continue }
				if card[i][j] { alice = true; break }
				ii := firstrow[j]; if (ii < i) { for jjj:=0;jjj<j;jjj++ { winning[i][jjj] = true } }
				jj := firstcol[i]; if (jj < j) { for iii:=0;iii<i;iii++ { winning[iii][j] = true } }
			}
			if alice { break }
		}
		ans := "NO"; if alice { ans = "YES" }
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

