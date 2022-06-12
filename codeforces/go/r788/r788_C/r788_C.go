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
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
const MOD = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); A := gis(N); for i:=0;i<N;i++ { A[i]-- }; B := gis(N); for i:=0;i<N;i++ { B[i]-- }; D := gis(N); for i:=0;i<N;i++ { D[i]-- }
		Arev := make([]int,N); for i:=0;i<N;i++ { Arev[A[i]] = i }
		C := iai(N,-1); ans := 1
		cycleChase := func (idx int) {
			targ := A[idx]; C[idx] = targ;
			for B[idx] != targ { idx = Arev[B[idx]]; C[idx] = A[idx] }
		}
		for i:=0;i<N;i++ { 
			if C[i] == -1 && A[i] == B[i] { C[i] = A[i]; continue }
			if C[i] == -1 && D[i] != -1 { cycleChase(i) }
		}
		for i:=0;i<N;i++ {
			if C[i] != -1 { continue }
			ans *= 2; ans %= MOD
			cycleChase(i)
		}
		fmt.Fprintln(wrtr,ans)
	}
}

