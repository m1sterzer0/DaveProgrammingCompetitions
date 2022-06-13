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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	T := gi()
	for tt:=1;tt<=T;tt++ {
		N := gi(); P := gis(N); for i:=0;i<N;i++ { P[i]-- }
		ansarr := make([]int,N*N)
		for b:=1;b<=N-3;b++ {
			pb := P[b]
			numg := 0
			if P[N-1] < pb { numg++ }
			for c:=N-2;c>b;c-- {
				id := N*b+c
				ansarr[id] = numg
				if P[c] < pb { numg++ }
			}
		}
		for c:= N-2;c>=2;c-- {
			pc := P[c]
			numg := 0
			if P[0] < pc { numg++ }
			for b:=1;b<c;b++ {
				id := N*b+c
				ansarr[id] *= numg
				if P[b] < pc { numg++ }
			}
		}
		ans := int64(0); for _,a := range ansarr { ans += int64(a) }
		fmt.Fprintln(wrtr,ans)
	}
}

