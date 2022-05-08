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
	N := gi(); A := gis(N)
	sb := make([]int,200001); for _,a := range A { sb[a]++ }
	ans := 0
	// Four cases: 1*1=1, 1*x=x,x*x=y,x*y=z
	ans += sb[1]*sb[1]*sb[1] // 1*1=1
	if sb[1] >= 0 { for x:=2;x<=200000;x++ { if sb[x] >= 1 { ans += 2*sb[1]*sb[x]*sb[x] } } } // 1*x==x
	for i:=2;i*i<=200000;i++ { if sb[i] >= 1 && sb[i*i] >= 1 { ans += sb[i*i]*sb[i]*sb[i] } } // x*x==y
	for x:=2;x<=200000;x++ { 
		if sb[x] == 0 { continue }
		for y:=x+1;x*y<=200000;y++ {
			if sb[y] == 0 { continue }
			if sb[x*y] == 0 { continue }
			ans += 2*sb[x]*sb[y]*sb[x*y]
		}
	}
	fmt.Println(ans)
}

