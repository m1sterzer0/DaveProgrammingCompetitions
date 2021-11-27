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
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE  (Straight up transcription of a passing answer -- need to dive into the )
	A := gis(9)
	if A[0] + A[2] + A[4] + A[6] + A[8] != A[1] + A[3] + A[5] + A[7] { fmt.Println("NO"); return }
	for a:=0; a <= min(A[0]-1,A[3]-1); a++ {
		for b:=0; b <= min(A[6]-1,A[7]-1); b++ {
			for c:=0; c <= min(A[8]-1,A[5]-1); c++ {
				for d:=0; d <= min(A[2]-1,A[1]-1); d++ {
					if A[3]-a>=A[6]-b && A[7]-b>=A[8]-c && A[5]-c>=A[2]-d && A[1]-d>=A[0]-a {
						barr := []byte{}
						dn := func() { barr = append(barr,'D') }
						up := func() { barr = append(barr,'U') }
						lf := func() { barr = append(barr,'L') }
						rt := func() { barr = append(barr,'R') }
						e,f,g,h := A[6]-b-1,A[8]-c-1,A[2]-d-1,A[0]-a-1
						w,x,y,z := A[3]-a-e-1,A[7]-b-f-1,A[5]-c-g-1,A[1]-d-h-1
						dn()
						for i:=0;i<a;i++ { up(); dn() }
						for i:=0;i<w;i++ { rt(); lf() }
						dn()
						for i:=0;i<e;i++ { up(); dn() }
						rt()
						for i:=0;i<b;i++ { lf(); rt() }
						for i:=0;i<x;i++ { up(); dn() }
						rt()
						for i:=0;i<f;i++ { lf(); rt() }
						up()
						for i:=0;i<c;i++ { dn(); up() }
						for i:=0;i<y;i++ { lf(); rt() }
						up()
						for i:=0;i<g;i++ { dn(); up() }
						lf()
						for i:=0;i<d;i++ { rt(); lf() }
						for i:=0;i<z;i++ { dn(); up() }
						lf()
						for i:=0;i<h;i++ { rt(); lf() }
						ans := string(barr)
						fmt.Println(ans)
						return
					 }
				}
			}
		}
	}
	fmt.Println("NO")
}

