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
type pair struct { x,y int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	gi();gi(); N := gi()
	sx,sy,gx,gy := gi(),gi(),gi(),gi()
	s := pair{sx,sy}
	X,Y := fill2(N)
	R := make([]pair,N); for i:=0;i<N;i++ { R[i] = pair{X[i],Y[i]}}
	sort.Slice(R, func(i,j int) bool { return R[i].x < R[j].x || R[i].x == R[j].x && R[i].y < R[j].y } )
	byrow := make(map[int][]int)
	bycol := make(map[int][]int)
	for _,r := range R { byrow[r.x] = append(byrow[r.x],r.y); bycol[r.y] = append(bycol[r.y],r.x) }
	findlt := func(a []int, t int) int {
		la := len(a); if la == 0 || t <= a[0] { return -1 }
		if a[la-1] < t { return a[la-1] }
		l,r := 0,la-1; for r-l > 1 { m := (r+l)>>1; if a[m] < t { l = m } else { r = m }}
		return a[l]
	}
	findgt := func(a []int, t int) int {
		la := len(a); if la == 0 || a[la-1] <= t { return -1 }
		if a[0] > t { return a[0] }
		l,r := 0,la-1; for r-l > 1 { m := (r+l)>>1; if a[m] > t { r = m } else { l = m }}
		return a[r]
	}
	dmap := make(map[pair]int)
	q := make([]pair,0)
	// offset distances by 1 to make the default map work well
	dmap[s] = 1; q = append(q,s)
	for len(q) > 0 {
		p := q[0]; q = q[1:]
		// There is a factor of 4 here if i need it, but i'm being lazy
		l,r := findlt(byrow[p.x],p.y),findgt(byrow[p.x],p.y)
		if l >= 0 { p2 := pair{p.x,l+1}; if dmap[p2] == 0 { dmap[p2] = dmap[p]+1; q = append(q,p2) } }
		if r >= 0 { p2 := pair{p.x,r-1}; if dmap[p2] == 0 { dmap[p2] = dmap[p]+1; q = append(q,p2) } }
		l,r = findlt(bycol[p.y],p.x),findgt(bycol[p.y],p.x)
		if l >= 0 { p2 := pair{l+1,p.y}; if dmap[p2] == 0 { dmap[p2] = dmap[p]+1; q = append(q,p2) } }
		if r >= 0 { p2 := pair{r-1,p.y}; if dmap[p2] == 0 { dmap[p2] = dmap[p]+1; q = append(q,p2) } }
	}
	d,_ := dmap[pair{gx,gy}]; d--; fmt.Println(d)
}
