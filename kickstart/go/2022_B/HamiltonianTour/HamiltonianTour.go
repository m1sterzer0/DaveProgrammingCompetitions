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
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		R,C := gi(),gi(); B := make([]string,R); for i:=0;i<R;i++ { B[i] = gs() }
		visited := make([][]bool,R); for i:=0;i<R;i++ { visited[i] = make([]bool,C) }
		ansarr := make([]byte,0)
		var dfs func(i,j,ii,jj int,src byte)
		dfs = func(i,j,ii,jj int,src byte) {
			visited[i][j] = true
			ac := func(c byte) { ansarr = append(ansarr,c) }
			checkWest := func() { if j > 0    && !visited[i][j-1] && B[i][j-1] != '#' { ac('W'); dfs(i,j-1,ii,1,'E') } else if ii == 0 { ac('S') } else { ac('N') }; ii = 1-ii }
			checkEast := func() { if j+1 < C  && !visited[i][j+1] && B[i][j+1] != '#' { ac('E'); dfs(i,j+1,ii,0,'W') } else if ii == 0 { ac('S') } else { ac('N') }; ii = 1-ii }
			checkNorth := func() { if i > 0   && !visited[i-1][j] && B[i-1][j] != '#' { ac('N'); dfs(i-1,j,1,jj,'S') } else if jj == 0 { ac('E') } else { ac('W') }; jj = 1-jj }
			checkSouth := func() { if i+1 < R && !visited[i+1][j] && B[i+1][j] != '#' { ac('S'); dfs(i+1,j,0,jj,'N') } else if jj == 0 { ac('E') } else { ac('W') }; jj = 1-jj }
			if ii == 0 && jj == 0 && src == 'N' { checkWest(); checkSouth(); checkEast();  ansarr = append(ansarr,'N'); return }
			if ii == 0 && jj == 0 && src == 'W' { checkNorth(); checkEast(); checkSouth(); ansarr = append(ansarr,'W'); return }
			if ii == 0 && jj == 1 && src == 'N' { checkEast(); checkSouth(); checkWest();  ansarr = append(ansarr,'N'); return }
			if ii == 0 && jj == 1 && src == 'E' { checkNorth(); checkWest(); checkSouth(); ansarr = append(ansarr,'E'); return }
			if ii == 1 && jj == 0 && src == 'S' { checkWest(); checkNorth(); checkEast();  ansarr = append(ansarr,'S'); return }
			if ii == 1 && jj == 0 && src == 'W' { checkSouth(); checkEast(); checkNorth(); ansarr = append(ansarr,'W'); return }
			if ii == 1 && jj == 1 && src == 'S' { checkEast(); checkNorth(); checkWest();  ansarr = append(ansarr,'S'); return }
			if ii == 1 && jj == 1 && src == 'E' { checkSouth(); checkWest(); checkNorth(); ansarr = append(ansarr,'E'); return }
		}
		dfs(0,0,0,0,'N')
		ansarr[len(ansarr)-1] = 'W' // patch up the loop
		numempty := 0
		for i:=0;i<R;i++ { for j:=0;j<C;j++ { if B[i][j] == '*' { numempty++ } } }
		if len(ansarr) != 4*numempty {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,"IMPOSSIBLE")
		} else {
			fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,string(ansarr))
		}
	}
}
