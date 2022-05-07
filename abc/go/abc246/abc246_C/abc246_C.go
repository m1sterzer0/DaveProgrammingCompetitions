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
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,K,X := gi(),gi(),gi(); A := gis(N)
	full := 0; for _,a := range(A) { full += a/X }
	ans := sumarr(A)
	if full >= K { 
		ans -= X*K
	} else {
		ans -= X*full; K -= full
		r := make([]int,N)
		for i:=0;i<N;i++ { r[i] = A[i] % X }
		sort.Slice(r,func(i,j int) bool { return r[i] > r[j] } )
		for i:=0;i<N && K>0;i++ { ans -= r[i]; K-- }
	}
	fmt.Println(ans)
}

