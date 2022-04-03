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
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
type wnum struct { i,w int }
func roundDown(f float64) int { if f < 0 { return int(f-1.0) }; return int(f) }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); X := make([]float64,N); Y := make([]float64,N); W := make([]float64,N)
		for i:=0;i<N;i++ { X[i],Y[i],W[i] = gf(),gf(),gf() }
		U := make([]wnum,N); V := make([]wnum,N)

		findWeightedMedian := func(ww []wnum) int {
			sort.Slice(ww,func(i,j int) bool { return ww[i].i < ww[j].i } )
			sumweights := 0; for _,w := range ww { sumweights += w.w }
			remweights := sumweights
			for _,w := range ww { remweights -= w.w; if 2 * remweights <= sumweights { return w.i } }
			return 0 // Shouldn't get here
		}

		for i:=0;i<N;i++ {
			x := roundDown(0.5 + 100 * X[i])
			y := roundDown(0.5 + 100 * Y[i])
			w := roundDown(0.5 + 100 * W[i]) 
			U[i] = wnum{ x+y, w }
			V[i] = wnum{ x-y, w }
		}
		umed := findWeightedMedian(U)
		vmed := findWeightedMedian(V)
		xmed := 0.005 * float64(umed+vmed)
		ymed := 0.005 * float64(umed-vmed)
		ans := float64(0.0)
		for i:=0;i<N;i++ {
			xdist := xmed - X[i]; if xdist < 0 { xdist *= -1.0 }
			ydist := ymed - Y[i]; if ydist < 0 { ydist *= -1.0 }
			dist := xdist; if ydist > xdist { dist = ydist }
			ans += dist * W[i]
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

