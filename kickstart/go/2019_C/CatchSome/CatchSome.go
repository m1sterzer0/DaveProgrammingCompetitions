package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func min(a,b int) int { if a > b { return b }; return a }
const inf int = 2000000000000000000
type dog struct { a,p int }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,K := gi2(); P := gis(N); A := gis(N); for i:=0;i<N;i++ { A[i]-- }
		dogs := make([]dog,0)
		for i:=0;i<N;i++ { dogs = append(dogs,dog{A[i],P[i]}) }
		sort.Slice(dogs,func(i,j int) bool { return dogs[i].p < dogs[j].p} )
		cdogs := make([][]int,1000)
		for _,d := range dogs { cdogs[d.a] = append(cdogs[d.a],d.p) }
		olddp1 := iai(N+1,inf); olddp2 := iai(N+1,inf); dp1 := iai(N+1,inf); dp2 := iai(N+1,inf)
		olddp1[0] = 0; olddp2[0] = 0; dp1[0] = 0; dp2[0] = 0; dogssofar := 0
		for c:=0;c<1000;c++ {
			if len(cdogs[c]) == 0 { continue }
			newdogs := len(cdogs[c])
			olddp1,olddp2,dp1,dp2 = dp1,dp2,olddp1,olddp2
			for i:=0;i<=dogssofar;i++ { dp1[i] = olddp1[i]; dp2[i] = olddp2[i] }
			for i:=1;i<=newdogs;i++ {
				for j:=0;j<=dogssofar;j++ {
					dp1[i+j] = min(dp1[i+j],olddp1[j]+2*cdogs[c][i-1])
					dp2[i+j] = min(dp2[i+j],olddp2[j]+2*cdogs[c][i-1])
					dp2[i+j] = min(dp2[i+j],olddp1[j]+cdogs[c][i-1])
				}
			}
			dogssofar += newdogs
		}
		ans := dp2[K]
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

