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
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
type pair struct {i,j int}

func solve(R,C int, S [][]int) int {
	lj,rj,ui,di := twodi(R,C,-1),twodi(R,C,-1),twodi(R,C,-1),twodi(R,C,-1)
	for i:=0;i<R;i++ {
		for j:=0;j<C;j++ {
			if i > 0 { ui[i][j] = i-1 }
			if i+1 < R { di[i][j] = i+1 }
			if j > 0 { lj[i][j] = j-1 }
			if j+1 < C { rj[i][j] = j+1 }
		}
	}
	ans := 0
	sumarr := 0; for i:=0;i<R;i++ { for j:=0;j<C;j++ { sumarr += S[i][j] } }
	evalset := make(map[pair]bool); for i:=0;i<R;i++ { for j:=0;j<C;j++ { evalset[pair{i,j}]=true } }
	evallist := make([]pair,0)
	for {
		ans += sumarr
		evallist = evallist[:0]; for k := range evalset { evallist = append(evallist,k) }
		elimset := make(map[pair]bool)
		for _,p := range evallist {
			if S[p.i][p.j] == 0 { continue }
			numneigbors,sumneighbors := 0,0
			if lj[p.i][p.j] != -1 { numneigbors++; sumneighbors += S[p.i][lj[p.i][p.j]] }
			if rj[p.i][p.j] != -1 { numneigbors++; sumneighbors += S[p.i][rj[p.i][p.j]] }
			if ui[p.i][p.j] != -1 { numneigbors++; sumneighbors += S[ui[p.i][p.j]][p.j] }
			if di[p.i][p.j] != -1 { numneigbors++; sumneighbors += S[di[p.i][p.j]][p.j] }
			if sumneighbors > numneigbors * S[p.i][p.j] { elimset[p] = true }
		}
		if len(elimset) == 0 { break }
		evalset = make(map[pair]bool)
		for p := range elimset {
			sumarr -= S[p.i][p.j]
			if lj[p.i][p.j] != -1 { ii,jj := p.i,lj[p.i][p.j]; evalset[pair{ii,jj}] = true; rj[ii][jj] = rj[p.i][p.j] }
			if rj[p.i][p.j] != -1 { ii,jj := p.i,rj[p.i][p.j]; evalset[pair{ii,jj}] = true; lj[ii][jj] = lj[p.i][p.j] } 
			if ui[p.i][p.j] != -1 { ii,jj := ui[p.i][p.j],p.j; evalset[pair{ii,jj}] = true; di[ii][jj] = di[p.i][p.j] }
			if di[p.i][p.j] != -1 { ii,jj := di[p.i][p.j],p.j; evalset[pair{ii,jj}] = true; ui[ii][jj] = ui[p.i][p.j] }
			S[p.i][p.j] = 0
		}
	}
	return ans
}

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C := gi(),gi()
		S := make([][]int,R); for i:=0;i<R;i++ { S[i] = make([]int,C) }
		for i:=0;i<R;i++ { for j:=0;j<C;j++ { S[i][j] = gi() } }
		ans := solve(R,C,S)
		fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
	}
}

