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
func max(a,b int) int { if a > b { return a }; return b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()

	// Prework 1 -- create a factor list for every integer in N log N
	//fmt.Fprintf(os.Stderr,"HERE1\n")
	factors := make([][]int,1000001)
	for i:=0;i<=1000000;i++ { factors[i] = make([]int,0) }
	for i:=1;i<=1000000;i++ {
		for j:=i;j<=1000000;j+=i {
			factors[j] = append(factors[j],i)
		}
	}
	//fmt.Fprintf(os.Stderr,"HERE2\n")

	// Prework 2 -- create maximal depth of matrygons
	best2matrygon := make([]int,1000001)
	best3matrygon := make([]int,1000001)
	best2matrygon[2] = 1
	for i:=3;i<=1000000;i++ {
		best2,best3 := 1,1
		for _,f := range factors[i] {
			if f == 1 || f == i { continue }
			b := i/f-1
			v := 1+best2matrygon[b]
			best2 = max(best2,v)
			if f > 2 { best3 = max(best3,v) }
		}
		best2matrygon[i] = best2
		best3matrygon[i] = best3
	}
	//fmt.Fprintf(os.Stderr,"HERE3\n")

    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi(); ans := best3matrygon[N]
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

