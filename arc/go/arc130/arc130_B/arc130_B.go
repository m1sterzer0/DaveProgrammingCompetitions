package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	H,W,C,Q := gi4(); T,N,CC := fill3(Q)
	rowcolor := make(map[int]int)
	colcolor := make(map[int]int)
	sb := ia(C+1)
	for i:=0;i<Q;i++ { if T[i] == 1 { rowcolor[N[i]] = CC[i] } else { colcolor[N[i]] = CC[i] } }
	freerows := H - len(rowcolor)
	freecols := W - len(colcolor)
	for _,v := range rowcolor { sb[v] += freecols }
	for _,v := range colcolor { sb[v] += freerows }
	rowadder := W - freecols; coladder := H - freerows
	for i:=Q-1;i>=0;i-- {
		if T[i] == 1 && rowcolor[N[i]] == CC[i] {
			sb[CC[i]] += rowadder; coladder -= 1; rowcolor[N[i]] = -1
		} else if T[i] == 2 && colcolor[N[i]] == CC[i] {
			sb[CC[i]] += coladder; rowadder -= 1; colcolor[N[i]] = -1
		}
	}
	ans := vecintstring(sb[1:])
	fmt.Println(ans)
}
