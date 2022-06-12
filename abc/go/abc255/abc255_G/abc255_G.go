package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type badMove struct { x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi(),gi(); A := gis(N); X,Y := fill2(M)
	badMoves := make([]badMove,M); for i:=0;i<M;i++ { badMoves[i] = badMove{X[i],Y[i]} }
	sort.Slice(badMoves,func(i,j int) bool { return badMoves[i].x < badMoves[j].x } )
	special := make([]int,0)
	getNonSpecial := func(p int) int {
		if len(special) == 0 || p < special[0] { return p }
		if p >= special[len(special)-1] { return p - len(special) }
		l,r := 0,len(special)-1; for r-l > 1 { m := (r+l)>>1; if p >= special[m] { l = m } else { r = m } }
		return p - l - 1
	}
	mm := 0
	extra := make(map[int]int)
	buf := make([]int,0)
	specialVal := make(map[int]int)
	for mm < M {
		x := badMoves[mm].x
		buf = buf[:0]
		baseline := getNonSpecial(x)
		working := baseline
		for mm < M && badMoves[mm].x == x {
			y := badMoves[mm].y
			z := x-y
			v1,ok := specialVal[z]
			if !ok { v1 = getNonSpecial(z) }
			v2,ok2 := extra[v1]
			if ok2 && v2 > 0  { extra[v1]--; buf = append(buf,v1) }
			if ok2 && v2 == 0 || !ok2 { working = min(working,v1) }
			mm++
		}
		if working != baseline {
			special = append(special,x)
			specialVal[x] = working
			extra[working]++
		}
		for _,b := range buf { extra[b]++ }
	}
	nimber := 0
	for _,a := range A {
		x1,ok := specialVal[a]
		if !ok { x1 = getNonSpecial(a) }
		nimber = nimber ^ x1
	}
	ans := "Takahashi"; if nimber == 0 { ans = "Aoki" }
	fmt.Println(ans)
}
