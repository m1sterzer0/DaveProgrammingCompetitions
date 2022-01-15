package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE

		// DP -- calculate
		// V1 : Best value if my lighthouse is on
		// V2 : Best value if my lighthouse is off but parent lighthouse is on
		// V3 : Best value if neither parent nor my lighthouse is on, but I am illuminated
		// V4 : Best value if I am not illuminated

		// V1 = self + sum(max(v1child,v2child))
		// V2 = self + sum(max(v1child,v3child,v4child))
		// V4 = sum(max(v3child,v4child))
		// V3 = self + sum(max(v1child,v3child,v4child)) + correctionTerm if needed

		V := gi(); B := gis(V); X,Y := fill2(V-1); for i:=0;i<V-1;i++ { X[i]--; Y[i]-- }
		gr := make([][]int,V)
		for i:=0;i<V-1;i++ { x,y := X[i],Y[i]; gr[x] = append(gr[x],y); gr[y] = append(gr[y],x) }
		v1 := ia(V); v2 := ia(V); v3 := ia(V); v4 := ia(V)
		var traverse func(n,p int)
		traverse = func(n,p int) {
			for _,n2 := range gr[n] {
				if n2 == p { continue }
				traverse(n2,n)
			}
			// Leaf cell case
			if n != 0 && len(gr[n]) == 1 {
				v1[n],v2[n],v3[n],v4[n] = B[n],B[n],-10000000000, 0
			} else {
				betterFlag := false; minDelta := inf
				v1[n]=B[n]; v2[n] = B[n]; v3[n] = B[n]
				for _,n2 := range gr[n] {
					if n2 == p { continue }
					v12 := max(v1[n2],v2[n2])
					v34 := max(v3[n2],v4[n2])
					v134 := max(v1[n2],v34)
					v1[n] += v12; v2[n] += v134; v4[n] += v34; v3[n] += v134;
					delta := v1[n2] - v34
					if delta >= 0 { betterFlag = true } else { minDelta = min(minDelta, v34 - v1[n2]) }
				}
				if !betterFlag { v3[n] -= minDelta }
			}
		}
		traverse(0,-1)
		ans := max(max(v1[0],v3[0]),v4[0])
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

