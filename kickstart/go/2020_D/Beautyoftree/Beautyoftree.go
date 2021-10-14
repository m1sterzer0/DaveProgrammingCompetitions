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
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,A,B := gi3(); P := gis(N-1); for i:=0;i<N-1;i++ { P[i]-- }
		gr := make([][]int,N)
		for i,p := range P { gr[p] = append(gr[p],i+1) }
		arra := ia(N); arrb := ia(N); st := ia(0)
		var dfs func(n int)
		dfs = func(n int) {
			st = append(st,n)
			for _,c := range gr[n] { dfs(c) }
			arra[n]++; arrb[n]++
			if len(st) > A { p := st[len(st)-1-A]; arra[p] += arra[n] }
			if len(st) > B { p := st[len(st)-1-B]; arrb[p] += arrb[n] }
			st = st[:len(st)-1]
		}
		dfs(0)
		num := N * sumarr(arra) + N * sumarr(arrb)
		for i:=0;i<N;i++ { num -= arra[i]*arrb[i] }
		ans := float64(num) / float64(N*N)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

