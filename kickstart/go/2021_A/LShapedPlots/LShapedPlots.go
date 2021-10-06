package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func min(a,b int) int { if a > b { return b }; return a }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C := gi2(); bd := make([][]int,R); for i:=0;i<R;i++ { bd[i] = gis(C) }
		lf := twodi(R,C,0)
		rt := twodi(R,C,0)
		up := twodi(R,C,0)
		dn := twodi(R,C,0)
		for i:=0;i<R;i++ {
			for x,j:=0,0;j<C;j++    { if bd[i][j] == 0 { x = 0 } else { x++ }; lf[i][j] = x }
			for x,j:=0,C-1;j>=0;j-- { if bd[i][j] == 0 { x = 0 } else { x++ }; rt[i][j] = x }
		}
		for j:=0;j<C;j++ {
			for x,i:=0,0;i<R;i++    { if bd[i][j] == 0 { x = 0 } else { x++ }; up[i][j] = x }
			for x,i:=0,R-1;i>=0;i-- { if bd[i][j] == 0 { x = 0 } else { x++ }; dn[i][j] = x }
		}

		cnt := func(a,b int) int {
			ans := 0
			if a >= 2 && b >= 4 { ans += min(a-1,b/2-1) }
			if b >= 2 && a >= 4 { ans += min(b-1,a/2-1) }
			return ans
		}
		ans := 0
		for i:=0;i<R;i++ {
			for j:=0;j<C;j++ {
				ans += cnt(lf[i][j],up[i][j])
				ans += cnt(lf[i][j],dn[i][j])
				ans += cnt(rt[i][j],up[i][j])
				ans += cnt(rt[i][j],dn[i][j])
			}
		}
       fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}