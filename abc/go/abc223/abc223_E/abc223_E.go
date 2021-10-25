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
func next_permutation(a []int) bool {
	la := len(a); var i,j int
	for i=la-2;i>=0;i-- { if a[i] < a[i+1] { break } }
	if i<0 { i,j = 0,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- } ; return false }
	for j=la-1;j>=0;j-- { if a[i] < a[j] { break } }
	a[i],a[j] = a[j],a[i]
	i,j = i+1,la-1; for i<j { a[i],a[j] = a[j],a[i]; i++; j-- }
	return true
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	X,Y,A,B,C := gi(),gi(),gi(),gi(),gi()
	check := func(a,b,c,x,y int) bool {
		// two ways -- a,b,c stacked on top, or a on bottom and b,c side by side
		y1,y2,y3 := (a+x-1)/x,(b+x-1)/x,(c+x-1)/x
		if y1+y2+y3 <= y { return true }
		if y1 >= y { return false }
		yleft := y-y1
		x1,x2 := (b+yleft-1)/yleft,(c+yleft-1)/yleft
		return x1+x2 <= x
	}
	good := false 
	a := []int{A,B,C}
	sort.Slice(a,func(i,j int) bool { return a[i] < a[j] })
	for i:=0;i<6;i++ {
		good = good || check(a[0],a[1],a[2],X,Y)
		good = good || check(a[0],a[1],a[2],Y,X)
		next_permutation(a)
	}
	ans := "No"; if good { ans = "Yes" }
	fmt.Println(ans)
}

