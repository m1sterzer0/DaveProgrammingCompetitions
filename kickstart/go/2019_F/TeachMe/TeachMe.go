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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,_ := gi2(); A := make([][]int,N)
		for i:=0;i<N;i++ { c := gi(); a := gis(c); sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } ); A[i] = a }
		sb := make(map[int]int)
		encode := func(a []int, mask int) int {
			res := 0; base := 1001; m := 1
			for i:=0;i<len(a);i++ {
				if mask & (1<<uint(i)) == 0 { continue }
				res += m * a[i]; m *= base
			}
			return res
		}
		for i:=0;i<N;i++ {
			numelem := len(A[i])
			for mask:=0;mask<1 << uint(numelem); mask++ {
				enc := encode(A[i],mask)
				sb[enc]++
			} 
		}
		ans := 0
		for i:=0;i<N;i++ { e := encode(A[i], 1 << uint(len(A[i])) - 1); ncands := N - sb[e]; ans += ncands }
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

