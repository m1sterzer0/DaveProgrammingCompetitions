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
		n := gi(); a := gis(n)
		lastl,lastr := 1000000,1000000
		ans := "YES"
		for i:=n-1;i>=0;i-=2 {
			if i == 0 {
				if a[i] > lastl || a[i] > lastr { ans = "NO" }
				break
			}
			if (a[i] > lastl || a[i-1] > lastr || a[i] > lastr || a[i-1] > lastl ) { ans = "NO"; break }
			lastl,lastr = a[i],a[i-1] 
		}
		fmt.Fprintln(wrtr,ans)
	}
}

