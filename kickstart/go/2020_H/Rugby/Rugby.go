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
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func abs(a int) int { if a < 0 { return -a }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X,Y := fill2(N)
		// Do Y first just the median
		sort.Slice(Y,func (i,j int) bool { return Y[i] < Y[j]})
		posy := Y[N/2]
		ysteps := 0
		for _,y := range Y { ysteps += abs(y-posy) }
		// Now we do the X
		sort.Slice(X, func(i,j int) bool { return X[i] < X[j]})
		posx := X[0]
		delx := ia(N); for i,x := range X { delx[i] = x - (posx+i) }
		sort.Slice(delx, func(i,j int) bool { return delx[i] < delx[j] } )
		posx += delx[N/2]
		xsteps := 0
		for i,x := range X { xsteps += abs(x - (posx+i)) }
		ans := xsteps + ysteps
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

