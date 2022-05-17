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
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
type pair struct {x,y int}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C,Rs,Cs,S := gi(),gi(),gi(),gi(),gi()
		P,Q := gf(),gf()
		gr := make([][]byte,R); for i:=0;i<R;i++ { gr[i] = make([]byte,C) }
		for i:=0;i<R;i++ { for j:=0;j<C;j++ { s := gs(); gr[i][j] = byte(s[0]) } }
		dd := []pair{{-1,0},{1,0},{0,-1},{0,1}}
		PP := make([]float64,10); cc := 1.0; for i:=0;i<10;i++ { PP[i] = cc * P; cc *= (1.0-P) }
		QQ := make([]float64,10); cc = 1.0;  for i:=0;i<10;i++ { QQ[i] = cc * Q; cc *= (1.0-Q) }
		count := twodi(R,C,0)
		var dfs func(x,y,s int) float64
		dfs = func(x,y,s int) float64 {
			if s == 0 { return 0.00 }
			best := 0.00
			for _,d := range dd {
				xx,yy := x+d.x,y+d.y
				if xx < 0 || yy < 0 || xx >= R || yy >= C { continue }
				lev := 0.0; if gr[xx][yy] == 'A' { lev += PP[count[xx][yy]] } else { lev += QQ[count[xx][yy]] }
				count[xx][yy]++; rev := dfs(xx,yy,s-1); count[xx][yy]--
				if lev+rev > best { best = lev+rev }
			}
			return best
		}
		ans := dfs(Rs,Cs,S)
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

