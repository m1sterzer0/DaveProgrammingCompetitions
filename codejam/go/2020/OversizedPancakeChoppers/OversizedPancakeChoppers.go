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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,D := gi(),gi(); A := gis(N)
		cnts := make(map[int]int); for _,a := range A { cnts[a]++ }
		best := D
		for d := 1; d<=D; d++ {
			l,u := 1,360000000001
			for u-l > 1 { // Binsearch for max slize
				m := (u+l)>>1
				numslices := 0
				for _,a := range A { numslices += a*d/m }
				if numslices < D { u = m } else { l = m }
			}
			maxsize := l
			for a := range cnts {
				if a > maxsize { continue }
				cuts := 0; donesofar := 0
				for k:=d;k<=D && cuts < best && donesofar < D;k++ {
					if k*a % d != 0 { continue }
					v := k*a/d
					if cnts[v] == 0 { continue }
					if k * cnts[v] + donesofar < D { 
						cuts += (k-1) * cnts[v]; donesofar += k * cnts[v]
					} else {
						veff := (D-donesofar)/k
						cuts += (k-1) * veff; donesofar += k * veff
						cuts += (D-donesofar); donesofar = D
					}
				}
				cuts += D-donesofar
				best = min(best,cuts)
			}
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best)
	}
}
