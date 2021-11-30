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
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
type fedge struct { x int; left bool }
func solvepos(xedges []fedge) int {
	sort.Slice(xedges,func(i,j int) bool { return xedges[i].x < xedges[j].x })
	numright := len(xedges) / 2; numleft := 0
	for _,e := range xedges { if e.left { numright-- } else { numleft++ }; if numleft == numright { return e.x } }
	return -1 //Shouldn't get here
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		K := gi(); X1,Y1,X2,Y2 := fill4(K)
		xedges := []fedge{}; yedges := []fedge{}
		for i:=0;i<K;i++ { 
			xedges = append(xedges,fedge{X1[i],true})
			xedges = append(xedges,fedge{X2[i],false})
			yedges = append(yedges,fedge{Y1[i],true})
			yedges = append(yedges,fedge{Y2[i],false})
		}
		x := solvepos(xedges)
		y := solvepos(yedges)
        fmt.Fprintf(wrtr,"Case #%v: %v %v\n",tt,x,y)
    }
}

