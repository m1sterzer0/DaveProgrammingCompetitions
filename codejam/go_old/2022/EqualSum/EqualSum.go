package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N := gi()
		A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			       31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
				   59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 128, 256, 512, 1024, 2048, 4096, 8192,
				   16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728,
				   268435456, 536870912 }
		astr := vecintstring(A)
		fmt.Fprintf(wrtr,"%v\n",astr); wrtr.Flush()
		B := gis(N)
		C := make([]int,0)
		for _,a := range A { C = append(C,a) }
		for _,b := range B { C = append(C,b) }
		sort.Slice(C,func(i,j int) bool { return C[i] > C[j] } )
		X := make([]int,0)
		Y := make([]int,0)
		xsum,ysum := 0,0
		for _,c := range C {
			if xsum <= ysum { X = append(X,c); xsum += c } else { Y = append(Y,c); ysum += c }
		}
		if xsum != ysum { fmt.Fprintf(os.Stderr,"MISMATCHED SUMS.  SOMETHING BAD HAPPENED\n"); os.Exit(1) }
		ans := vecintstring(X)
		fmt.Fprintf(wrtr,"%v\n",ans); wrtr.Flush() 
    }
}

