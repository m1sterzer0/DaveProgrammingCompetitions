package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func ia(m int) []int { return make([]int,m) }
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); A := gis(N); B := gis(N); Q := gi(); X,Y := fill2(Q); for i:=0;i<Q;i++ { X[i]--; Y[i]-- }
	aset,bset,diffhash := make(map[int]bool),make(map[int]bool),make(map[int]bool)
	diffcnt,aptr,bptr,acnt,bcnt := 0,0,0,0,0
	cnta,cntb,good := make([]int,N),make([]int,N),make([]bool,N+1)
	for {
		for aptr < N && aset[A[aptr]] { cnta[aptr] = acnt; aptr++ }
		for bptr < N && bset[B[bptr]] { cntb[bptr] = bcnt; bptr++ }
		if aptr < N { acnt++; aset[A[aptr]] = true }
		if bptr < N { bcnt++; bset[B[bptr]] = true }
		if aptr == N && bptr == N { break }
		if aptr < N && bptr < N && acnt == bcnt {
			a,b := A[aptr],B[bptr]
			if diffhash[a] { diffcnt--; diffhash[a] = false } else { diffcnt++; diffhash[a] = true }
			if diffhash[b] { diffcnt--; diffhash[b] = false } else { diffcnt++; diffhash[b] = true }
			if diffcnt == 0 { good[acnt] = true }
		}
	}
	for i:=0;i<Q;i++ {
		x,y := X[i],Y[i]
		a,b := cnta[x],cntb[y]
		ans := "No"; if a == b && good[a] { ans = "Yes" }
		fmt.Fprintln(wrtr,ans)
	}

}

