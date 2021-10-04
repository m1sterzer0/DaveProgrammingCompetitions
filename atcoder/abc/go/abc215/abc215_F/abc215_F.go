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
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func abs(a int) int { if a < 0 { return -a }; return a }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
type PI struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N)
	pts := make([]PI,N); for i:=0;i<N;i++ { pts[i] = PI{X[i],Y[i]} }
	sort.Slice(pts,func(i,j int) bool { return pts[i].x < pts[j].x || pts[i].x == pts[j].x && pts[i].y < pts[j].y } )
	XX := ia(N); for i:=0;i<N;i++ { XX[i] = pts[i].x }
	YY := ia(N); for i:=0;i<N;i++ { YY[i] = pts[i].y }
	suffixmin := ia(N); suffixmax := ia(N)
	suffixmin[N-1] = YY[N-1]; suffixmax[N-1] = YY[N-1]
	for i:=N-2;i>=0;i-- { suffixmin[i] = min(suffixmin[i+1],YY[i]) }
	for i:=N-2;i>=0;i-- { suffixmax[i] = max(suffixmax[i+1],YY[i]) }
	// Binary search
	check := func(d int) bool { 
		ptr := 0
		for i:=0;i<N;i++ {
			for ptr < N && (ptr <= i || XX[ptr] < XX[i]+d ) { ptr++ }
			if ptr == N { return false }
			a,b := suffixmin[ptr],suffixmax[ptr]
			if abs(YY[i]-a) >= d || abs(YY[i]-b) >= d { return true }
		}
		return false // Shouldn't get here
	}
	l,u := 0,1_000_000_001
	for u-l > 1 {
		m := (u+l) >> 1
		if check(m) { l = m  } else {u = m }
	}
	fmt.Println(l)
}

