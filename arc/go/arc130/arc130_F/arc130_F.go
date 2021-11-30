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
type pt2 struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	// Solution transcription
	N := gi(); A := gis(N); tmp := 1<<60; pos := 0
	for i,a := range A { if a < tmp { tmp,pos = a,i } }
	b,c := ia(pos+1),ia(N-pos) //b iu non-increasing part, c is non-decreasing part
	for i:=0;i<pos+1;i++ { b[i] = A[pos-i] } 
	for i:=0;i<N-pos;i++ { c[i] = A[i+pos] } 
	ccw := func(a,b,c pt2) int { return (b.x-a.x)*(c.y-a.y) - (c.x-a.x)*(b.y-a.y) } 
	solve := func(v []int) int {
		m := len(v); zan := []int{}
		for p:=0;p<m;p++ {
			for len(zan) >= 2 {
				sz := len(zan)
				a := pt2{zan[sz-2],v[zan[sz-2]]}
				b := pt2{zan[sz-1],v[zan[sz-1]]}
				c := pt2{p,v[p]}
				if ccw(a,b,c) <= 0 { zan = zan[:sz-1] } else { break }
			}
			zan = append(zan,p)
		}
		diff := []int{}
		for i:=1;i<len(zan);i++ {
			dx := zan[i] - zan[i-1]
			dy := v[zan[i]] - v[zan[i-1]]
			for i:=0;i<dx;i++ {
				vv := dy/dx
				if dy%dx > i { vv++ }
				diff = append(diff,vv)
			}
		}
		sort.Slice(diff,func(i,j int) bool { return diff[i] < diff[j]} )
		ret,tmp := v[0],v[0]
		for _,a := range diff {
			tmp += a; ret += tmp
		}
		return ret
	}
	ans := solve(b) + solve(c) - tmp
	fmt.Println(ans)
}

