package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func fsieve(n int) []int {
	fs := make([]int,n+1); fs[0] = 0; fs[1] = 1
	for i:=2;i<=n;i++ { fs[i] = -1 };
	for i:=2;i<=n;i+=2 { fs[i] = 2 }
	for i:=3;i<=n;i+=2 { if fs[i] == -1 { fs[i] = i; inc := 2*i; for k:=i*i;k<=n;k+=inc { if fs[k] == -1 { fs[k] = i } } } }
	return fs
}
type pstate struct { x1,x2,h uint64 }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,Q := gi(),gi(); A := gis(N); L,R := fill2(Q); for i:=0;i<Q;i++ { L[i]--; R[i]-- }
	f := fsieve(1000010); r := uint64(0)
	hasher := make(map[int]pstate)
	rand.Seed(8675309)
	cumharr := make([]uint64,N)
	addPrime := func(n int) {
		v,ok := hasher[n]
		if !ok { v = pstate{rand.Uint64(),rand.Uint64(),uint64(0)} }
		if v.h == 0 { r = r ^ v.x1; v.h++ } else if v.h == 1 { r = r ^ v.x1 ^ v.x2; v.h++ } else { r = r ^ v.x2; v.h = 0 }
		hasher[n] = v
	}
	for i:=0;i<N;i++ {
		a := A[i]
		for a != 1 { p := f[a]; addPrime(p); a /= p }
		cumharr[i] = r
	}
	for i:=0;i<Q;i++ {
		l,r := L[i],R[i]; hval := cumharr[r]; if l > 0 { hval = hval ^ cumharr[l-1] }
		if hval == 0 { fmt.Fprintln(wrtr,"Yes") } else { fmt.Fprintln(wrtr,"No") }
	}
}
