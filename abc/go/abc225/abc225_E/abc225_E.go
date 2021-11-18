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
type pt struct {idx,x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); X,Y := fill2(N)
	bgn,end := []pt{},[]pt{}
	for i:=0;i<N;i++ { bgn = append(bgn,pt{i,X[i],Y[i]-1}); end = append(end,pt{i,X[i]-1,Y[i]}) }
	sort.Slice(bgn, func(i,j int) bool { return bgn[i].y * bgn[j].x < bgn[i].x * bgn[j].y } )
	sort.Slice(end, func(i,j int) bool { return end[i].y * end[j].x < end[i].x * end[j].y } )
	used := make([]bool,N)
	ans := 0
	bidx,eidx := 0,0
	for eidx < N {
		if used[end[eidx].idx] { eidx++; continue }
		ans++; refx,refy := end[eidx].x, end[eidx].y
		for bidx < N {
			if bgn[bidx].y * refx < bgn[bidx].x * refy { used[bgn[bidx].idx] = true; bidx++; continue }
			break
		}
		for eidx < N {
			if end[eidx].y * refx <= end[eidx].x * refy { eidx++; continue }
			break
		}
	}
	fmt.Println(ans)
}

