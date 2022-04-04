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
		N := gi(); bd := make([]string,N); for i:=0;i<N;i++ { bd[i] = gs() }

		getRow := func(i int) []int {
			ans := make([]int,0)
			for j:=0;j<N;j++ { if bd[i][j] == 'X' { ans = append(ans,j)} }
			return ans
		}

		getCol := func(j int) []int {
			ans := make([]int,0)
			for i:=0;i<N;i++ { if bd[i][j] == 'X' { ans = append(ans,i)} }
			return ans
		}

		checkBoard := func() bool {
			rowsb := make([]bool,N)
			colsb := make([]bool,N)
			centerFound := false

			for i:=0;i<N;i++ {
				if rowsb[i] { continue }
				r1 := getRow(i)
				if len(r1) == 1 {
					c1 := getCol(r1[0])
					if centerFound || colsb[r1[0]] || len(c1) != 1 { return false }
					rowsb[i] = true; colsb[r1[0]] = true; centerFound = true
				} else if len(r1) == 2 {
					c1 := getCol(r1[0]); c2 := getCol(r1[1])
					if len(c1) != 2 || len(c2) != 2 || c1[0] != c2[0] || c1[1] != c2[1] || rowsb[c1[0]] || rowsb[c1[1]] { return false }
					rr1 := getRow(c1[0]); rr2 := getRow(c1[1])
					if len(rr1) != 2 || len(rr2) != 2 || rr1[0] != rr2[0] || rr1[1] != rr2[1] || colsb[rr1[0]] || colsb[rr1[1]] { return false }
					rowsb[c1[0]] = true; rowsb[c1[1]] = true; colsb[rr1[0]] = true; colsb[rr1[1]] = true
				} else {
					return false
				}
			}
			return true
		}
		ans := "IMPOSSIBLE"
		if checkBoard() { ans = "POSSIBLE"}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,ans)
    }
}

