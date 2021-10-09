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
func gi2() (int,int) { return gi(),gi() }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	type op struct {a,c int}
	N,M := gi2(); A,C := fill2(M)
	ops := make([]op,M)
	for i:=0;i<M;i++ { ops[i] = op{A[i],C[i]} }
	sort.Slice(ops,func(i,j int)bool{return ops[i].c < ops[j].c})
	ans := 0
	for _,xx := range ops {
		inc := xx.a % N
		if inc == 0 { continue }
		g := gcd(inc,N)
		ans,N = ans+xx.c*(N-g),g
	}
	if N == 1 { fmt.Println(ans) } else { fmt.Println(-1) }
}

