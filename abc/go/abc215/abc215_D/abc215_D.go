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
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }

func fsieve(n int) []int {
	s := iai(n+1,-1)
	s[0] = 0; s[1] = 1; for i:=2;i<=n;i+=2 { s[i] = 2 }
	for i:=3;i<=n;i+=2 { 
		if s[i] != -1 { continue }
		s[i] = i
		for k:=i*i;k<=n;k+=2*i { if s[k] == -1 { s[k] = i } }
	}
	return s
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N,M := gi2(); A := gis(N)
	f := fsieve(100_000)
	pfactors := make([]bool,100_001)
	for _,a := range A {
		aa := a
		for aa != 1 {
			p := f[aa];
			pfactors[p] = true
			for aa % p == 0 { aa /= p } 
		}
	}
	sb := make([]bool,100_001); ans := []int{1}; sb[1] = true
	for i:=2;i<=M;i++ { 
		p := f[i]
		if !pfactors[p] && sb[i/p] { ans = append(ans,i); sb[i] = true }
	}
	fmt.Fprintln(wrtr,len(ans))
	for _,a := range ans { fmt.Fprintln(wrtr,a) }
}
